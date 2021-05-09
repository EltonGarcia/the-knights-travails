package main

type Path struct {
	Moves []Square
}

func newPath(s Square) Path {
	return Path{
		Moves: []Square{s},
	}
}

func (p *Path) Distance() int {
	return len(p.Moves)
}

func (p *Path) LastPosition() Square {
	return p.Moves[p.Distance()-1]
}

func (p *Path) HasReachedTarget(target Square) bool {
	position := p.LastPosition()
	//TODO: evaluate square comparison improvements
	return position.X == target.X && position.Y == target.Y
}

func (p *Path) Move(m Move, boardSize int) (Path, bool) {
	position := p.LastPosition()
	square := Square{X: position.X + m.X, Y: position.Y + m.Y}
	if !isSquareValid(square, boardSize) {
		return Path{}, false
	}
	newPath := Path{
		Moves: append(p.Moves, square),
	}
	return newPath, true
}

func isSquareValid(s Square, boardSize int) bool {
	return s.X >= 1 &&
		s.Y >= 1 &&
		s.X <= boardSize &&
		s.Y <= boardSize
}
