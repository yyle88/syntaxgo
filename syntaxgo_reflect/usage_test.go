package syntaxgo_reflect

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateTypeUsageCode(t *testing.T) {
	require.Equal(t, "syntaxgo_reflect.Example", GenerateTypeUsageCode(reflect.TypeOf(Example{})))
	require.Equal(t, "syntaxgo_reflect.Example", GenerateTypeUsageCode(GetType(Example{})))
	require.Equal(t, "syntaxgo_reflect.Example", GenerateTypeUsageCode(GetTypeV2[Example]()))
	require.Equal(t, "syntaxgo_reflect.Example", GenerateTypeUsageCode(GetTypeV3(&Example{})))
}
