package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	doThing("input.txt", false, false)
	doThing("input.txt", true, false)
	doPartTwo("input.txt")
}

func doThing(filename string, partTwo bool, debug bool) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	lines := make([]string, 0)

	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	numbers := make([]int, len(lines))

	// make a digits regex
	digitsOnly := regexp.MustCompile(`\d`)

	for i, line := range lines {
		if partTwo {
			line = replaceStringNumbers(line)
		}
		digits := digitsOnly.FindAllString(line, -1)
		// Skip bugged lines
		if len(digits) == 0 {
			fmt.Print("BUGGED", "\n") // This is here because the test input fails on the first version
			continue
		}
		str := digits[0] + digits[len(digits)-1]
		num, err := strconv.Atoi(str)
		if err == nil {
			numbers[i] = num
		}
	}

	total := sumSlice(numbers)
	fmt.Print(total, "\n")
}

func sumSlice(nums []int) int {
	total := 0
	for _, v := range nums {
		total = total + v
	}
	return total
}

func replaceStringNumbers(str string) string {
	r := strings.NewReplacer(
		// Handle the ones that overlap at the end of strings
		"twoneeight", "218",
		// And these
		"oneight", "18",
		"twone", "21",
		"eightwo", "82",
		"threeight", "38",
		"nineight", "98",
		"sevenine", "79",
		"fiveight", "58",
		// And then just singles
		"one", "1",
		"two", "2",
		"three", "3",
		"four", "4",
		"five", "5",
		"six", "6",
		"seven", "7",
		"eight", "8",
		"nine", "9",
	)
	return r.Replace(str)
}

// doPartTwo is just a different approach
func doPartTwo(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	lines := make([]string, 0)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	numbers := make([]int, len(lines))

	// make a digits regex
	digitsOnly := regexp.MustCompile(`\d`)

	for i, line := range lines {
		// Replace words with numbers
		forward := replaceStringsForward(line)
		backward := replaceStringsBackwards(reverseString(line))
		// Now extract only the digits using regex
		fDigits := digitsOnly.FindAllString(forward, -1)
		bDigits := digitsOnly.FindAllString(backward, -1)

		str := fDigits[0] + bDigits[0]
		num, err := strconv.Atoi(str)
		if err == nil {
			numbers[i] = num
		}
	}
	total := sumSlice(numbers)
	fmt.Print(total, "\n")
}

func reverseString(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func replaceStringsForward(str string) string {
	r := strings.NewReplacer(
		"one", "1",
		"two", "2",
		"three", "3",
		"four", "4",
		"five", "5",
		"six", "6",
		"seven", "7",
		"eight", "8",
		"nine", "9",
	)
	return r.Replace(str)
}

func replaceStringsBackwards(str string) string {
	r := strings.NewReplacer(
		"eno", "1",
		"owt", "2",
		"eerht", "3",
		"ruof", "4",
		"evif", "5",
		"xis", "6",
		"neves", "7",
		"thgie", "8",
		"enin", "9",
	)
	return r.Replace(str)
}
