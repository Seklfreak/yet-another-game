package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Seklfreak/yet-another-game/actions/enemy_ship"
	"github.com/Seklfreak/yet-another-game/actions/move"
	"github.com/Seklfreak/yet-another-game/actions/nothing"
	"github.com/Seklfreak/yet-another-game/actions/quit"
	"github.com/Seklfreak/yet-another-game/actions/restore"
	"github.com/Seklfreak/yet-another-game/actions/setup"
	"github.com/Seklfreak/yet-another-game/models"
	"github.com/manifoldco/promptui"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Welcome!")

	// prepare state with default values
	state := &models.State{
		Health:  100,
		Credits: 100,

		PositionX: 0,
		PositionY: 0,

		Encountered: map[int]map[int]bool{0: {0: true}},

		ActionContext: make(map[string]string),
	}

	// restore game from save game if possible
	(&restore.Action{}).Do(state)
	// launch game setup, eg ship name
	(&setup.Action{}).Do(state)

	// all actions of the general game loop
	loopActions := models.Actions{
		&move.Action{},
		&quit.Action{},
	}
	loopActionKeys := loopActions.Keys()

	// all encounters in space, with a configurable probability
	type encounter struct {
		Chance int
		Action models.Action
	}
	encounters := []encounter{
		{Chance: 15, Action: &nothing.Action{}},
		{Chance: 10, Action: &enemy_ship.Action{}},
	}
	var encountersChanceSum int
	for _, encounter := range encounters {
		encountersChanceSum += encounter.Chance
	}

	// game loop
	for {
		// if the health is 0 or lower, we died :( saving game and exiting
		if state.Health <= 0 {
			(&quit.Action{}).Do(state)
		}

		// ask and perform action
		_, result, _ := (&promptui.Select{
			Label: "What do you want to do?",
			Items: loopActionKeys,
		}).Run()
		if loopActions.Do(state, result) {
			break
		}

		// discover new space tile if undiscovered
		if discover(state) {
			fmt.Printf("Discovering X %d Y %d…\n", state.PositionX, state.PositionY)

			// find and do random encounter
			discovery := rand.Intn(encountersChanceSum)
			for _, encounter := range encounters {
				if discovery < encounter.Chance {
					encounter.Action.Do(state)
					break
				}

				discovery -= encounter.Chance
			}
		}
	}

}

// discover returns true if a tile has not been discovered yet
// it will mark the tile as discovered
func discover(state *models.State) bool {
	if state.Encountered[state.PositionX][state.PositionY] {
		return false
	}

	// mark position as encountered
	if state.Encountered[state.PositionX] == nil {
		state.Encountered[state.PositionX] = make(map[int]bool)
	}
	state.Encountered[state.PositionX][state.PositionY] = true

	return true
}
