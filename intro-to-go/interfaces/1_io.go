package main

import (
	"io"
	"io/ioutil"
	"os"
	"strings"
)
func duplicate(in io.Reader, out io.Writer, n int) {
	data, _ := ioutil.ReadAll(in)

	for i := 0; i < n; i++ {
		out.Write(data)
	}
}

func main() {
	// input is an os.File
	f, _ := os.Open("./helloworld.txt")
	duplicate(f, os.Stdout, 5)

	// input is strings.NewReader
	reader := strings.NewReader("from my reader\n")
	duplicate(reader, os.Stdout, 5)
}
