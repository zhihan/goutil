package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func isHidden(name string) bool {
	return name[0] == '.'
}

func ext(name string) string {
	var parts = strings.Split(name, ".")
	if len(parts) == 1 {
		return ""
	} else {
		return parts[len(parts)-1]
	}
}

func isEmacsBackup(name string) bool {
	return name[len(name)-1] == '~'
}

type sourceFile struct {
	Is func(string) bool
	Total bool
}


func visitDir(workDir string, sourceCounter sourceFile) {
	var files, _ = ioutil.ReadDir(workDir)
	
	for _, file := range files {
		// Skip hidden files
		if isHidden(file.Name()) {
			continue
		}

		if sourceCounter.Is(file.Name()) {
			sourceCounter.Add(1)
		}

		if file.IsDir() {
			visitDir(filepath.Join(workDir, file.Name()),
			sourceCounter)
		} else {
			fmt.Printf("%s:%s\n", workDir, file.Name())
		}
	}
}

func main() {
	var cwd, _ = os.Getwd()

	var goCounter = Go{0}
	visitDir(cwd, &goCounter)
	fmt.Printf("Source counter %v\n", goCounter)

}
