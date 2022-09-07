package main

import (
	"fmt"
	"math"
)

func main() {
	pow := make([]int, 64)
	for i := range pow {
		pow[i] = 1 << uint64(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
	fmt.Println(uint64(math.Pow(2, 64) - 1))
}
