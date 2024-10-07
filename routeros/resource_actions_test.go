package routeros

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"reflect"
	"strconv"
	"testing"
)

func Test_defaultResourceCreateContext(t *testing.T) {
	//a:= DefaultCreate[client.GRE]()
	rs := ResourceInterfaceGre()
	rd := rs.TestResourceData()

	m := map[string]string{}

	for k, v := range rs.Schema {
		if v.Computed {
			continue
		}

		// TODO Check this!
		if !rd.HasChange(k) && v.Optional {
			continue
		}

		switch reflect.TypeOf(rd.Get(k)).Kind() {
		case reflect.String:
			m[SnakeToKebab(k)] = rd.Get(k).(string)
		case reflect.Int:
			m[SnakeToKebab(k)] = strconv.Itoa(rd.Get(k).(int))
		case reflect.Bool:
			m[SnakeToKebab(k)] = BoolToMikrotikJSON(rd.Get(k).(bool))
		}
	}

	b, _ := json.Marshal(m)
	log.Println(string(b))

	s := `{".id":"*39","allow-fast-path":"true","clamp-tcp-mss":"true","comment":"New comment","disabled":"true",` +
		`"dont-fragment":"","dscp":"","ipsec-secret":"12321","keepalive":"","local-address":"","mtu":"1472",` +
		`"name":"gre","remote-address":""}`
	m2 := map[string]string{}

	err := json.Unmarshal([]byte(s), &m2)
	if err != nil {
		t.Errorf(err.Error())
	}

	for k, v := range m2 {
		if KebabToSnake(k) == ".id" {
			rd.SetId(v)
			continue
		}

		if _, ok := rs.Schema[KebabToSnake(k)]; !ok {
			t.Errorf("Lost field: %v", KebabToSnake(k))
		}
		log.Println(KebabToSnake(k))

		switch rs.Schema[KebabToSnake(k)].Type {
		case schema.TypeString:
			_ = rd.Set(KebabToSnake(k), v)
		case schema.TypeInt:
			i, _ := strconv.Atoi(v)
			_ = rd.Set(KebabToSnake(k), i)
		case schema.TypeBool:
			_ = rd.Set(KebabToSnake(k), BoolFromMikrotikJSON(v))
		default:
			t.Error(rs.Schema[KebabToSnake(k)].Type)
		}
	}
}
