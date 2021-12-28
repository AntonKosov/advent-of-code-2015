package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestArrayWithoutRed(t *testing.T) {
	actual := process("[1,2,3]")
	require.Equal(t, 6, actual)
}

func TestArrayWithRed(t *testing.T) {
	actual := process(`[1,"red",5]`)
	require.Equal(t, 6, actual)
}

func TestArrayWithSubRed(t *testing.T) {
	actual := process(`[1,{"c":"red","b":2},3]`)
	require.Equal(t, 4, actual)
}

func TestAllIgnored(t *testing.T) {
	actual := process(`{"d":"red","e":[1,2,3,4],"f":5}`)
	require.Equal(t, 0, actual)
}

func TestNestedArray(t *testing.T) {
	actual := process((`[[[3]]]`))
	require.Equal(t, 3, actual)
}
