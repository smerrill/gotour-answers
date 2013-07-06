package main

import (
	"io"
	"os"
	"strings"

//    "fmt"
)

type rot13Reader struct {
	r io.Reader
}

func (x rot13Reader) Read(p []byte) (int, error) {
	bytes, err := x.r.Read(p)
	for key, val := range p {
		p[key] = x.rot13(val)
	}

	return bytes, err
}

func (x rot13Reader) rot13(b byte) byte {
	if b >= 65 && b <= 90 {
		// Process capital letters.
		if b+13 > 90 {
			return b - 13
		}
		return b + 13
	} else if b >= 97 && b <= 122 {
		// Process lowercase letters.
		if b+13 > 122 {
			return b - 13
		}
		return b + 13
	}
	return b
}

func main() {
	s := strings.NewReader(
		"Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
