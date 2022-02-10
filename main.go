package main

import "github.com/nelsonlpco/transactions/api/adapter/rest"

func main() {
	server := rest.NewServer()

	server.Start()
}
