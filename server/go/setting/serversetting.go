package setting

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

// ServerSettingType is a struct has properties in setting.yml
type ServerSettingType struct {
	Port        string `yaml:"Port"`
	Debug       bool   `yaml:"Debug"`
	TextdataDir string `yaml:"TextdataDir"`
	DBHost      string `yaml:"DBHost"`
	DBPort      string `yaml:"DBPort"`
	DBUser      string `yaml:"DBUser"`
	DBPass      string `yaml:"DBPass"`
	DBName      string `yaml:"DBName"`
	AuthDisabled bool  `yaml:"AuthDisabled"`
}

// ServerSetting has properties in setting.yml
var ServerSetting *ServerSettingType

// LoadSetting is a function that load properties from setting.yml to ServerSetting
func LoadSetting() error {
	bytes, err := ioutil.ReadFile(`./setting.yml`)
	if err != nil {
		return err
	}

	s := ServerSettingType{}
	err = yaml.Unmarshal(bytes, &s)
	if err != nil {
		return err
	}
	ServerSetting = &s
	return nil
}

