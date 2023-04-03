package routeros

import (
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var (
	testResource = schema.Resource{
		Schema: map[string]*schema.Schema{
			MetaResourcePath: PropResourcePath("/test/resource"),
			MetaId:           PropId(Id),
			"string": {
				Type: schema.TypeString,
			},
			"int": {
				Type: schema.TypeInt,
			},
			"bool": {
				Type: schema.TypeBool,
			},
			"computed": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}

	testDatasource = schema.Resource{
		Schema: map[string]*schema.Schema{
			"test_name": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						MetaResourcePath: PropResourcePath("/test/resource"),
						MetaId:           PropId(Id),
						"string": {
							Type: schema.TypeString,
						},
						"int": {
							Type: schema.TypeInt,
						},
						"bool": {
							Type: schema.TypeBool,
						},
					},
				},
			},
		},
	}
)

func Test_mikrotikResourceDataToTerraform(t *testing.T) {

	testItem := MikrotikItem{".id": "*39", "string": "string12345", "int": "10", "bool": "true"}

	testResourceData := testResource.TestResourceData()
	expectedRes := map[string]interface{}{"string": "string12345", "int": 10, "bool": true}

	err := MikrotikResourceDataToTerraform(testItem, testResource.Schema, testResourceData)
	if err != nil {
		t.Errorf("decoding err: %v", err)
	}

	for key, expected := range expectedRes {
		actual := testResourceData.Get(key)

		if !reflect.DeepEqual(actual, expected) {
			t.Fatalf("bad: expected:%#v\nactual:%#v", expected, actual)
		}
	}

}

func Test_terraformResourceDataToMikrotik(t *testing.T) {

	expected := MikrotikItem{"string": "string12345", "int": "10", "bool": "yes"}

	testResourceData := testResource.TestResourceData()
	testResourceData.SetId("*39")
	testResourceData.Set("string", "string12345")
	testResourceData.Set("int", 10)
	testResourceData.Set("bool", true)

	actual, _ := TerraformResourceDataToMikrotik(testResource.Schema, testResourceData)

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("bad: expected:%#v\nactual:%#v", expected, actual)
	}
}

func Test_mikrotikResourceDataToTerraformDatasource(t *testing.T) {
	testItems := []MikrotikItem{
		{"string": "string12345", "int": "10", "bool": "yes"},
		{"string": "12345string", "int": "20", "bool": "no"},
	}

	testResourceData := testDatasource.TestResourceData()
	expectedRes := []map[string]interface{}{
		{MetaResourcePath: "", MetaId: 0, "string": "string12345", "int": 10, "bool": true},
		{MetaResourcePath: "", MetaId: 0, "string": "12345string", "int": 20, "bool": false},
	}

	err := MikrotikResourceDataToTerraformDatasource(&testItems, "test_name", testDatasource.Schema, testResourceData)
	if err != nil {
		t.Errorf("decoding err: %v", err)
	}

	for i, rec := range testResourceData.Get("test_name").([]interface{}) {
		for key, actual := range rec.(map[string]interface{}) {
			if !reflect.DeepEqual(actual, expectedRes[i][key]) {
				t.Fatalf("bad: (key: %v) expected:%#v\tactual:%#v", key, expectedRes[i][key], actual)
			}
		}
	}
}

func Test_loadTransformSet(t *testing.T) {
	testData := []struct {
		s       string
		reverse bool
	}{
		{`"channel":"channel.config","datapath":"datapath.config"`, false},
		{`"mikrotik-field-name":"schema-field-name"`, false},
		{`"channel":"channel.config","datapath":"datapath.config"`, true},
		{`"mikrotik-field-name":"schema-field-name"`, true},
	}

	expected := []map[string]string{
		{"channel": "channel.config", "datapath": "datapath.config"},
		{"mikrotik-field-name": "schema-field-name"},
		{"channel.config": "channel", "datapath.config": "datapath"},
		{"schema-field-name": "mikrotik-field-name"},
	}

	for i, actual := range testData {
		if !reflect.DeepEqual(loadTransformSet(actual.s, actual.reverse), expected[i]) {
			t.Fatalf("bad: (item: %v) expected:%#v\tactual:%#v", i, expected[i], loadTransformSet(actual.s, actual.reverse))
		}
	}
}
