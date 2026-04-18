package model

import (
	"fmt"
	"strconv"
	"strings"
)

type State struct {
	Board  [4][4]int
	EmptyX int
	EmptyY int
}

func NewSolvedState() State {
	return State{
		Board: [4][4]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
			{13, 14, 15, 0},
		},
		EmptyX: 3,
		EmptyY: 3,
	}
}

func (s *State) swap(x, y int) {
	s.Board[s.EmptyX][s.EmptyY] = s.Board[x][y]
	s.Board[x][y] = 0

	s.EmptyX = x
	s.EmptyY = y
}

func (s *State) Move(input string) error {
	num, err := strconv.Atoi(input)
	if err != nil {
		return fmt.Errorf("tile must be a number")
	}

	var tx, ty int
	found := false
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if s.Board[i][j] == num {
				tx, ty = i, j
				found = true
			}
		}
	}
	if !found {
		return fmt.Errorf("tile %d not found", num)
	}

	if (tx == s.EmptyX && (ty == s.EmptyY+1 || ty == s.EmptyY-1)) ||
		(ty == s.EmptyY && (tx == s.EmptyX+1 || tx == s.EmptyX-1)) {
		s.swap(tx, ty)
		return nil
	}

	return fmt.Errorf("tile %d is not neighboring to empty space", num)
}

func (s *State) Clone() State {
	return *s
}

func (s State) String() string {
	var sb strings.Builder

	for _, row := range s.Board {
		for j, cell := range row {
			if cell == 0 {
				sb.WriteString("  _")
			} else {
				sb.WriteString(fmt.Sprintf("%3d", cell))
			}
			if j < 3 {
				sb.WriteString(" |")
			}
		}
		sb.WriteString("\n")
		sb.WriteString("----+----+----+----\n")
	}

	return sb.String()
}

func (s *State) MoveByDirection(direction string) error {
	switch direction {
	case "up":
		if s.EmptyX == 3 {
			return fmt.Errorf("cannot move up")
		}
		s.swap(s.EmptyX+1, s.EmptyY)
	case "down":
		if s.EmptyX == 0 {
			return fmt.Errorf("cannot move down")
		}
		s.swap(s.EmptyX-1, s.EmptyY)
	case "left":
		if s.EmptyY == 3 {
			return fmt.Errorf("cannot move left")
		}
		s.swap(s.EmptyX, s.EmptyY+1)
	case "right":
		if s.EmptyY == 0 {
			return fmt.Errorf("cannot move right")
		}
		s.swap(s.EmptyX, s.EmptyY-1)
	default:
		return fmt.Errorf("unknown direction")
	}
	return nil
}
