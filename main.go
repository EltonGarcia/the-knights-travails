package main

import (
	"fmt"
	"regexp"
	"strings"
)

// Regex to validate chessNotation
var chessNotation = regexp.MustCompile(`^[aA-hH][1-8]$`)

// Prompts the user the start and target position and prints the shortest path.
func main() {
	s := promptPosition("Please insert the starting position: ")
	start := convertToSquare(s)
	t := promptPosition("Please insert the target position: ")
	target := convertToSquare(t)
	if start == target{
		fmt.Printf("The start and target positions must be different.")
		return
	}
	result := findMoves(start, target)
	fmt.Printf("The shortest path from %s to %s is: \n",
		squareToChessNotation(start), squareToChessNotation(target))
	printPath(result)
}

// Prints every move of a path in chess notation.
func printPath(path Path) {
	// Skip the start square
	moves := path.Moves[1:]
	for _, square := range moves {
		fmt.Printf("%s ", squareToChessNotation(square))
	}
	fmt.Println()
}

// Converts a Square to a chess notation string.
func squareToChessNotation(s Square) string {
	x := rune(s.X - 1 + 'A')
	y := rune(s.Y + '0')
	return string([]rune{x, y})
}

// Converts a chess notation string to a Square.
func convertToSquare(position string) Square {
	position = strings.ToUpper(position)
	x := int(position[0]-'A') + 1
	y := int(position[1] - '0')
	return Square{X: x, Y: y}
}

// Prompts the user the input of a chess position.
func promptPosition(message string) string {
	var input string
	for {
		fmt.Print(message)
		if _, err := fmt.Scan(&input); err != nil {
			println("Invalid input!")
			continue
		}
		if !isValidChessNotation(input) {
			fmt.Printf("%s is not a valid chess position.\n", input)
			continue
		}
		return input
	}
}

// Validates if a string is a valid chess notation.
func isValidChessNotation(input string) bool {
	return chessNotation.MatchString(input)
}
