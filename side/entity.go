package side

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/pkg/errors"
)

type Side struct {
	ID      string   `json:"id"`
	Version string   `json:"version"`
	Name    string   `json:"name"`
	URL     string   `json:"url"`
	Tests   []Test   `json:"tests"`
	Suites  []Suite  `json:"suites"`
	URLs    []string `json:"urls"`
	Plugins []string `json:"plugins"`
}

type Test struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Commands []Command `json:"commands"`
}

type Command struct {
	ID      string     `json:"id"`
	Comment string     `json:"comment"`
	Command string     `json:"command"`
	Target  string     `json:"target"`
	Targets [][]string `json:"targets"`
	Value   string     `json:"value"`
}

type Suite struct {
	ID             string   `json:"id"`
	Name           string   `json:"name"`
	PersistSession bool     `json:"persistSession"`
	Parallel       bool     `json:"parallel"`
	Timeout        int      `json:"timeout"`
	Tests          []string `json:"tests"`
}

func (s Side) Save(filePath string) error {
	sideOutput, err := json.Marshal(s)
	if err != nil {
		err = errors.Wrap(err, "Failed to Marshal side")
		return err
	}
	if err := ioutil.WriteFile(filePath, sideOutput, 0777); err != nil {
		err = errors.Wrap(err, "Failed to save side")
		return err
	}
	return nil
}

func Load(filePath string) *Side {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	// decode json
	side := new(Side)
	if err := json.Unmarshal(bytes, &side); err != nil {
		log.Fatal(err)
	}
	return side
}
