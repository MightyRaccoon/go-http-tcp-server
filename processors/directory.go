package processors

import (
	"log"
	"lowlevelserver/utils"
	"net"
	"strconv"
	"time"
)

func ProcesssDirectory(conn net.Conn, workerId int, path string, sendBody bool) {
	fullPath := path + "/index.html"

	//Если файла не существует, то отдает 403
	if !utils.CheckFileExists(fullPath) {
		headers := map[string]string{
			"Server: ":     strconv.Itoa(workerId),
			"Date: ":       time.Now().String(),
			"Connection: ": "close",
		}
		utils.Response403(conn, headers)
		return
	}

	// Если файл существует, то пытаемся обработать
	body, err := utils.ReadTextContent(fullPath)
	if err != nil {
		// Вообще, технически тут должна быть какая-нибудь 5xx
		log.Println("Worker ", workerId, "Can't read file: ", err)
	}
	ProcessText(conn, workerId, body, "text/html", true)
}
