package routeros

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var reMetadataFields = regexp.MustCompile(`^___\S+___$`)
var reTransformSet = regexp.MustCompile(`"\s*?(\S+?)\s*?\s*?:\s*?\s*?(\S+?)\s*?"`)
var reSkipFields = regexp.MustCompile(`"\s*?(\S+?)\s*?"\s*?`)

// GetMetadata Get item metadata fields from resource schema.
func GetMetadata(s map[string]*schema.Schema) *MikrotikItemMetadata {
	meta := &MikrotikItemMetadata{}
	// Schema map iteration.
	for terraformSnakeName, terraformMetadata := range s {
		if reMetadataFields.MatchString(terraformSnakeName) {
			switch terraformSnakeName {
			case MetaId:
				meta.IdType = IdType(terraformMetadata.Default.(int))
			case MetaResourcePath:
				meta.Path = terraformMetadata.Default.(string)
			default:
				if meta.Meta == nil {
					meta.Meta = make(map[string]string)
				}
				meta.Meta[terraformSnakeName] = terraformMetadata.Default.(string)
			}
		}
	}
	return meta
}

func isEmpty(propName string, schemaProp *schema.Schema, d *schema.ResourceData, confValue cty.Value) bool {
	v := d.Get(propName)
	switch schemaProp.Type {
	case schema.TypeString:
		if schemaProp.Default != nil {
			return v.(string) == "" && schemaProp.Default.(string) == ""
		}
		return v.(string) == "" && confValue.IsNull()
	case schema.TypeFloat, schema.TypeInt:
		return confValue.IsNull() && schemaProp.Default == nil
	case schema.TypeBool:
		// If true, it is always not empty:
		if v.(bool) {
			return false
		}
		// Use the default value:
		if schemaProp.Default != nil {
			return false
		}
		return confValue.IsNull()
	case schema.TypeList:
		if confValue.Type().ElementType().IsObjectType() {
			return len(v.([]interface{})) == 0
		}
		return len(v.([]interface{})) == 0 && confValue.IsNull()
	case schema.TypeSet:
		return v.(*schema.Set).Len() == 0 && confValue.IsNull()
	case schema.TypeMap:
		return len(v.(map[string]interface{})) == 0
	default:
		panic("[isEmpty] wrong resource type: " + schemaProp.Type.String())
	}
}

// loadTransformSet Converting the metadata of the 'MetaTransformSet' field into a map for forward and reverse
// transformation. This map should be applied before the logical processing of fields at the beginning of
// serialization/deserialization.
// Forward transformation for use in the 'MikrotikResourceDataToTerraform' function, reverse transformation for use
// in the 'TerraformResourceDataToMikrotik' function.
// s: "channel.config: channel","datapath.config: datapath"` in the Mikrotik (kebab) notation!
func loadTransformSet(s string, reverse bool) (m map[string]string) {
	m = make(map[string]string)
	for _, b := range reTransformSet.FindAllStringSubmatch(s, -1) {
		if !reverse {
			m[b[1]] = b[2]
		} else {
			m[b[2]] = b[1]
		}
	}
	return
}

// loadSkipFields A list of fields that will not be serialized and transferred to Mikrotik.
func loadSkipFields(s string) (m map[string]struct{}) {
	m = make(map[string]struct{})
	for _, b := range reSkipFields.FindAllStringSubmatch(s, -1) {
		m[b[1]] = struct{}{}
	}
	return
}

// ListToString Convert List and Set to a delimited string.
func ListToString(v any) (res string) {
	for i, elem := range v.([]interface{}) {
		if i > 0 {
			res += fmt.Sprintf(",%v", elem)
		} else {
			res = fmt.Sprint(elem)
		}
	}
	return
}

