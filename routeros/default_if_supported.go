package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var setSupportedMode = false
var setSupportedValue bool

func SetSupported(schemaDefaultFunc schema.SchemaDefaultFunc, value bool) {
	setSupportedMode = true
	setSupportedValue = value
	schemaDefaultFunc()
	setSupportedMode = false
}

func DefaultIfSupported(value interface{}) schema.SchemaDefaultFunc {
	var supported = false

	return func() (interface{}, error) {
		if setSupportedMode {
			supported = setSupportedValue
		}

		if (!supported) {
			return nil, nil
		}

		return value, nil
	}
}
