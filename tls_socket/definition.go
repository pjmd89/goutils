package tls_socket

import (
	"net"
)

type TlsMessage struct {
	Message string
}
type TlsSocket struct {
	Port                 int
	Protocol             string
	UseBroadcast         bool
	HandleReceiveMessage func(TlsMessage, string)
}
type TlsConnection struct {
	id            string
	conn          net.Conn
	handleReceive func(TlsMessage, string)
}
type TlsSocketServer struct {
	TlsSocket
	CommonName         string
	Country            string
	Organization       string
	OrganizationalUnit string
	ListenTo           string
	CurrentConections  map[string]*TlsConnection
}

type TlsSocketClient struct {
	TlsSocket
	TlsConnection
	ServerIP string
}
