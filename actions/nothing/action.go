package nothing

import (
	"fmt"

	"github.com/Seklfreak/yet-another-game/models"
)

type Action struct {
}

func (a *Action) Key() string {
	return "Nothing"
}

func (a *Action) Do(state *models.State) bool {
	fmt.Println("There is nothing here")

	return false
}
