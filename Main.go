package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

//GameContext holds the context for everything going on
type GameContext struct {
	Player1 player
	Player2 player
	Board   board
}

type player struct {
	turn   bool
	tokens [6]int
}

type board struct {
	winCondition bool
}

func main() {
	Game := new(GameContext)
	setup(Game)

	for Game.Board.winCondition == false {
		rollValue := roll()
		validMoves := checkValidMoves(rollValue, Game)
		movePiece(rollValue, validMoves, Game)

		swapPlayers(Game)

		fmt.Println(Game)
	}
}

func swapPlayers(Game *GameContext) {
	if Game.Player1.turn == true {
		Game.Player1.turn = false
		Game.Player2.turn = true
	} else {
		Game.Player2.turn = false
		Game.Player1.turn = true
	}
}

func movePiece(rollValue int, validMoves []int, Game *GameContext) {
	if len(validMoves) == 0 {
		return
	}
	fmt.Println("Valid moves include token in space", validMoves)
	fmt.Println("Move piece:")
	reader := bufio.NewReader(os.Stdin)
	input, _, err := reader.ReadLine()
	check(err)
	inputStr, err := strconv.Atoi(string((input)))
	check(err)
	token, err := strconv.Atoi(strings.Trim(fmt.Sprintln(inputStr), "\n"))
	check(err)
	if Game.Player1.turn == true {
		Game.Player1.tokens[token] = rollValue
	} else {
		Game.Player2.tokens[token] = rollValue
	}
}

func checkValidMoves(roll int, Game *GameContext) []int {
	var validMoves []int
	if roll == 0 {
		fmt.Println("Rolled a 0, no turn!")
		return validMoves
	}
	if Game.Player1.turn == true {
		for x := 0; x < 6; x++ {
			if Game.Player1.tokens[x] == 0 {
				validMoves = append(validMoves, x)
				break
			}
		}
	} else {
		for x := 0; x < 6; x++ {
			if Game.Player1.tokens[x] == 0 {
				validMoves = append(validMoves, x)
				break
			}
		}
	}
	return validMoves
}

func check(e error) {
	if e != nil {
		panic(1)
	}
}

func roll() int {
	roll := random(0, 5)
	return roll
}

func setup(Game *GameContext) {
	if random(1, 3) == 1 {
		Game.Player1.turn = true
		fmt.Println("Player1 goes first!")
	} else {
		Game.Player2.turn = true
		fmt.Println("Player 2 goes first!")
	}
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
