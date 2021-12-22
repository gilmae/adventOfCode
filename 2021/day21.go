package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Player struct {
	Position, Score int
}

func (p Player) Copy() Player {
	return Player{Position: p.Position, Score: p.Score}
}

type d100 struct {
	Value int
}

type Game struct {
	Player1, Player2 Player
}

func (g Game) Copy() Game {
	return Game{Player1: g.Player1.Copy(), Player2: g.Player2.Copy()}
}

func (d *d100) Roll() int {
	result := d.Value
	d.Value++
	if d.Value == 101 {
		d.Value = 1
	}
	return result
}

var inputFile = flag.String("inputFile", "inputs/day21.input", "Relative file path to use as input.")
var part = flag.String("part", "a", "Which part to solve")

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	lines := strings.Split(contents, "\n")

	p1 := Player{0, 0}
	p1.Position, _ = strconv.Atoi(strings.Split(lines[0], ": ")[1])
	p2 := Player{0, 0}
	p2.Position, _ = strconv.Atoi(strings.Split(lines[1], ": ")[1])

	if *part == "a" {
		die := d100{1}
		rolls := 0
		won := false

		for {
			won = p1.TakeTurn(die.Roll() + die.Roll() + die.Roll())
			rolls += 3
			if won {
				fmt.Printf("Player 1 won on %d after %d rolls\n", p2.Score*rolls, rolls)
				break
			}

			won = p2.TakeTurn(die.Roll() + die.Roll() + die.Roll())

			rolls += 3
			if won {
				fmt.Printf("Player 2 won on %d after %d rolls\n", p1.Score*rolls, rolls)
				break
			}

		}
	} else if *part == "b" {
		cache := make(map[string][2]int64)
		p1wins, p2wins := playDiracDice([2]int{0, 0}, [2]int{p1.Position, p2.Position}, 3, true, cache)

		if p1wins > p2wins {
			fmt.Println(p1wins)
		} else {
			fmt.Println(p2wins)
		}
	}
}

func (p *Player) TakeTurn(roll int) bool {
	p.Position = ((p.Position - 1 + roll) % 10) + 1

	p.Score += p.Position

	return p.Score >= 1000
}

func playDiracDice(playerScores [2]int, playerPositions [2]int, rollsLeftForPlayer int, player1Turn bool, cache map[string][2]int64) (int64, int64) {
	player1Wins, player2Wins := int64(0), int64(0)

	// Have we been here before?
	key := fmt.Sprint(playerScores, playerPositions, rollsLeftForPlayer, player1Turn)
	if v, ok := cache[key]; ok {
		return v[0], v[1]
	}

	// Check if the current player has won
	var player int
	if player1Turn {
		player = 0
	} else {
		player = 1
	}

	newScores := [2]int{playerScores[0], playerScores[1]}
	if rollsLeftForPlayer == 0 {
		// No more rolls of the dice, check their score
		newScores[player] += playerPositions[player]
		if newScores[player] >= 21 {
			if player1Turn {
				return 1, 0
			} else {
				return 0, 1
			}
		}

		player1Turn = !player1Turn
		rollsLeftForPlayer = 3
		player = (player + 1) % 2
	}

	for d := 1; d < 4; d++ {
		newPositions := [2]int{playerPositions[0], playerPositions[1]}
		newPositions[player] += d
		if newPositions[player] > 10 {
			newPositions[player] -= 10
		}

		p1w, p2w := playDiracDice(newScores, newPositions, rollsLeftForPlayer-1, player1Turn, cache)
		player1Wins += p1w
		player2Wins += p2w
	}

	cache[key] = [2]int64{player1Wins, player2Wins}
	return player1Wins, player2Wins

}

//270005289024391 (/)
