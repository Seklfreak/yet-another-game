package delete

import (
	"os"

	"github.com/Seklfreak/yet-another-game/models"
)

type Action struct {
}

func (a *Action) Key() string {
	return "Delete game"
}

func (a *Action) Do(state *models.State) bool {
	os.Remove("./savegame.json")

	return true
}
