package move

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Seklfreak/yet-another-game/models"
	"github.com/manifoldco/promptui"
)

type Action struct {
}

func (a *Action) Key() string {
	return "Move"
}

func (a *Action) Do(state *models.State) bool {
	var options []string
	for xI := -1; xI <= 1; xI++ {
		for yI := -1; yI <= 1; yI++ {
			newX := state.PositionX + xI
			newY := state.PositionY + yI
			if newX == state.PositionX && newY == state.PositionY {
				continue
			}

			options = append(options, fmt.Sprintf("X %d Y %d", newX, newY))
		}
	}

	prompt := promptui.Select{
		Label: fmt.Sprintf("You are at X %d Y %d, where do you want to go?", state.PositionX, state.PositionY),
		Items: options,
		Size:  len(options),
	}
	_, result, _ := prompt.Run()

	params := strings.Split(result, " ")
	state.PositionX, _ = strconv.Atoi(params[1])
	state.PositionY, _ = strconv.Atoi(params[3])
	fmt.Printf("moved to X %d Y %d\n", state.PositionX, state.PositionY)

	return false
}
