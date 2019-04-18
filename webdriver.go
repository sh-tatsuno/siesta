package siesta

import (
	"log"

	"github.com/sclevine/agouti"
	"github.com/sh-tatsuno/siesta/yaml"
)

type SiestaPage struct {
	Page *agouti.Page
	conf *yaml.Yaml
}

func (s *SiestaPage) Run() error {
	var err error
	for _, test := range s.conf.Tests {
		for _, task := range test.Tasks {
			err = s.Commander(task)
		}
	}
	return err
}

func CreateSiesta(y *yaml.Yaml) SiestaPage {
	options := []string{"--headless"}
	driver := agouti.ChromeDriver(agouti.ChromeOptions("args", options))

	if err := driver.Start(); err != nil {
		log.Fatalf("Failed to start driver:%v", err)
	}

	page, err := driver.NewPage(agouti.Browser("chrome"))
	if err != nil {
		log.Fatalf("Failed to open page:%v", err)
	}

	s := SiestaPage{
		Page: page,
		conf: y,
	}

	return s
}

// TODO:
// side読み込んでyaml保存
// yaml読み込んでside保存

// yaml読み込んで実行
