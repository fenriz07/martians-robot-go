package entity

import "errors"

type Robot struct {
	Coordinates
	Orientation
	Instructions
	Lost bool
}

type Instructions struct {
	Instructions []string
}

func CreateNewInstructions(instructions string) (*Instructions, error) {

	letters := []string{}

	instructionsAllowed := map[string]bool{"L": true, "R": true, "F": true}

	for _, c := range instructions {

		letter := string(c)

		if instructionsAllowed[letter] {
			letters = append(letters, letter)
			continue
		}

		return nil, errors.New("there is a letter is not allowed")

	}

	instructionsEntity := Instructions{
		Instructions: letters,
	}

	return &instructionsEntity, nil

}
