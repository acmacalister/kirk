package main

import (
	"github.com/acmacalister/skittles"
	"github.com/kylelemons/go-gypsy/yaml"
	"log"
	"strings"
)

type configuration struct {
	servers  []string
	username string
	password string
	pemfile  string
	taskList []string
}

func loadConfig(environment string) configuration {
	var config configuration
	conf, err := yaml.ReadFile("kirk.yml")

	if err != nil {
		log.Fatal(skittles.Red(err))
	}

	serverStr, _ := conf.Get(environment + ".servers")
	config.servers = strings.Split(serverStr, ",")
	config.password, _ = conf.Get(environment + ".password")
	config.username, _ = conf.Get(environment + ".username")
	config.pemfile, _ = conf.Get(environment + ".pemfile")
	taskListStr, _ := conf.Get(environment + ".task_list")
	config.taskList = strings.Split(taskListStr, ",")

	return config
}
