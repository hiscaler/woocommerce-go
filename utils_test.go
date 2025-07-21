package woocommerce

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToISOTimeString(t *testing.T) {
	testCases := []struct {
		tag      string
		date     string
		addMin   bool
		addMax   bool
		expected string
	}{
		{"min", "2020-01-01", true, false, "2020-01-01T00:00:00Z"},
		{"has time", "2020-01-01 01:02:03", true, false, "2020-01-01T01:02:03Z"},
		{"bad format", "2020-01-0101:02:03", true, false, "2020-01-0101:02:03"},
	}
	for _, testCase := range testCases {
		s := ToISOTimeString(testCase.date, testCase.addMin, testCase.addMax)
		assert.Equal(t, testCase.expected, s, testCase.tag)
	}
}
