package leapyears

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsLeapYear(t *testing.T) {

	t.Run("400で割り切れる年はすべて閏年である", func(t *testing.T) {
		assert.True(t, IsLeapYear(1600))
		assert.True(t, IsLeapYear(2000))
	})

	t.Run("100で割り切れるが400で割り切れない年はすべて閏年ではない", func(t *testing.T) {
		assert.False(t, IsLeapYear(1700))
		assert.False(t, IsLeapYear(1800))
		assert.False(t, IsLeapYear(1900))
		assert.False(t, IsLeapYear(2100))
	})

	t.Run("4で割り切れるが100で割り切れない年はすべて閏年である", func(t *testing.T) {
		assert.True(t, IsLeapYear(2008))
		assert.True(t, IsLeapYear(2012))
		assert.True(t, IsLeapYear(2016))
	})

	t.Run("4で割り切れない年はすべて閏年ではない", func(t *testing.T) {
		assert.False(t, IsLeapYear(2017))
		assert.False(t, IsLeapYear(2018))
		assert.False(t, IsLeapYear(2019))
	})

}
