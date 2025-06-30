// made by baidu ai
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	//"net/url"
	//"strconv"
	//"strings"
	"time"
)

type HttpClient struct {
	host       string
	port       int
	conn       net.Conn
	timeout    time.Duration
	keepAlive  bool
}

func NewClient(host string, port int, timeoutSec int) (*HttpClient, error) {
	client := &HttpClient{
		host:      host,
		port:      port,
		timeout:   time.Duration(timeoutSec) * time.Second,
		keepAlive: true,
	}

	err := client.connect()
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (c *HttpClient) connect() error {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", c.host, c.port), c.timeout)
	if err != nil {
		return err
	}
	c.conn = conn
	return nil
}

func (c *HttpClient) Request(method, path string, headers map[string]string, payload interface{}) (string, error) {
	if c.conn == nil {
		if err := c.connect(); err != nil {
			return "", err
		}
	}

	var body []byte
	if payload != nil {
		var err error
		body, err = json.Marshal(payload)
		if err != nil {
			return "", err
		}
	}

	req := bytes.NewBufferString(fmt.Sprintf("%s %s HTTP/1.1\r\n", method, path))
	req.WriteString(fmt.Sprintf("Host: %s\r\n", c.host))
	req.WriteString(fmt.Sprintf("Connection: %s\r\n", map[bool]string{true: "keep-alive", false: "close"}[c.keepAlive]))

	for k, v := range headers {
		req.WriteString(fmt.Sprintf("%s: %s\r\n", k, v))
	}

	if body != nil {
		req.WriteString(fmt.Sprintf("Content-Length: %d\r\n", len(body)))
		req.WriteString("Content-Type: application/json\r\n")
	}

	req.WriteString("\r\n")
	if body != nil {
		req.Write(body)
	}

	_, err := c.conn.Write(req.Bytes())
	if err != nil {
		return "", err
	}

	resp, err := ioutil.ReadAll(c.conn)
	if err != nil {
		return "", err
	}

	if !c.keepAlive {
		c.conn.Close()
		c.conn = nil
	}

	return string(resp), nil
}

func (c *HttpClient) Close() {
	if c.conn != nil {
		c.conn.Close()
	}
}

func main() {
    // python3 -m http.server 9999
	client, err := NewClient("127.0.0.1", 9999, 10)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// GET请求示例
	resp, err := client.Request("GET", "/", map[string]string{
		"User-Agent": "GoSocketClient/1.0",
	}, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("GET Response:\n", resp)

	// POST请求示例
	resp, err = client.Request("POST", "/post", map[string]string{
		"Accept": "application/json",
	}, map[string]interface{}{
		"name":  "John Doe",
		"email": "john@example.com",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("POST Response:\n", resp)
}
