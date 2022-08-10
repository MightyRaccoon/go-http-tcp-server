package processors

import (
	"log"
	"lowlevelserver/utils"
	"net"
	"strconv"
	"time"
)

func ProcesssDirectory(conn net.Conn, workerId int, full_path string, sendBody bool) {
	full_path = full_path + "/index.html"

	headers := map[string]string{
		"Server: ":     strconv.Itoa(workerId),
		"Date: ":       time.Now().String(),
		"Connection: ": "close",
	}

	//Если файла не существует, то отдает 403
	if !utils.CheckFileExists(full_path) {
		utils.Response403(conn, headers)
		return
	}

	// Если файл существует, то пытаемся обработать
	body, err := utils.ReadTextContent(full_path)
	if err != nil {
		// Вообще, технически тут должна быть какая-нибудь 5xx
		log.Println("Worker ", workerId, "Can't read file: ", err)
	}
	ProcessText(conn, workerId, body, "text/html", true)
}
