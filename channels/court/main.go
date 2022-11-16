package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	rand.Seed(time.Now().UnixNano())

	court := make(chan int, 1)

	wg.Add(2)

	court <- 1

	go player("Alex", court)
	go player("Ken", court)

	wg.Wait()
}

func player(name string, court chan int) {
	defer wg.Done()
	for {
		ball, ok := <-court
		if !ok {
			fmt.Println("Победил игрок", name)
			return
		}

		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Println("Игрок не ударил по мячу", name)

			close(court)
			return
		}

		fmt.Println("Игрок ударил по мячу", name, ball)

		ball++

		court <- ball
	}
}
