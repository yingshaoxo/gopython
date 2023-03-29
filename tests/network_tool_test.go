package main

import (
	"testing"

	"github.com/yingshaoxo/gopython/network_tool"
)

func Test_ip_port_forward(t *testing.T) {
	network_tool.IP_port_forward("127.0.0.1:9999", "127.0.0.1:5551")
}
