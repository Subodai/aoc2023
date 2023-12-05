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
	data := new(Data)
	processInput("input.txt", data)
	//fmt.Printf("\n\nRAW DATA: %+v\n", data)
	processData(data)
	//fmt.Printf("\n\nComputed DATA: %+v\n", data)

	var (
		red, _   = strconv.Atoi(os.Args[1])
		green, _ = strconv.Atoi(os.Args[2])
		blue, _  = strconv.Atoi(os.Args[3])
	)
	//fmt.Printf("Red: %d, Green: %d, Blue: %d\n", red, green, blue)
	validateData(data, red, green, blue)
	fmt.Printf("\n\nValidated DATA: %+v\n", data.Stats)
	total, power := getTotal(data)
	fmt.Printf("Total: %d, Power: %d\n", total, power)
}

type Cube struct {
	Colour string
	Count  int
}

type Set struct {
	Cubes []Cube
}

type Game struct {
	ID   int
	Sets []Set
}

type GameStats struct {
	ID    int
	Red   int
	Green int
	Blue  int
	Power int
	Valid bool
}

type Data struct {
	Games []Game
	Stats []GameStats
}

// processInput takes our file and dumps it into the data object for later processing
func processInput(filename string, data *Data) {
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

	digitsOnly := regexp.MustCompile(`\d`)
	lettersOnly := regexp.MustCompile(`[A-Za-z]`)

	for _, line := range lines {
		// Need to get the Game ID
		// get position of semicolon
		start := strings.Index(line, ":")
		id, _ := strconv.Atoi(line[5:start])
		//fmt.Printf("Game ID: %d \n", id)
		game := new(Game)
		game.ID = id

		// Need to get the content
		content := line[start+2:]
		//fmt.Print(content, "\n")

		// Need to split the content into sets
		sets := strings.SplitN(content, "; ", -1)

		// Need to split the sets into cubes
		for _, set := range sets {
			s := new(Set)
			//fmt.Print(set, "\n")
			cubes := strings.SplitN(set, ", ", -1)

			for _, cube := range cubes {
				c := new(Cube)
				//fmt.Print(cube, "\n")
				c.Colour = strings.Join(lettersOnly.FindAllString(cube, -1), "")
				c.Count, _ = strconv.Atoi(strings.Join(digitsOnly.FindAllString(cube, -1), ""))
				//fmt.Printf("%+v\n", c)
				s.Cubes = append(s.Cubes, *c)
			}
			//fmt.Printf("%+v\n", s)
			game.Sets = append(game.Sets, *s)
		}
		data.Games = append(data.Games, *game)
	}
}

// processData processes our list of raw data and gets the max colour values for each game
func processData(data *Data) {
	for _, game := range data.Games {
		stat := new(GameStats)
		stat.ID = game.ID
		stat.Red = 0
		stat.Green = 0
		stat.Blue = 0

		for _, set := range game.Sets {
			for _, cube := range set.Cubes {
				switch cube.Colour {
				case "red":
					if stat.Red < cube.Count {
						stat.Red = cube.Count
					}
				case "green":
					if stat.Green < cube.Count {
						stat.Green = cube.Count
					}
				case "blue":
					if stat.Blue < cube.Count {
						stat.Blue = cube.Count
					}
				}
			}
		}

		stat.Power = stat.Red * stat.Green * stat.Blue

		// Append back to our data
		data.Stats = append(data.Stats, *stat)
	}
}

func validateData(data *Data, red int, green int, blue int) {
	for i, stat := range data.Stats {
		valid := true
		if stat.Red > red || stat.Green > green || stat.Blue > blue {
			valid = false
		}
		data.Stats[i].Valid = valid
	}
}

func getTotal(data *Data) (total int, totalPower int) {
	total = 0
	totalPower = 0
	for _, stat := range data.Stats {
		totalPower += stat.Power
		if stat.Valid {
			total += stat.ID
		}
	}
	return
}
