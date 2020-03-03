package main

import (
	"io/ioutil"
	"log"
	"os"

	tpl "github.com/ironcore864/tftpl/template"
	"github.com/ironcore864/tftpl/utils"
)

func main() {
	input, outputDir, templateDir := utils.SetupArgs()

	context, err := tpl.NewTemplateContext(*input)
	if err != nil {
		log.Println("Parse input: ", err)
		os.Exit(1)
	}

	items, _ := ioutil.ReadDir(*templateDir)
	for _, item := range items {
		if !item.IsDir() {
			outputFileName := utils.GetOutputFilenameBasedOnTemplateFilename(item.Name())
			tpl.Render(context, *templateDir+"/"+item.Name(), *outputDir, outputFileName)
		}
	}
}
