package tls_socket

import (
	"encoding/gob"
	"log"
	"net"
)

func CreateConnection(conn net.Conn, handle func(TlsMessage, string), id string) *TlsConnection {
	//defer conn.Close()
	return &TlsConnection{
		id:            id,
		conn:          conn,
		handleReceive: handle,
	}
}

func (o *TlsConnection) receiveMessage() {
	for {
		if o.conn == nil {
			log.Println("no esta conectado")
			return
		}
		var msg TlsMessage
		err := gob.NewDecoder(o.conn).Decode(&msg)
		if err != nil {
			log.Printf("error en recibido: %v", err)
			return
		}
		o.handleReceive(msg, o.id)
	}
}

func (o *TlsConnection) SendMessage(message TlsMessage) {
	if o.conn == nil {
		log.Println("no esta conectado")
		return
	}
	gob.NewEncoder(o.conn).Encode(message)
}
