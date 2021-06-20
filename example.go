package main

import (
	"fmt"
	"log"

	"client-api/clientapi"
)

func main() {
	client := clientapi.CreateClient()

	status, err := client.Health.Check()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Healthcheck status: %v\n", status.Status)
}
