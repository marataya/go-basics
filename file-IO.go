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
	fmt.Printf("Wrote file with %v characters\n", length)
	defer file.Close()
	defer readFile("./file.txt")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func readFile(fileName string) {
	data, err := os.ReadFile(fileName)
	checkError(err)
	fmt.Println("Data read from file: ", string(data))
}
