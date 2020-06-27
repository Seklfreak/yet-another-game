package models

type Action interface {
	Key() string
	Do(state *State) bool
}
