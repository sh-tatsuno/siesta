package yaml

import (
	"io/ioutil"

	"github.com/pkg/errors"
	yaml2 "gopkg.in/yaml.v2"
)

type Yaml struct {
	Name    string `yaml:"name"`
	URL     string `yaml:"url"`
	Tests   []Test `yaml:"tests"`
	SaveDir string `yaml:"save_dir"`
	TimeOut int    `yaml:"timeout"`
}

type Test struct {
	Name  string `yaml:"name"`
	Tasks []Task `yaml:"tasks"`
}

type Task struct {
	Name    string   `yaml:"name"`
	Command string   `yaml:"command"`
	Target  string   `yaml:"target,omitempty"`
	Targets []string `yaml:"targets,omitempty"`
	Value   string   `yaml:"value"`
}

func (y Yaml) Save(filePath string) error {
	yamlOutput, err := yaml2.Marshal(y)
	if err != nil {
		err = errors.Wrap(err, "Failed to Marshal yaml")
		return err
	}
	if err := ioutil.WriteFile(filePath, yamlOutput, 0777); err != nil {
		err = errors.Wrap(err, "Failed to save yaml")
		return err
	}
	return nil
}
