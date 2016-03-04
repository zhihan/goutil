package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path"
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
	is    func(string) bool
	Lang  string
	Total int32
	Len   int32
}

func (me *sourceFile) Process(filename string) {
	basename := path.Base(filename)
	var lineSep = []byte{'\n'}
	if me.is(basename) {
		me.Total = me.Total + 1

		content, _ := ioutil.ReadFile(filename)
		
		me.Len = me.Len + int32(bytes.Count(content, lineSep))
	}
}

func visitDir(workDir string, sourceCounters []*sourceFile) {
	var files, _ = ioutil.ReadDir(workDir)

	for _, file := range files {
		// Skip hidden files
		if isHidden(file.Name()) {
			continue
		}

		if file.IsDir() {
			visitDir(filepath.Join(workDir, file.Name()),
				sourceCounters)
		} else {
			for _, counter := range sourceCounters {
				counter.Process(
					filepath.Join(workDir, file.Name()))
			}
		}
	}
}

func isGoSource(name string) bool {
	return ext(name) == "go"
}

func isScalaSource(name string) bool {
	return ext(name) == "scala"
}

func isPythonSource(name string) bool {
	return ext(name) == "py"
}

func isJavaSource(name string) bool {
	return ext(name) == "java"
}

func isGroovySource(name string) bool {
	return ext(name) == "groovy"
}

func main() {
	var cwd, _ = os.Getwd()

	var goCounter = sourceFile{isGoSource, "go", 0, 0}
	var scalaCounter = sourceFile{isScalaSource, "scala", 0, 0}
	var pythonCounter = sourceFile{isPythonSource, "python", 0, 0}
	var javaCounter = sourceFile{isJavaSource, "java", 0, 0}
	var groovyCounter = sourceFile{isGroovySource, "groovy", 0, 0}
	
	var counters = []*sourceFile{&goCounter,
		&scalaCounter,
		&pythonCounter,
		&javaCounter,
		&groovyCounter}
	visitDir(cwd, counters)

	for _, counter := range counters {
		fmt.Printf("%s source code:\n", counter.Lang)
		fmt.Printf("  Total %d files, %d lines. \n",
			counter.Total, counter.Len)
	}

}
