package main

import (
	"fmt"
	"log"
	"os"

	"go-smbshell/config"
	"go-smbshell/listener"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(fmt.Sprintf("Usage: %s <IP> <Key>", os.Args[0]))
		os.Exit(1)
	}
	config.Key = os.Args[1]
	log.Println("Starting the listener...")
	if err := listener.Start(); err != nil {
		log.Panic(err)
	}
}
