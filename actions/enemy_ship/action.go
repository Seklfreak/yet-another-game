package enemy_ship

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/Seklfreak/yet-another-game/actions/enemy_ship_fight"
	"github.com/Seklfreak/yet-another-game/actions/enemy_ship_pay"
	"github.com/Seklfreak/yet-another-game/models"
	"github.com/manifoldco/promptui"
)

type Action struct {
}

func (a *Action) Key() string {
	return "Enemy Ship"
}

func (a *Action) Do(state *models.State) bool {
	loopActions := []models.Action{
		&enemy_ship_pay.Action{},
		&enemy_ship_fight.Action{},
	}

	var items []string
	for _, loopAction := range loopActions {
		items = append(items, loopAction.Key())
	}

	fee := rand.Intn(25) + 1

	state.ActionContext["enemy_ship_fee"] = strconv.Itoa(fee)

	fmt.Println("Oh no, you encounter an enemy ship!")
	fmt.Printf("They ask you to pay a fee of %d credits to move on.\n", fee)
	fmt.Printf("You have %d credits.\n", state.Credits)

GameLoop:
	for {
		prompt := promptui.Select{
			Label: "What do you want to do?",
			Items: items,
		}
		_, result, _ := prompt.Run()

		for _, loopAction := range loopActions {
			if result != loopAction.Key() {
				continue
			}

			if loopAction.Do(state) {
				break GameLoop
			}
		}
	}

	return false
}
