package main

import (
	"os"
	"fmt"
	"log"
  "errors"
	"net/http"

  "github.com/MuhammadHasbiAshshiddieqy/Golang-RESTful-API/service"
)

func main() {
	err := checkEnv()
	if err != nil {
		log.Fatal(err)
	}

  router := service.HandleRequests()

	fmt.Printf("Starting server at port 8080\n")
  log.Fatal(http.ListenAndServe(":8080", router))
}

func checkEnv() error {
	log.Println("Checking environment.....")

	_, ok := os.LookupEnv("DATA_SOURCE")
	if !ok {
		return errors.New("env:data_source does not exist")
	}

	log.Println("Environment is complete.....")
	return nil
}
