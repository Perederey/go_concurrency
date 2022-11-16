package main

import "net/http"

func main() {

	go func() {
		user, err := http.Get("/api/v2/user/1") // 200ms
	}()

	go func() {
		vendor, err := http.Get("/api/v5/vendor/28282") // 400ms
	}()

	go func() {
		config, err := http.Get("/xml/v1/config/28282") // 200ms
	}

	orderInfo := struct {
		user, vendor, config
	}{}
}
