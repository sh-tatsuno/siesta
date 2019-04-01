package siesta

import (
	"fmt"
	"path"
	"regexp"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"github.com/sclevine/agouti"
	"github.com/sh-tatsuno/siesta/yaml"
)

type SiestaPage struct {
	Page *agouti.Page
	conf yaml.Yaml
}

func (s SiestaPage) Open(task yaml.Task) error {
	url := task.Target
	// check base url & join if target means path
	if len(s.conf.URL) > 0 && len(url) > 0 {
		if string([]rune(url)[0]) == "/" {
			url = path.Join(s.conf.URL, url)
		}
	}
	if err := s.Page.Navigate(url); err != nil {
		err = errors.Wrap(err, "Failed to navigate")
		return err
	}

	return nil
}

func (s SiestaPage) SetWindowSize(task yaml.Task) error {
	arr := regexp.MustCompile("x").Split(task.Target, -1)
	if len(arr) != 2 {
		err := fmt.Errorf("Uncompatible array length: expected 2 actual %d", len(arr))
		return err
	}

	width, err := strconv.Atoi(arr[0])
	if err != nil {
		err = fmt.Errorf("cannot cast width to int of %s", arr[0])
		return err
	}

	height, err := strconv.Atoi(arr[1])
	if err != nil {
		err = fmt.Errorf("cannot cast height to int of %s", arr[1])
		return err
	}

	if err := s.Page.Size(width, height); err != nil {
		return err
	}

	return nil
}

func (s SiestaPage) Click(task yaml.Task) error {
	// TODO
	return nil
}

func (s SiestaPage) Type(task yaml.Task) error {
	// TODO
	return nil
}

func (s SiestaPage) MouseOver(task yaml.Task) error {
	// TODO
	return nil
}

func (s SiestaPage) SendKeys(task yaml.Task) error {
	// TODO
	return nil
}

func (s SiestaPage) MouseOut(task yaml.Task) error {
	// TODO
	return nil
}

func (s SiestaPage) RunScript(task yaml.Task) error {
	// TODO
	return nil
}

func (s SiestaPage) Pause(task yaml.Task) error {
	ms, err := strconv.Atoi(task.Target)
	if err != nil {
		err = fmt.Errorf("cannot cast pause to int of %s", task.Target)
		return err
	}
	time.Sleep(time.Duration(ms) * time.Millisecond)
	return nil
}

func (s SiestaPage) SelectWindow(task yaml.Task) error {
	// TODO
	return nil
}

func (s SiestaPage) Commander(task yaml.Task) error {
	switch task.Command {
	case "open":
		s.Open(task)
	case "setWindowSize":
		s.SetWindowSize(task)
	case "click":
		s.Click(task)
	case "type":
		s.Type(task)
	case "mouseOver":
		s.MouseOver(task)
	case "sendKeys":
		s.SendKeys(task)
	case "mouseOut":
		s.MouseOut(task)
	case "pause":
		s.Pause(task)
	case "selectWindow":
		s.SelectWindow(task)
	case "runScript":
		s.RunScript(task) // default runscript
	default:
		return errors.New("no match command")
	}
	return nil
}

/* TODO */
// runscript

/* capture screenshot*/
// runscript
// window.captureEntirePageScreenshot
//

// "id": "52554b60-15b4-42ec-b939-e4236039fcd7",
// "comment": "",
// "command": "type",
// "target": "name=q",
// "targets": [
//   ["name=q", "name"],
//   ["css=.gLFyf", "css:finder"],
//   ["xpath=//input[@name='q']", "xpath:attributes"],
//   ["xpath=//form[@id='tsf']/div[2]/div/div/div/div/input", "xpath:idRelative"],
//   ["xpath=//div/div/input", "xpath:position"]
// ],
// "value": "hello"
// }, {
// "id": "ab952c1a-541d-4105-b5c6-0f1945456b75",
// "comment": "",
// "command": "sendKeys",
// "target": "name=q",
// "targets": [
//   ["name=q", "name"],
//   ["css=.gLFyf", "css:finder"],
//   ["xpath=//input[@name='q']", "xpath:attributes"],
//   ["xpath=//form[@id='tsf']/div[2]/div/div/div/div/input", "xpath:idRelative"],
//   ["xpath=//div/div/input", "xpath:position"]
// ],
// "value": "${KEY_ENTER}"
// }, {
// "id": "227807fc-2ddf-4724-8541-bf202f3ea595",
// "comment": "",
// "command": "pause",
// "target": "5000",
// "targets": [],
// "value": ""
