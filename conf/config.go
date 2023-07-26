package conf

import (
	"gopkg.in/yaml.v3"
	"os"
)

func MustLoad(path string, v any) {
	yamlFile, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(yamlFile, v)
	if err != nil {
		panic(err)
	}
}
