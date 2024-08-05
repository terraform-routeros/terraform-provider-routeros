// go run tools/boilerplate/main.go routeros_interface_lte_apn
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"text/template"
	"unicode"

	"github.com/fatih/color"
)

var (
	reNewItemName = regexp.MustCompile(`^routeros_[a-z_]+$`)
	// isDS          = flag.Bool("ds", false, "This is a datasource")
	isSystem = flag.Bool("system", false, "This is a system resource")
	csvTable = flag.Bool("table", false, "Extracting attributes from the WIKI table")
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

	if *csvTable {
		if _, err := os.Stat(flag.Args()[0]); err != nil {
			Fatalf("CSV file %v not found", flag.Args()[0])
		}
		extractAttributes(flag.Args()[0])
		os.Exit(0)
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

	os.MkdirAll("routeros", os.ModePerm)

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

	var flags int = os.O_WRONLY | os.O_APPEND
	if _, err := os.Stat(filepath.Join("routeros", "provider.go")); err != nil {
		flags |= os.O_CREATE
	}

	f, err = os.OpenFile(filepath.Join("routeros", "provider.go"), flags, os.ModePerm)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(f, "\"%v\":    %v(),\n", resName, Resource.String()+goName)
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

var attribute = `    "{{.Attribute}}": {
        Type: schema.Type{{.Type}},
        Optional: true,
        Description: "{{.Description}}",
        {{- if .Slice }}
	ValidateFunc: validation.StringInSlice([]string{ "{{.Slice}}" }, false),{{ end }}
        {{- if .DiffSuppress }}
	DiffSuppressFunc: AlwaysPresentNotUserProvided,{{ end }}
    },
`

var (
	reCSV         = regexp.MustCompile(`(?m)"(.*?)"(?:,|$)`)
	reAttrName    = regexp.MustCompile(`[a-z-]+`)
	reAttrDefault = regexp.MustCompile(`(?m)Default:?\s*(""|\w+)`)
	reAttrEnum    = regexp.MustCompile(`(?m)\(\s*([\w-| ]+);`)
	enumReplacer  = strings.NewReplacer(" ", "", `"`, "`", "'", "`", "|", `", "`)
)

func extractAttributes(filename string) {
	tmpl, err := template.New("attr").Parse(attribute)
	if err != nil {
		panic(err)
	}
	tmpl.Option()

	file, err := os.Open(filename)
	if err != nil {
		Fatalf("[extractAttributes] %v", err)
	}
	defer file.Close()

	w := os.Stdout

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		rec := reCSV.FindAllStringSubmatch(row, -1)
		if len(rec) != 2 {
			fmt.Fprintln(w, row)
			continue
		}

		r1, r2 := rec[0][1], rec[1][1]

		if len(r1) > 0 && r1[0] == '"' {
			r1 = r1[1:]
		}
		if len(r1) > 0 && r1[len(r1)-1] == '"' {
			r1 = r1[:len(r1)-1]
		}

		if len(r2) > 0 && r2[0] == '"' {
			r2 = r2[1:]
		}
		if len(r2) > 0 && r2[len(r2)-1] == '"' {
			r2 = r2[:len(r2)-1]
		}

		// [ ["Property", Property] ["Description" Description] ]
		if r1 == "Property" && r2 == "Description" {
			continue
		}

		var diffSuppress bool
		attrType := "String"
		if res := reAttrDefault.FindStringSubmatch(r1); len(res) > 1 {
			switch res[1] {
			// src-address (Default:"")
			case `""`:
			// use-network-apn (yes | no; Default: yes)
			case "yes", "no":
				attrType = "Bool"
				diffSuppress = true
			// startup-delay (Default: 5m)
			default:
				diffSuppress = true
				if _, err := strconv.Atoi(res[1]); err == nil {
					attrType = "Int"
				}
			}
		}

		var validate string
		for _, match := range reAttrEnum.FindAllStringSubmatch(r1, -1) {
			validate = enumReplacer.Replace(match[1])
		}

		ww := os.Stdout

		tmpl.Execute(ww, struct {
			Attribute    string
			Type         string
			Description  string
			Slice        string
			DiffSuppress bool
		}{
			Attribute:    strings.ReplaceAll(reAttrName.FindString(r1), "-", "_"),
			Type:         attrType,
			Description:  strings.ReplaceAll(r2, `"`, "`"),
			Slice:        validate,
			DiffSuppress: diffSuppress,
		})

		if r1 == "type" {
			os.Exit(0)
		}

		if err != nil {
			Fatalf("%v", err)
		}
	}
}
