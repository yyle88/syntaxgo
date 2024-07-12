package syntaxgo_tag

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExtractTagValue(t *testing.T) {
	const tag = `gorm:"column:name; primaryKey;" json:"name"`

	value := ExtractTagValue(tag, "gorm")
	t.Log(value)
	require.Equal(t, "column:name; primaryKey;", value)
}

func TestExtractTagField(t *testing.T) {
	const tmp = "column:name; primaryKey;"

	field := ExtractTagField(tmp, "column", true)
	t.Log(field)
	require.Equal(t, "name", field)
}

func TestExtractTagField_2(t *testing.T) {
	const tmp = "column: name; primaryKey;"

	field := ExtractTagField(tmp, "column", false)
	t.Log(field)
	require.Equal(t, " name", field)
}

func TestExtract(t *testing.T) {
	const tag = `gorm:"column:name; primaryKey;" json:"name"`

	value := ExtractTagValue(tag, "gorm")
	t.Log(value)
	require.Equal(t, "column:name; primaryKey;", value)

	field := ExtractTagField(value, "column", false)
	t.Log(field)
	require.Equal(t, "name", field)
}

func TestExtractTagValueIndex(t *testing.T) {
	const tag = `gorm:"column:name; primaryKey;" json:"name"`

	value, sdx, edx := ExtractTagValueIndex(tag, "gorm")
	t.Log(value, sdx, edx)
	require.Equal(t, "column:name; primaryKey;", value)
	sub := tag[sdx:edx]
	require.Equal(t, value, sub)
}

func TestExtractTagFieldIndex(t *testing.T) {
	const tmp = "column:name; primaryKey;"

	field, sdx, edx := ExtractTagFieldIndex(tmp, "column", true)
	t.Log(field, sdx, edx)
	require.Equal(t, "name", field)
	sub := tmp[sdx:edx]
	require.Equal(t, field, sub)
}

func TestExtractTagFieldIndex_2(t *testing.T) {
	const tmp = "column: name; primaryKey;"

	field, sdx, edx := ExtractTagFieldIndex(tmp, "column", false)
	t.Log(field, sdx, edx)
	require.Equal(t, " name", field)
	sub := tmp[sdx:edx]
	require.Equal(t, field, sub)
}
