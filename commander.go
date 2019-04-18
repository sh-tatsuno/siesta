package siesta

import (
	"github.com/sh-tatsuno/siesta/yaml"
)

// TODO: commentの拡充
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
