package syntaxgo_reflect

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetUsageCode(t *testing.T) {
	require.Equal(t, "syntaxgo_reflect.Example", GetUsageCode(reflect.TypeOf(Example{})))
	require.Equal(t, "syntaxgo_reflect.Example", GetUsageCode(GetType(Example{})))
	require.Equal(t, "syntaxgo_reflect.Example", GetUsageCode(GetTypeV2[Example]()))
	require.Equal(t, "syntaxgo_reflect.Example", GetUsageCode(GetTypeV3(&Example{})))
}
