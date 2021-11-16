//Package keyboard captures user input from keyboard
package keyboard

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//UserInput function catures user input from keyboard prompting them with the given msg
// Returns user input as string, error
func UserInput(msg string) (string, error) {
	
	fmt.Print(msg, " :")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	// clean up input
	input = strings.TrimSpace(input)
	return input, nil
}


