package processors

import (
	"lowlevelserver/utils"
	"net"
	"strconv"
	"time"
)

func ProcesssNotSupportedMethod(conn net.Conn, workerId int) {
	headers := map[string]string{
		"Server":     strconv.Itoa(workerId),
		"Date":       time.Now().String(),
		"Connection": "close",
	}
	utils.Response405(conn, headers)
}
