package mytests

import (
	"log"
	"testing"

	port_scanner "github.com/yingshaoxo/gopython/network"
)

func TestGetPortsFromAHost(t *testing.T) {
	urls := port_scanner.ScanPorts("localhost", 0, 65535)
	log.Println(urls)
	if len(urls) == 0 {
		t.Fatalf("there should has open ports")
	}
}

func TestGetAllHostsFromNetwork(t *testing.T) {
	//hosts := get_all_hosts_of_a_network("192.168.0.0/16")
	hosts := port_scanner.GetAllHostsOfANetwork("192.168.50.0/24")
	log.Println(hosts)
	if len(hosts) == 0 {
		t.Fatalf("we should get a lot of hosts here")
	}
}

func TestGetAllReachablePortsFromNetwork(t *testing.T) {
	urls := port_scanner.ScanPortsOfANetwork("192.168.50.0/24", 5000, 5010)
	log.Println(urls)
	/*
		if len(urls) == 0 {
			t.Fatalf("we should get a lot of hosts here")
		}
	*/
}

func TestGetAllReachablePortsFromNetwork2(t *testing.T) {
	port_scanner.SetScanningTimeOut(100)
	urls := port_scanner.ScanPortsOfANetwork("192.168.50.0/16", 5000, 5010)
	log.Println(urls)
	/*
		if len(urls) < 1 {
			t.Fatalf("we should get a lot of hosts here")
		}
	*/
}

func TestFakePing(t *testing.T) {
	if port_scanner.FakePing("localhost:80", 100) == false {
		t.Fatalf("localhost:80 should have something")
	}
}
