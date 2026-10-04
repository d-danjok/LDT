package common

import (
	"bufio"
	"os"
	"strings"
)

func ReadLineWithSpaces() string {
	reader := bufio.NewReader(os.Stdin)
	inputLine, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}

	inputLine = inputLine[:len(inputLine)-1]
	inputLine = strings.Split(inputLine, "\r")[0] //for windows format new line

	return inputLine
}
