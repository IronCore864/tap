package utils

import (
	"flag"
	"log"
	"os"
	"path"
	"strings"
)

// SetupArgs reads CLI parameters and return them as three strings
func SetupArgs() (*string, *string, *string) {
	var cmd = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	input := cmd.String("inputFile", "./input/eks.yaml", "input file path and name, example: ./input/eks.yaml")
	outputDir := cmd.String("outputDir", ".", "output directory, example: ./output")
	templateDir := cmd.String("templateDir", "./templates", "directory containing all templates to be rendered, example: ./templates")
	if err := cmd.Parse(os.Args[1:]); err != nil {
		log.Println("Parse cmd params: ", err)
		os.Exit(1)
	}
	return input, outputDir, templateDir
}

// GetOutputFilenameBasedOnTemplateFilename gets "test.var.tpl" and returns "test.var" as string(removing .tpl)
func GetOutputFilenameBasedOnTemplateFilename(filename string) string {
	return strings.TrimSuffix(filename, path.Ext(filename))
}