// TerraformResourceDataToMikrotik Marshal Mikrotik resource from TF resource schema.
func TerraformResourceDataToMikrotik(s map[string]*schema.Schema, d *schema.ResourceData) (MikrotikItem, *MikrotikItemMetadata) {
	item := MikrotikItem{}
	meta := &MikrotikItemMetadata{}
	rawConfig := d.GetRawConfig()
	var transformSet map[string]string
	var skipFields, setUnsetFields map[string]struct{}

	// {"channel.config: channel", "datapath.config: datapath", "schema-field-name": "mikrotik-field-name"}
	if ts, ok := s[MetaTransformSet]; ok {
		transformSet = loadTransformSet(ts.Default.(string), false)
	}

	// "field_first", "field_second", "field_third"
	if sf, ok := s[MetaSkipFields]; ok {
		skipFields = loadSkipFields(sf.Default.(string))
	}
	if suf, ok := s[MetaSetUnsetFields]; ok {
		setUnsetFields = loadSkipFields(suf.Default.(string))
	}

	// Schema map iteration.
	for terraformSnakeName, terraformMetadata := range s {
		// Fill in the metadata fields.
		if reMetadataFields.MatchString(terraformSnakeName) {
			switch terraformSnakeName {
			case MetaId:
				meta.IdType = IdType(terraformMetadata.Default.(int))
			case MetaResourcePath:
				meta.Path = terraformMetadata.Default.(string)
			case MetaTransformSet, MetaSkipFields, MetaSetUnsetFields, MetaDropByValue:
				continue
			default:
				meta.Meta[terraformSnakeName] = terraformMetadata.Default.(string)
			}
			continue
		}

		// Skip the fields specified in the schema.
		if _, ok := skipFields[terraformSnakeName]; ok {
			continue
		}

		// Skip all read-only properties.
		if terraformMetadata.Computed && !terraformMetadata.Optional {
			continue
		}

		/*
			Skip all empty Optional fields.
			This logic may be broken, but I don't have enough examples to test it.

			All fields checked in Winbox must be Optional: true & Computed: true.
			Otherwise, there will be an error: "After applying this test step, the plan was not empty."

			        Terraform will perform the following actions:

			          # routeros_interface_bridge.test_bridge will be updated in-place
			          ~ resource "routeros_interface_bridge" "test_bridge" {
			              - fast_forward        = true -> null
			                id                  = "*DD"
			                name                = "test_bridge"
			                # (22 unchanged attributes hidden)
			            }
		*/
		/*
			old, new := d.GetChange(terraformSnakeName)
			conf := d.GetRawConfig().GetAttr(terraformSnakeName).IsNull()
			fmt.Println(rawConfig.GetAttr(terraformSnakeName).IsKnown())
			fmt.Printf("%25s - old: '%10v', new: '%10v', isNull: %v", terraformSnakeName, old, new, conf)
		*/
		if terraformMetadata.Optional && !d.HasChange(terraformSnakeName) &&
			isEmpty(terraformSnakeName, s[terraformSnakeName], d, rawConfig.GetAttr(terraformSnakeName)) {
			// fmt.Println(" ... skipped")
			continue
		}
		// fmt.Println()

		// terraformSnakeName = fast_forward, schemaPropData = true
		// NewMikrotikItem.Fields["fast-forward"] = "true"
		mikrotikKebabName := SnakeToKebab(terraformSnakeName)
		value := d.Get(terraformSnakeName)

		// WiFi basic_rates_ag -> basic-rates-a/g
		if transformSet != nil && terraformMetadata.Type != schema.TypeMap {
			if new, ok := transformSet[terraformSnakeName]; ok {
				mikrotikKebabName = SnakeToKebab(new)
			}
		}

		switch terraformMetadata.Type {
		case schema.TypeString:
			if _, ok := setUnsetFields[terraformSnakeName]; ok && value.(string) == "" {
				// Unset
				item["!"+mikrotikKebabName] = ""
				continue
			}
			item[mikrotikKebabName] = value.(string)
		case schema.TypeFloat:
			item[mikrotikKebabName] = strconv.FormatFloat(value.(float64), 'f', -1, 64)
		case schema.TypeInt:
			item[mikrotikKebabName] = strconv.Itoa(value.(int))
		case schema.TypeBool:
			// true:  {...,"interfaces":"ether3","passive":"","priority":"128",...}
			// false: {...,"interfaces":"ether3",             "priority":"128",...}
			if _, ok := setUnsetFields[terraformSnakeName]; ok {
				if value.(bool) {
					item[mikrotikKebabName] = ""
				} else {
					// Unset
					item["!"+mikrotikKebabName] = ""
				}
				continue
			}
			item[mikrotikKebabName] = BoolToMikrotikJSON(value.(bool))
		// Used to represent an ordered collection of items.
		case schema.TypeList:

			switch terraformMetadata.Elem.(type) {
			case *schema.Schema:

				item[mikrotikKebabName] = ListToString(value)

			case *schema.Resource:

				// skip if object is empty
				if value.([]interface{})[0] == nil {
					continue
				}

				list := value.([]interface{})[0].(map[string]interface{})
				ctyList := rawConfig.GetAttr(terraformSnakeName).AsValueSlice()[0]

				for fieldName, value := range list {
					// "output.0.affinity"
					filedNameInState := fmt.Sprintf("%v.%v.%v", terraformSnakeName, 0, fieldName)
					fieldSchema := terraformMetadata.Elem.(*schema.Resource).Schema[fieldName]

					// Skip all read-only properties.
					if fieldSchema.Computed && !fieldSchema.Optional {
						continue
					}

					if fieldSchema.Optional && !d.HasChange(filedNameInState) &&
						isEmpty(filedNameInState, fieldSchema, d, ctyList.GetAttr(fieldName)) {
						continue
					}
					fieldName = SnakeToKebab(mikrotikKebabName + "." + fieldName)

					switch value := value.(type) {
					case string:
						item[fieldName] = value
					case int:
						item[fieldName] = strconv.Itoa(value)
					case bool:
						item[fieldName] = BoolToMikrotikJSON(value)
					}
				}
			}
			// Used to represent an unordered collection of items.
		case schema.TypeSet:
			if _, ok := setUnsetFields[terraformSnakeName]; ok {
				// policy = ["api", "read", "winbox"] -> ["api", "read"]
				// old = ... read, !telnet, !rommon, api, !local, winbox
				// new = api, read
				var res []string
				for _, v := range d.GetRawConfig().GetAttr(terraformSnakeName).AsValueSet().Values() {
					res = append(res, v.AsString())
				}

				old, _ := d.GetChange(terraformSnakeName)
				for _, v := range old.(*schema.Set).List() {
					elem := v.(string)
					if len(elem) > 0 && elem[0] != '!' && !slices.Contains(res, elem) {
						res = append(res, "!"+elem)
					}
				}
				item[mikrotikKebabName] = strings.Join(res, ",")
				continue
			}
			item[mikrotikKebabName] = ListToString(value.(*schema.Set).List())
		case schema.TypeMap:
			for k, v := range value.(map[string]interface{}) {
				// channel + "." + config
				k = SnakeToKebab(mikrotikKebabName + "." + k)
				// Field transformation: "channel.config" ---> "channel".
				if transformSet != nil {
					if new, ok := transformSet[k]; ok {
						k = new
					}
				}

				// Conversion of boolean values.
				s := BoolToMikrotikJSONStr(v.(string))

				item[k] = s
			}
		default:
			panic(fmt.Sprintf("[TerraformResourceDataToMikrotik] resource type not implemented: %v for '%v'",
				terraformMetadata.Type, terraformSnakeName))
		}

	}

	return item, meta
}

