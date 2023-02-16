package confirmop

import (
	"fmt"
	"strings"
)

// ConfirmOperation prompts the user to confirm the operation and
// returns false if the user enters "n" or "no" and true if the user enters "y" or "yes".
// If the user enters an invalid option and retry is true, the user will be prompted
// until they enter a valid option. If retry is false, the function will return false.
//
// desc is a short description of the operation. Can be empty.
// confirmText is the question that the user should answer. Defaults to "confirm" if empty string is passed.
// retry determines if the user should be prompted repeatedly until they enter a valid option.
//
// example:
// userChoice := ConfirmOperation("Do you want to delete this file","proceed",true)
//
//	if userChoice{
//		//Safe to delete the file
//	}
func ConfirmOperation(desc, confirmText string, retry bool, defaultChoice bool) bool {
	var answer string
	if confirmText == "" {
		confirmText = "confirm"
	}

	// this is to ensure that user is prompted at least once
	var tries int

	for retry || tries == 0 {
		tries++
		if desc == "" {
			fmt.Printf("%s? (y/n): ", confirmText)
		} else {
			fmt.Printf("%s \n%s? (y/n): ", desc, confirmText)

		}

		if _, err := fmt.Scanf("%s", &answer); err != nil {
			if err.Error() == "unexpected newline" {
				return defaultChoice
			}
			fmt.Printf("invalid answer : expected (y or n) got (%s) :\n%v\n", answer, err)
			if !retry {
				return false
			}
			continue
		}
		fmt.Println("answer:", answer)
		answer = strings.TrimSpace(answer)
		answer = strings.ToLower(answer)

		switch answer {
		case "y", "yes":
			return true
		case "n", "no":
			return false
		default:
			fmt.Printf("invalid answer : expected (y or n) got (%s)\n", answer)
			if !retry {
				return false
			}
			continue
		}
	}
	return false
}
