package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Person struct {
	Name string
	Age int32
}

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello world")
}

func getBruce(writer http.ResponseWriter, request *http.Request) {
	var bruce = Person{"Bruce", 2}
	var bytes, _ = json.Marshal(bruce)
	writer.Write(bytes)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/bruce", getBruce)
	http.ListenAndServe(":1234", nil)
}
