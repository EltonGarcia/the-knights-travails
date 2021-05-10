# The Knight’s Travails

Solution for The Knight’s Travails challenge written in go using Breadth-First Search.

## Description

Given a standard 8x8 chessboard, accept two squares identified by algebraic chess notation. The first square is the starting position, and the second square is the ending position. Find the shortest sequence of valid moves to take a Knight piece from the starting position to the ending position. Each move must be a legal move by a Knight. For any two squares there may be more than one valid solution. There are no pieces other than the Knight on the board.

**Algebraic chess notation** identifies each square with a letter from A to H and a number from 1 to 8. The columns are labelled with letters, and the rows are numbered. The lower left is A1.

A **Knight** moves two steps in a straight line from its starting position, turns 90 degrees to the left or right and then moves one square. A Knight can jump over other pieces.

**Input**

Must be two squares, identified in algebraic chess notation representing the starting and ending positions of the Knight, separated by a space.

**Output**

Must be a list of squares through which the Knight passes, in algebraic chess notation. This must include the ending position, but exclude the starting position.

**Example**

Test Input:
`A8 B7`

Expected Output:
`C7 B5 D6 B7`

## Instructions to execute

* Install [Go](https://golang.org/doc/install).

* Clone this repository.

* Navigate to the project directory:
``cd the-knights-travails/``
  
* Build the project:
  ``go build``

* Execute:
``./the-knights-travails``

* Execute the unit-tests:
``go test -v``

**Result**

![Result](/screenshot/result.png?raw=true "result")
  
## Dependencies
[Testify](https://github.com/stretchr/testify) employed to assert unit tests.

[Copier](https://github.com/jinzhu/copier) employed for deep struct copy.