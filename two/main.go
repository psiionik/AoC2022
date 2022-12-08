package main

import (
	"bufio"
	"fmt"
	"os"
)

func sum[T int | uint | byte](arr []T) T {
	var result T = 0
	for _, value := range arr {
		result += value
	}

	return result
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

// 	var totalScore int = 0

// 	fileScanner := bufio.NewScanner(readFile)
// 	fileScanner.Split(bufio.ScanLines)

// 	for fileScanner.Scan() {
// 		token := fileScanner.Text()

// 		b := []byte(token)
// 		x := int(sum(b))

// 		first_value := int(b[0])
// 		switch x {
// 		case 185:
// 			totalScore += 4
// 			break
// 		case 186:
// 			if first_value == 65 {
// 				totalScore += 8
// 			} else if first_value == 66 {
// 				totalScore += 1
// 			}
// 			break
// 		case 187:
// 			if first_value == 65 {
// 				totalScore += 3
// 			} else if first_value == 66 {
// 				totalScore += 5
// 			} else {
// 				totalScore += 7
// 			}
// 			break
// 		case 188:
// 			if first_value == 66 {
// 				totalScore += 9
// 			} else if first_value == 67 {
// 				totalScore += 2
// 			}
// 		case 189:
// 			totalScore += 6
// 			break
// 		}
// 	}

// 	fmt.Println(totalScore)

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

	var totalScore int = 0

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		token := fileScanner.Text()

		b := []byte(token)
		x := int(sum(b))

		first_value := int(b[0])
		switch x {
		// A X
		case 185:
			totalScore += 3
			break
		case 186:
			// A Y
			if first_value == 65 {
				totalScore += 4
				// B X
			} else if first_value == 66 {
				totalScore += 1
			}
			break
		case 187:
			// A Z
			if first_value == 65 {
				totalScore += 8
				// B Y
			} else if first_value == 66 {
				totalScore += 5
				// C X
			} else {
				totalScore += 2
			}
			break
		case 188:
			// B Z
			if first_value == 66 {
				totalScore += 9
				// C Y
			} else if first_value == 67 {
				totalScore += 6
			}
		case 189:
			// C Z
			totalScore += 7
			break
		}
	}

	fmt.Println(totalScore)

	readFile.Close()
}
