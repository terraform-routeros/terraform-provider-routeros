package main

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/terraform-routeros/terraform-provider-routeros/routeros"
)

var (
	log, _ = NewLogger(context.Background())

	reUserResources = regexp.MustCompile(`(?m)^(/.*?)\sadd\s(.*?)[\r\n]+`)
	// (/interface ethernet) set ([ find default-name=ether1 ] disable-running-check=no)
	// (/ip service) set (api-ssl) (certificate=ssl)
	reSystemResources = regexp.MustCompile(`(?m)^(/.*?)\sset\s(?:([\w-]+)\s)?(.*?)[\r\n]+`)
	reAttributes      = regexp.MustCompile(`\S+=\S+`)
	// .id=*1;available=101;name=dhcp;ranges=192.168.88.100-192.168.88.200;total=101;used=0
	reId = regexp.MustCompile(`\.id=(\S+?);`)

	resourceTempate = `resource "%v" "%v" {
  %v
}
  
`
	providerTemplate = `terraform {
  required_providers {
    routeros = {
      source  = "terraform-routeros/routeros"
      version = "~> 1"
    }
  }
}

provider "routeros" {
  hosturl  = "https://%v"
  username = "%v"
  # password = "" # env ROS_PASSWORD or MIKROTIK_PASSWORD
  insecure = true
}

`
)

