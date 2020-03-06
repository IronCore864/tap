package test

import (
	"os"
	"testing"

	"github.com/ironcore864/tap/utils"
)

func TestSetupArgs(t *testing.T) {
	tests := map[string]map[string][]string{
		"args": map[string][]string{
			"normal":  []string{"cmd", "-inputFile", "in.yaml", "-outputDir", "out", "-template", "templates"},
			"file":    []string{"cmd", "-inputFile", "in.yaml", "-outputDir", "out", "-template", "a.tpl"},
			"default": []string{"cmd"},
		},
		"expected": map[string][]string{
			"normal":  []string{"in.yaml", "out", "templates"},
			"file":    []string{"in.yaml", "out", "a.tpl"},
			"default": []string{"./input/eks.yaml", ".", "./templates"},
		},
	}
	for name, args := range tests["args"] {
		os.Args = args
		in, out, template := utils.SetupArgs()
		res := []string{in, out, template}
		for i := 0; i < len(res); i++ {
			if res[i] != tests["expected"][name][i] {
				t.Errorf("Error getting CLI parameters, got: %s, want: %s.", res[i], tests["expected"][name][i])
			}
		}
	}
}

func TestGetOutputFilenameBasedOnFilename(t *testing.T) {
	tests := map[string]string{
		"test.var.tpl":          "test.var",
		"/path/to/test.var.tpl": "test.var",
	}
	for tpl, expected := range tests {
		got := utils.GetOutputFilenameBasedOnFilename(tpl)
		if expected != got {
			t.Errorf("Error getting filename based on template, expected: %s, got %s", expected, got)
		}
	}
}

func TestIsDirectory(t *testing.T) {
	tests := map[string]bool{
		"input":         true,
		"utils.test.go": false,
	}
	for file, expected := range tests {
		got, _ := utils.IsDirectory(file)
		if got != expected {
			t.Errorf("Error detecting if input is directory or not,  expected: %t, got %t", expected, got)
		}
	}
}
