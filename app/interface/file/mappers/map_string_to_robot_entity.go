package mappers

import (
	"errors"
	"martians/app/domain/entity"
	"strconv"
	"strings"
)

func MapStringToRobotEntity(robotWithInstruction [][]string) ([]entity.Robot, error) {

	robots := []entity.Robot{}

	for _, l := range robotWithInstruction {

		c, o, err := stringToCoordinatesAndOrientation(l[0])

		if err != nil {
			return nil, err
		}

		i, err := entity.CreateNewInstructions(l[1])

		if err != nil {
			return nil, err
		}

		robot := entity.Robot{
			Coordinates:  *c,
			Orientation:  *o,
			Instructions: *i,
			Lost:         false,
		}

		robots = append(robots, robot)
	}

	return robots, nil
}

func stringToCoordinatesAndOrientation(line string) (*entity.Coordinates, *entity.Orientation, error) {
	mapCoordinatesAndOrientation := strings.Split(line, " ")

	if len(mapCoordinatesAndOrientation) != 3 {
		return nil, nil, errors.New("error in the line to robot details")
	}

	x, err := strconv.Atoi(mapCoordinatesAndOrientation[0])
	if err != nil {
		return nil, nil, err
	}

	y, err := strconv.Atoi(mapCoordinatesAndOrientation[1])
	if err != nil {
		return nil, nil, err
	}

	orientation, err := entity.NewOrientationtByString(mapCoordinatesAndOrientation[2])

	if err != nil {
		return nil, nil, err
	}

	return &entity.Coordinates{X: x, Y: y}, &orientation, nil

}
