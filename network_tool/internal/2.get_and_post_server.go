// made by baidu ai
package main

import (
	"bytes"
	//"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const (
	CRLF        = "\r\n"
	SERVER_NAME = "DebugServer/1.0"
)

func main() {
	port := "8080"
	if len(os.Args) > 1 {
		port = os.Args[1]
	}

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal("Listen error: ", err)
	}
	defer listener.Close()

	log.Printf("Debug server started on port %s (%s)\n", port, runtime.GOOS)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Accept error:", err)
			continue
		}

		if runtime.GOOS == "windows" {
			go handleConnection(conn)
		} else {
			go unixHandleConn(conn)
		}
	}
}

func printRequestInfo(conn net.Conn, method, rawURL string, headers map[string]string, body []byte) {
	u, err := url.Parse(rawURL)
	if err != nil {
		log.Println("URL parse error:", err)
		return
	}

	query, _ := url.ParseQuery(u.RawQuery)

	fmt.Println("\n=== Request Details ===")
	fmt.Printf("host: %s\n", u.Host)
	fmt.Printf("method: %s\n", method)
	fmt.Printf("url: %s\n", rawURL)
	fmt.Printf("url_arguments: %v\n", query)
	fmt.Printf("headers: %v\n", headers)

	if method == "POST" && len(body) > 0 {
		fmt.Printf("payload: %v\n", body)
		fmt.Println("// To convert payload:")
		fmt.Println("// 1. bytes -> string: string(body)")
		fmt.Println("// 2. string -> dict: json.Unmarshal([]byte(jsonStr), &result)")
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	handleRequest(conn)
}

func unixHandleConn(conn net.Conn) {
	defer conn.Close()
	if tcpConn, ok := conn.(*net.TCPConn); ok {
		file, err := tcpConn.File()
		if err != nil {
			log.Println("File error:", err)
			return
		}
		defer file.Close()
		handleFileConnection(file)
	}
}

func handleFileConnection(file *os.File) {
	conn, err := net.FileConn(file)
	if err != nil {
		log.Println("FileConn error:", err)
		return
	}
	defer conn.Close()
	handleRequest(conn)
}

func handleRequest(conn net.Conn) {
	conn.SetDeadline(time.Now().Add(30 * time.Second))

	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		log.Println("Read error:", err)
		return
	}

	request := string(buf[:n])
	lines := strings.Split(request, CRLF)
	if len(lines) < 1 {
		sendError(conn, 400, "Bad Request")
		return
	}

	parts := strings.Split(lines[0], " ")
	if len(parts) < 3 {
		sendError(conn, 400, "Bad Request")
		return
	}
	method, path := parts[0], parts[1]

	headers := make(map[string]string)
	bodyStart := 0
	for i, line := range lines[1:] {
		if line == "" {
			bodyStart = i + 2
			break
		}
		headerParts := strings.SplitN(line, ":", 2)
		if len(headerParts) == 2 {
			headers[strings.TrimSpace(headerParts[0])] = strings.TrimSpace(headerParts[1])
		}
	}

	var body []byte
	if method == "POST" && bodyStart > 0 && bodyStart < len(lines) {
		body = []byte(strings.Join(lines[bodyStart:], CRLF))
	}

	printRequestInfo(conn, method, path, headers, body)

	if method == "GET" {
		serveFile(conn, path)
	} else {
		sendResponse(conn, 200, "OK", []byte("Request logged"))
	}
}

func serveFile(conn net.Conn, path string) {
	if path == "/" {
		path = "/index.html"
	}

	filePath := filepath.Join(".", path)
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		sendError(conn, 404, "Not Found")
		return
	}

	contentType := "text/plain"
	if strings.HasSuffix(path, ".html") {
		contentType = "text/html"
	} else if strings.HasSuffix(path, ".css") {
		contentType = "text/css"
	} else if strings.HasSuffix(path, ".js") {
		contentType = "application/javascript"
	}

	headers := map[string]string{
		"Content-Type":   contentType,
		"Content-Length": strconv.Itoa(len(data)),
		"Server":         SERVER_NAME,
	}

	sendResponse(conn, 200, "OK", data, headers)
}

func sendError(conn net.Conn, code int, message string) {
	body := fmt.Sprintf("%d %s", code, message)
	headers := map[string]string{
		"Content-Type":   "text/plain",
		"Content-Length": strconv.Itoa(len(body)),
		"Server":         SERVER_NAME,
	}
	sendResponse(conn, code, message, []byte(body), headers)
}

func sendResponse(conn net.Conn, code int, status string, body []byte, customHeaders ...map[string]string) {
	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("HTTP/1.1 %d %s%s", code, status, CRLF))

	headers := map[string]string{
		"Content-Length": strconv.Itoa(len(body)),
		"Server":         SERVER_NAME,
	}

	if len(customHeaders) > 0 {
		for k, v := range customHeaders[0] {
			headers[k] = v
		}
	}

	for k, v := range headers {
		buffer.WriteString(fmt.Sprintf("%s: %s%s", k, v, CRLF))
	}

	buffer.WriteString(CRLF)
	buffer.Write(body)

	conn.Write(buffer.Bytes())
}
