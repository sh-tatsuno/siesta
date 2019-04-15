package siesta

import (
	"fmt"
	"path"
	"regexp"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"github.com/sh-tatsuno/siesta/yaml"
)

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
	selector, err := s.Selector(task)
	if err != nil {
		return err
	}
	err = selector.Click()
	return err
}

func (s SiestaPage) Type(task yaml.Task) error {
	selector, err := s.Selector(task)
	if err != nil {
		return err
	}
	err = selector.Fill(task.Value)
	return err
}

func (s SiestaPage) MouseOver(task yaml.Task) error {
	selector, err := s.Selector(task)
	if err != nil {
		return err
	}
	err = selector.Fill(task.Value)
	return err
}

func (s SiestaPage) SendKeys(task yaml.Task) error {
	selector, err := s.Selector(task)
	if err != nil {
		return err
	}
	err = selector.SendKeys(task.Value)
	return err
}

func (s SiestaPage) MouseOut(task yaml.Task) error {
	// TODO
	return nil
}

func (s SiestaPage) RunScript(task yaml.Task) error {
	err := s.Page.RunScript(task.Target, nil, nil)
	return err
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
