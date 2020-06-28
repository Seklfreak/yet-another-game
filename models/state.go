package models

type State struct {
	Name string

	Health  int
	Credits int
	Exp     int

	PositionX int
	PositionY int

	Encountered map[int]map[int]bool // map[x]map[y]bool

	ActionContext map[string]string `json:"-"`
}

func (s *State) GetLevel() int {
	level := s.Exp / 100

	return level + 1
}
