package file

import (
	"martians/app/domain/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStartInstructions(t *testing.T) {
	t.Parallel()

	type testStartInstructions struct {
		Title         string
		FileUbication string
		WantErr       error
	}

	table := []testStartInstructions{
		{
			Title:         ErrReadingToFile.Error(),
			FileUbication: "anywhere",
			WantErr:       ErrReadingToFile,
		},
		{
			Title:         ErrInstructionNotCompleted.Error(),
			FileUbication: "../../../test/interface/file/less_that_3",
			WantErr:       ErrInstructionNotCompleted,
		},
		{
			Title:         ErrAnInstructionIsMissing.Error(),
			FileUbication: "../../../test/interface/file/pair",
			WantErr:       ErrAnInstructionIsMissing,
		},
		{
			Title:         ErrMappingStringToSurfaceEntity.Error(),
			FileUbication: "../../../test/interface/file/err_surface",
			WantErr:       ErrMappingStringToSurfaceEntity,
		},
		{
			Title:         ErrMappingRobot.Error(),
			FileUbication: "../../../test/interface/file/err_robot",
			WantErr:       ErrMappingRobot,
		},
		{
			Title:         "Success",
			FileUbication: "../../../test/interface/file/success",
			WantErr:       nil,
		},
	}

	i := NewInputInterface(func(robots []entity.Robot, surface entity.Surface) []entity.Robot {
		return []entity.Robot{
			{
				Coordinates: entity.Coordinates{X: 1, Y: 1},
				Orientation: entity.N,
				Lost:        true,
			},
			{
				Coordinates: entity.Coordinates{X: 2, Y: 1},
				Orientation: entity.N,
				Lost:        false,
			},
		}
	})

	for _, test := range table {

		t.Run(test.Title, func(t *testing.T) {
			gotErr := i.StartInstructions(test.FileUbication)

			assert.Equal(t, test.WantErr, gotErr)
		})
	}
}
