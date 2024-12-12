package main

import (
	"fmt"
	"io"
)

type MyReader struct{}

func (r MyReader) Read(p []byte) (n int, err error) {

	for i := range p {
		p[i] = 'A'
	}
	return len(p), nil
}
func main() {
	//reader.Validate(MyReader{})
	r := MyReader{}

	buf := make([]byte, 10)

	for i := 0; i < 5; i++ {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println("Error:", err)
			break
		}

		fmt.Printf("%s", buf[:n])
	}
}
