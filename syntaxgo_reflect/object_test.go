package syntaxgo_reflect

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetObject(t *testing.T) {
	a := GetObject[int]()
	require.Equal(t, 0, a)
}

func TestNewObject(t *testing.T) {
	p := NewObject[int]()
	require.Equal(t, 0, *p)
}
