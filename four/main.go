package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findPartialOverlap(string_list []string) int {
	pair_one := strings.Split(string_list[0], "-")
	pair_two := strings.Split(string_list[1], "-")

	range_start_one, _ := strconv.Atoi(pair_one[0])
	range_end_one, _ := strconv.Atoi(pair_one[1])
	range_start_two, _ := strconv.Atoi(pair_two[0])
	range_end_two, _ := strconv.Atoi(pair_two[1])

	if range_end_one >= range_start_two && range_start_one <= range_end_two {
		return 1
	}

	if range_end_two >= range_start_one && range_start_two <= range_end_one {
		return 1
	}

	return 0
}

func findCompleteOverlap(string_list []string) int {
	pair_one := strings.Split(string_list[0], "-")
	pair_two := strings.Split(string_list[1], "-")

	range_start_one, _ := strconv.Atoi(pair_one[0])
	range_end_one, _ := strconv.Atoi(pair_one[1])
	range_start_two, _ := strconv.Atoi(pair_two[0])
	range_end_two, _ := strconv.Atoi(pair_two[1])

	if range_start_one <= range_start_two && range_end_one >= range_end_two {
		return 1
	}

	if range_start_one >= range_start_two && range_end_one <= range_end_two {
		return 1
	}

	return 0
}

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

// 	total_overlaps := 0

// 	fileScanner := bufio.NewScanner(readFile)
// 	fileScanner.Split(bufio.ScanLines)
// 	for fileScanner.Scan() {
// 		token := fileScanner.Text()

// 		string_split_comma := strings.Split(token, ",")

// 		total_overlaps += findCompleteOverlap(string_split_comma)
// 	}

// 	fmt.Println(total_overlaps)

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

	total_overlaps := 0

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		token := fileScanner.Text()

		string_split_comma := strings.Split(token, ",")

		total_overlaps += findPartialOverlap(string_split_comma)
	}

	fmt.Println(total_overlaps)

	readFile.Close()
}
