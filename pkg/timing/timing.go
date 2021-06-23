package timing

import (
	"math"
)

func Dit(wpm float32) int {

	msPerDit := int(math.Round(float64(60/(50*wpm)) * 1000))

	return msPerDit
}

func Dah(wpm float32) int {

	msPerDah := Dit(wpm) * 3

	return msPerDah
}

func CharacterSpace(wpm float32) int {

	msPerSpace := Dah(wpm)

	return msPerSpace
}

func WordSpace(wpm float32) int {

	msPerSpace := Dit(wpm) * 7

	return msPerSpace
}

func CharacterSegmentSpace(wpm float32) int {

	return Dit(wpm)
}
