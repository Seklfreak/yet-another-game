package models

// Action represents an action
type Action interface {
	Key() string
	Do(state *State) bool
}

// Actions is a list of Actions with helper methods
type Actions []Action

// Keys returns all keys of a list of actions
func (a Actions) Keys() []string {
	var loopActionKeys []string
	for _, loopAction := range a {
		loopActionKeys = append(loopActionKeys, loopAction.Key())
	}

	return loopActionKeys
}

// Do performs an action of a list of actions using the key
func (a Actions) Do(state *State, key string) bool {
	for _, loopAction := range a {
		if key != loopAction.Key() {
			continue
		}

		return loopAction.Do(state)
	}

	return false
}
