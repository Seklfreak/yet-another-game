package gain_exp

import (
	"fmt"
	"strconv"

	"github.com/Seklfreak/yet-another-game/models"
	"github.com/manifoldco/promptui"
)

type Action struct {
}

func (a *Action) Key() string {
	return "Gain EXP"
}

func (a *Action) Do(state *models.State) bool {
	amount, _ := strconv.Atoi(state.ActionContext["gain_exp_amount"])

	currentLevel := state.GetLevel()

	state.Exp += amount

	newLevel := state.GetLevel()

	if newLevel > currentLevel {
		fmt.Printf(promptui.Styler(promptui.FGRed)("Level up! You are now level %d.\n"), newLevel)
	} else {
		fmt.Printf("You gained "+
			promptui.Styler(promptui.FGYellow)("%d EXP")+
			".\n", amount)
	}

	return false
}
