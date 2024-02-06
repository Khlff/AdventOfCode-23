package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(p1("input.txt"))
	fmt.Println(p2("input.txt"))
}

func p1(filepath string) int32 {
	fh, _ := os.Open(filepath)
	defer fh.Close()

	const redCubes = 12
	const greenCubes = 13
	const blueCubes = 14

	re := regexp.MustCompile(`(\d+) (red|blue|green)`)
	var sum int32

	var gameNumber int32 = 1
	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		s := scanner.Text()
		sets := strings.Split(s, ";")

		var flag = true
		for _, set := range sets {
			if !flag {
				break
			}
			matches := re.FindAllStringSubmatch(set, -1)
			for _, match := range matches {
				num, _ := strconv.Atoi(match[1])
				if match[2] == "red" {
					if num > redCubes {
						flag = false
						break
					}
				} else if match[2] == "blue" {
					if num > blueCubes {
						flag = false
						break
					}
				} else {
					if num > greenCubes {
						flag = false
						break
					}
				}
			}
		}

		if flag {
			sum += gameNumber
		}

		gameNumber++
	}

	return sum
}

func p2(filepath string) int64 {
	fh, _ := os.Open(filepath)
	defer fh.Close()

	re := regexp.MustCompile(`(\d+) (red|blue|green)`)
	var sum int64

	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		var minRed = 0
		var minBlue = 0
		var minGreen = 0

		s := scanner.Text()
		sets := strings.Split(s, ";")

		for _, set := range sets {
			matches := re.FindAllStringSubmatch(set, -1)
			for _, match := range matches {
				num, _ := strconv.Atoi(match[1])
				if match[2] == "red" {
					if num > minRed {
						minRed = num
					}
				} else if match[2] == "blue" {
					if num > minBlue {
						minBlue = num
					}
				} else {
					if num > minGreen {
						minGreen = num
					}
				}
			}
		}

		sum += int64(minGreen * minBlue * minRed)
	}

	return sum
}
