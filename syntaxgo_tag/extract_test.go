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

	field := ExtractTagField(tmp, "column", EXCLUDE_WHITESPACE_PREFIX)
	t.Log(field)
	require.Equal(t, "name", field)
}

func TestExtractTagField_2(t *testing.T) {
	const tmp = "column: name; primaryKey;"

	field := ExtractTagField(tmp, "column", INCLUDE_WHITESPACE_PREFIX)
	t.Log(field)
	require.Equal(t, " name", field)
}

func TestExtract(t *testing.T) {
	const tag = `gorm:"column:name; primaryKey;" json:"name"`

	value := ExtractTagValue(tag, "gorm")
	t.Log(value)
	require.Equal(t, "column:name; primaryKey;", value)

	field := ExtractTagField(value, "column", INCLUDE_WHITESPACE_PREFIX)
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

	field, sdx, edx := ExtractTagFieldIndex(tmp, "column", EXCLUDE_WHITESPACE_PREFIX)
	t.Log(field, sdx, edx)
	require.Equal(t, "name", field)
	sub := tmp[sdx:edx]
	require.Equal(t, field, sub)
}

func TestExtractTagFieldIndex_2(t *testing.T) {
	const tmp = "column: name; primaryKey;"

	field, sdx, edx := ExtractTagFieldIndex(tmp, "column", INCLUDE_WHITESPACE_PREFIX)
	t.Log(field, sdx, edx)
	require.Equal(t, " name", field)
	sub := tmp[sdx:edx]
	require.Equal(t, field, sub)
}

func TestExtractNoValueTagFieldNameIndex(t *testing.T) {
	const tmp = "column:name;index;"

	sdx, edx := ExtractNoValueTagFieldNameIndex(tmp, "index")
	t.Log(sdx, edx)
	require.Equal(t, "index", tmp[sdx:edx])
}

func TestExtractNoValueTagFieldNameIndex_2(t *testing.T) {
	const tmp = "column:name;index:abc;"

	sdx, edx := ExtractNoValueTagFieldNameIndex(tmp, "index")
	require.Equal(t, -1, sdx)
	require.Equal(t, -1, edx)
}

func TestExtractNoValueTagFieldNameIndex_3(t *testing.T) {
	const tmp = "column:name;index;field2"

	sdx, edx := ExtractNoValueTagFieldNameIndex(tmp, "index")
	t.Log(sdx, edx)
	require.Equal(t, "index", tmp[sdx:edx])
}

func TestExtractNoValueTagFieldNameIndex_4(t *testing.T) {
	const tmp = "column:name;index:value;field2"

	sdx, edx := ExtractNoValueTagFieldNameIndex(tmp, "index")
	require.Equal(t, -1, sdx)
	require.Equal(t, -1, edx)
}

func TestExtractNoValueTagFieldNameIndex_5(t *testing.T) {
	const tmp = "column:name; index;field2"

	sdx, edx := ExtractNoValueTagFieldNameIndex(tmp, "index")
	t.Log(sdx, edx)
	require.Equal(t, "index", tmp[sdx:edx])
}

func TestExtractNoValueTagFieldNameIndex_6(t *testing.T) {
	const tmp = "column:name;index;field2"

	sdx, edx := ExtractNoValueTagFieldNameIndex(tmp, "field2")
	t.Log(sdx, edx)
	require.Equal(t, "field2", tmp[sdx:edx])
}

func TestExtractNoValueTagFieldNameIndex_7(t *testing.T) {
	const tmp = "column:name;index;field2;"

	sdx, edx := ExtractNoValueTagFieldNameIndex(tmp, "field2")
	t.Log(sdx, edx)
	require.Equal(t, "field2", tmp[sdx:edx])
}

func TestExtractNoValueTagFieldNameIndex_8(t *testing.T) {
	const tmp = "column:name;index;field2:value;"

	sdx, edx := ExtractNoValueTagFieldNameIndex(tmp, "field2")
	require.Equal(t, -1, sdx)
	require.Equal(t, -1, edx)
}

func TestExtractNoValueTagFieldNameIndex_9(t *testing.T) {
	const tmp = "column:name;index;field2: value;"

	sdx, edx := ExtractNoValueTagFieldNameIndex(tmp, "field2")
	require.Equal(t, -1, sdx)
	require.Equal(t, -1, edx)
}

func TestExtractNoValueTagFieldNameIndex_10(t *testing.T) {
	const tmp = "column:name;index :value;field2"

	sdx, edx := ExtractNoValueTagFieldNameIndex(tmp, "index")
	require.Equal(t, -1, sdx)
	require.Equal(t, -1, edx)
}
