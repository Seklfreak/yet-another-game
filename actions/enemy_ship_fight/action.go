package enemy_ship_fight

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/Seklfreak/yet-another-game/actions/gain_exp"
	"github.com/Seklfreak/yet-another-game/color"
	"github.com/Seklfreak/yet-another-game/models"
)

type Action struct {
}

func (a *Action) Key() string {
	return "Fight"
}

func (a *Action) Do(state *models.State) bool {
	// with higher levels we take less damage
	damage := (rand.Intn(50) + 1) / state.GetLevel()

	lootCredits := rand.Intn(100) + 1
	exp := 20 + rand.Intn(50) + 1

	state.Health -= damage

	if state.Health < 0 {
		fmt.Printf("%sYou did not survive the fight, game over!%s\n", color.Red, color.Reset)
		fmt.Printf("%sThanks for playing.%s\n", color.Red, color.Reset)
		return true
	}

	state.Credits += lootCredits

	fmt.Printf("%sYou destroyed the other ship and took %s%d damage%s. You have %s%d health%s left.%s\n",
		color.Red, color.Yellow, damage, color.Red, color.Yellow, state.Health, color.Red, color.Reset,
	)
	fmt.Printf("After examining the ship you found %s%d credits%s. Now you have %s%d credits%s.\n",
		color.Yellow, lootCredits, color.Reset, color.Yellow, state.Credits, color.Reset,
	)

	state.ActionContext["gain_exp_amount"] = strconv.Itoa(exp)
	(&gain_exp.Action{}).Do(state)

	return true
}
