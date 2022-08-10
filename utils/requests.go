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
	if len(splitted) < 3 {
		return "", ""
	}
	//Есть пробелы -> в splitted 3 элемента, значит пробелы заменились на %20
	if len(splitted) == 3 {
		return splitted[0], strings.ReplaceAll(splitted[1], "%20", " ")
	}
	path := strings.Join(splitted[1:len(splitted)-1], " ")
	return splitted[0], path
}
