package routeros

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var reMetadataFields = regexp.MustCompile(`^___\S+___$`)
var reTransformSet = regexp.MustCompile(`"\s*?(\S+?)\s*?"\s*?:\s*?"\s*?(\S+?)\s*?"`)
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

func isEmpty(propName string, schemaProp *schema.Schema, d *schema.ResourceData) bool {
	v := d.Get(propName)
	switch schemaProp.Type {
	case schema.TypeString:
		if schemaProp.Default != nil {
			return v.(string) == "" && schemaProp.Default.(string) == ""
		}
		return v.(string) == ""
	case schema.TypeInt:
		return !d.HasChange(propName)
	case schema.TypeBool:
		if schemaProp.Default != nil {
			return schemaProp.Default.(bool) == v.(bool)
		}
		return !v.(bool) // false == isEmpty
	case schema.TypeList:
		return len(v.([]interface{})) == 0
	case schema.TypeSet:
		return v.(*schema.Set).Len() == 0
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
// s: `"channel":"channel.config","datapath":"datapath.config"` in the Mikrotik (kebab) notation!
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
	var transformSet map[string]string
	var skipFields map[string]struct{}

	// {"channel.config": "channel", "schema-field-name": "mikrotik-field-name"}
	if ts, ok := s[MetaTransformSet]; ok {
		transformSet = loadTransformSet(ts.Default.(string), true)
	}

	// "field_first", "field_second", "field_third"
	if sf, ok := s[MetaSkipFields]; ok {
		skipFields = loadSkipFields(sf.Default.(string))
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
			case MetaTransformSet, MetaSkipFields:
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
		if terraformMetadata.Optional && !d.HasChange(terraformSnakeName) &&
			isEmpty(terraformSnakeName, s[terraformSnakeName], d) {
			continue
		}

		// terraformSnakeName = fast_forward, schemaPropData = true
		// NewMikrotikItem.Fields["fast-forward"] = "true"
		mikrotikKebabName := SnakeToKebab(terraformSnakeName)
		value := d.Get(terraformSnakeName)

		switch terraformMetadata.Type {
		case schema.TypeString:
			item[mikrotikKebabName] = value.(string)
		case schema.TypeInt:
			item[mikrotikKebabName] = strconv.Itoa(value.(int))
		case schema.TypeBool:
			item[mikrotikKebabName] = BoolToMikrotikJSON(value.(bool))
		// Used to represent an ordered collection of items.
		case schema.TypeList:
			item[mikrotikKebabName] = ListToString(value)
		// Used to represent an unordered collection of items.
		case schema.TypeSet:
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

	// {"channel": "channel.config", "mikrotik-field-name": "schema-field-name"}
	if ts, ok := s[MetaTransformSet]; ok {
		transformSet = loadTransformSet(ts.Default.(string), false)
	}

	// TypeMaps initialization information storage.
	// map["channel"] = bool
	var maps = make(map[string]bool)

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

		// Mikrotik fields transformation: "channel" ---> "channel.config".
		// For further use in the map.
		if transformSet != nil {
			if new, ok := transformSet[mikrotikKebabName]; ok {
				mikrotikKebabName = new
			}
		}

		// field-name => field_name
		terraformSnakeName := KebabToSnake(mikrotikKebabName)

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

		case schema.TypeInt:
			i, e := strconv.Atoi(mikrotikValue)
			if e != nil {
				diags = diag.Errorf("%v for '%v' field", err, terraformSnakeName)
				break
			}
			err = d.Set(terraformSnakeName, i)

		case schema.TypeBool:
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
			if mikrotikValue != "" {
				for _, v := range strings.Split(mikrotikValue, ",") {
					if s[terraformSnakeName].Elem.(*schema.Schema).Type == schema.TypeInt {
						i, err := strconv.Atoi(v)
						if err != nil {
							diags = diag.Errorf("%v for '%v' field", err, terraformSnakeName)
							continue
						}

						l = append(l, i)
					} else {
						l = append(l, v)
					}
				}

				if err != nil {
					break // case
				}
			}

			if s[terraformSnakeName].Type == schema.TypeList {
				err = d.Set(terraformSnakeName, l)
			} else {
				err = d.Set(terraformSnakeName,
					schema.NewSet(schema.HashSchema(s[terraformSnakeName].Elem.(*schema.Schema)), l))
			}

		case schema.TypeMap:
			// "yes" -> "true"; "no" -> "false"
			mikrotikValue = BoolFromMikrotikJSONStr(mikrotikValue)

			if _, ok := maps[terraformSnakeName]; !ok {
				// Create a new map when processing the first incoming field.
				maps[terraformSnakeName] = true
				d.Set(terraformSnakeName, map[string]interface{}{subFieldSnakeName: mikrotikValue})
			} else {
				m := d.Get(terraformSnakeName).(map[string]interface{})
				m[subFieldSnakeName] = mikrotikValue
				d.Set(terraformSnakeName, m)
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

	return diags
}

func MikrotikResourceDataToTerraformDatasource(items *[]MikrotikItem, resourceDataKeyName string, s map[string]*schema.Schema, d *schema.ResourceData) diag.Diagnostics {
	var diags diag.Diagnostics
	var dsItems []map[string]interface{}

	// Checking the schema.
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

	// Array of Mikrotik items iteration.
	for _, item := range *items {

		dsItem := map[string]interface{}{}

		// Incoming map iteration.
		for mikrotikKebabName, mikrotikValue := range item {
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

			case schema.TypeInt:
				i, err := strconv.Atoi(mikrotikValue)
				if err != nil {
					diags = append(diags, diag.Errorf("%v for '%v' field", err, terraformSnakeName)...)
					continue
				}
				propValue = i

			case schema.TypeBool:
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
	if err := d.Set(resourceDataKeyName, dsItems); err != nil {
		diags = append(diags, diag.FromErr(err)...)
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
