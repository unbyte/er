package er

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRand(t *testing.T) {
	a := assert.New(t)
	for i := 0; i < 10; i++ {
		v := RandRepeat(1<<2, 1<<6)
		a.True(v >= (1<<2) && v <= (1<<6), v)
	}

	for i := 0; i < 10; i++ {
		v := RandRepeat(1<<2, -1)
		a.True(v >= (1<<2) && v < (1<<2)+defaultMax, v)
	}

	for i := 0; i < 10; i++ {
		v := RandInRange(1<<2, 1<<6)
		a.True(v >= (1<<2) && v <= (1<<6), v)
	}

	for i := 0; i < 10; i++ {
		v := RandInt(1 << 4)
		a.True(v >= 0 && v < (1<<4), v)
	}

	for i := 0; i < 10; i++ {
		v := RandAscii()
		a.True(v > 0x20 && v < 0x7f, v)
	}

	for i := 0; i < 10; i++ {
		v := RandRune([]rune{'a', 'z'})
		a.True(v >= 'a' && v <= 'z', v)

		v = RandRune([]rune{'A', 'Z'})
		a.True(v >= 'A' && v <= 'Z', v)

		v = RandRune([]rune{'0', '9'})
		a.True(v >= '0' && v <= '9', v)
	}

	same := true
	for i := 0; i < 20; i++ {
		same = same && RandBool()
	}
	a.False(same)

	a.Zero(RandInt(0))
}
