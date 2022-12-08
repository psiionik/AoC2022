package main

import (
	"bufio"
	"fmt"
	"os"
)

func checkUniqueCharacters(stream string) bool {

	for i := 0; i < len(stream); i++ {
		for j := i + 1; j < len(stream); j++ {
			if stream[i] == stream[j] {
				return false
			}
		}
	}

	return true
}

// Part One
// func main() {

// 	debug := false

// 	if len(os.Args) == 2 {
// 		if os.Args[1] == "debug" {
// 			debug = true
// 		}
// 	}

// 	var readFile *os.File
// 	var err error

// 	if debug {
// 		readFile, err = os.Open("sample.txt")
// 	} else {
// 		readFile, err = os.Open("input.txt")
// 	}

// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}

// 	position := 0

// 	fileScanner := bufio.NewScanner(readFile)
// 	fileScanner.Split(bufio.ScanLines)
// 	for fileScanner.Scan() {
// 		token := fileScanner.Text()

// 		for i := 0; i+4 < len(token); i++ {
// 			if checkUniqueCharacters(token[i : i+4]) {
// 				position = i + 4
// 				break
// 			}
// 		}
// 	}

// 	fmt.Println(position)

// 	readFile.Close()
// }

// Part Two
func main() {

	debug := false

	if len(os.Args) == 2 {
		if os.Args[1] == "debug" {
			debug = true
		}
	}

	var readFile *os.File
	var err error

	if debug {
		readFile, err = os.Open("sample.txt")
	} else {
		readFile, err = os.Open("input.txt")
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	position := 0

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		token := fileScanner.Text()

		for i := 0; i+14 < len(token); i++ {
			if checkUniqueCharacters(token[i : i+14]) {
				position = i + 14
				break
			}
		}
	}

	fmt.Println(position)

	readFile.Close()
}
