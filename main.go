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

	state := &models.State{
		Health:  100,
		Credits: 100,

		PositionX: 0,
		PositionY: 0,

		Encountered: map[int]map[int]bool{0: {0: true}},

		ActionContext: make(map[string]string),
	}

	(&restore.Action{}).Do(state)
	(&setup.Action{}).Do(state)

	loopActions := []models.Action{
		&move.Action{},
		&quit.Action{},
	}

	type encounter struct {
		Chance int
		Action models.Action
	}
	encounters := []encounter{
		{Chance: 1, Action: &nothing.Action{}},
		{Chance: 1, Action: &enemy_ship.Action{}},
	}
	var encountersChanceSum int
	for _, encounter := range encounters {
		encountersChanceSum += encounter.Chance
	}

	// game loop
GameLoop:
	for {
		if state.Health <= 0 {
			(&quit.Action{}).Do(state)
			break GameLoop
		}

		fmt.Println("--- loop")

		var items []string
		for _, loopAction := range loopActions {
			items = append(items, loopAction.Key())
		}

		_, result, _ := (&promptui.Select{
			Label: "What do you want to do?",
			Items: items,
		}).Run()

		for _, loopAction := range loopActions {
			if result != loopAction.Key() {
				continue
			}

			if loopAction.Do(state) {
				break GameLoop
			}
		}

		if !state.Encountered[state.PositionX][state.PositionY] {
			// mark position as encountered
			if state.Encountered[state.PositionX] == nil {
				state.Encountered[state.PositionX] = make(map[int]bool)
			}
			state.Encountered[state.PositionX][state.PositionY] = true

			fmt.Printf("Discovering X %d Y %d…\n", state.PositionX, state.PositionY)

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
