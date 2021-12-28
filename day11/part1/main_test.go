package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestExample1(t *testing.T) {
	actual := process("abcdefgh")
	require.Equal(t, "abcdffaa", actual)
}

func TestExample2(t *testing.T) {
	actual := process("ghijklmn")
	require.Equal(t, "ghjaabcc", actual)
}

func TestHasPairs(t *testing.T) {
	actual := hasPairs([]rune("abcdffaa"))
	require.True(t, actual)
}
