package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	content := "Hello from Go!"
	file, err := os.Create("./file.txt")
	checkError(err)
	length, err := io.WriteString(file, content)
	checkError(err)
	fmt.Println("Wrote file with %v characters\n", length)
	defer file.Close()

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}

}
