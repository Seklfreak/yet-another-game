package enemy_ship

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/Seklfreak/yet-another-game/actions/enemy_ship_fight"
	"github.com/Seklfreak/yet-another-game/actions/enemy_ship_pay"
	"github.com/Seklfreak/yet-another-game/color"
	"github.com/Seklfreak/yet-another-game/input"
	"github.com/Seklfreak/yet-another-game/models"
)

type Action struct {
}

func (a *Action) Key() string {
	return "Enemy Ship"
}

func (a *Action) Do(state *models.State) bool {
	loopActions := models.Actions{
		&enemy_ship_pay.Action{},
		&enemy_ship_fight.Action{},
	}
	loopActionKeys := loopActions.Keys()

	fee := rand.Intn(25) + 1

	state.ActionContext["enemy_ship_fee"] = strconv.Itoa(fee)

	fmt.Printf("%sYou encounter another spaceship!%s\n", color.Red, color.Reset)
	fmt.Printf("%s\"Good to meet you travellers of %s.\"%s\n", color.Purple, state.Name, color.Reset)
	fmt.Printf("%s\"We do not want any complication to your endeavour!\"%s\n", color.Purple, color.Reset)
	fmt.Printf("%s\"If you pay us %s%d credits%s we will move right along.\"\n%s",
		color.Purple, color.Yellow, fee, color.Purple, color.Reset,
	)
	fmt.Printf("You have %s%d credits%s.\n", color.Yellow, state.Credits, color.Reset)

	for {
		result := input.Choose("What do you want to do?", loopActionKeys)
		if loopActions.Do(state, result) {
			break
		}
	}

	return false
}
