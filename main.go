package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Seklfreak/yet-another-game/actions/delete"
	"github.com/Seklfreak/yet-another-game/actions/enemy_ship"
	"github.com/Seklfreak/yet-another-game/actions/move"
	"github.com/Seklfreak/yet-another-game/actions/nothing"
	"github.com/Seklfreak/yet-another-game/actions/quit"
	"github.com/Seklfreak/yet-another-game/actions/restore"
	"github.com/Seklfreak/yet-another-game/actions/setup"
	"github.com/Seklfreak/yet-another-game/actions/shop"
	"github.com/Seklfreak/yet-another-game/color"
	"github.com/Seklfreak/yet-another-game/input"
	"github.com/Seklfreak/yet-another-game/models"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Printf("%sWelcome!%s\n", color.Yellow, color.Reset)

	state := models.NewState()

	// restore game from save game if possible
	(&restore.Action{}).Do(state)
	// launch game setup, eg ship name
	(&setup.Action{}).Do(state)

	// all actions of the general game loop
	// new possible actions of the main game loop can be added here
	loopActions := models.Actions{
		&move.Action{},
		&quit.Action{},
	}
	loopActionKeys := loopActions.Keys()

	// all encounters in space, with a configurable probability
	// new encounters can be added here
	type encounter struct {
		Chance int
		Action models.Action
	}
	encounters := []encounter{
		{Chance: 15, Action: &nothing.Action{}},
		{Chance: 10, Action: &enemy_ship.Action{}},
		{Chance: 5, Action: &shop.Action{}},
	}
	var encountersChanceSum int
	for _, encounter := range encounters {
		encountersChanceSum += encounter.Chance
	}

	// game loop
	for {
		// if the health is 0 or lower, we died :( saving game and exiting
		if state.Health <= 0 {
			(&delete.Action{}).Do(state)
			break
		}

		// ask and perform action
		result := input.Choose("What do you want to do?", loopActionKeys)
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
