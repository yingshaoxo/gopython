// made by baidu ai

/*
//package main
//
//import (
//	"net/http"
//	"log"
//	"os"
//)
//
//func main() {
//	port := "8080"
//	if len(os.Args) > 1 {
//		port = os.Args[1]
//	}
//
//	dir, err := os.Getwd()
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fs := http.FileServer(http.Dir(dir))
//	http.Handle("/", logRequest(fs))
//
//	log.Printf("Serving %s on HTTP port: %s\n", dir, port)
//	log.Fatal(http.ListenAndServe(":"+port, nil))
//}
//
//func logRequest(handler http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
//		handler.ServeHTTP(w, r)
//	})
//}
*/


/*
You did good, the final thing you need to do for me is:


when user use get request, you print out host, method, url, url_arguments as dict, headers as dict


when you use post to send post request, you print out host, method, url, url_arguments as dict, headers as dict, payload as bytes(we normally could convert that bytes to json_string, then convert to a dict, you can add this part as comments for future look)


host: str
method: str
url: str
url_arguments: dict
headers: dict
payload: bytes
*/

package main

import (
	"net"
	"os"
	"runtime"
	"path/filepath"
	"strings"
	"time"
	"io/ioutil"
	"log"
	"strconv"
)

const (
	CRLF        = "\r\n"
	SERVER_NAME = "GoStaticServer/1.1"
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

	log.Printf("Server started on port %s (%s)\n", port, runtime.GOOS)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Accept error:", err)
			continue
		}

		if runtime.GOOS == "windows" {
			if tcpConn, ok := conn.(*net.TCPConn); ok {
				tcpConn.SetKeepAlive(true)
				tcpConn.SetKeepAlivePeriod(3 * time.Minute)
			}
			go handleConnection(conn)
		} else {
			go unixHandleConn(conn)
		}
	}
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

func handleConnection(conn net.Conn) {
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
		sendErrorResponse(conn, 400, "Bad Request")
		return
	}

	parts := strings.Split(lines[0], " ")
	if len(parts) != 3 {
		sendErrorResponse(conn, 400, "Bad Request")
		return
	}

	method, path := parts[0], parts[1]
	if method != "GET" {
		sendErrorResponse(conn, 501, "Not Implemented")
		return
	}

	serveFile(conn, path)
}

func serveFile(conn net.Conn, path string) {
	wd, _ := os.Getwd()
	fullPath := filepath.Join(wd, filepath.Clean(path))

	if !strings.HasPrefix(fullPath, wd) {
		sendErrorResponse(conn, 403, "Forbidden")
		return
	}

	info, err := os.Stat(fullPath)
	if err != nil {
		sendErrorResponse(conn, 404, "Not Found")
		return
	}

	if info.IsDir() {
		handleDirectory(conn, fullPath)
	} else {
		handleRegularFile(conn, fullPath)
	}
}

func handleDirectory(conn net.Conn, path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		sendErrorResponse(conn, 500, "Internal Error")
		return
	}

	response := "HTTP/1.1 200 OK" + CRLF
	response += "Server: " + SERVER_NAME + CRLF
	response += "Date: " + time.Now().Format(time.RFC1123) + CRLF
	response += "Content-Type: text/html; charset=utf-8" + CRLF + CRLF

	response += "<html><head><title>Index of " + path + "</title></head><body>"
	response += "<h1>Index of " + path + "</h1><hr><pre>"

	for _, file := range files {
		name := file.Name()
		if file.IsDir() {
			name += "/"
		}
		response += "<a href=\"" + name + "\">" + name + "</a>" + CRLF
	}

	response += "</pre><hr></body></html>"
	conn.Write([]byte(response))
}

func handleRegularFile(conn net.Conn, path string) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		sendErrorResponse(conn, 500, "Internal Error")
		return
	}

	response := "HTTP/1.1 200 OK" + CRLF
	response += "Server: " + SERVER_NAME + CRLF
	response += "Date: " + time.Now().Format(time.RFC1123) + CRLF
	response += "Content-Length: " + strconv.Itoa(len(content)) + CRLF
	response += "Content-Type: " + getContentType(path) + CRLF + CRLF

	conn.Write([]byte(response))
	conn.Write(content)
}

func getContentType(path string) string {
	ext := strings.ToLower(filepath.Ext(path))
	switch ext {
	case ".html", ".htm":
		return "text/html"
	case ".css":
		return "text/css"
	case ".js":
		return "application/javascript"
	default:
		return "application/octet-stream"
	}
}

func sendErrorResponse(conn net.Conn, code int, message string) {
	status := strconv.Itoa(code) + " " + message
	response := "HTTP/1.1 " + status + CRLF
	response += "Server: " + SERVER_NAME + CRLF
	response += "Date: " + time.Now().Format(time.RFC1123) + CRLF
	response += "Content-Type: text/plain" + CRLF + CRLF
	response += status

	conn.Write([]byte(response))
}
