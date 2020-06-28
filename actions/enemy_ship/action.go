package enemy_ship

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/Seklfreak/yet-another-game/actions/enemy_ship_fight"
	"github.com/Seklfreak/yet-another-game/actions/enemy_ship_pay"
	"github.com/Seklfreak/yet-another-game/input"
	"github.com/Seklfreak/yet-another-game/models"
	"github.com/manifoldco/promptui"
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

	fmt.Println(promptui.Styler(promptui.FGRed)("You encounter another spaceship!"))
	fmt.Printf(promptui.Styler(promptui.FGMagenta)("\"Good to meet you travellers of %s.\"\n"), state.Name)
	fmt.Println(promptui.Styler(promptui.FGMagenta)("\"We do not want any complication to your endeavour!\""))
	fmt.Printf(
		promptui.Styler(promptui.FGMagenta)("\"If you pay us ")+
			promptui.Styler(promptui.FGYellow)("%d credits ")+
			promptui.Styler(promptui.FGMagenta)("we will move right along.\"\n"),
		fee,
	)
	fmt.Printf("You have %s credits.\n", promptui.Styler(promptui.FGYellow)(fmt.Sprintf("%d credits", state.Credits)))

	for {
		result := input.Choose("What do you want to do?", loopActionKeys)
		if loopActions.Do(state, result) {
			break
		}
	}

	return false
}
