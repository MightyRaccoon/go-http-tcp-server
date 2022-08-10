package processors

import (
	"lowlevelserver/utils"
	"net"
	"strconv"
	"time"
)

func ProcesssNotExistedContent(conn net.Conn, workerId int) {
	headers := map[string]string{
		"Server: ":     strconv.Itoa(workerId),
		"Date: ":       time.Now().String(),
		"Connection: ": "close",
	}
	utils.Response404(conn, headers)
}
