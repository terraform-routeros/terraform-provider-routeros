package routeros

type TransportType int

// Using numbering from 1 to control type values.
const (
	TransportAPI TransportType = 1 + iota
	TransportREST
)

type IdType int

const (
	Id IdType = 1 + iota
	Name
	// ... and maybe ssh?!
)

type ItemId struct {
	Type  IdType
	Value string
}

func (t IdType) String() string {
	switch t {
	case Id:
		return ".id"
	case Name:
		return "name"
	}
	return "error: undefined id type"
}

// MikrotikItemMetadata This information must travel from the schema to the resource polling function.
type MikrotikItemMetadata struct {
	IdType IdType            // The field contains ID.
	Path   string            // Resource URL.
	Meta   map[string]string // Additional metadata that may be present in the schema.
}

// MikrotikItem Contains only data.
type MikrotikItem map[string]string

func (m MikrotikItem) GetID(t IdType) string {
	switch t {
	case Id:
		// REST
		if id, ok := m[".id"]; ok {
			return id
		}
		// API
		if id, ok := m["ret"]; ok {
			return id
		}
	case Name:
		if id, ok := m["name"]; ok {
			return id
		}
	default:
		panic("[MikrotikItem.GetID] wrong IdType")
	}
	return ""
}

func (m *MikrotikItem) replace(swap any) {
	switch t := swap.(type) {
	case *MikrotikItem:
		*m = *t
	default:
		panic("not the same type")
	}
}

// KebabToSnake Convert Mikrotik JSON names to TF schema names: some-filed to some_field.
func KebabToSnake(name string) string {
	res := []byte(name)
	for i := range res {
		if res[i] == '-' {
			res[i] = '_'
		}
	}
	return string(res)
}

// SnakeToKebab Convert IF schema names to Mikrotik JSON names: some_filed to some-field.
func SnakeToKebab(name string) string {
	res := []byte(name)
	for i := range res {
		if res[i] == '_' {
			res[i] = '-'
		}
	}
	return string(res)
}

func BoolToMikrotikJSON(b bool) string {
	if b {
		return "yes"
	}
	return "no"
}

func BoolFromMikrotikJSON(s string) bool {
	if s == "true" || s == "yes" {
		return true
	}
	return false
}

// Map helpers.

func BoolToMikrotikJSONStr(s string) string {
	if s == "true" {
		return "yes"
	}
	if s == "false" {
		return "no"
	}
	return s
}

func BoolFromMikrotikJSONStr(s string) string {
	if s == "yes" {
		return "true"
	}
	if s == "no" {
		return "false"
	}
	return s
}
