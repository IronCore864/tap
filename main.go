package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/ghodss/yaml"
)

func newTemplateContext(file string) (map[string]interface{}, error) {
	ctx := make(map[string]interface{})
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("unable to Read configuration file: %s, error: %s", file, err)
	}
	ctx = make(map[string]interface{})
	if err := yaml.Unmarshal(content, &ctx); err != nil {
		return nil, fmt.Errorf("unable decode the configuration file: %s, error: %v", file, err)
	}
	return ctx, nil
}

func render(ctx interface{}, tpl, outputPath, outputFile string) {
	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		os.MkdirAll(outputPath, os.ModePerm)
	}
	output, err := os.Create(fmt.Sprintf("%s/%s", outputPath, outputFile))
	if err != nil {
		log.Println("create outputFile: ", err)
		return
	}
	t, err := template.ParseFiles(tpl)
	if err != nil {
		log.Println("parse template: ", err)
		return
	}
	e := t.Execute(output, ctx)
	if e != nil {
		log.Println("executing tpl:", err)
	}
}

func setupArgs() (*string, *string, *string) {
	input := flag.String("inputFile", "./input/eks.yaml", "input file path and name, example: ./input/eks.yaml")
	outputDir := flag.String("outputDir", ".", "output directory, example: ./output")
	templateDir := flag.String("templateDir", "./templates", "directory containing all templates to be rendered, example: ./templates")
	flag.Parse()
	return input, outputDir, templateDir
}

func main() {
	input, outputDir, templateDir := setupArgs()

	context, err := newTemplateContext(*input)
	if err != nil {
		log.Println("parse input: ", err)
		os.Exit(1)
	}

	items, _ := ioutil.ReadDir(*templateDir)
	for _, item := range items {
		if !item.IsDir() {
			outputFileName := strings.TrimSuffix(item.Name(), path.Ext(item.Name()))
			render(context, "templates/"+item.Name(), *outputDir, outputFileName)
		}
	}
}
