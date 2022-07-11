package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func GetYaml(filename string, model interface{}) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	if err = yaml.Unmarshal(file, model); err != nil {
		return err
	}
	return nil
}
