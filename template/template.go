package template

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/template"

	"github.com/ghodss/yaml"
)

// NewTemplateContext reads input file and returns a context used for rendering
func NewTemplateContext(file string) (map[string]interface{}, error) {
	ctx := make(map[string]interface{})

	content, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("Unable to read configuration file: %s, error: %s", file, err)
	}

	if err := yaml.Unmarshal(content, &ctx); err != nil {
		return nil, fmt.Errorf("Unable decode the configuration file: %s, error: %v", file, err)
	}

	return ctx, nil
}

// Render takes context and the template and renders the template and write to output file
func Render(ctx interface{}, tpl, outputPath, outputFile string) error {
	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		os.MkdirAll(outputPath, os.ModePerm)
	}

	output, err := os.Create(fmt.Sprintf("%s/%s", outputPath, outputFile))
	if err != nil {
		log.Println(err)
		return fmt.Errorf("Create outputFile: %s", err)
	}

	t, err := template.ParseFiles(tpl)
	if err != nil {
		log.Println("Parse template: ", err)
		return fmt.Errorf("Parse template: %s", err)
	}

	e := t.Execute(output, ctx)
	if e != nil {
		return fmt.Errorf("Executing tpl: %s", err)
	}

	return nil
}
