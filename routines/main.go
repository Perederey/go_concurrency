package main

import (
	"fmt"
	"runtime"
	"sync"
)

// goroutines
// wg
//

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		lowerCase()
	}()

	go func() {
		defer wg.Done()
		upperCase()
	}()

	wg.Wait()

	fmt.Println("Exit")

}

func lowerCase() {
	for i := 0; i < 3; i++ {
		for char := 'a'; char < 'a'+26; char++ {
			fmt.Printf("%c ", char)
		}
	}
}

func upperCase() {
	for i := 0; i < 3; i++ {
		for char := 'A'; char < 'A'+26; char++ {
			fmt.Printf("%c ", char)
		}
	}
}
