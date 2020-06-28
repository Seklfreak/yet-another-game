package enemy_ship_pay

import (
	"fmt"
	"strconv"

	"github.com/Seklfreak/yet-another-game/color"
	"github.com/Seklfreak/yet-another-game/models"
)

type Action struct {
}

func (a *Action) Key() string {
	return "Pay ransom"
}

func (a *Action) Do(state *models.State) bool {
	fee, _ := strconv.Atoi(state.ActionContext["enemy_ship_fee"])

	if fee > state.Credits {
		fmt.Println("You do not have enough credits to pay the fee…")
		return false
	}

	state.Credits -= fee

	fmt.Printf("You paid %s%d credits%s to get rid of the enemy ship.\n", color.Yellow, fee, color.Reset)

	return true
}
