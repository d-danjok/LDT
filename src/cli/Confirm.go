package cli

import (
	"fmt"
)

func Confirm(msg string) (bool, error) {
	fmt.Printf("\nDo you want to %s [confirm/deny] \n: ", msg)

	var confirmation string
	_, err := fmt.Scanln(&confirmation)
	if err != nil {
		return false, fmt.Errorf("invalid input")
	}
	switch confirmation {
	case "c", "conf", "confirm":
		fmt.Printf("\nConfirmed\n")
		return true, nil
	case "d", "deny":
		return false, nil
	default:
		return false, fmt.Errorf("%s is not a valid confirmation", confirmation)
	}
}
