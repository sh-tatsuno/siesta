package yaml

import (
	"log"

	"github.com/sclevine/agouti"
)

// TODO:
// generate agouti from yaml
func (y *Yaml) Parse() {
	options := []string{"--headless"}
	driver := agouti.ChromeDriver(agouti.ChromeOptions("args", options))

	if err := driver.Start(); err != nil {
		log.Fatalf("Failed to start driver:%v", err)
	}

}
