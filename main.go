package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/template"

	"github.com/spf13/viper"
)

func readInput(input string) {
	viper.SetConfigName(filepath.Base(input))
	viper.SetConfigType("yaml")
	viper.AddConfigPath(filepath.Dir(input))
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
}

func render(file, tempplate, outputPath string) {
	m := map[string]interface{}{
		"terraform.tfvars": buildTerraformTfvarsStruct(),
		"config.tf":        buildConfigTfStruct(),
	}
	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		os.MkdirAll(outputPath, os.ModePerm)
	}
	output, err := os.Create(fmt.Sprintf("%s/%s", outputPath, file))
	if err != nil {
		log.Println("create file: ", err)
		return
	}
	t, err := template.ParseFiles(tempplate)
	if err != nil {
		log.Println("parse file: ", err)
		return
	}
	e := t.Execute(output, m[file])
	if e != nil {
		log.Println("executing template:", err)
	}

}

func setupArgs() (*string, *string, *string, *string) {
	input := flag.String("inputFile", "eks.yaml", "input file path and name, example: ./input/eks.yaml")
	output := flag.String("outputDir", ".", "output directory, example: ./output")
	tfvarsTpl := flag.String("tfvarsTemplate", "./templates/terraform.tfvars.tpl", "terraform.tfvars template path and name")
	configTpl := flag.String("configTemplate", "./templates/config.tf.tpl", "config.tf template path and name")
	flag.Parse()
	return input, tfvarsTpl, configTpl, output
}

func main() {
	input, tfvarsTpl, configTpl, output := setupArgs()

	readInput(*input)

	toBeRendered := map[string]string{
		"terraform.tfvars": *tfvarsTpl,
		"config.tf":        *configTpl,
	}

	for file, template := range toBeRendered {
		render(file, template, *output)
	}
}
