package main

import (
	"bufio"
	"bytes"
	"code.google.com/p/go.crypto/ssh"
	"fmt"
	"github.com/acmacalister/skittles"
	"log"
	"os"
	"time"
)

func runTaskList(config *configuration) {
	sshConfig := &ssh.ClientConfig{
		User: config.username,
		Auth: []ssh.AuthMethod{
			ssh.Password(config.password),
		},
	}

	c := make(chan int)
	for _, server := range config.servers {
		client, err := ssh.Dial("tcp", server, sshConfig)
		if err != nil {
			log.Fatal(skittles.Red("Failed to dial: " + err.Error()))
		}

		go executeTaskList(client, config.taskList, server, c)
	}

	for i, _ := range config.servers {
		fmt.Sprintln("%d, %d", <-c, i)
	}
}

func executeTaskList(client *ssh.Client, taskList []string, server string, c chan int) {
	for _, task := range taskList {
		file, err := os.Open(task)
		if err != nil {
			log.Fatal(skittles.Red(err))
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			session, err := client.NewSession()
			if err != nil {
				log.Fatal(skittles.Red("Failed to create session: " + err.Error()))
			}
			var b bytes.Buffer
			session.Stdout = &b
			if err := session.Run(scanner.Text()); err != nil {
				log.Fatal(skittles.Red(err))
			}

			fmt.Printf(b.String())
			session.Close()
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(skittles.Red(err))
		}
		c <- 0
	}
}
