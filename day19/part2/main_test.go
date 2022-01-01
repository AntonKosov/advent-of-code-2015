package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestABCD(t *testing.T) {
	actual := process(parseFormula("ABCD"))
	require.Equal(t, 3, actual)
}

func TestOneBracket(t *testing.T) {
	actual := process(parseFormula("ORnFArF"))
	require.Equal(t, 2, actual)
}

func TestNestedBracket(t *testing.T) {
	actual := process(parseFormula("ORnORnFArArF"))
	require.Equal(t, 3, actual)
}

func TestBracketWithMultipleValues(t *testing.T) {
	actual := process(parseFormula("ORnABYBArF"))
	require.Equal(t, 3, actual)
}
