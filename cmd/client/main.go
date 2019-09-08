package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"go-smbshell/agent"
	"go-smbshell/config"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println(fmt.Sprintf("Usage: %s <IP> <Key>", os.Args[0]))
		os.Exit(1)
	}
	target := os.Args[1]
	config.Key = os.Args[2]
	if err := agent.Connect(os.Args[1]); err != nil {
		log.Panic(err)
	}
	log.Println(fmt.Sprintf("Connected to %s", target))
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("go-smbshell> ")
		userInput, _ := reader.ReadString('\n')
		userInput = strings.TrimSpace(strings.Replace(userInput, "\n", "", -1))
		if len(userInput) <= 1 {
			continue
		}
		if userInput == "exit" {
			os.Exit(1)
		}
		agent.Command(userInput)
	}
}
