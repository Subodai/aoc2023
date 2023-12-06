package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Race struct {
	Time         int
	Distance     int
	PossibleWins int
}

type Data struct {
	Races []Race
}

func main() {
	lines := fileToLines("input2.txt")
	data := new(Data)
	makeRaceData(lines, data)
	processRacePossibilities(data)
	total := getTotal(data)
	fmt.Print("\nTotal: ", total, "\n")

}

func fileToLines(filename string) (lines []string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return
}

func makeRaceData(lines []string, data *Data) {
	times := strings.SplitN(lines[0][6:], " ", -1)
	distances := strings.SplitN(lines[1][10:], " ", -1)
	for i := range times {
		race := new(Race)
		race.Time, _ = strconv.Atoi(times[i])
		race.Distance, _ = strconv.Atoi(distances[i])
		data.Races = append(data.Races, *race)
	}

}

func processRacePossibilities(data *Data) {
	for index, race := range data.Races {
		for i := 1; i < race.Time; i++ {
			if i*(race.Time-i) > race.Distance {
				race.PossibleWins += 1
			}
		}
		data.Races[index] = race
	}
}

func getTotal(data *Data) (total int) {
	total = 1
	for _, race := range data.Races {
		total *= race.PossibleWins
	}
	return
}
