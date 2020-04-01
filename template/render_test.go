package template

import (
	"fmt"
	"testing"
)

type RenderTest struct {
	context        map[string]interface{}
	template       string
	outputDir      string
	outputFilename string
	error          error
}

func TestRender(t *testing.T) {
	tests := []*RenderTest{
		&RenderTest{
			map[string]interface{}{
				"vpc_id": "vpc-025c50ddacb9519fd",
			},
			rootDir() + "/template/test/templates/test.tfvars.tpl",
			rootDir() + "/template/test/output/",
			"test.tfvars",
			nil,
		},
		&RenderTest{
			map[string]interface{}{
				"vpc_id": "vpc-025c50ddacb9519fd",
			},
			rootDir() + "/template/test/templates_err/err.tpl",
			rootDir() + "/template/test/output/",
			"err",
			fmt.Errorf("Parse template: template: err.tpl:1: missing value for command"),
		},
	}
	for _, test := range tests {
		err := Render(test.context, test.template, test.outputDir, test.outputFilename)
		if test.error == nil && err == nil {
			continue
		} else if (test.error == nil && err != nil) || test.error.Error() != err.Error() {
			t.Errorf("Expected err: %s, got: %s", test.error.Error(), err.Error())
		}
	}
}

type RenderAllTest struct {
	context     map[string]interface{}
	outputDir   string
	template    string
	isDirectory bool
}

func TestRenderAll(t *testing.T) {
	tests := []*RenderAllTest{
		&RenderAllTest{
			map[string]interface{}{
				"vpc_id": "vpc-025c50ddacb9519fd",
			},
			rootDir() + "/template/test/output/",
			rootDir() + "/template/test/templates/test.tfvars.tpl",
			false,
		},
		&RenderAllTest{
			map[string]interface{}{
				"vpc_id": "vpc-025c50ddacb9519fd",
			},
			rootDir() + "/template/test/output/",
			rootDir() + "/template/test/templates/",
			true,
		},
	}
	for _, test := range tests {
		err := RenderAll(test.context, test.outputDir, test.template, test.isDirectory)
		if err != nil {
			t.Errorf("Err: %s", err.Error())
		}
	}
}
