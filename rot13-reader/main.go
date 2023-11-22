package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

// Read implements instance of rot13Reader, which returns its own reader.
func (r rot13Reader) Read(p []byte) (n int, err error) {
	// get reader from referenced rot13Reader
	n, err = r.r.Read(p)
	for i := 0; i < n; i++ {
		p[i] = translateCharacterUsingRot13(p[i])
	}
	return n, err
}

func translateCharacterUsingRot13(in byte) byte {
	out := in
	switch {
	// if char is > a and < z (i.e we're lowercase)
	case 'a' <= in && in <= 'z':
		// find the 13th next character in the range (mod 26 so that we're still in alphanumeric range)
		out = (in-'a'+13)%26 + 'a'
		// if char is > A and < Z (i.e we're upperxase)
	case 'A' <= in && in <= 'Z':
		out = (in-'A'+13)%26 + 'A'
	default:
	}
	return out
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
