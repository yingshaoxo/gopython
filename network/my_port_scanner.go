// We use this package to do port scanning for a host or many hosts
package port_scanner

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"time"
)

var timeout time.Duration = time.Duration(1000 * 1000 * 500) // 0.5 second

// We can set timeout for each scanning
func SetTimeOut(millisecond int) {
	timeout = time.Duration(1000 * 1000 * millisecond)
}

func worker(address chan string, results chan string) {
	for uri := range address {
		connection, err := net.DialTimeout("tcp", uri, timeout)
		//fmt.Println(err)
		if err != nil {
			results <- ""
			continue
		}
		connection.Close()
		//fmt.Println(uri)
		results <- uri
	}
}

func scan_ports(hosts []string, startPort int, endPort int) []string {
	//1-65535
	urls := make([]string, 0)

	//address := make(chan string, 10000)
	address := make(chan string, 65535)
	results := make(chan string)

	for i := 0; i < cap(address); i++ {
		go worker(address, results) // now we have 10000 workers
	}

	/*
		for _, host := range hosts {
			go func(host_ string) {
				for i := startPort; i <= endPort; i++ {
					address <- fmt.Sprintf("%s:%d", host_, i)
				}
			}(host)
		}
	*/
	go func() {
		for _, host := range hosts {
			for i := startPort; i <= endPort; i++ {
				address <- fmt.Sprintf("%s:%d", host, i)
			}
		}
	}()

	for _, _ = range hosts {
		for i := startPort; i <= endPort; i++ {
			uri := <-results
			if uri != "" {
				urls = append(urls, uri)
			}
		}
	}

	close(address)
	close(results)

	return urls
}

// We can scan ports for a single host
func ScanPorts(host string, startPort int, endPort int) string {
	var hosts = []string{host}
	urls := scan_ports(hosts, startPort, endPort)
	json_result, err := json.Marshal(urls)
	if err != nil {
		return ""
	} else {
		return string(json_result)
	}
}

// Get a range of hosts in a given network
func GetAllHostsOfANetwork(network string) []string {
	hosts := make([]string, 0)

	// convert string to IPNet struct
	_, ipv4Net, err := net.ParseCIDR(network) // 192.168.0.0/16
	if err != nil {
		return hosts
	}

	// convert IPNet struct mask and address to uint32
	// network is BigEndian
	mask := binary.BigEndian.Uint32(ipv4Net.Mask)
	start := binary.BigEndian.Uint32(ipv4Net.IP)

	// find the final address
	finish := (start & mask) | (mask ^ 0xffffffff)

	// loop through addresses as uint32
	for i := start; i <= finish; i++ {
		// convert back to net.IP
		ip := make(net.IP, 4)
		binary.BigEndian.PutUint32(ip, i)
		hosts = append(hosts, ip.String())
	}

	return hosts
}

// We can scan ports for a whole network, such as:
// 	192.168.1.1/24
func ScanPortsOfANetwork(network string, startPort int, endPort int) string {
	hosts := GetAllHostsOfANetwork(network)

	urls := scan_ports(hosts, startPort, endPort)

	json_result, err := json.Marshal(urls)
	if err != nil {
		return ""
	} else {
		return string(json_result)
	}
}

func FakePing(address string) bool {
	connection, err := net.Dial("tcp", address)
	if err != nil {
		return false
	}
	connection.Close()
	return true
}