// MikrotikResourceDataToTerraform Unmarshal Mikrotik resource (incoming data: JSON, etc.) to TF resource schema.
func MikrotikResourceDataToTerraform(item MikrotikItem, s map[string]*schema.Schema, d *schema.ResourceData) diag.Diagnostics {
	var diags diag.Diagnostics
	var err error
	var transformSet map[string]string
	var setUnsetFields, skipFields, dropByValue map[string]struct{}

	// {"channel": "channel.config", "mikrotik-field-name": "schema-field-name"}
	if ts, ok := s[MetaTransformSet]; ok {
		transformSet = loadTransformSet(ts.Default.(string), true)
	}

	// "field_first", "field_second", "field_third"
	if suf, ok := s[MetaSetUnsetFields]; ok {
		setUnsetFields = loadSkipFields(suf.Default.(string))
	}
	if sf, ok := s[MetaSkipFields]; ok {
		skipFields = loadSkipFields(sf.Default.(string))
	}

	if dbv, ok := s[MetaDropByValue]; ok {
		dropByValue = loadSkipFields(dbv.Default.(string))
	}

	// TypeMap,TypeSet initialization information storage.
	var maps = make(map[string]map[string]interface{})
	var nestedLists = make(map[string]map[string]interface{})

	// Incoming map iteration.
	for mikrotikKebabName, mikrotikValue := range item {
		// Set the ID.
		//if mikrotikKebabName == ".id" {
		//	if err = d.Set(KeyId, mikrotikValue); err != nil {
		//		diags = append(diags, diag.FromErr(err)...)
		//	}
		//	continue
		//}

		// Skip all service fields (i.e. `.id`, `.nextid`, `ret` ...).
		if mikrotikKebabName[0:1] == "." || mikrotikKebabName == "ret" {
			continue
		}

		if _, ok := dropByValue[mikrotikValue]; ok {
			continue
		}

		// Mikrotik fields transformation: "channel" ---> "channel.config".
		// For further use in the map.
		if transformSet != nil {
			if new, ok := transformSet[mikrotikKebabName]; ok {
				mikrotikKebabName = new
			}
		}

		// field-name => field_name
		terraformSnakeName := KebabToSnake(mikrotikKebabName)

		if skipFields != nil {
			if _, ok := skipFields[terraformSnakeName]; ok {
				continue
			}
		}

		// Composite fields.
		var subFieldSnakeName string
		if strings.Contains(terraformSnakeName, ".") {
			f := strings.SplitN(terraformSnakeName, ".", 2)
			terraformSnakeName, subFieldSnakeName = f[0], f[1]
		}

		if _, ok := s[terraformSnakeName]; !ok {
			// For development.
			// panic("[MikrotikResourceDataToTerraform] The field was lost during the Schema development: " + terraformSnakeName)
			diags = append(diags, diag.Diagnostic{
				// TODO Waiting for TestStep.ExpectWarning https://github.com/hashicorp/terraform-plugin-testing/pull/17
				// The test response to Warnings has not yet been implemented.
				Severity: diag.Warning,
				Summary:  "Field '" + terraformSnakeName + "' not found in the schema",
				Detail: fmt.Sprintf("[MikrotikResourceDataToTerraform] The field was lost during the Schema development: ▷ '%s': '%s' ◁",
					terraformSnakeName, mikrotikValue),
			})
			// Catch all fields.
			continue
		}

		switch s[terraformSnakeName].Type {
		case schema.TypeString:
			err = d.Set(terraformSnakeName, mikrotikValue)

		case schema.TypeFloat:
			f, e := strconv.ParseFloat(mikrotikValue, 64)
			if e != nil {
				diags = diag.Errorf("%v for '%v' field", e, terraformSnakeName)
				break
			}
			err = d.Set(terraformSnakeName, f)

		case schema.TypeInt:
			i, e := strconv.Atoi(mikrotikValue)
			if e != nil {
				diags = diag.Errorf("%v for '%v' field", e, terraformSnakeName)
				break
			}
			err = d.Set(terraformSnakeName, i)

		case schema.TypeBool:
			if _, ok := setUnsetFields[terraformSnakeName]; ok {
				err = d.Set(terraformSnakeName, true)
				break
			}
			err = d.Set(terraformSnakeName, BoolFromMikrotikJSON(mikrotikValue))

		case schema.TypeList, schema.TypeSet:
			var l []interface{}

			// Don't fill in empty strings (preventing a non-empty plan).
			// |   # routeros_interface_wireguard_peer.wg_peer will be updated in-place
			// |   ~ resource "routeros_interface_wireguard_peer" "wg_peer" {
			// |       ~ allowed_address       = [
			// |           - "",
			// |         ]
			// |         id                    = "*2"
			// |         # (7 unchanged attributes hidden)
			// |     }

			// Flat Lists & Sets:
			if _, ok := s[terraformSnakeName].Elem.(*schema.Schema); mikrotikValue != "" && ok {
				for _, v := range strings.Split(mikrotikValue, ",") {
					switch s[terraformSnakeName].Elem.(*schema.Schema).Type {
					case schema.TypeFloat:
						f, err := strconv.ParseFloat(v, 64)
						if err != nil {
							diags = diag.Errorf("%v for '%v' field", err, terraformSnakeName)
							continue
						}
						l = append(l, f)

					case schema.TypeInt:
						i, err := strconv.Atoi(v)
						if err != nil {
							diags = diag.Errorf("%v for '%v' field", err, terraformSnakeName)
							continue
						}
						l = append(l, i)

					default:
						l = append(l, v)
					}
				}

				if err != nil {
					break // case
				}
			}

			if s[terraformSnakeName].Type == schema.TypeList {
				switch s[terraformSnakeName].Elem.(type) {
				case *schema.Schema:
					err = d.Set(terraformSnakeName, l)
				case *schema.Resource:
					var v any

					if _, ok := s[terraformSnakeName].Elem.(*schema.Resource).Schema[subFieldSnakeName]; !ok {
						diags = append(diags, diag.Diagnostic{
							Severity: diag.Warning,
							Summary:  "Field '" + terraformSnakeName + "." + subFieldSnakeName + "' not found in the schema",
							Detail: fmt.Sprintf("[MikrotikResourceDataToTerraform] the datasource Schema sub-field was lost during development: ▷ '%s.%s' ◁",
								terraformSnakeName, subFieldSnakeName),
						})
						continue
					}

					switch s[terraformSnakeName].Elem.(*schema.Resource).Schema[subFieldSnakeName].Type {
					case schema.TypeString:
						v = mikrotikValue
					case schema.TypeFloat:
						v, err = strconv.ParseFloat(mikrotikValue, 64)
						if err != nil {
							diags = diag.Errorf("%v for '%v.%v' field", err, terraformSnakeName, subFieldSnakeName)
						}
					case schema.TypeInt:
						v, err = strconv.Atoi(mikrotikValue)
						if err != nil {
							diags = diag.Errorf("%v for '%v.%v' field", err, terraformSnakeName, subFieldSnakeName)
						}
					case schema.TypeBool:
						v = BoolFromMikrotikJSON(mikrotikValue)
					}

					if err != nil {
						break
					}

					if list, ok := nestedLists[terraformSnakeName]; !ok {
						nestedLists[terraformSnakeName] = map[string]interface{}{subFieldSnakeName: v}
					} else {
						list[subFieldSnakeName] = v
					}
				}
			} else {
				err = d.Set(terraformSnakeName,
					schema.NewSet(schema.HashSchema(s[terraformSnakeName].Elem.(*schema.Schema)), l))
			}

		case schema.TypeMap:
			// "yes" -> "true"; "no" -> "false"
			mikrotikValue = BoolFromMikrotikJSONStr(mikrotikValue)

			if m, ok := maps[terraformSnakeName]; !ok {
				// Create a new map when processing the first incoming field.
				maps[terraformSnakeName] = map[string]interface{}{subFieldSnakeName: mikrotikValue}
			} else {
				m[subFieldSnakeName] = mikrotikValue
			}

		default:
			// For development.
			//panic(fmt.Sprintf("[MikrotikResourceDataToTerraform] resource type not implemented: %v for '%v'",
			//	s[terraformSnakeName].Type.String(), mikrotikValue))
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Warning,
				Summary:  "Can't fill the schema field",
				Detail: fmt.Sprintf("Resource type not implemented: '%v' for '%v'",
					s[terraformSnakeName].Type.String(), terraformSnakeName),
			})
		}

		if err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}

	// Lists processing.
	for name, list := range nestedLists {
		if err = d.Set(name, []interface{}{list}); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	// Maps processing.
	for name, m := range maps {
		if err = d.Set(name, m); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}

	return diags
}

