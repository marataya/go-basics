package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(mean([]int{1, 2, 3}))
	fmt.Println(mean([]int{1, 2, 3, 4}))
	fmt.Println(median([]float64{1, 2, 3, 4, 57, 57, 1, 4, 5}))
	fmt.Println(frequency(moby))
	ExampleCommaOK()
}

func mean(nums []int) float64 {
	s := sum(nums)
	return float64(s) / float64(len(nums))
}

func sum(nums []int) int {
	total := 0
	for _, x := range nums {
		total += x
	}

	return total
}

func median(nums []float64) float64 {
	vals := make([]float64, len(nums))
	copy(vals, nums)
	sort.Float64s(vals)

	fmt.Println(vals)

	i := len(vals) / 2
	if len(vals)%2 == 1 {
		return vals[i]
	}

	return (vals[i-1] + vals[i]) / 2

}

func frequency(words []string) map[string]int {
	freq := make(map[string]int)
	for _, w := range words {
		freq[w]++
	}

	return freq

}

func ExampleCommaOK() {
	prices := map[string]int{
		"Banana": 0, // Banana's are free!
	}

	price, ok := prices["Banana"]
	if ok {
		fmt.Printf("The price of Banana is $%d\n", price)
	} else {
		fmt.Printf("We don't have Bananas")
	}

	price, ok = prices["Apple"]
	if ok {
		fmt.Printf("The price of Apple is $%d\n", price)
	} else {
		fmt.Printf("We don't have Apples")
	}

	// Output:
	// The price of Banana is $0
	// We don't have Apples
}

var moby = []string{
	"call", "me", "ishmael", "some", "years", "ago", "never", "mind", "how",
	"long", "precisely", "having", "little", "or", "no", "money", "in", "my",
	"purse", "and", "nothing", "particular", "to", "interest", "me", "on",
	"shore", "i", "thought", "i", "would", "sail", "about", "a", "little",
	"and", "see", "the", "watery", "part", "of", "the", "world", "it", "is",
	"a", "way", "i", "have", "of", "driving", "off", "the", "spleen", "and",
	"regulating", "the", "circulation", "whenever", "i", "find", "myself",
	"growing", "grim", "about", "the", "mouth", "whenever", "it", "is", "a",
	"damp", "drizzly", "november", "in", "my", "soul", "whenever", "i", "find",
	"myself", "involuntarily", "pausing", "before", "coffin", "warehouses",
	"and", "bringing", "up", "the", "rear", "of", "every", "funeral", "i",
	"meet", "and", "especially", "whenever", "my", "hypos", "get", "such",
	"an", "upper", "hand", "of", "me", "that", "it", "requires", "a", "strong",
	"moral", "principle", "to", "prevent", "me", "from", "deliberately",
	"stepping", "into", "the", "street", "and", "methodically", "knocking",
	"people", "s", "hats", "off", "then", "i", "account", "it", "high", "time",
	"to", "get", "to", "sea", "as", "soon", "as", "i", "can", "this", "is",
	"my", "substitute", "for", "pistol", "and", "ball", "with", "a",
	"philosophical", "flourish", "cato", "throws", "himself", "upon", "his",
	"sword", "i", "quietly", "take", "to", "the", "ship", "there", "is",
	"nothing", "surprising", "in", "this", "if", "they", "but", "knew", "it",
	"almost", "all", "men", "in", "their", "degree", "some", "time", "or",
	"other", "cherish", "very", "nearly", "the", "same", "feelings", "towards",
	"the", "ocean", "with", "me",
}
