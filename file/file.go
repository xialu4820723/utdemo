package file

import (
	"io/ioutil"
	"strings"
)

func loadLines(filePath string) []string {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return []string{}
	}
	content := string(data)
	return strings.Split(content, "\n")
}
