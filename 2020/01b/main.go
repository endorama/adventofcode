package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// https://gist.github.com/manorie/20874a3c59e9fdfb4e184cac4130944d
func HeapPermutation(a []int, size int) {
	if size == 1 {
		fmt.Println(a)
	}

	for i := 0; i < size; i++ {
		HeapPermutation(a, size-1)

		if size%2 == 1 {
			a[0], a[size-1] = a[size-1], a[0]
		} else {
			a[i], a[size-1] = a[size-1], a[i]
		}
	}
}

func main() {
	// reading file content in slice of strings by \n
	content, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(content), "\n")
	log.Printf("File contents: %s\n", lines)

	numbers := []int{}
	for _, l := range lines {
		num, _ := strconv.Atoi(l)
		numbers = append(numbers, num)
	}

	HeapPermutation(numbers, len(numbers))
}