func main() {

	resHcl := bytes.NewBufferString(fmt.Sprintf(providerTemplate, "192.168.180.171", "admin"))
	resImport := bytes.NewBuffer(nil)

	provider := routeros.NewProvider()

	var providerResources = make(map[string][]string)
	for k, v := range provider.ResourcesMap {
		path := v.Schema[routeros.MetaResourcePath].Default.(string)
		// Don't add aliases
		if _, ok := providerAliasesMap[k]; ok {
			continue
		}
		providerResources[path] = append(providerResources[path], k)
	}

	conn, err := NewSsh("192.168.180.171:22", "admin", "1")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// /interface ethernet set [ find default-name=ether2 ] disable-running-check=no
	// /interface wireguard add listen-port=1829 mtu=1420 name=wg1
	// /interface list add name=list
	// ...
	config, err := GetMikrotikConfig(conn)
	if err != nil {
		log.Fatal(err)
	}

	if len(config) == 0 {
		log.Fatal("Config is empty")
	}

	// ip-pool: N, interface-wireguard: N, ...
	var hclNames = make(map[string]int)

	// /interface wireguard >>> add <<< listen-port=1829 mtu=1420 name=wg1
	for _, ss := range reUserResources.FindAllStringSubmatch(config, -1) {
		// "/interface wireguard", "listen-port=1829 mtu=1420 name=wg1"
		path, attributes := ss[1], ss[2]

		// routeros_interface_wireguard
		hclSection, err := GetResourceSection(hclNames, providerResources, path)
		if err != nil {
			log.Error(err)
			continue
		}

		// HCL file
		hclAttributes, required := GetAttributes(provider, hclSection.ResourceName, attributes)
		fmt.Fprintf(resHcl, resourceTempate, hclSection.ResourceName, hclSection.HCLName, strings.Join(hclAttributes, "\n  "))

		// Import script
		id := GetResourceId(conn, path, required)
		fmt.Fprintf(resImport, "terraform import %v.%v %v\n", hclSection.ResourceName, hclSection.HCLName, id)
	}

	// /interface ethernet >>>set<<< [ find default-name=ether1 ] disable-running-check=no
	// /ip service set api-ssl certificate=ssl
	for _, ss := range reSystemResources.FindAllStringSubmatch(config, -1) {
		// "/interface ethernet",  "", "[ find default-name=ether1 ] disable-running-check=no"
		// "/ip service", "api-ssl", "certificate=ssl"
		path, name, attributes := ss[1], ss[2], ss[3]
		if name == "" {
			name = "."
		}

		// routeros_ip_service
		hclSection, err := GetResourceSection(hclNames, providerResources, path)
		if err != nil {
			log.Print(err)
			continue
		}

		switch hclSection.ResourceName {
		case "routeros_ip_service":
			// Add the Required attribute
			attributes += " numbers=" + name
		}

		hclAttributes, required := GetAttributes(provider, hclSection.ResourceName, attributes)

		switch hclSection.ResourceName {
		case "routeros_interface_ethernet":
			// Get Id
			name = GetResourceId(conn, path, required)
		}

		// HCL file
		fmt.Fprintf(resHcl, resourceTempate, hclSection.ResourceName, hclSection.HCLName, strings.Join(hclAttributes, "\n  "))
		// Import script
		fmt.Fprintf(resImport, "terraform import %v.%v %v\n", hclSection.ResourceName, hclSection.HCLName, name)
	}

	baseName := fmt.Sprintf("autoimport-%v", time.Now().Format("20060102-1504"))

	// HCL file
	tfFile, err := os.OpenFile(baseName+".tf", os.O_CREATE|os.O_TRUNC|os.O_RDWR, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer tfFile.Close()

	_, err = resHcl.WriteTo(tfFile)
	if err != nil {
		log.Fatal(err)
	}

	// Import script
	importFile, err := os.OpenFile(baseName+".sh", os.O_CREATE|os.O_TRUNC|os.O_RDWR, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer importFile.Close()

	_, err = resImport.WriteTo(importFile)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Complete")
}

// resource "routeros_bgp_connection" "test" {
// }
// HCLResource "ResourceName" "HCLName" {
// }
type HCLResource struct {
	ResourceName string
	HCLName      string
}

// The function returns the name of the resource in provider notation and the unique identifier of the resource.
func GetResourceSection(hclNames map[string]int, providerResources map[string][]string, path string) (*HCLResource, error) {
	// /ip pool => /ip/pool
	path = strings.Replace(path, " ", "/", -1)
	// /ip/pool => routeros_ip_pool
	resNames, ok := providerResources[path]
	if !ok {
		return nil, fmt.Errorf("No resource was found for '%v' path", path)
	}

	if len(resNames) == 0 {
		return nil, fmt.Errorf("Path '%v' exists, but no resource name was found for it", path)
	}

	// Several resources can have the same path.
	resourceName := resNames[0]

	// routeros_ip_pool => ip-pool
	hclName := strings.ReplaceAll(strings.TrimPrefix(resourceName, "routeros_"), "_", "-")
	if !strings.Contains(hclName, "system-") {
		if _, ok := hclNames[hclName]; !ok {
			hclNames[hclName] = 1
		} else {
			hclNames[hclName]++
		}
		hclName = fmt.Sprintf("%v-%v", hclName, hclNames[hclName])
	}

	// {"routeros_ip_pool", "ip-pool-N"}
	return &HCLResource{resourceName, hclName}, nil
}

// routeros_interface_ethernet default_name=ether1 ===> factory_name=ether1, name=ether1
var TransformMap = map[string][]string{
	"routeros_interface_ethernet:default_name": []string{"factory_name", "name"},
}

// The function returns a slice of attributes in HCL format and a slice of Required fields.
func GetAttributes(provider *schema.Provider, resourceName, attributes string) (hclAttributes, required []string) {
	// routeros_interface_ethernet
	resource := provider.ResourcesMap[resourceName]

	// factory_name, name
	var attrRequired = make(map[string]schema.ValueType)
	for k, v := range resource.Schema {
		if v.Required {
			attrRequired[k] = v.Type
		}
	}

	// Padding
	var maxNameLength = 0

	// default-name=ether1
	if pairs := reAttributes.FindAllString(attributes, -1); len(pairs) > 0 {
		// Get the longest length of attribute names
		for _, p := range pairs {
			pp := strings.Split(p, "=")
			if len(pp[0]) > maxNameLength {
				maxNameLength = len(pp[0])
			}
		}

		for _, p := range pairs {
			// default-name=ether1
			pp := strings.Split(p, "=")
			// default_name, ether1
			attrName, attrVaule := routeros.KebabToSnake(pp[0]), pp[1]
			if attrs, ok := TransformMap[resourceName+":"+attrName]; ok {
				for _, name := range attrs {
					// + factory_name=ether1, + name=ether1
					pairs = append(pairs, fmt.Sprintf("%v=%v", name, attrVaule))
				}
			}
		}

		for _, p := range pairs {
			// default-name=ether1
			pp := strings.Split(p, "=")
			attrName, attrVaule := routeros.KebabToSnake(pp[0]), pp[1]

			// "default_name" : {
			// 		Type:        schema.TypeString,
			// 		Computed:    true,
			// 		Description: description,
			// }
			schemaAttr, ok := resource.Schema[attrName]
			if !ok {
				log.Warnf("Attribute '%v' not found", attrName)
				continue
			}

			// Computed:    true,
			if schemaAttr.Computed {
				continue
			}

			switch schemaAttr.Type {
			case schema.TypeString:
				// key=value => key = "value"
				attrVaule = `"` + attrVaule + `"`
			case schema.TypeBool:
				// key=yes => key = true
				attrVaule = routeros.BoolFromMikrotikJSONStr(attrVaule)
			case schema.TypeSet, schema.TypeList:
				// key=a,b,c => key = ["a", "b", "c"]
				switch schemaAttr.Elem.(*schema.Schema).Type {
				case schema.TypeString:
					attrVaule = `["` + strings.Join(strings.Split(attrVaule, ","), `", "`) + `"]`
				default:
					attrVaule = `[` + strings.Join(strings.Split(attrVaule, ","), `,`) + `]`
				}
			}
			// Add padding
			hclAttributes = append(hclAttributes, fmt.Sprintf("%v%v = %v", attrName, strings.Repeat(" ", maxNameLength-len(attrName)), attrVaule))

			// Remove the Required field from the general list
			if schemaAttr.Required {
				required = append(required, p)
				delete(attrRequired, attrName)
			}
		}
	}

	// Add all required fields with the error-causing value
	for attrName, attrType := range attrRequired {
		var value = "?"
		if attrType == schema.TypeString {
			value = `"?"`
		}
		hclAttributes = append(hclAttributes, fmt.Sprintf(`%v%v = %v`, attrName, strings.Repeat(" ", maxNameLength-len(attrName)), value))
	}

	return
}
