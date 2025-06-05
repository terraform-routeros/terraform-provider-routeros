package main

import (
	"compress/gzip"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"slices"
	"sort"
	"strings"

	"github.com/fatih/color"
	godiffpatch "github.com/sourcegraph/go-diff-patch"
	"github.com/terraform-routeros/terraform-provider-routeros/routeros"
)

// https://tikoci.github.io/restraml/$v/inspect.json" -o- | gzip > ros-$v.json.gz

type Schema map[string]any

var (
	versionRange = flag.String("r", "", "Version range <7.18:7.19>")
	filter       = flag.String("f", "", "Filter [/routing/bgp]")
	all          = flag.Bool("all", false, "Show all changes")
	markdown     = flag.Bool("markdown", false, "Markdown")
)

func main() {
	flag.Parse()

	if *versionRange == "" {
		log.Fatalln("'-r' must not be empty")
	}

	vv := strings.Split(*versionRange, ":")
	if len(vv) != 2 {
		log.Fatal("Versions should be specified in 'X:Y' format")
	}

	f1, err := os.Open("ros-" + vv[0] + ".json.gz")
	if err != nil {
		log.Fatal(err)
	}
	defer f1.Close()

	gz1, err := gzip.NewReader(f1)
	if err != nil {
		log.Fatal(err)
	}
	defer gz1.Close()

	f2, err := os.Open("ros-" + vv[1] + ".json.gz")
	if err != nil {
		log.Fatal(err)
	}
	defer f2.Close()

	gz2, err := gzip.NewReader(f2)
	if err != nil {
		log.Fatal(err)
	}
	defer gz2.Close()

	if *markdown {
		fmt.Println("```diff")
	}

	fmt.Printf("%v <=> %v\n", color.RedString("-"+vv[0]), color.GreenString("+"+vv[1]))

	var schema = Schema{}
	err = json.NewDecoder(gz1).Decode(&schema)
	if err != nil {
		log.Fatal(err)
	}

	var localFilter []string
	if *filter != "" {
		localFilter = []string{*filter}
	} else {
		localFilter = []string{}
		for _, v := range routeros.Provider().ResourcesMap {
			localFilter = append(localFilter, v.Schema[routeros.MetaResourcePath].Default.(string))
		}

		sort.Strings(localFilter)
	}

	if *all {
		localFilter = []string{}
	}

	oldSchema := Go(schema, "", localFilter)
	sort.Strings(oldSchema)
	oldSchema = slices.Compact(oldSchema)

	schema = Schema{}
	err = json.NewDecoder(gz2).Decode(&schema)
	if err != nil {
		log.Fatal(err)
	}

	newSchema := Go(schema, "", localFilter)
	sort.Strings(newSchema)
	newSchema = slices.Compact(newSchema)

	diff := godiffpatch.GeneratePatch("schema.json", strings.Join(oldSchema, "\n"), strings.Join(newSchema, "\n"))

	var resource string
	for _, v := range strings.Split(diff, "\n") {
		if len(v) < 2 || (v[0] != '-' && v[0] != '+') || v[:2] == "--" || v[:2] == "++" {
			continue
		}

		// +/routing/bgp/template:input.filter-communities
		ss := strings.SplitN(v, ":", 2)
		if resource != ss[0][1:] {
			fmt.Println("\n" + ss[0][1:])
		}

		switch v[0] {
		case '-':
			if *markdown {
				fmt.Printf("-%v\n", color.RedString(ss[1]))
			} else {
				fmt.Printf("   %v\n", color.RedString(ss[1]))
			}
		case '+':
			if *markdown {
				fmt.Printf("+%v\n", color.GreenString(ss[1]))
			} else {
				fmt.Printf("   %v\n", color.GreenString(ss[1]))
			}
		}
		resource = ss[0][1:]
	}
	if *markdown {
		fmt.Println("```")
	}
}

var (
	rr = strings.NewReplacer("/add", "", "/set", "", "/reset", "")
)

func filterContains(base string, filter []string) bool {
	base = rr.Replace(base)
	for _, v := range filter {
		if base == v {
			return true
		}
	}
	return false
}

func Go(s Schema, base string, filter []string) []string {
	var res []string

	for k, v := range s {
		// if k == "_type" && v == "dir" {
		// 	fmt.Printf("%2d %v\n", level, base)
		// }
		if k != "_type" && (path.Base(base) == "add" || path.Base(base) == "set" || path.Base(base) == "reset") {
			if len(filter) > 0 {
				if filterContains(base, filter) {
					res = append(res, fmt.Sprintf("%v:%v", path.Dir(base), k))
				}
			} else {
				res = append(res, fmt.Sprintf("%v:%v", path.Dir(base), k))
			}
		}

		switch vv := v.(type) {
		case map[string]any:
			res = append(res, Go(vv, base+"/"+k, filter)...)
		default:
		}
	}

	return res
}
