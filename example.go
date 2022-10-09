package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/sidwebworks/gofetch-client/gofetch"
)

var typicodeApiClient = getTypicodeApiClient()

func main() {
	logger := log.Default()

	typicodeApiClient.SetMaxIdleConnections(1)

	typicodeApiClient.SetConnectionTimeout(time.Second * 2)

	res, err := typicodeApiClient.Get("https://jsonplaceholder.typicode.com/todos/1", nil)

	if err != nil {
		logger.Println(err)
		panic(err)
	}

	defer res.Body.Close()

	bytes, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func getTypicodeApiClient() gofetch.Client {

	headers := make(http.Header)

	headers.Set("Authorization", "Bearer token")

	client := gofetch.NewBuilder().Build()

	return client
}
