package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/sidwebworks/gofetch-client/gofetch"
)

var fetch = getTypicodeApiClient()

var logger = log.Default()

type Thing struct {
	userId    int
	id        int
	title     string
	completed bool
}

func getData() {
	res, err := fetch.Get("https://jsonplaceholder.typicode.com/todos/1", nil)

	if err != nil {
		logger.Println(err)
		panic(err)
	}

	thing := Thing{}

	err = res.Json(&thing)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Stuff %v", thing)
}

func main() {
	getData()
}

func getTypicodeApiClient() gofetch.Client {

	h := make(http.Header)

	h.Set("Authorization", "Bearer token")

	clientB := gofetch.New()

	clientB.SetHeaders(h)

	clientB.SetConnectionTimeout(time.Second * 20)

	return clientB.Build()
}
