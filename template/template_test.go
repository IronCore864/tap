package template

import (
	"fmt"
	"path"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

func rootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}

type TemplateTest struct {
	inputFile string
	context   map[string]interface{}
	err       string
}

func TestTemplateContext(t *testing.T) {
	tests := []*TemplateTest{
		&TemplateTest{
			rootDir() + "/template/test/input/test.yaml",
			map[string]interface{}{
				"vpc_id": "vpc-025c50ddacb9519fd",
			},
			""},
		&TemplateTest{
			rootDir() + "/template/test/input/test.json",
			map[string]interface{}{
				"vpc_id": "vpc-025c50ddacb9519fd",
			},
			""},
		&TemplateTest{
			rootDir() + "/template/test/input/not_exist.yaml",
			nil,
			"Unable to read configuration file"},
		&TemplateTest{
			rootDir() + "/template/test/input/error.yaml",
			nil,
			"Unable decode the configuration file"},
	}

	for _, test := range tests {
		fmt.Println(test.inputFile)
		ctx, err := NewTemplateContext(test.inputFile)
		if test.err == "" && err != nil {
			t.Errorf("Expected err: %s, got: %s", test.err, err.Error())
		}
		if test.err != "" && !strings.Contains(err.Error(), test.err) {
			t.Errorf("Expected err: %s, got: %s", test.err, err.Error())
		}
		if !reflect.DeepEqual(test.context, ctx) {
			t.Errorf("Expected ctx: %s, got: %s", test.context, ctx)
		}
	}
}
