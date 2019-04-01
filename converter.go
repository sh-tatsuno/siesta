package siesta

import (
	"github.com/sh-tatsuno/siesta/side"
	"github.com/sh-tatsuno/siesta/yaml"
)

func ConvertSideToYaml(s *side.Side) *yaml.Yaml {

	ts := s.Tests
	tests := []yaml.Test{}
	for _, t := range ts {
		test := yaml.Test{
			Name:  t.Name,
			Tasks: []yaml.Task{},
		}
		cmds := t.Commands
		for _, c := range cmds {
			task := yaml.Task{
				Description: c.Comment,
				Command:     c.Command,
				Target:      c.Target,
				Value:       c.Value,
			}
			// TODO: jump if runscript
			test.Tasks = append(test.Tasks, task)
		}
		tests = append(tests, test)
	}

	suites := s.Suites
	y := &yaml.Yaml{
		Name:    s.Name,
		URL:     s.URL,
		Tests:   tests,
		TimeOut: suites[0].Timeout, // TODO: suitesごとにyaml分割できるように修正
	}

	return y
}

// TODO:
//func ConvertYamlToSide(ys []*yaml.Yaml) *side.Side {
//}
