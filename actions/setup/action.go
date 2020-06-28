package setup

import (
	"errors"

	"github.com/Seklfreak/yet-another-game/models"
	"github.com/manifoldco/promptui"
)

type Action struct {
}

func (a *Action) Key() string {
	return "Setup"
}

func (a *Action) Do(state *models.State) bool {
	if state.Name != "" {
		return false
	}

	result, _ := (&promptui.Prompt{
		Label: "Name your ship",
		Validate: func(input string) error {
			if len(input) <= 0 {
				return errors.New("too short")
			}

			return nil
		},
	}).Run()
	state.Name = result

	return false
}
