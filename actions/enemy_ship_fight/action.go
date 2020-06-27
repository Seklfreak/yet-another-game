package enemy_ship_fight

import (
	"fmt"
	"math/rand"

	"github.com/Seklfreak/yet-another-game/models"
)

type Action struct {
}

func (a *Action) Key() string {
	return "Fight"
}

func (a *Action) Do(state *models.State) bool {
	damage := rand.Intn(50) + 1
	lootCredits := rand.Intn(100) + 1

	state.Health -= damage

	if state.Health < 0 {
		fmt.Println("You did not survive the fight, game over!")
		fmt.Println("Thanks for playing.")
		return true
	}

	state.Credits += lootCredits

	fmt.Printf("You destroyed the other ship and took %d damage. You have %d health left.\n",
		damage,
		state.Health,
	)
	fmt.Printf("After examining the ship you found %d credits. Now you have %d credits.\n",
		lootCredits,
		state.Credits,
	)
	return true
}
