package processors

import (
	"log"
	"lowlevelserver/utils"
	"net"
	"strconv"
	"time"
)

func ProcesssApplication(conn net.Conn, workerId int, file_path string, contentType string, sendBody bool) {
	log.Println("Worker ", workerId, "Run Application Processor")
	content, err := utils.ReadByteContent(file_path)
	if err != nil {
		log.Println("Worker ", workerId, " ", err)
	}

	headers := map[string]string{
		"Content-Length: ": strconv.Itoa(len(content)),
		"Content-Type: ":   contentType,
		"Server: ":         strconv.Itoa(workerId),
		"Date: ":           time.Now().String(),
		"Connection:":      "close",
	}

	utils.Response200Bytes(conn, content, headers, sendBody)
}
