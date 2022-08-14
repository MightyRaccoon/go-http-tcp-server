package utils

import (
	"bufio"
	"net"
	"strings"
)

func ReadRequest(conn net.Conn) (string, error) {
	reader := bufio.NewReader(conn)
	buf, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return buf, nil
}

func ParseRequest(request string) (string, string) {
	splitted := strings.Split(request, " ")

	switch {
	case len(splitted) < 3:
		return "", ""
	case len(splitted) == 3: //Есть пробелы -> в splitted 3 элемента, значит пробелы заменились на %20
		return splitted[0], "./data" + strings.ReplaceAll(splitted[1], "%20", " ")
	default:
		return splitted[0], "./data" + strings.Join(splitted[1:len(splitted)-1], " ")
	}
}
