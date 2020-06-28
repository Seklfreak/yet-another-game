package shop

import (
	"fmt"
	"strconv"

	"github.com/Seklfreak/yet-another-game/actions/shop_quit"
	"github.com/Seklfreak/yet-another-game/actions/shop_repair"
	"github.com/Seklfreak/yet-another-game/models"
	"github.com/manifoldco/promptui"
)

type Action struct {
}

func (a *Action) Key() string {
	return "Shop"
}

func (a *Action) Do(state *models.State) bool {
	loopActions := models.Actions{
		&shop_repair.Action{},
		&shop_quit.Action{},
	}
	loopActionKeys := loopActions.Keys()

	repairFee := 100 - state.Health
	state.ActionContext["shop_repair_fee"] = strconv.Itoa(repairFee)

	fmt.Println("You meet a shop!")
	fmt.Printf("\"Welcome %s.\"\n", state.Name)
	if repairFee > 0 {
		fmt.Printf("\"We can repair your ship for %d.\"\n", repairFee)
	}
	fmt.Printf("You have %d credits.\n", state.Credits)

	for {
		repairFee = 100 - state.Health
		state.ActionContext["shop_repair_fee"] = strconv.Itoa(repairFee)

		_, result, _ := (&promptui.Select{
			Label: "What do you want to do?",
			Items: loopActionKeys,
		}).Run()
		if loopActions.Do(state, result) {
			break
		}
	}

	return false
}
