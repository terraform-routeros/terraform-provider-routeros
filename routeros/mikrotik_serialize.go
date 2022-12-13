package routeros

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var reMetadataFields = regexp.MustCompile(`^___\S+___$`)

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
				meta.Meta[terraformSnakeName] = terraformMetadata.Default.(string)
			}
		}
	}
	return meta
}

func isEmpty(propName string, s map[string]*schema.Schema, d *schema.ResourceData) bool {
	v := d.Get(propName)
	switch reflect.TypeOf(v).Kind() {
	case reflect.String:
		if s[propName].Default != nil {
			return v.(string) == "" && s[propName].Default.(string) == ""
		}
		return v.(string) == ""
	case reflect.Int:
		return !d.HasChange(propName)
	case reflect.Bool:
		if s[propName].Default != nil {
			return s[propName].Default.(bool) == v.(bool)
		}
		return v.(bool)
	case reflect.Slice:
		return len(v.([]interface{})) == 0
	default:
		panic("[isEmpty] wrong resource type: " + reflect.TypeOf(v).Kind().String())
	}
}

// TerraformResourceDataToMikrotik Marshal Mikrotik resource from TF resource schema.
func TerraformResourceDataToMikrotik(s map[string]*schema.Schema, d *schema.ResourceData) (MikrotikItem, *MikrotikItemMetadata) {
	item := MikrotikItem{}
	meta := &MikrotikItemMetadata{}

	// Schema map iteration.
	for terraformSnakeName, terraformMetadata := range s {
		// Fill in the metadata fields.
		if reMetadataFields.MatchString(terraformSnakeName) {
			switch terraformSnakeName {
			case MetaId:
				meta.IdType = IdType(terraformMetadata.Default.(int))
			case MetaResourcePath:
				meta.Path = terraformMetadata.Default.(string)
			default:
				meta.Meta[terraformSnakeName] = terraformMetadata.Default.(string)
			}
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
		if terraformMetadata.Optional && !d.HasChange(terraformSnakeName) && isEmpty(terraformSnakeName, s, d) {
			continue
		}

		// terraformSnakeName = fast_forward, schemaPropData = true
		// NewMikrotikItem.Fields["fast-forward"] = "true"
		mikrotikKebabName := SnakeToKebab(terraformSnakeName)
		value := d.Get(terraformSnakeName)

		switch reflect.TypeOf(value).Kind() {
		case reflect.String:
			item[mikrotikKebabName] = value.(string)
		case reflect.Int:
			item[mikrotikKebabName] = strconv.Itoa(value.(int))
		case reflect.Bool:
			item[mikrotikKebabName] = BoolToMikrotikJSON(value.(bool))
		case reflect.Slice:
			var ss []string
			for _, iface := range value.([]interface{}) {
				ss = append(ss, fmt.Sprint(iface))
			}
			item[mikrotikKebabName] = strings.Join(ss, ",")
		default:
			panic(fmt.Sprintf("[TerraformResourceDataToMikrotik] resource type not implemented: %v for '%v'",
				reflect.TypeOf(value).Kind().String(), terraformSnakeName))
		}

	}

	return item, meta
}

// MikrotikResourceDataToTerraform Unmarshal Mikrotik resource (incoming data: JSON, etc.) to TF resource schema.
func MikrotikResourceDataToTerraform(item MikrotikItem, s map[string]*schema.Schema, d *schema.ResourceData) diag.Diagnostics {
	var diags diag.Diagnostics
	var err error

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

		// field-name => field_name
		terraformSnakeName := KebabToSnake(mikrotikKebabName)
		if _, ok := s[terraformSnakeName]; !ok {
			// For development.
			// panic("[MikrotikResourceDataToTerraform] The field was lost during the Schema development: " + terraformSnakeName)
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Warning,
				Summary:  "Field not found",
				Detail: "[MikrotikResourceDataToTerraform] The field was lost during the Schema development: " +
					terraformSnakeName + " <<<= " + mikrotikValue,
			})
			// Or panic.
			return diags
		}

		switch s[terraformSnakeName].Type {
		case schema.TypeString:
			// Don't fill in empty strings.
			if s[terraformSnakeName].Optional && mikrotikValue == "" {
				continue
			}
			err = d.Set(terraformSnakeName, mikrotikValue)

		case schema.TypeInt:
			i, err := strconv.Atoi(mikrotikValue)
			if err != nil {
				diags = diag.Errorf("%v for '%v' field", err, terraformSnakeName)
				break
			}
			err = d.Set(terraformSnakeName, i)

		case schema.TypeBool:
			err = d.Set(terraformSnakeName, BoolFromMikrotikJSON(mikrotikValue))

		case schema.TypeList:
			// Don't fill in empty strings.
			if s[terraformSnakeName].Optional && mikrotikValue == "" {
				continue
			}

			var l []interface{}
			for _, s := range strings.Split(mikrotikValue, ",") {
				l = append(l, s)
			}
			err = d.Set(terraformSnakeName, l)

		default:
			// For development.
			//panic(fmt.Sprintf("[MikrotikResourceDataToTerraform] resource type not implemented: %v for '%v'",
			//	s[terraformSnakeName].Type.String(), mikrotikValue))
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Warning,
				Summary:  "Can't fill the schema field",
				Detail: fmt.Sprintf("Resource type not implemented: %v for '%v'",
					s[terraformSnakeName].Type.String(), mikrotikValue),
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
			Summary:  "Field not found",
			Detail: "[MikrotikResourceDataToTerraformDatasource] the datasource Schema field was lost during " +
				"development: " + resourceDataKeyName,
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
					Summary:  "Field not found",
					Detail: "[MikrotikResourceDataToTerraformDatasource] the field was lost during the Schema development: " +
						terraformSnakeName + " <<<= " + mikrotikValue,
				})
				// Or panic.
				return diags
			}

			var propValue interface{}

			switch s[terraformSnakeName].Type {
			case schema.TypeString:
				// Don't fill in empty strings.
				if s[terraformSnakeName].Optional && mikrotikValue == "" {
					continue
				}
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
				// Don't fill in empty strings.
				if s[terraformSnakeName].Optional && mikrotikValue == "" {
					continue
				}

				var l []interface{}
				for _, s := range strings.Split(mikrotikValue, ",") {
					l = append(l, s)
				}
				propValue = l

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

	d.SetId(resource.UniqueId())
	if err := d.Set(resourceDataKeyName, dsItems); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	return diags
}
