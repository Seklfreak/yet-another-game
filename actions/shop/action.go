package shop

import (
	"fmt"
	"strconv"

	"github.com/Seklfreak/yet-another-game/actions/shop_quit"
	"github.com/Seklfreak/yet-another-game/actions/shop_repair"
	"github.com/Seklfreak/yet-another-game/color"
	"github.com/Seklfreak/yet-another-game/input"
	"github.com/Seklfreak/yet-another-game/models"
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
	fmt.Printf("%s\"Welcome %s.\"%s\n", color.Purple, state.Name, color.Reset)
	if repairFee > 0 {
		fmt.Printf("%s\"We can repair your ship for %s%d credits%s.\"%s\n",
			color.Purple, color.Yellow, repairFee, color.Purple, color.Reset,
		)
	}
	fmt.Printf("You have %s%d credits%s.\n", color.Yellow, state.Credits, color.Reset)

	for {
		repairFee = 100 - state.Health
		state.ActionContext["shop_repair_fee"] = strconv.Itoa(repairFee)

		result := input.Choose("What do you want to do?", loopActionKeys)
		if loopActions.Do(state, result) {
			break
		}
	}

	return false
}
