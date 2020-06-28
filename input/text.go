package input

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Text waits for an user input of a text and returns the text
func Text(description string) string {
	for {
		fmt.Print(description)

		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(strings.Replace(text, "\n", "", -1))

		if len(text) <= 0 {
			fmt.Println("text too short")
			continue
		}

		return text
	}
}
