package main

import (
	"log"
	"sync"
)

func main() {
	err := worker()
	if err != nil {
		log.Println(err)
	}
}

func worker() error {
	const nGoroutines = 2
	var wg sync.WaitGroup
	wg.Add(nGoroutines)
	errCh := make(chan error, nGoroutines)
	vendorDto := struct{}{}
	go func() {
		defer wg.Done()
		errCh <- makeRequest(ctx, "/path", &vendorDto)
	}()

	vendorConfigDto := struct{}{}
	go func() {
		defer wg.Done()
		errCh <- makeRequest(ctx, "/another_url", &vendorConfigDto)
	}()

	wg.Wait()
	close(errCh)
	for err := range errCh {
		if err != nil {
			return err
		}
	}
}
