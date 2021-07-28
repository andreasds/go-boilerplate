package configs_test

const ConfigNormalPath = "./data/normal.env"

var normalInstance *ConfigNormal

type ConfigNormal struct {
	Foo struct {
		Level1 string `mapstructure:"LEVEL1"`
		Level2 string `mapstructure:"LEVEL2"`
	} `mapstructure:"FOO"`

	Bar struct {
		Level1 int `mapstructure:"LEVEL1"`
		Level2 int `mapstructure:"LEVEL2"`
	} `mapstructure:"BAR"`
}

const ConfigOmitemptyPath = "./data/omitempty.env"

var omitemptyInstance *ConfigOmitempty

type ConfigOmitempty struct {
	Foo struct {
		Level1 string `mapstructure:"LEVEL1"`
		Level2 int    `mapstructure:",omitempty"`
	} `mapstructure:"FOO"`

	Bar struct {
		Level1 float64 `mapstructure:",omitempty"`
		Level2 bool    `mapstructure:",omitempty"`
	} `mapstructure:"BAR"`
}

const ConfigRemainPath = "./data/remain.env"

var remainInstance *ConfigRemain

type ConfigRemain struct {
	Foo struct {
		Level1 string                 `mapstructure:"LEVEL1"`
		Level2 map[string]interface{} `mapstructure:",remain"`
	} `mapstructure:"FOO"`

	Bar struct {
		Level1 map[string]interface{} `mapstructure:",remain"`
	} `mapstructure:"BAR"`
}
