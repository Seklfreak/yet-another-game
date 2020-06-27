package quit

import (
	"encoding/json"
	"io/ioutil"

	"github.com/Seklfreak/yet-another-game/models"
)

type Action struct {
}

func (a *Action) Key() string {
	return "Save and Quit"
}

func (a *Action) Do(state *models.State) bool {
	data, err := json.Marshal(state)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("./savegame.json", data, 0644)
	if err != nil {
		panic(err)
	}

	return true
}
