package entity

import "errors"

const (
	L                   string = "L"
	R                   string = "R"
	F                   string = "F"
	MaxCharactersLength int    = 100
)

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

	instructionsAllowed := map[string]bool{L: true, R: true, F: true}

	for k, c := range instructions {

		if k > MaxCharactersLength {
			return nil, errors.New("all instruction strings will be less than 100 characters in length")
		}

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
