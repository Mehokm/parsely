package stl

import (
	"fmt"
	"os"
	"testing"
)

func TestParseAscii(t *testing.T) {
	aFile, err := os.Open("../examples/cube_ascii.stl")
	if err != nil {
		fmt.Println(err)
	}

	a := Parse(aFile)

	if "MYSOLID" != a.Name {
		t.Errorf("incorrect solid name: %v", a.Name)
	}

	if 12 != len(a.Facets) {
		t.Errorf("incorrect number of facets: %v", len(a.Facets))
	}
}
