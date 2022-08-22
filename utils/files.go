package utils

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/pkg/errors"
)

func DefineContentType(path string) (string, string) {
	contentMap := map[string][2]string{
		"html": {"text", "text/html"},
		"css":  {"text", "text/css"},
		"js":   {"application", "application/javascript"},
		"jpg":  {"image", "image/jpeg"},
		"jpeg": {"image", "image/jpeg"},
		"png":  {"image", "image/png"},
		"gif":  {"image", "image/gif"},
		"swf":  {"application", "application/x-shockwave-flash"},
		"txt":  {"text", "text/plain"},
	}
	// Костыль, чтоб понять, что запрос к директории
	if []rune(path)[len([]rune(path))-1] == '/' {
		return "directory", "text/html"
	}
	splitted := strings.Split(path, ".")
	fileExtension := strings.ToLower(splitted[len(splitted)-1])
	return contentMap[fileExtension][0], contentMap[fileExtension][1]
}

func ReadTextContent(path string) (string, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return "", errors.Wrap(err, "ReadTextContent")
	}
	return string(content), nil
}

func ReadByteContent(path string) ([]byte, error) {
	image, err := ioutil.ReadFile(path)
	if err != nil {
		return []byte{}, errors.Wrap(err, "ReadByteContent")
	}
	return image, nil
}

func CheckFileExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
