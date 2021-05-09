package main

const boardSize = 8

type Square struct {
	X int
	Y int
}

type Move struct {
	X int
	Y int
}

type Path struct {
	Moves []Move
	Distance int
}

// Specifies all possible moves for a chess knight.
var moves = []Move{
	{ -2, -1 },
	{ -1, -2 },
	{ 1, -2 },
	{ 2, -1 },
	{ -2, 1 },
	{ -1, 2 },
	{ 1, 2 },
	{ 2, 1 },
}

func findMoves(start, target Square){

}

func main(){

}