func MikrotikResourceDataToTerraformDatasource(items *[]MikrotikItem, resourceDataKeyName string, s map[string]*schema.Schema, d *schema.ResourceData) diag.Diagnostics {
	var diags diag.Diagnostics
	var dsItems []map[string]interface{}
	// System resource have an empty 'resourceDataKeyName'.
	var isSystemDatasource bool = (resourceDataKeyName == "")
	var skipFields map[string]struct{}

	if sf, ok := s[MetaSkipFields]; ok {
		skipFields = loadSkipFields(sf.Default.(string))
	}

	// Checking the schema.
	if !isSystemDatasource {
		sv, ok := s[resourceDataKeyName]
		if !ok {
			// For development.
			//panic("[MikrotikResourceDataToTerraformDatasource] the datasource Schema field was lost during development: " + resourceDataKeyName)
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Warning,
				Summary:  "Field '" + resourceDataKeyName + "' not found in the schema",
				Detail: fmt.Sprintf("[MikrotikResourceDataToTerraformDatasource] the datasource Schema field was lost during development: ▷ '%s' ◁",
					resourceDataKeyName),
			})
			// Or panic.
			return diags
		}
		s = sv.Elem.(*schema.Resource).Schema
	}

	if isSystemDatasource && len(*items) != 1 {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "System resources should not return an array of values",
			Detail: fmt.Sprintf("[MikrotikResourceDataToTerraformDatasource] system resource '%s' polling returned %d values",
				s[MetaResourcePath].Default.(string), len(*items)),
		})
		return diags
	}

	// Array of Mikrotik items iteration.
	for _, item := range *items {

		dsItem := map[string]interface{}{}

		// Incoming map iteration.
		for mikrotikKebabName, mikrotikValue := range item {
			// MT can return the field name in uppercase format.
			mikrotikKebabName = strings.ToLower(mikrotikKebabName)
			// In this case the ID must be a string.
			if mikrotikKebabName == ".id" {
				dsItem["id"] = mikrotikValue
				continue
			}

			// Skip all service fields.
			if mikrotikKebabName[0:1] == "." {
				continue
			}

			// field-name => field_name
			terraformSnakeName := KebabToSnake(mikrotikKebabName)

			if skipFields != nil {
				if _, ok := skipFields[terraformSnakeName]; ok {
					continue
				}
			}

			if _, ok := s[terraformSnakeName]; !ok {
				// For development.
				//panic("[MikrotikResourceDataToTerraformDatasource] the field was lost during development.: " + terraformSnakeName)
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Warning,
					Summary:  "Field '" + terraformSnakeName + "' not found in the schema",
					Detail: fmt.Sprintf("[MikrotikResourceDataToTerraformDatasource] the field was lost during the Schema development: ▷ '%s': '%s' ◁",
						terraformSnakeName, mikrotikValue),
				})
				// Catch all fields.
				continue
			}

			var propValue interface{}

			switch s[terraformSnakeName].Type {
			case schema.TypeString:
				propValue = mikrotikValue

			case schema.TypeFloat:
				f, err := strconv.ParseFloat(mikrotikValue, 64)
				if err != nil {
					diags = append(diags, diag.Errorf("%v for '%v' field", err, terraformSnakeName)...)
					continue
				}
				propValue = f

			case schema.TypeInt:
				i, err := strconv.Atoi(mikrotikValue)
				if err != nil {
					diags = append(diags, diag.Errorf("%v for '%v' field", err, terraformSnakeName)...)
					continue
				}
				propValue = i

			case schema.TypeBool:
				// TODO Add support for set/unset fields?
				propValue = BoolFromMikrotikJSON(mikrotikValue)

			case schema.TypeList:
				var l []interface{}
				if mikrotikValue != "" {
					for _, s := range strings.Split(mikrotikValue, ",") {
						l = append(l, s)
					}
				}
				propValue = l

			// TODO Add processing of missing types: List(int), Set, Map
			case schema.TypeSet:
				var l []interface{}
				if mikrotikValue != "" {
					for _, s := range strings.Split(mikrotikValue, ",") {
						l = append(l, s)
					}
				}
				// String sets only (schema.HashString)!
				propValue = schema.NewSet(schema.HashString, l)

			default:
				// For development.
				//panic(fmt.Sprintf("Resource type not implemented: %v for '%v'",
				//	s[terraformSnakeName].Type.String(), mikrotikValue))
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Warning,
					Summary:  "Can't fill the schema field",
					Detail: fmt.Sprintf("Resource type not implemented: %v for '%v'",
						s[terraformSnakeName].Type.String(), mikrotikValue),
				})

			}

			dsItem[terraformSnakeName] = propValue
		}

		dsItems = append(dsItems, dsItem)
	}

	d.SetId(UniqueId())
	if !isSystemDatasource {
		if err := d.Set(resourceDataKeyName, dsItems); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	} else {
		for k, v := range dsItems[0] {
			if err := d.Set(k, v); err != nil {
				diags = append(diags, diag.FromErr(err)...)
			}
		}
	}

	return diags
}

// Copied from terraform-plugin-testing@v1.2.0/helper/resource/id.go
// Because this functionality is marked deprecated.
const UniqueIdPrefix = `terraform-`

// idCounter is a monotonic counter for generating ordered unique ids.
var idMutex sync.Mutex
var idCounter uint32

func UniqueId() string {
	return PrefixedUniqueId(UniqueIdPrefix)
}

func PrefixedUniqueId(prefix string) string {
	// Be precise to 4 digits of fractional seconds, but remove the dot before the
	// fractional seconds.
	timestamp := strings.Replace(
		time.Now().UTC().Format("20060102150405.0000"), ".", "", 1)

	idMutex.Lock()
	defer idMutex.Unlock()
	idCounter++
	return fmt.Sprintf("%s%s%08x", prefix, timestamp, idCounter)
}
