package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func readerFunc() {
	reader := strings.NewReader("This is a demo reader function")
	p := make([]byte, 4)

	for {
		n, err := reader.Read(p)
		if err != nil {
			if err == io.EOF {
				fmt.Println(string(p[:n])) //should handle any remainding bytes.
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(p[:n]))
	}
}

func writerFunc() {
	proverbs := []string{
		"This is a demo writer function",
		"Go just go",
		"Don't panic",
	}
	var writer bytes.Buffer

	for _, p := range proverbs {
		n, err := writer.Write([]byte(p))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if n != len(p) {
			fmt.Println("failed to write data")
			os.Exit(1)
		}
	}

	fmt.Println(writer.String())
}

func main() {
	readerFunc()
	writerFunc()
}
