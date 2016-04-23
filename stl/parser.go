package stl

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"
)

var ASCIIBytes = []byte("solid")

func Parse(file io.Reader) {
	br := bufio.NewReader(file)

	testBytes, _ := br.Peek(5)
	if bytes.Equal(testBytes, ASCIIBytes) {
		fmt.Println(parseASCII(br))
	} else {
		parseBinary(br)
	}

}

func parseASCII(r io.Reader) *StlObject {
	scanner := bufio.NewScanner(r)

	so := &StlObject{}
	so.Facets = make([]Facet, 0)

	for scanner.Scan() {
		tokens := getTokensFromString(scanner.Text())

		if tokens[0] == "solid" && len(tokens) > 1 {
			so.Name = tokens[1]
		}

		if tokens[0] == "facet" && tokens[1] == "normal" {
			so.Facets = append(so.Facets, Facet{Normal: getVec3FromStringSlice(tokens[2:])})
		}

		var vs [3]Vec3
		if tokens[0] == "outer" && tokens[1] == "loop" {
			for i := 0; i < 3; i++ {
				scanner.Scan()
				tokens = getTokensFromString(scanner.Text())

				if tokens[0] == "vertex" {
					vs[i] = getVec3FromStringSlice(tokens[1:])
				}
			}
			so.Facets[len(so.Facets)-1].Vertices = vs
		}
	}

	return so
}

func getTokensFromString(s string) []string {
	return strings.Split(strings.TrimSpace(s), " ")
}

func getVec3FromStringSlice(ss []string) Vec3 {
	var x, y, z float64
	if len(ss) == 3 {
		if a, err := strconv.ParseFloat(ss[0], 64); err == nil {
			x = a
		}
		if b, err := strconv.ParseFloat(ss[1], 64); err == nil {
			y = b
		}
		if c, err := strconv.ParseFloat(ss[2], 64); err == nil {
			z = c
		}

		return Vec3{x, y, z}
	}

	return Vec3{}
}

func parseBinary(file io.Reader) {

}
