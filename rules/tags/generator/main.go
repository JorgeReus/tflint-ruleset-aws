// +build generators

package main

import (
	"bytes"
	"go/format"
	"log"
	"os"
	"sort"
	"text/template"

	utils "github.com/terraform-linters/tflint-ruleset-aws/rules/generator-utils"
)

const filename = "resources.go"

type TemplateData struct {
	Resources []string
}

func main() {
	provider := utils.LoadProviderSchema("../../tools/provider-schema/schema.json")
	resources := make([]string, 0)

	for name, resource := range provider.ResourceSchemas {
		if _, ok := resource.Block.Attributes["tags"]; ok {
			resources = append(resources, name)
		}
	}

	sort.Strings(resources)

	tpl, err := template.New("tagged").Parse(templateBody)
	if err != nil {
		log.Fatalf("error parsing template: %v", err)
	}

	var buffer bytes.Buffer
	err = tpl.Execute(&buffer, &TemplateData{
		Resources: resources,
	})
	if err != nil {
		log.Fatalf("error executing template: %v", err)
	}

	formatted, err := format.Source(buffer.Bytes())
	if err != nil {
		log.Fatalf("error formatting generated file: %v", err)
	}

	f, err := os.Create(filename)
	if err != nil {
		log.Fatalf("error creating file (%s): %v", filename, err)
	}
	defer f.Close()

	_, err = f.Write(formatted)
	if err != nil {
		log.Fatalf("error writing to file (%s): %v", filename, err)
	}

}

const templateBody = `
// Code generated by generator/main.go; DO NOT EDIT.
package tags
var Resources = []string{
	{{- range .Resources }}
	"{{ . }}",
	{{- end }}
}
`
