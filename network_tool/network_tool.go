// We handle network related problems here
package network_tool

import (
	"io"
	"net"
)

func copy_IO(src, dest net.Conn) {
	defer src.Close()
	defer dest.Close()
	io.Copy(src, dest)
}

func IP_port_forward(from_ip_port string, to_ip_port string) {
	// port_bridge
	// from_ip_port: 127.0.0.1:80
	// to_ip_port: 127.0.0.1:8080
	ln, err := net.Listen("tcp", from_ip_port)
	if err != nil {
		panic(err)
	}

	var request_function = func(conn net.Conn) {
		proxy, err := net.Dial("tcp", to_ip_port)
		if err != nil {
			panic(err)
		}

		go copy_IO(conn, proxy)
		go copy_IO(proxy, conn)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		go request_function(conn)
	}
}
