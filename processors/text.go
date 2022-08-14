package processors

import (
	"log"
	"lowlevelserver/utils"
	"net"
	"strconv"
	"time"
)

func ProcessText(conn net.Conn, workerId int, path string, contentType string, sendBody bool) {
	content, err := utils.ReadTextContent(path)
	if err != nil {
		log.Println("Worker ", workerId, " ", err)
	}

	headers := map[string]string{
		"Content-Length: ": strconv.Itoa(len(content)),
		"Content-Type: ":   contentType,
		"Server: ":         strconv.Itoa(workerId),
		"Date: ":           time.Now().String(),
		"Connection: ":     "close",
	}

	utils.Response200Text(conn, content, headers, sendBody)

}
