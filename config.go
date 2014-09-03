package main

import (
	"github.com/acmacalister/skittles"
	"gopkg.in/yaml.v1"
	"io/ioutil"
	"log"
	"os"
)

type Configuration struct {
	Servers  []string `yaml:"servers"`
	Username string   `yaml:"username"`
	Password string   `yaml:"password"`
	Pemfile  string   `yaml:"pemfile"`
	TaskList []string `yaml:"task_list"`
}

func loadConfig(environment string) Configuration {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(skittles.Red(err))
	}
	buffer, err := ioutil.ReadFile(dir + "/kirk.yml")

	if err != nil {
		log.Fatal(skittles.Red(err))
	}

	m := make(map[string]Configuration)
	if err := yaml.Unmarshal(buffer, &m); err != nil {
		log.Fatal(skittles.Red(err))
	}

	return m[environment]
}
