//Package keyboard captures user input from keyboard
package keyboard

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//UserInput function catures user input from keyboard prompting them with the given msg
func UserInput(msg string) (string, error) {
	//types.String
	fmt.Print(msg, " :")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	// clean up input and see if it is valid grade
	input = strings.TrimSpace(input)
	return input, nil
}


