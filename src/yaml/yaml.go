package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

const yamlFile = "conf.yaml"

type Message struct {
	Environments map[string]models `yaml:"environments"`
	hello        string
}
type models map[string][]Model
type Model struct {
	AppType     string `yaml:"app-type"`
	ServiceType string `yaml:"service-type"`
	Enable      bool   `yaml: "enable"`
}

func main() {
	y := Message{}

	f, err := ioutil.ReadFile(yamlFile)
	if err != nil {
		fmt.Println(err)
	}

	err = yaml.Unmarshal(f, &y)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	y.hello = "world"
	fmt.Println(y)
	fmt.Printf("%+v\n", y.Environments)

}
