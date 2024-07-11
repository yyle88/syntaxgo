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

	field := ExtractTagField(tmp, "column")
	t.Log(field)
	require.Equal(t, "name", field)
}

func TestExtract(t *testing.T) {
	const tag = `gorm:"column:name; primaryKey;" json:"name"`

	value := ExtractTagValue(tag, "gorm")
	t.Log(value)
	require.Equal(t, "column:name; primaryKey;", value)

	field := ExtractTagField(value, "column")
	t.Log(field)
	require.Equal(t, "name", field)
}
