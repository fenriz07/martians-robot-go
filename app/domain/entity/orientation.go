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
