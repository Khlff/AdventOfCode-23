package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

/*
	First day solutions
*/

func main() {
	fmt.Println(p1("input_1.txt"))
	fmt.Println(p2("input_2.txt"))
}

func p1(filepath string) int32 {
	fh, _ := os.Open(filepath)
	defer fh.Close()

	var sum int32
	var digits []int32

	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		digits = digits[:0]
		for _, e := range scanner.Text() {
			if unicode.IsDigit(e) {
				digits = append(digits, e-48)
			}
		}
		sum += digits[0]*10 + digits[len(digits)-1]
	}

	return sum
}

func p2(filepath string) int32 {
	fh, _ := os.Open(filepath)
	defer fh.Close()

	digitsLetters := map[int32]string{
		1: "one",
		2: "two",
		3: "three",
		4: "four",
		5: "five",
		6: "six",
		7: "seven",
		8: "eight",
		9: "nine",
	}

	var sum int32
	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		var digits []int32
		inputString := scanner.Text()

		for len(inputString) > 0 {
			if unicode.IsDigit(rune(inputString[0])) {
				digits = append(digits, rune(inputString[0])-48)
				inputString = inputString[1:]
				continue
			}

			for num, s := range digitsLetters {
				if strings.HasPrefix(inputString, s) {
					digits = append(digits, num)
				}
			}
			inputString = inputString[1:]
		}

		sum += digits[0]*10 + digits[len(digits)-1]
	}
	return sum
}
