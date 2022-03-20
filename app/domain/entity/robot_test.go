package entity

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCreateNewInstructions struct {
	Title   string
	Word    string
	WantErr error
}

func TestCreateNewInstructions(t *testing.T) {

	t.Parallel()

	table := []testCreateNewInstructions{
		{
			Title:   "Success",
			Word:    "RFRFRFRF",
			WantErr: nil,
		},
		{
			Title:   "there is a letter is not allowed",
			Word:    "RFRFRFRFJ",
			WantErr: errors.New("there is a letter is not allowed"),
		},
		{
			Title:   "all instruction strings will be less than 100 characters in length",
			Word:    "RFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRFRF",
			WantErr: errors.New("all instruction strings will be less than 100 characters in length"),
		},
	}

	for _, test := range table {
		t.Run(test.Title, func(t *testing.T) {
			_, err := CreateNewInstructions(test.Word)

			assert.Equal(t, err, test.WantErr)
		})
	}

}
