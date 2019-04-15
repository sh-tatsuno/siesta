package siesta

import (
	"fmt"
	"regexp"

	"github.com/sclevine/agouti"
	"github.com/sh-tatsuno/siesta/yaml"
)

// Selector : select element from task
// premise to use targets, not target
func (s *SiestaPage) Selector(task yaml.Task) (*agouti.Selection, error) {

	if task.Targets == nil {
		return nil, nil
	}

	var selector *agouti.Selection
	for ind, t := range task.Targets {
		arr := regexp.MustCompile("=").Split(t, 2)
		if len(arr) != 2 {
			err := fmt.Errorf("must be over 1")
			return nil, err
		}
		// have to test when len(arr) == 1
		tag := arr[0]
		query := arr[1]

		if ind == 0 {
			switch tag {
			case "id":
				selector = s.Page.FindByID(query)

			case "css":
				selector = s.Page.FindByClass(query)

			case "name":
				selector = s.Page.FindByName(query)

			case "xpath":
				selector = s.Page.FindByXPath(query)

			case "linkText":
				selector = s.Page.FindByLink(query)

			default:
				return nil, nil
			}
		} else {
			switch tag {
			case "id":
				selector = selector.FindByID(query)

			case "css":
				selector = selector.FindByClass(query)

			case "name":
				selector = selector.FindByName(query)

			case "xpath":
				selector = selector.FindByXPath(query)

			case "linkText":
				selector = selector.FindByLink(query)

			default:
				continue
			}
		}
	}

	return selector, nil
}
