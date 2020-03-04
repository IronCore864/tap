package test

import (
	"testing"

	tpl "github.com/ironcore864/unitet/template"
)

type RenderTest struct {
	context        map[string]interface{}
	template       string
	outputDir      string
	outputFilename string
	error          string
}

func TestRender(t *testing.T) {
	tests := []*RenderTest{
		&RenderTest{
			map[string]interface{}{
				"vpc_id": "vpc-025c50ddacb9519fd",
			},
			rootDir() + "/test/templates/test.tfvars.tpl",
			rootDir() + "/test/output/",
			"test.tfvars",
			"",
		},
	}
	for _, test := range tests {
		err := tpl.Render(test.context, test.template, test.outputDir, test.outputFilename)
		if err != nil {
			t.Errorf("Err: %s", err.Error())
		}
	}
}
