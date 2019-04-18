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
		var task yaml.Task
		for _, c := range cmds {
			ts := c.Targets
			if len(ts) < 2 {
				task = yaml.Task{
					Name:    c.Comment,
					Command: c.Command,
					Target:  c.Target,
					Value:   c.Value,
				}
			} else {
				var targets []string
				for _, target := range ts {
					if len(target) > 1 {
						targets = append(targets, target[0])
					}
				}
				task = yaml.Task{
					Name:    c.Comment,
					Command: c.Command,
					Targets: targets,
					Value:   c.Value,
				}
			}
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
