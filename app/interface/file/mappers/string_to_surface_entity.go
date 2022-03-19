package mappers

import (
	"errors"
	"martians/app/domain/entity"
	"martians/app/share/constants"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
)

const (
	StringToSurfaceEntityNameFunction = "StringToSurfaceEntity"
)

func StringToSurfaceEntity(coordinates string) (*entity.Surface, error) {
	mapCoordinates := strings.Split(coordinates, " ")

	if len(mapCoordinates) != 2 {

		e := errors.New("error mapping string to surface entity")

		log.Error().Err(e).
			Str(constants.Function, StringToSurfaceEntityNameFunction).
			Msg(e.Error())

		return nil, e
	}

	x, err := strconv.Atoi(mapCoordinates[0])

	if err != nil {
		log.Error().Err(err).
			Str(constants.Function, StringToSurfaceEntityNameFunction).
			Msg("error conv str to ing coordinate x")
		return nil, err
	}

	y, err := strconv.Atoi(mapCoordinates[1])

	if err != nil {
		log.Error().Err(err).
			Str(constants.Function, StringToSurfaceEntityNameFunction).
			Msg("error conv str to ing coordinate y")
		return nil, err
	}

	return &entity.Surface{
		Coordinates: entity.Coordinates{
			X: x,
			Y: y,
		},
	}, nil

}
