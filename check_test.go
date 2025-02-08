package nagiosplugin

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheck(t *testing.T) {
	nSpaceMonkeys := float64(200000)
	maxSpaceMonkeys := float64(1 << 32)

	t.Run("Init state", func(t *testing.T) {
		assert := assert.New(t)

		c := NewCheck("CHECK TEST")
		assert.Equal("CHECK TEST UNKNOWN: ", c.String(), "Initial state should be UNKNOWN")
	})

	t.Run("Allow UNKNOWN to CRITICAL", func(t *testing.T) {
		assert := assert.New(t)

		c := NewCheck("CHECK TEST")

		_ = c.AddPerfDatum("space_monkeys", "c", nSpaceMonkeys, 0, maxSpaceMonkeys, 10000, 100000)
		c.AddResult(CRITICAL, fmt.Sprintf("%v terrifying space monkeys in the engineroom", nSpaceMonkeys))

		assert.Equal(
			"CHECK TEST CRITICAL: 200000 terrifying space monkeys in the engineroom | space_monkeys=200000c;10000;100000;0;4294967296",
			c.String(),
		)
	})

	t.Run("Allow UNKNOWN to WARNING", func(t *testing.T) {
		assert := assert.New(t)

		c := NewCheck("CHECK TEST")

		_ = c.AddPerfDatum("space_monkeys", "c", nSpaceMonkeys, 0, maxSpaceMonkeys, 10000, 100000)
		c.AddResult(WARNING, fmt.Sprintf("%v slightly annoying space monkeys in the engineroom", nSpaceMonkeys))

		assert.Equal(
			"CHECK TEST WARNING: 200000 slightly annoying space monkeys in the engineroom | space_monkeys=200000c;10000;100000;0;4294967296",
			c.String(),
		)
	})

	t.Run("Disallow CRITICAL to WARNING", func(t *testing.T) {
		assert := assert.New(t)

		c := NewCheck("CHECK TEST")

		_ = c.AddPerfDatum("space_monkeys", "c", nSpaceMonkeys, 0, maxSpaceMonkeys, 10000, 100000)

		c.AddResult(CRITICAL, fmt.Sprintf("%v terrifying space monkeys in the engineroom", nSpaceMonkeys))
		c.AddResult(WARNING, fmt.Sprintf("%v slightly annoying space monkeys in the engineroom", nSpaceMonkeys))

		assert.Equal(
			"CHECK TEST CRITICAL: 200000 terrifying space monkeys in the engineroom | space_monkeys=200000c;10000;100000;0;4294967296",
			c.String(),
		)
	})

	t.Run("Disallow CRITICAL to UNKNOWN", func(t *testing.T) {
		assert := assert.New(t)

		c := NewCheck("CHECK TEST")

		_ = c.AddPerfDatum("space_monkeys", "c", nSpaceMonkeys, 0, maxSpaceMonkeys, 10000, 100000)

		c.AddResult(CRITICAL, fmt.Sprintf("%v terrifying space monkeys in the engineroom", nSpaceMonkeys))
		c.AddResult(UNKNOWN, fmt.Sprintf("%v unknown annoying space monkeys in the engineroom", nSpaceMonkeys))

		assert.Equal(
			"CHECK TEST CRITICAL: 200000 terrifying space monkeys in the engineroom | space_monkeys=200000c;10000;100000;0;4294967296",
			c.String(),
		)
	})
}
