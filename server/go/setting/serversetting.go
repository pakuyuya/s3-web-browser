package setting

import (
	"io/ioutil"
	"path"

	yaml "gopkg.in/yaml.v2"
)

type ServerSettingType struct {
	Port        string `yaml:"Port"`
	Debug       bool   `yaml:"Debug"`
	TextdataDir string `yaml:"TextdataDir"`
	DBHost      string `yaml:"DBHost"`
	DBPort      string `yaml:"DBPort"`
	DBUser      string `yaml:"DBUser"`
	DBPass      string `yaml:"DBPass"`
	DBName      string `yaml:"DBName"`
}

var ServerSetting *ServerSettingType

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

func ResolveTextdatasPath(filename string) (string, error) {
	return path.Join(ServerSetting.TextdataDir, filename), nil
}
