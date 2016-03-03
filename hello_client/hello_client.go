package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/zhihan/goutil/person"
)

func main() {
	resp, _ := http.Get("http://localhost:1234")
	content, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Printf("Get: %s\n", content)

	resp, _ = http.Get("http://localhost:1234/bruce")
	content, _ = ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	var bruce person.Person
	_ = json.Unmarshal(content, &bruce)
	fmt.Printf("Bruce is: %v\n", bruce)
}
