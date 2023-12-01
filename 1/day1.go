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
	doThing("testinput.txt", true, false)
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

	for i, line := range lines {
		// make a digits regex
		digitsOnly := regexp.MustCompile(`\d`)
		if debug {
			fmt.Print(line, " : ")
		}
		if partTwo {
			line = replaceStringNumbers(line)
		}
		if debug {
			fmt.Print(line, " : ")
		}
		digits := digitsOnly.FindAllString(line, -1)
		// Skip bugged lines
		if len(digits) == 0 {
			fmt.Print("BUGGED", "\n") // This is here because the test input fails on the first version
			continue
		}
		if debug {
			fmt.Print(digits, " : ")
		}
		str := digits[0] + digits[len(digits)-1]
		num, err := strconv.Atoi(str)
		if debug {
			fmt.Print(num, "\n")
		}
		if err == nil {
			numbers[i] = num
		}
		// fmt.Print(ln[0], ln[len(ln)-1], "\n")
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

	r := strings.NewReplacer("one", "1",
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
