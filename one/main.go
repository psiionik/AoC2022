package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// Part 1
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

// 	fileScanner := bufio.NewScanner(readFile)

// 	fileScanner.Split(bufio.ScanLines)

// 	max_val := -1
// 	current_val := 0

// 	for fileScanner.Scan() {
// 		token := fileScanner.Text()

// 		if token == "" {
// 			if max_val < current_val {
// 				max_val = current_val
// 			}
// 			current_val = 0
// 		} else {
// 			val, _ := strconv.Atoi(token)
// 			current_val += val
// 		}
// 	}

// 	fmt.Println(max_val)

// 	readFile.Close()
// }

// Part 2
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

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	value_list := make([]int, 0)
	current_val := 0

	for fileScanner.Scan() {
		token := fileScanner.Text()

		if token == "" {
			value_list = append(value_list, current_val)
			current_val = 0
		} else {
			val, _ := strconv.Atoi(token)
			current_val += val
		}
	}

	value_list = append(value_list, current_val)

	sort.Sort(sort.Reverse(sort.IntSlice(value_list)))

	fmt.Println(value_list[0] + value_list[1] + value_list[2])

	readFile.Close()
}
