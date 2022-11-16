package main

import (
	"fmt"
	"log"
	"sync"
)

var counter int

func main() {
	n := 1000
	var wg sync.WaitGroup
	wg.Add(n)

	for g := 0; g < n; g++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 2; i++ {
				value := counter
				value++
				log.Println("log....")
				counter = value

			}
		}()
	}

	wg.Wait()

	fmt.Println(counter)
}
