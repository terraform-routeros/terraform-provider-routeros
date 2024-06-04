// go run tools/boilerplate/main.go routeros_interface_lte_apn
package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"unicode"

	"github.com/fatih/color"
)

var (
	reNewItemName = regexp.MustCompile(`^routeros_[a-z_]+$`)
	// isDS          = flag.Bool("ds", false, "This is a datasource")
	isSystem = flag.Bool("system", false, "This is a system resource")
)

func Fatalf(format string, a ...any) {
	_, _ = fmt.Fprintf(os.Stderr, format, a...)
	_, _ = fmt.Fprintln(os.Stderr, "")
	os.Exit(1)
}

type ItemType byte

const (
	Resource ItemType = iota + 1
	Datasource
)

func (t ItemType) String() string {
	switch t {
	case Resource:
		return "Resource"
	case Datasource:
		return "Datasource"
	default:
		panic("Unknown ItemType")
	}
}

func (t ItemType) HCL() string {
	switch t {
	case Resource:
		return "resource"
	case Datasource:
		return "datasource"
	default:
		panic("Unknown ItemType")
	}
}

func main() {
	flag.Parse()

	if len(flag.Args()) < 1 {
		Fatalf("Usage: go run tools/bolerplate/main.go routeros_new_resource")
	}

	resName := flag.Args()[0]
	if !reNewItemName.MatchString(resName) {
		Fatalf("The resource name must be in the format: 'routeros_[a-z_]+', got '%v'", resName)
	}

	itemType := Resource
	// if *isDS {
	// 	itemType = Datasource
	// }
	itemCrud := ""
	if *isSystem {
		itemCrud = "System"
	}
	log.Printf("Creating a template for '%v' %v (%v)", resName, color.YellowString(itemType.String()), color.YellowString(itemCrud))

	goName := Capitalize(resName)

	// if !*isDS {
	fName := fmt.Sprintf("%v_%v", Resource.HCL(), strings.TrimPrefix(resName, "routeros_"))
	f, err := os.OpenFile(filepath.Join("routeros", fName+".go"), os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}

	tmpl, err := template.New("res").Parse(resourceFile)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(f, struct {
		GoResourceName string
		System         bool
	}{Resource.String() + goName, *isSystem})
	if err != nil {
		panic(err)
	}
	f.Close()

	f, err = os.OpenFile(filepath.Join("routeros", fName+"_test.go"), os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}

	tmpl, err = template.New("test").Parse(testFile)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(f, struct {
		GoResourceName string
		ResourceName   string
	}{goName, resName})
	if err != nil {
		panic(err)
	}
	f.Close()

	os.MkdirAll(filepath.Join("examples", "resources", resName), os.ModePerm)

	f, err = os.OpenFile(filepath.Join("examples", "resources", resName, "import.sh"), os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	tmpl, err = template.New("ex_import").Parse(exampleImportFile)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(f, struct {
		ResourceName string
	}{resName})
	if err != nil {
		panic(err)
	}
	f.Close()

	f, err = os.OpenFile(filepath.Join("examples", "resources", resName, "resource.tf"), os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	tmpl, err = template.New("ex_res").Parse(exampleResourceFile)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(f, struct {
		ResourceName string
	}{resName})
	if err != nil {
		panic(err)
	}
	f.Close()

	f, err = os.OpenFile(filepath.Join("routeros", "provider.go"), os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(f, `"%v":    %v(),\n`, resName, Resource.String()+goName)
	f.Close()
	// }
}

var exampleImportFile = `#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/ get [print show-ids]]
terraform import {{.ResourceName}}.test *3`

var exampleResourceFile = `
resource "{{.ResourceName}}" "test" {
}`

var testFile = `
package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const test{{.GoResourceName}} = "{{.ResourceName}}.test"

func TestAcc{{.GoResourceName}}Test_basic(t *testing.T) {
	t.Parallel()
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				Steps: []resource.TestStep{
					{
						Config: testAcc{{.GoResourceName}}Config(""),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(test{{.GoResourceName}}),
							resource.TestCheckResourceAttr(test{{.GoResourceName}}, "", ""),
						),
					},
					{
						Config: testAcc{{.GoResourceName}}Config(""),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(test{{.GoResourceName}}),
							resource.TestCheckResourceAttr(test{{.GoResourceName}}, "", ""),
						),
					},
				},
			})

		})
	}
}

func testAcc{{.GoResourceName}}Config(param string) string {
	return fmt.Sprintf(` + "`" + `%v

resource "{{.ResourceName}}" "test" {
}
` + "`" + `, providerConfig, param)
}
`

var resourceFile = `
package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
REST JSON
*/

// https://help.mikrotik.com/docs/display/ROS/
func {{.GoResourceName}}() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/"),
		MetaId:           PropId(Id),

	}

	return &schema.Resource{
		CreateContext: Default{{- if .System }}System{{end}}Create(resSchema),
		ReadContext:   Default{{- if .System }}System{{end}}Read(resSchema),
		UpdateContext: Default{{- if .System }}System{{end}}Update(resSchema),
		DeleteContext: Default{{- if .System }}System{{end}}Delete(resSchema),

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resSchema,
	}
}`

func Capitalize(s string) (res string) {
	s = strings.TrimPrefix(s, "routeros")
	var uc = false

	for _, c := range s {
		if c == '_' {
			uc = true
			continue
		}

		if uc {
			res += string(unicode.ToUpper(c))
		} else {
			res += string(c)
		}
		uc = false
	}

	return
}
