package utils

import (
	"io/ioutil"
	"os"
	"strings"
)

func DefineContentType(path string) string {
	content_map := map[string]string{
		"html": "text/html",
		"css":  "text/css",
		"js":   "application/javascript",
		"jpg":  "image/jpeg",
		"jpeg": "image/jpeg",
		"png":  "image/png",
		"gif":  "image/gif",
		"swf":  "application/x-shockwave-flash",
		"txt":  "text/plain",
	}
	// Костыль, чтоб понять, что запрос к директории
	if []rune(path)[len([]rune(path))-1] == '/' {
		return "directory"
	}
	splitted := strings.Split(path, ".")
	file_extension := strings.ToLower(splitted[len(splitted)-1])
	return content_map[file_extension]
}

func ReadTextContent(path string) (data string, err error) {

	content, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func ReadByteContent(path string) ([]byte, error) {
	image, err := ioutil.ReadFile(path)
	if err != nil {
		return []byte{}, err
	}

	return image, err

}

func CheckFileExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
