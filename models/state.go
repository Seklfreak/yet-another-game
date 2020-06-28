package models

// State represents the shared state of the game
type State struct {
	Name string

	Health  int
	Credits int
	Exp     int

	PositionX int
	PositionY int

	Encountered map[int]map[int]bool // map[x]map[y]bool

	ActionContext map[string]string
}

// NewState creates a new state object with sensible default values
func NewState() *State {
	return &State{
		Health:  100,
		Credits: 100,

		PositionX: 0,
		PositionY: 0,

		Encountered: map[int]map[int]bool{0: {0: true}},

		ActionContext: make(map[string]string),
	}
}

// GetLevel calculates the current level based on the experience in the state
func (s *State) GetLevel() int {
	level := s.Exp / 100

	return level + 1
}
