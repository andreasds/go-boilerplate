package configs_test

import (
	"testing"

	"github.com/andreasds/go-boilerplate/configs"
	"github.com/stretchr/testify/assert"
)

type InputConfig struct {
	FilePath      string
	Configuration interface{}
}

func TestConfig(t *testing.T) {
	t.Run("getConfig", getConfigTestCase)
}

func getConfigTestCase(t *testing.T) {
	tests := []struct {
		name     string
		input    InputConfig
		expected interface{}
	}{
		{
			name: "success normal",
			input: InputConfig{
				FilePath:      ConfigNormalPath,
				Configuration: normalInstance,
			},
			expected: &ConfigNormal{
				Foo: struct {
					Level1 string "mapstructure:\"LEVEL1\""
					Level2 string "mapstructure:\"LEVEL2\""
				}{
					Level1: "FooLevel1",
					Level2: "Foo Level 2",
				},
				Bar: struct {
					Level1 int "mapstructure:\"LEVEL1\""
					Level2 int "mapstructure:\"LEVEL2\""
				}{
					Level1: 9999,
					Level2: -9999,
				},
			},
		},
		{
			name: "success omitempty",
			input: InputConfig{
				FilePath:      ConfigOmitemptyPath,
				Configuration: omitemptyInstance,
			},
			expected: &ConfigOmitempty{
				Foo: struct {
					Level1 string "mapstructure:\"LEVEL1\""
					Level2 int    "mapstructure:\",omitempty\""
				}{
					Level1: "",
					Level2: 0,
				},
				Bar: struct {
					Level1 float64 "mapstructure:\",omitempty\""
					Level2 bool    "mapstructure:\",omitempty\""
				}{
					Level1: 0.0,
					Level2: false,
				},
			},
		},
		{
			name: "success remain",
			input: InputConfig{
				FilePath:      ConfigRemainPath,
				Configuration: remainInstance,
			},
			expected: &ConfigRemain{
				Foo: struct {
					Level1 string                 "mapstructure:\"LEVEL1\""
					Level2 map[string]interface{} "mapstructure:\",remain\""
				}{
					Level1: "FooLevel1",
					Level2: map[string]interface{}{
						"level2": "Foo Level 2",
					},
				},
				Bar: struct {
					Level1 map[string]interface{} "mapstructure:\",remain\""
				}{
					Level1: map[string]interface{}{
						"level1": "BarLevel1",
						"level2": "Bar Level 2",
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			configs.GetConfig(test.input.FilePath, &test.input.Configuration)
			got := test.input.Configuration
			assert.EqualValues(t, got, test.expected)
		})
	}
}
