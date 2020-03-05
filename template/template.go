package template

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/template"

	"github.com/ghodss/yaml"
	"github.com/ironcore864/unitet/utils"
)

// NewTemplateContext reads YAML input file and returns a context used for rendering
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

// Render renders a template with given context and output to a file
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

// RenderAll renders all the templates (if given input is a directory) or a single file (if given input is a file)
func RenderAll(context map[string]interface{}, outputDir, template string, isDirectory bool) error {
	if isDirectory {
		items, _ := ioutil.ReadDir(template)
		for _, item := range items {
			if !item.IsDir() {
				outputFileName := utils.GetOutputFilenameBasedOnFilename(item.Name())
				err := Render(context, template+"/"+item.Name(), outputDir, outputFileName)
				if err != nil {
					return err
				}
			}
		}
	} else {
		outputFileName := utils.GetOutputFilenameBasedOnFilename(template)
		err := Render(context, template, outputDir, outputFileName)
		if err != nil {
			return err
		}
	}
	return nil
}
