package main

import (
	"fmt"
	"io"
	"net/http"
)


func main() {
  response, err := http.Get("http://localhost:8888")
	if err != nil {
		panic(err)
	}
  defer response.Body.Close()

  byteArray, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
  fmt.Println(string(byteArray))
}
