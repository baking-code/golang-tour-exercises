package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// Read implements io.Reader.
func (MyReader) Read(p []byte) (n int, err error) {
	for i := 0; i < len(p); i++ {
		// write each byte for the input slice as 'A'
		p[i] = 'A'
	}
	// return number of bytes written, plus no error
	return len(p), nil
}

func main() {
	reader.Validate(MyReader{})
}
