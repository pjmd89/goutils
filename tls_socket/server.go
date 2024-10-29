package tls_socket

import (
	"crypto/tls"
	"fmt"
	"log"

	"github.com/google/uuid"
)

func (o *TlsSocketServer) Start() {
	// Start worker socket
	listenTo := ""
	port := 1029
	protocol := "tcp"

	if o.ListenTo != "" {
		listenTo = o.ListenTo
	}

	if o.Port != 0 {
		port = o.Port
	}
	address := fmt.Sprintf("%s:%d", listenTo, port)

	cert, err := o.get_certificate()
	if err != nil {
		log.Println(err)
		return
	}

	if o.Protocol != "" {
		protocol = o.Protocol
	}
	ln, err := tls.Listen(protocol, address, cert)

	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		conn, err := ln.Accept()
		defer conn.Close()
		id := uuid.New()
		o.CurrentConections = make(map[string]*TlsConnection)
		client := CreateConnection(conn, o.HandleReceiveMessage, id.String())
		o.CurrentConections[id.String()] = client

		if err != nil {
			log.Println(err)
			continue
		}
		message := TlsMessage{Message: "connected!"}
		o.CurrentConections[id.String()].SendMessage(message)
		go o.CurrentConections[id.String()].receiveMessage()
	}
}
