package stl

import (
	"fmt"
	"os"
	"testing"
)

func TestExample(t *testing.T) {
	aFile, err := os.Open("../examples/cube_binary.stl")
	if err != nil {
		fmt.Println(err)
	}

	Parse(aFile)
}
