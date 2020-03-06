package utils

import (
	"flag"
	"log"
	"os"
	"path"
	"strings"
)

// SetupArgs reads CLI parameters and return them as three strings
func SetupArgs() (string, string, string) {
	var cmd = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	input := cmd.String("inputFile", "", "input file path and name, example: ./input/input.yaml")
	outputDir := cmd.String("outputDir", ".", "output directory, example: ./output, default current directory")
	template := cmd.String("template", "", "if given a file, the file of the template to be rendered; if given a directory, all templates under the directory will be rendered")
	if err := cmd.Parse(os.Args[1:]); err != nil {
		log.Println("Parse cmd params: ", err)
		os.Exit(1)
	}
	return *input, *outputDir, *template
}

// GetOutputFilenameBasedOnFilename gets "test.var.tpl" and returns "test.var" as string(removing .tpl)
func GetOutputFilenameBasedOnFilename(filename string) string {
	names := strings.Split(filename, "/")
	f := names[len(names)-1]
	return strings.TrimSuffix(f, path.Ext(f))
}

// IsDirectory returns true if a given path is a directory rather than file
func IsDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return fileInfo.IsDir(), err
}
