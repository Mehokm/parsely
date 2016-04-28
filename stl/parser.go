package stl

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"io"
	"math"
	"strconv"
	"strings"
)

var asciiBytes = []byte("solid")

func Parse(file io.Reader) *Solid {
	br := bufio.NewReader(file)

	testBytes, _ := br.Peek(5)
	if bytes.Equal(testBytes, asciiBytes) {
		return parseASCII(br)
	} else {
		return parseBinary(br)
	}
}

func parseASCII(r io.Reader) *Solid {
	scanner := bufio.NewScanner(r)

	s := &Solid{}
	s.Facets = make([]Facet, 0)

	for scanner.Scan() {
		tokens := getTokensFromString(scanner.Text())

		if tokens[0] == "solid" && len(tokens) > 1 {
			s.Name = tokens[1]
		}

		if len(tokens) == 5 && tokens[0] == "facet" && tokens[1] == "normal" {
			s.Facets = append(s.Facets, Facet{Normal: getVec3FromStringSlice(tokens[2:])})
		}

		if tokens[0] == "outer" && tokens[1] == "loop" {
			var vs [3]Vec3

			for i := 0; i < 3; i++ {
				scanner.Scan()
				tokens = getTokensFromString(scanner.Text())

				if len(tokens) == 4 && tokens[0] == "vertex" {
					vs[i] = getVec3FromStringSlice(tokens[1:])
				}
			}
			s.Facets[len(s.Facets)-1].Vertices = vs
		}
	}

	return s
}

func getTokensFromString(s string) []string {
	return strings.Fields(s)
}

func getVec3FromStringSlice(ss []string) Vec3 {
	var v Vec3
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

		v.X = x
		v.Y = y
		v.Z = z
	}

	return v
}

func parseBinary(r io.Reader) *Solid {
	s := &Solid{}
	binaryHeader := make([]byte, 84)

	r.Read(binaryHeader)

	s.Name = strings.TrimSpace(string(binaryHeader[:80]))

	facets := make([]Facet, binary.LittleEndian.Uint32(binaryHeader[80:]))
	for i := 0; i < len(facets); i++ {
		binaryFacet := make([]byte, 50)

		r.Read(binaryFacet)

		normal := getVec3FromByteSlice(binaryFacet, 0)
		vs := [3]Vec3{
			getVec3FromByteSlice(binaryFacet, 1),
			getVec3FromByteSlice(binaryFacet, 2),
			getVec3FromByteSlice(binaryFacet, 3),
		}

		facets[i] = Facet{normal, vs}
	}
	s.Facets = facets

	return s
}

func uint32ToFloat32(u []byte) float32 {
	return math.Float32frombits(binary.LittleEndian.Uint32(u))
}

func getVec3FromByteSlice(b []byte, offset int) Vec3 {
	start := 12 * offset
	return Vec3{
		float64(uint32ToFloat32(b[start : start+4])),
		float64(uint32ToFloat32(b[start : start+8])),
		float64(uint32ToFloat32(b[start : start+12])),
	}
}
