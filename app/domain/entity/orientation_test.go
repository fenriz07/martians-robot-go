package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrientationString(t *testing.T) {
	want := "N"
	assert.Equal(t, want, N.String())
}

func TestOrientationEnumIndex(t *testing.T) {
	var want uint8 = 1
	assert.Equal(t, want, N.EnumIndex())
}

func TestNewOrientationtByString(t *testing.T) {
	t.Parallel()

	type testNewOrientationtByString struct {
		Letter string
		Want   Orientation
	}

	table := []testNewOrientationtByString{
		{
			Letter: "N",
			Want:   N,
		},
		{
			Letter: "E",
			Want:   E,
		},
		{
			Letter: "S",
			Want:   S,
		},
		{
			Letter: "W",
			Want:   W,
		},
		{
			Letter: "K",
			Want:   N,
		},
	}

	for _, test := range table {
		t.Run(test.Letter, func(t *testing.T) {
			got, _ := NewOrientationtByString(test.Letter)

			assert.Equal(t, test.Want, got)
		})
	}
}

func TestNewOrientationtByInt(t *testing.T) {
	t.Parallel()

	type testNewOrientationtByInt struct {
		Title string
		Index uint8
		Want  Orientation
	}

	table := []testNewOrientationtByInt{
		{
			Title: "N",
			Index: 1,
			Want:  N,
		},
		{
			Title: "E",
			Index: 2,
			Want:  E,
		},
		{
			Title: "S",
			Index: 3,
			Want:  S,
		},
		{
			Title: "W",
			Index: 4,
			Want:  W,
		},
		{
			Title: "BAD",
			Index: 5,
			Want:  N,
		},
	}

	for _, test := range table {
		t.Run(test.Title, func(t *testing.T) {
			got := NewOrientationtByInt(test.Index)

			assert.Equal(t, test.Want, got)
		})
	}
}
