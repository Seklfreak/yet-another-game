package models

type State struct {
	Health  int
	Credits int

	PositionX int
	PositionY int

	Encountered map[int]map[int]bool // map[x]map[y]bool

	ActionContext map[string]string `json:"-"`
}
