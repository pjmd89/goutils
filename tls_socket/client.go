package tls_socket

import (
	"crypto/tls"
	"fmt"
	"log"
)

func (o *TlsSocketClient) Start() {
	serverIP := ""
	port := 1029
	protocol := "tcp"
	config := &tls.Config{InsecureSkipVerify: true}
	var err error

	if o.ServerIP != "" {
		serverIP = o.ServerIP
	}

	if o.Port != 0 {
		port = o.Port
	}
	address := fmt.Sprintf("%s:%d", serverIP, port)

	if o.Protocol != "" {
		protocol = o.Protocol
	}
	conn, err := tls.Dial(protocol, address, config)
	o.conn = conn
	o.handleReceive = o.HandleReceiveMessage
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
	message := TlsMessage{Message: "Hello!"}
	o.SendMessage(message)
	o.receiveMessage()

}
