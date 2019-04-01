package siesta

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/sh-tatsuno/siesta/side"
	"github.com/sh-tatsuno/siesta/yaml"
	yaml2 "gopkg.in/yaml.v2"
)

func ReadSide(filePath string) *side.Side {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	// decode json
	s := new(side.Side)
	if err := json.Unmarshal(bytes, &s); err != nil {
		log.Fatal(err)
	}
	return s
}

func ReadYaml(filePath string) *yaml.Yaml {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	// decode yaml
	y := new(yaml.Yaml)
	if err := yaml2.Unmarshal(bytes, &y); err != nil {
		log.Fatal(err)
	}
	return y
}
