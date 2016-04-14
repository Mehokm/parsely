package stl

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

var ASCIIBytes = []byte("solid")

func Parse(file io.Reader) {
	testBuff := make([]byte, 5)

	bufio.NewReader(file).Read(testBuff)
	fmt.Println(bytes.Equal(testBuff, ASCIIBytes))
	if bytes.Equal(testBuff, ASCIIBytes) {
		parseASCII(file)
	} else {
		parseBinary(file)
	}

}

func parseASCII(file io.Reader) {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(strings.TrimSpace(scanner.Text()))
	}
}

func parseBinary(file io.Reader) {

}
