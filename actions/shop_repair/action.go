package shop_repair

import (
	"fmt"
	"strconv"

	"github.com/Seklfreak/yet-another-game/models"
)

type Action struct {
}

func (a *Action) Key() string {
	return "Repair ship"
}

func (a *Action) Do(state *models.State) bool {
	fee, _ := strconv.Atoi(state.ActionContext["shop_repair_fee"])

	if fee <= 0 {
		fmt.Println("Your ship does not need to be repaired")
		return false
	}
	if fee > state.Credits {
		fmt.Println("You do not have enough credits to pay for the repair.")
		return false
	}

	state.Credits -= fee
	state.Health = 100

	fmt.Printf("You paid %d credits to repair your ship\n", fee)

	return false
}
