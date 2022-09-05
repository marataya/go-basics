package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Value 1: ")
	input, _ := reader.ReadString('\n')
	float1, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
		panic("Value 1 must be a number")
	}
	fmt.Print("Value 2: ")
	input, _ = reader.ReadString('\n')
	float2, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
		panic("Value 2 must be a number")
	}
	fmt.Print("Select and operation: ")
	input, _ = reader.ReadString('\n')
	operation := strings.TrimSpace(input)
	result := 0.0
	switch operation {
	case "*":
		result = float1 * float2
	case "/":
		result = float1 / float2
	case "+":
		result = float1 + float2
	case "-":
		result = float1 - float2
	default:
		panic("Invalid operation")
	}

	result = math.Round(result*100) / 100
	fmt.Println("Result", result)

	fmt.Println("3 sec pause")
	time.Sleep(time.Second * 3)
}
