package main

import (
	"log"
	"os"

	tpl "github.com/ironcore864/unitet/template"
	"github.com/ironcore864/unitet/utils"
)

func main() {
	input, outputDir, template := utils.SetupArgs()

	context, err := tpl.NewTemplateContext(input)
	if err != nil {
		log.Println("Parse input: ", err)
		os.Exit(1)
	}

	isDirectory, err := utils.IsDirectory(template)
	if err != nil {
		log.Println("Error getting template: ", err)
		os.Exit(1)
	}

	err = tpl.RenderAll(context, outputDir, template, isDirectory)
	if err != nil {
		log.Println("Error rendering: ", err)
		os.Exit(1)
	}
}
