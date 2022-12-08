package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	uppercase_prio_diff = 38
	lowercase_prio_diff = 96
)

func calculatePriority(character rune) int {
	if character >= 97 && character <= 122 {
		return int(character) - lowercase_prio_diff
	} else if character >= 65 && character <= 90 {
		return int(character) - uppercase_prio_diff
	}

	return -1
}

func findDuplicate(token string) rune {
	token_length := len(token)
	first_compartment := token[:token_length/2]
	second_compartment := token[token_length/2:]

	character_counter := make(map[rune]int, 0)

	for _, character := range first_compartment {
		if _, ok := character_counter[character]; ok {
			character_counter[character] += 1
		} else {
			character_counter[character] = 1
		}
	}

	for _, character := range second_compartment {
		if _, ok := character_counter[character]; ok {
			return character
		}
	}

	return rune(0)
}

func findTriplet(strings_list []string) rune {
	character_counter_one := make(map[rune]bool, 0)
	character_counter_two := make(map[rune]bool, 0)
	character_counter_three := make(map[rune]bool, 0)

	first_string := strings_list[0]
	second_string := strings_list[1]
	third_string := strings_list[2]

	for _, character := range first_string {
		if _, ok := character_counter_one[character]; ok {
			character_counter_one[character] = true
		} else {
			character_counter_one[character] = true
		}
	}

	for _, character := range second_string {
		if _, ok := character_counter_two[character]; ok {
			character_counter_two[character] = true
		} else {
			character_counter_two[character] = true
		}
	}

	for _, character := range third_string {
		if _, ok := character_counter_three[character]; ok {
			character_counter_three[character] = true
		} else {
			character_counter_three[character] = true
		}
	}

	for key, value := range character_counter_one {
		if value == true && character_counter_two[key] && character_counter_three[key] {
			return key
		}
	}

	return rune(0)
}

// Part one
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

// 	var sum int = 0

// 	for fileScanner.Scan() {
// 		token := fileScanner.Text()
// 		strings.ReplaceAll(token, "\n", "")

// 		duplicate_char := findDuplicate(token)

// 		sum += calculatePriority(duplicate_char)
// 	}

// 	fmt.Println(sum)

// 	readFile.Close()
// }

// Part two
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

	var sum int = 0
	count := 0
	strings_list := make([]string, 0)

	for fileScanner.Scan() {
		token := fileScanner.Text()
		strings.ReplaceAll(token, "\n", "")

		if count == 2 {
			count = 0
			strings_list = append(strings_list, token)
			char := findTriplet(strings_list)
			sum += calculatePriority(char)
			strings_list = make([]string, 0)
		} else {
			strings_list = append(strings_list, token)
			count++
		}
	}

	fmt.Println(sum)

	readFile.Close()
}
