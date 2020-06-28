package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Choose ask the user to choose between the items and returns the chosen item
func Choose(description string, items []string) string {
	for {
		fmt.Println(description)
		for i, item := range items {
			fmt.Printf("%d: %s\n", i+1, item)
		}

		fmt.Printf("Please choose a number: ")

		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(strings.Replace(text, "\n", "", -1))

		selection, err := strconv.Atoi(text)
		if err != nil ||
			selection <= 0 ||
			len(items) < selection {
			fmt.Println("invalid input")
			continue
		}

		return items[selection-1]
	}
}
