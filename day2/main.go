package day2

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type Game struct {
	ID   int
	Sets []Set
}

type Set struct {
	Red   int
	Green int
	Blue  int
}

func Part1(input string) (sum int, err error) {
	games, err := parseGames(input)
	if err != nil {
		return 0, err
	}

	for _, game := range games {
		if game.isPossible() {
			sum += game.ID
		}
	}

	return sum, nil
}

func (g *Game) isPossible() bool {
	for _, set := range g.Sets {
		if set.Red > 12 || set.Green > 13 || set.Blue > 14 {
			return false
		}
	}
	return true
}

func Part2(input string) (sum int, err error) {
	games, err := parseGames(input)
	if err != nil {
		return 0, err
	}

	for _, game := range games {
		sum += game.calculatePower()
	}

	return sum, nil
}

func (g *Game) calculatePower() int {
	minSet := g.findMinimumSet()
	return minSet.Red * minSet.Green * minSet.Blue
}

func (g *Game) findMinimumSet() (minimumSet Set) {
	for _, set := range g.Sets {
		if set.Red > minimumSet.Red {
			minimumSet.Red = set.Red
		}
		if set.Green > minimumSet.Green {
			minimumSet.Green = set.Green
		}
		if set.Blue > minimumSet.Blue {
			minimumSet.Blue = set.Blue
		}
	}

	return minimumSet
}

func parseGames(input string) (games []Game, err error) {
	if input == "" {
		return games, errors.New("games string is empty")
	}
	for _, line := range strings.Split(input, "\n") {

		var game Game
		game, err = parseGame(line)
		if err != nil {
			return games, err
		}
		games = append(games, game)
	}
	return
}

var (
	numRegexp    = regexp.MustCompile("\\d+")
	colorsRegexp = regexp.MustCompile("red|green|blue")
)

func parseGame(line string) (parsedGame Game, err error) {
	if line == "" {
		return parsedGame, nil
	}
	lineSplit := strings.Split(line, ":")
	if len(lineSplit) < 2 {
		return Game{}, errors.New("Invalid format")
	}
	gameString := lineSplit[0]
	setsString := lineSplit[1]

	id, err := strconv.Atoi(numRegexp.FindString(gameString))
	if err != nil {
		return Game{}, err
	}
	parsedGame.ID = id

	if setsString == "" {
		return parsedGame, errors.New("Sets string is empty")
	}

	sets := strings.Split(setsString, ";")
	for _, setString := range sets {
		views := strings.Split(setString, ",")
		set := Set{}
		for _, view := range views {
			color := colorsRegexp.FindString(view)
			countStr := numRegexp.FindString(view)

			if countStr == "" {
				return Game{}, errors.New("empty count")
			}

			var count int
			count, err = strconv.Atoi(countStr)
			if err != nil {
				return Game{}, err
			}

			switch color {
			case "red":
				set.Red = count
			case "green":
				set.Green = count
			case "blue":
				set.Blue = count
			default:
				return Game{}, errors.New("Unknown color " + color)
			}
		}
		parsedGame.Sets = append(parsedGame.Sets, set)
	}

	return parsedGame, nil
}
