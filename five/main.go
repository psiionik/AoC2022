package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func reverseColumn(column []byte) []byte {
	new_list := make([]byte, 0)
	for j := len(column) - 1; j >= 0; j-- {
		new_list = append(new_list, column[j])
	}

	return new_list
}

func reverseStateMap(state_map [][]byte) {
	for i := 0; i < len(state_map); i++ {
		new_list := make([]byte, 0)
		for j := len(state_map[i]) - 1; j >= 0; j-- {
			new_list = append(new_list, state_map[i][j])
		}
		state_map[i] = new_list
	}
}

func moveStackPartOne(state_map [][]byte, instruction string) {
	instruction_split_space := strings.Split(instruction, " ")
	amount, _ := strconv.Atoi(instruction_split_space[1])
	source_stack, _ := strconv.Atoi(instruction_split_space[3])
	destination_stack, _ := strconv.Atoi(instruction_split_space[5])
	source_stack--
	destination_stack--

	state_map[destination_stack] = append(state_map[destination_stack], reverseColumn(state_map[source_stack][(len(state_map[source_stack])-amount):])...)
	state_map[source_stack] = state_map[source_stack][:len(state_map[source_stack])-amount]
}

func moveStackPartTwo(state_map [][]byte, instruction string) {
	instruction_split_space := strings.Split(instruction, " ")
	amount, _ := strconv.Atoi(instruction_split_space[1])
	source_stack, _ := strconv.Atoi(instruction_split_space[3])
	destination_stack, _ := strconv.Atoi(instruction_split_space[5])
	source_stack--
	destination_stack--

	state_map[destination_stack] = append(state_map[destination_stack], state_map[source_stack][(len(state_map[source_stack])-amount):]...)
	state_map[source_stack] = state_map[source_stack][:len(state_map[source_stack])-amount]
}

func printTopStacks(state_map [][]byte) {
	result := ""
	for i := 0; i < len(state_map); i++ {
		if len(state_map[i]) <= 0 {
			continue
		}
		result += string(state_map[i][len(state_map[i])-1])
	}

	fmt.Println(result)
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

// 	state_map := make([][]byte, 9)

// 	for s := range state_map {
// 		state_map[s] = make([]byte, 0)
// 	}

// 	fileScanner := bufio.NewScanner(readFile)
// 	fileScanner.Split(bufio.ScanLines)
// 	for fileScanner.Scan() {
// 		token := fileScanner.Text()
// 		if token == "" {
// 			break
// 		}

// 		for i := 0; i < len(token); i += 4 {
// 			letter := token[i+1]

// 			if letter >= 65 && letter <= 90 {
// 				index_to_state_map := i / 4
// 				state_map[index_to_state_map] = append(state_map[index_to_state_map], letter)
// 			}
// 		}
// 	}

// 	reverseStateMap(state_map)

// 	for fileScanner.Scan() {
// 		token := fileScanner.Text()

// 		moveStackPartOne(state_map, token)
// 	}

// 	printTopStacks(state_map)

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

	state_map := make([][]byte, 9)

	for s := range state_map {
		state_map[s] = make([]byte, 0)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		token := fileScanner.Text()
		if token == "" {
			break
		}

		for i := 0; i < len(token); i += 4 {
			letter := token[i+1]

			if letter >= 65 && letter <= 90 {
				index_to_state_map := i / 4
				state_map[index_to_state_map] = append(state_map[index_to_state_map], letter)
			}
		}
	}

	reverseStateMap(state_map)

	for fileScanner.Scan() {
		token := fileScanner.Text()

		moveStackPartTwo(state_map, token)
	}

	printTopStacks(state_map)

	readFile.Close()
}
