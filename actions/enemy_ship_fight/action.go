package enemy_ship_fight

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/Seklfreak/yet-another-game/actions/gain_exp"
	"github.com/Seklfreak/yet-another-game/models"
	"github.com/manifoldco/promptui"
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
		fmt.Println(promptui.Styler(promptui.FGRed)("You did not survive the fight, game over!"))
		fmt.Println(promptui.Styler(promptui.FGRed)("Thanks for playing."))
		return true
	}

	state.Credits += lootCredits

	fmt.Printf(promptui.Styler(promptui.FGRed)("You destroyed the other ship and took "+
		promptui.Styler(promptui.FGYellow)("%d damage")+
		promptui.Styler(promptui.FGRed)(". You have ")+
		promptui.Styler(promptui.FGYellow)("%d health ")+
		promptui.Styler(promptui.FGRed)("left.\n")),
		damage,
		state.Health,
	)
	fmt.Printf("After examining the ship you found "+
		promptui.Styler(promptui.FGYellow)("%d credits")+
		". Now you have "+
		promptui.Styler(promptui.FGYellow)("%d credits")+
		".\n",
		lootCredits,
		state.Credits,
	)

	state.ActionContext["gain_exp_amount"] = strconv.Itoa(exp)
	(&gain_exp.Action{}).Do(state)

	return true
}
