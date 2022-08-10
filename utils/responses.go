package utils

import (
	"fmt"
	"log"
	"net"
)

// *net.Conn -> net.Conn
func AddHeader(conn net.Conn, headers map[string]string) {
	for header, value := range headers {
		fmt.Fprint(conn, header, value, "\r\n")
	}
	fmt.Fprint(conn, "\r\n")
}

func Response404(conn net.Conn, headers map[string]string) {
	fmt.Fprint(conn, "HTTP/1.1 404 Not Found\r\n")
	AddHeader(conn, headers)

}

func Response403(conn net.Conn, headers map[string]string) {
	fmt.Fprint(conn, "HTTP/1.1 403 Forbidden\r\n")
	AddHeader(conn, headers)

}

func Response200Text(conn net.Conn, content string, headers map[string]string, sendBody bool) {
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	AddHeader(conn, headers)
	if sendBody {
		fmt.Fprint(conn, content)
	}
}

func Response200Bytes(conn net.Conn, content []byte, headers map[string]string, sendBody bool) {
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	AddHeader(conn, headers)
	if sendBody {
		_, err := conn.Write(content)
		if err != nil {
			log.Println(err)
		}
	}
}

func Response405(conn net.Conn, headers map[string]string) {
	fmt.Fprint(conn, "HTTP/1.1 405 Method Not Allowed\r\n")
	AddHeader(conn, headers)
	fmt.Fprint(conn, "\r\n")
}
