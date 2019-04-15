package siesta

import (
	"github.com/sclevine/agouti"
	"github.com/sh-tatsuno/siesta/yaml"
)

type SiestaPage struct {
	Page *agouti.Page
	conf yaml.Yaml
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
		s.RunScript(task)
		//default:
		//return errors.New("no match command")
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
