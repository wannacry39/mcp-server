package model

type State struct {
	PlayGround [4][4]int
	EmptyX     int
	EmptyY     int
}

func NewState() State {
	s := State{
		PlayGround: [4][4]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
			{13, 14, 15, 0},
		},
		EmptyX: 3,
		EmptyY: 3,
	}

	return s
}

func (s *State) MoveUp() {

}

func (s *State) MoveDown() {

}

func (s *State) MoveLeft() {

}

func (s *State) MoveRight() {

}
