package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRobotIsLost(t *testing.T) {
	t.Parallel()

	type tableRobotIsLost struct {
		Title string
		Robot Robot
		Want  bool
	}

	s := Surface{
		Coordinates: Coordinates{X: 5, Y: 5},
	}

	table := []tableRobotIsLost{
		{
			Title: "X",
			Robot: Robot{
				Coordinates: Coordinates{
					X: 6,
					Y: 1,
				},
			},
			Want: true,
		},
		{
			Title: "-X",
			Robot: Robot{
				Coordinates: Coordinates{
					X: -1,
					Y: 1,
				},
			},
			Want: true,
		},
		{
			Title: "Y",
			Robot: Robot{
				Coordinates: Coordinates{
					X: 2,
					Y: 6,
				},
			},
			Want: true,
		},
		{
			Title: "-Y",
			Robot: Robot{
				Coordinates: Coordinates{
					X: 2,
					Y: -1,
				},
			},
			Want: true,
		},
		{
			Title: "Not lost",
			Robot: Robot{
				Coordinates: Coordinates{
					X: 2,
					Y: 2,
				},
			},
			Want: false,
		},
	}

	for _, test := range table {
		t.Run(test.Title, func(t *testing.T) {
			got := s.RobotIsLost(test.Robot)

			assert.Equal(t, test.Want, got)
		})
	}
}
