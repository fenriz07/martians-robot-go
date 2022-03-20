package usecase

import (
	"martians/app/domain/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecuteSequenceUseCase(t *testing.T) {
	t.Parallel()

	type TestExecuteSequenceUseCase struct {
		Title string
		Input []entity.Robot
		Want  []entity.Robot
	}

	table := []TestExecuteSequenceUseCase{
		{
			Title: "Is lost for default",
			Input: []entity.Robot{
				{
					Coordinates: entity.Coordinates{X: 10, Y: 10},
					Orientation: entity.N,
					Instructions: entity.Instructions{
						Instructions: []string{"F"},
					},
					Lost: false,
				},
			},
			Want: []entity.Robot{
				{
					Coordinates: entity.Coordinates{X: 10, Y: 10},
					Orientation: entity.N,
					Instructions: entity.Instructions{
						Instructions: []string{"F"},
					},
					Lost: true,
				},
			},
		},
		{
			Title: "Success Move",
			Input: []entity.Robot{
				{
					Coordinates: entity.Coordinates{X: 1, Y: 1},
					Orientation: entity.E,
					Instructions: entity.Instructions{
						Instructions: []string{"R", "F", "R", "F", "R", "F", "R", "F"},
					},
					Lost: false,
				},
			},
			Want: []entity.Robot{
				{
					Coordinates: entity.Coordinates{X: 1, Y: 1},
					Orientation: entity.E,
					Instructions: entity.Instructions{
						Instructions: []string{"R", "F", "R", "F", "R", "F", "R", "F"},
					},
					Lost: false,
				},
			},
		},
		{
			Title: "Lost Robot",
			Input: []entity.Robot{
				{
					Coordinates: entity.Coordinates{X: 3, Y: 2},
					Orientation: entity.N,
					Instructions: entity.Instructions{
						Instructions: []string{"F", "R", "R", "F", "L", "L", "F", "F", "R", "R", "F", "L", "L"},
					},
					Lost: false,
				},
			},
			Want: []entity.Robot{
				{
					Coordinates: entity.Coordinates{X: 3, Y: 3},
					Orientation: entity.N,
					Instructions: entity.Instructions{
						Instructions: []string{"F", "R", "R", "F", "L", "L", "F", "F", "R", "R", "F", "L", "L"},
					},
					Lost: true,
				},
			},
		},
		{
			Title: "Complete use case",
			Input: []entity.Robot{
				{
					Coordinates: entity.Coordinates{X: 1, Y: 1},
					Orientation: entity.E,
					Instructions: entity.Instructions{
						Instructions: []string{"R", "F", "R", "F", "R", "F", "R", "F"},
					},
					Lost: false,
				},
				{
					Coordinates: entity.Coordinates{X: 3, Y: 2},
					Orientation: entity.N,
					Instructions: entity.Instructions{
						Instructions: []string{"F", "R", "R", "F", "L", "L", "F", "F", "R", "R", "F", "L", "L"},
					},
					Lost: false,
				},
				{
					Coordinates: entity.Coordinates{X: 0, Y: 3},
					Orientation: entity.W,
					Instructions: entity.Instructions{
						Instructions: []string{"L", "L", "F", "F", "F", "L", "F", "L", "F", "L"},
					},
					Lost: false,
				},
			},
			Want: []entity.Robot{
				{
					Coordinates: entity.Coordinates{X: 1, Y: 1},
					Orientation: entity.E,
					Instructions: entity.Instructions{
						Instructions: []string{"R", "F", "R", "F", "R", "F", "R", "F"},
					},
					Lost: false,
				},
				{
					Coordinates: entity.Coordinates{X: 3, Y: 3},
					Orientation: entity.N,
					Instructions: entity.Instructions{
						Instructions: []string{"F", "R", "R", "F", "L", "L", "F", "F", "R", "R", "F", "L", "L"},
					},
					Lost: true,
				},
				{
					Coordinates: entity.Coordinates{X: 2, Y: 3},
					Orientation: entity.S,
					Instructions: entity.Instructions{
						Instructions: []string{"L", "L", "F", "F", "F", "L", "F", "L", "F", "L"},
					},
					Lost: false,
				},
			},
		},
	}

	useCase := ExecuteSequenceUseCase{}

	surface := entity.Surface{
		Coordinates: entity.Coordinates{
			X: 5,
			Y: 3,
		},
		LostCoordinates: []entity.Coordinates{},
	}

	for _, test := range table {
		t.Run(test.Title, func(t *testing.T) {
			got := useCase.Launch(test.Input, surface)

			assert.Equal(t, test.Want, got)
		})
	}
}
