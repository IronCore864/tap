package test

import (
	"os"
	"testing"

	"github.com/ironcore864/unitet/utils"
)

func TestSetupArgs(t *testing.T) {
	tests := map[string]map[string][]string{
		"args": map[string][]string{
			"normal":  []string{"cmd", "-inputFile", "in", "-outputDir", "out", "-templateDir", "templates"},
			"default": []string{"cmd"},
		},
		"expected": map[string][]string{
			"normal":  []string{"in", "out", "templates"},
			"default": []string{"./input/eks.yaml", ".", "./templates"},
		},
	}
	for name, args := range tests["args"] {
		os.Args = args
		in, out, template := utils.SetupArgs()
		res := []string{*in, *out, *template}
		for i := 0; i < len(res); i++ {
			if res[i] != tests["expected"][name][i] {
				t.Errorf("Error getting CLI parameters, got: %s, want: %s.", res[i], tests["expected"][name][i])
			}
		}
	}
}

func TestGetOutputFilenameBasedOnTemplateFilename(t *testing.T) {
	expected := "test.var"
	got := utils.GetOutputFilenameBasedOnTemplateFilename("test.var.tpl")
	if expected != got {
		t.Errorf("Error getting filename based on template, expected: %s, got %s", expected, got)
	}
}
