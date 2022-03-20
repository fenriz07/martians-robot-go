package entity

import "errors"

type Orientation uint8

const (
	N Orientation = iota + 1
	E
	S
	W
)

func (orientation Orientation) String() string {
	return [...]string{"N", "E", "S", "W"}[orientation-1]
}

func (orientation Orientation) EnumIndex() uint8 {
	return uint8(orientation)
}

func NewOrientationtByString(orientation string) (Orientation, error) {

	switch orientation {
	case "N":
		return N, nil
	case "S":
		return S, nil
	case "E":
		return E, nil
	case "W":
		return W, nil
	}

	return N, errors.New("can't create a new orientation")
}

func NewOrientationtByInt(orientation uint8) Orientation {

	switch orientation {
	case 1:
		return N
	case 2:
		return E
	case 3:
		return S
	case 4:
		return W
	}

	return N
}
