package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func isHidden(name string) bool {
	return name[1] == '.'
}

func ext(name string) string {
	var parts = strings.Split(name, ".")
	if len(parts) == 1 {
		return ""
	} else {
		return parts[len(parts)-1]
	}
}

func main() {
	var cwd, _ = os.Getwd()

	
	fmt.Println("Workding dir " + cwd)
	var files, _ = ioutil.ReadDir(cwd)
	for _, file := range files {
		fmt.Println(file.Name())
	}

	fmt.Println(" ext1: " + ext(".ab"))
	fmt.Printf(" Ishidden %t\n", isHidden(".ab"))
	fmt.Println(" ext2: " + ext("ab"))

}
