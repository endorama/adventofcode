package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

// func doWork(int i, int n, lines []string) {}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}
func product(array []int) int {
	result := 1
	for _, v := range array {
		result *= v
	}
	return result
}

func main() {
	// reading file content in slice of strings by \n
	content, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(content), "\n")
	log.Printf("File contents: %s\n", lines)

	// Setting up value collector
	collectorCtx, collectorCancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer collectorCancel()
	ch := make(chan int)
	values := []int{}
	// this goroutine intercept values and store for future use
	go func() {
		for {
			select {
			case <-collectorCtx.Done():
				log.Println("collector timeout")
				return
			default:
			}
			found, ok := <-ch
			if !ok {
				log.Fatal("Error receiving from channel")
			}
			values = append(values, found)
		}
	}()

	var wg sync.WaitGroup
	// wg.Add(len(lines))

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Make sure it's called to release resources even if no errors

	for i, l := range lines {
		numberL, _ := strconv.Atoi(l)
		// log.Println(fmt.Sprintf("Starting %d:%d", i, numberL))

		wg.Add(1)
		go func(log *log.Logger, i, l int) {
			// log.Println("starting")
			for _, el := range lines {
				// check if the context has been called somewhere else
				select {
				case <-ctx.Done():
					wg.Done()
					return
				default: // must define to avoid blocking
				}
		
				// we skip ourselves
				// if i == j {
				//   continue
				// }
				// log.Println(fmt.Sprintf("%d %s", l, el))
				numberEl, _ := strconv.Atoi(el)
				val := l + numberEl
				// fmt.Println(val)
				if val == 2020 {
					ch <- l
					ch <- numberEl
					log.Println("EUREKA!")
					wg.Done()
					cancel()
					return
				}
			}
			wg.Done()
			return
		}(log.New(os.Stdout, fmt.Sprintf("%d-%s ", i, l), log.Lshortfile), i, numberL)
	}
	wg.Wait()
	valueSum := sum(values)
	valueProduct := product(values)
	fmt.Print(fmt.Sprintf("values: %v\t%d\t%d", values, valueSum, valueProduct))
}
