package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func isNum(char rune) bool {
	return char >= '0' && char <= '9'
}
func partOne() {
	fmt.Println("This is a test")

	// This is how to read an entire file into a single var
	// dat, _ := os.ReadFile("input.txt")
	// fmt.Println(string(dat))

	// using a scanner so we can read line by line
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	sumCal := 0
	for scanner.Scan() {
		numString := "empty"
		var lastNum rune
		for _, char := range scanner.Text() {
			if isNum(rune(char)) {
				if numString == "empty" {
					numString = string(char)
				}
				lastNum = char
			}
		}
		numString += string(lastNum)
		fmt.Println("numstring: " + numString)
		curNum, _ := strconv.Atoi(numString)
		sumCal += curNum

	}

	fmt.Println("sum: " + strconv.Itoa(sumCal))
}

func partTwo() {
	fmt.Println("This is a test")

	// This is how to read an entire file into a single var
	// dat, _ := os.ReadFile("input.txt")
	// fmt.Println(string(dat))

	// using a scanner so we can read line by line
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	sumCal := 0
	for scanner.Scan() {
		numString := "empty"
		var lastNum rune
		for _, char := range scanner.Text() {
			if isNum(rune(char)) {
				if numString == "empty" {
					numString = string(char)
				}
				lastNum = char
			}
		}
		numString += string(lastNum)
		fmt.Println("numstring: " + numString)
		curNum, _ := strconv.Atoi(numString)
		sumCal += curNum

	}

	fmt.Println("sum: " + strconv.Itoa(sumCal))
}

func main() {
	partOne()
	partTwo()
}
