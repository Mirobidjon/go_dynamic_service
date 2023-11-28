package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	pwd := os.Args[1] + "/genproto"
	readFolder(pwd)
}

func readFolder(pwd string) {
	files, err := os.ReadDir(pwd)
	if err != nil {
		log.Fatalf("fafsd %v", err)
	}

	for _, file := range files {
		name := pwd + "/" + file.Name()
		if file.IsDir() {
			readFolder(name)
			fmt.Println("added bson tag on " + file.Name() + " files")
		} else {
			err := updateFiles(name)
			if err != nil {
				log.Fatalf("fafsd %v", err)
			}
		}
	}
}

func updateFiles(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}

	text := ""

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "json:\"") {
			line = swap(line)
		}

		text += line + "\n"
	}

	file.Close()

	file, err = os.Create(fileName)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.WriteString(text)
	return err
}

func swap(s string) string {
	var response string
	arr := strings.Split(s, "json:\"")

	if len(arr) > 1 {
		response += arr[0]

		if q := strings.Split(arr[1], ","); len(q) > 0 {
			name := q[0]
			response += "json:\"" + name + "\"" + " bson:\"" + name + "\"" + "`"
		}
	}

	return response
}
