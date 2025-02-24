package inpututils

import (
	"bufio"
	"fmt"
	"os"
)

// GetInput prompts the user and returns input as a string
func GetInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return input[:len(input)-2] // Trim newline character
}
