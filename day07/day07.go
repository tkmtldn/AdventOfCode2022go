package main

import (
	"bufio"
	"fmt"
	"os"
)

func GetStrings(path string) []string {
	file, err := os.Open(path)
	check(err)
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	scan := bufio.NewScanner(file)

	var result []string

	for scan.Scan() {
		result = append(result, scan.Text())
	}

	check(scan.Err())

	return result
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data := GetStrings("./day05/input05.txt")
	fmt.Printf("First answer: %v \n", data)
	fmt.Printf("Second answer: %v \n", data)
}
