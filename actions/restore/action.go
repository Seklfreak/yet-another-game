package restore

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Seklfreak/yet-another-game/models"
)

type Action struct {
}

func (a *Action) Key() string {
	return "Restore"
}

func (a *Action) Do(state *models.State) bool {
	data, err := ioutil.ReadFile("./savegame.json")
	if err != nil {
		// ignore path error, (usually "no such file or directory")
		if _, ok := err.(*os.PathError); ok {
			return false
		}

		panic(err)
	}

	err = json.Unmarshal(data, &state)
	if err != nil {
		panic(err)
	}

	fmt.Println("restored previous game")
	return false
}
