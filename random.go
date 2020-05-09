package er

import (
	"math"
	"math/rand"
)

func RandRune(runeGroup []rune) rune {
	r := RandInt(len(runeGroup) / 2)
	return rune(
		RandInRange(
			int(runeGroup[2*r]),
			int(runeGroup[2*r+1])))
}

func RandAscii() rune {
	return rune(' ' + rand.Intn(0x5f))
}

// [min, max]
func RandInRange(min, max int) int {
	return min + rand.Intn(max-min+1)
}

// TODO how to handle max without boundary
const defaultMax = math.MaxInt8 >> 1

func RandRepeat(min, max int) int {
	if max < 0 {
		max = min + defaultMax
	}
	return RandInRange(min, max)
}

const halfMaxInt16 = math.MaxInt16 >> 1

func RandBool() bool {
	return rand.Intn(math.MaxInt16) > halfMaxInt16
}

func RandInt(max int) int {
	if max <= 0 {
		return 0
	}
	return rand.Intn(max)
}
