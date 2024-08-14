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

func TestExtractNoValueFieldNameIndex(t *testing.T) {
	const tmp = "column:name;index;"

	sdx, edx := ExtractNoValueFieldNameIndex(tmp, "index")
	t.Log(sdx, edx)
	require.Equal(t, "index", tmp[sdx:edx])
}

func TestExtractNoValueFieldNameIndex_2(t *testing.T) {
	const tmp = "column:name;index:abc;"

	sdx, edx := ExtractNoValueFieldNameIndex(tmp, "index")
	require.Equal(t, -1, sdx)
	require.Equal(t, -1, edx)
}

func TestExtractNoValueFieldNameIndex_3(t *testing.T) {
	const tmp = "column:name;index;field2"

	sdx, edx := ExtractNoValueFieldNameIndex(tmp, "index")
	t.Log(sdx, edx)
	require.Equal(t, "index", tmp[sdx:edx])
}

func TestExtractNoValueFieldNameIndex_4(t *testing.T) {
	const tmp = "column:name;index:value;field2"

	sdx, edx := ExtractNoValueFieldNameIndex(tmp, "index")
	require.Equal(t, -1, sdx)
	require.Equal(t, -1, edx)
}

func TestExtractNoValueFieldNameIndex_5(t *testing.T) {
	const tmp = "column:name; index;field2"

	sdx, edx := ExtractNoValueFieldNameIndex(tmp, "index")
	t.Log(sdx, edx)
	require.Equal(t, " index", tmp[sdx:edx])
}

func TestExtractNoValueFieldNameIndex_6(t *testing.T) {
	const tmp = "column:name;index;field2"

	sdx, edx := ExtractNoValueFieldNameIndex(tmp, "field2")
	t.Log(sdx, edx)
	require.Equal(t, "field2", tmp[sdx:edx])
}

func TestExtractNoValueFieldNameIndex_7(t *testing.T) {
	const tmp = "column:name;index;field2;"

	sdx, edx := ExtractNoValueFieldNameIndex(tmp, "field2")
	t.Log(sdx, edx)
	require.Equal(t, "field2", tmp[sdx:edx])
}

func TestExtractNoValueFieldNameIndex_8(t *testing.T) {
	const tmp = "column:name;index;field2:value;"

	sdx, edx := ExtractNoValueFieldNameIndex(tmp, "field2")
	require.Equal(t, -1, sdx)
	require.Equal(t, -1, edx)
}

func TestExtractNoValueFieldNameIndex_9(t *testing.T) {
	const tmp = "column:name;index;field2: value;"

	sdx, edx := ExtractNoValueFieldNameIndex(tmp, "field2")
	require.Equal(t, -1, sdx)
	require.Equal(t, -1, edx)
}

func TestExtractNoValueFieldNameIndex_10(t *testing.T) {
	const tmp = "column:name;index :value;field2"

	sdx, edx := ExtractNoValueFieldNameIndex(tmp, "index")
	require.Equal(t, -1, sdx)
	require.Equal(t, -1, edx)
}

func TestExtractNoValueFieldNameIndex_11(t *testing.T) {
	const tmp = "column:name;index;field2 ;"

	sdx, edx := ExtractNoValueFieldNameIndex(tmp, "field2")
	t.Log(sdx, edx)
	require.Equal(t, "field2 ", tmp[sdx:edx])
}

func TestExtractNoValueFieldNameIndex_12(t *testing.T) {
	const tmp = "column:name;index ;field2"

	sdx, edx := ExtractNoValueFieldNameIndex(tmp, "index")
	t.Log(sdx, edx)
	require.Equal(t, "index ", tmp[sdx:edx])
}

func TestExtractNoValueFieldNameIndex_13(t *testing.T) {
	const tmp = "column:name;  index  ;field2"

	sdx, edx := ExtractNoValueFieldNameIndex(tmp, "index")
	t.Log(sdx, edx)
	require.Equal(t, "  index  ", tmp[sdx:edx])
}

func TestExtractNoValueFieldNameIndex_14(t *testing.T) {
	const tmp = "column:name;index;field2:value"

	sdx, edx := ExtractNoValueFieldNameIndex(tmp, "index")
	t.Log(sdx, edx)
	require.Equal(t, "index", tmp[sdx:edx])
}

func TestExtractNoValueFieldNameIndex_15(t *testing.T) {
	const tmp = "column:name;index;field2 "

	sdx, edx := ExtractNoValueFieldNameIndex(tmp, "field2")
	t.Log(sdx, edx)
	require.Equal(t, "field2 ", tmp[sdx:edx])
}

func TestExtractFieldEqualsValueIndex(t *testing.T) {
	const tmp = "column:name;index:idx_myname;field2:value"

	sdx, edx := ExtractFieldEqualsValueIndex(tmp, "index", "idx_myname")
	t.Log(sdx, edx)
	require.Equal(t, "idx_myname", tmp[sdx:edx])
}

func TestExtractFieldEqualsValueIndex_2(t *testing.T) {
	const tmp = "column:name;index:idx_myname ;field2:value"

	sdx, edx := ExtractFieldEqualsValueIndex(tmp, "index", "idx_myname")
	t.Log(sdx, edx)
	require.Equal(t, "idx_myname ", tmp[sdx:edx])
}

func TestExtractFieldEqualsValueIndex_3(t *testing.T) {
	const tmp = "column:name;index: idx_myname;field2:value"

	sdx, edx := ExtractFieldEqualsValueIndex(tmp, "index", "idx_myname")
	t.Log(sdx, edx)
	require.Equal(t, " idx_myname", tmp[sdx:edx])
}

func TestExtractFieldEqualsValueIndex_4(t *testing.T) {
	const tmp = "column:name;index: idx_myname;index: idx_myname2;"

	sdx, edx := ExtractFieldEqualsValueIndex(tmp, "index", "idx_myname2")
	t.Log(sdx, edx)
	require.Equal(t, " idx_myname2", tmp[sdx:edx])
}

func TestExtractFieldEqualsValueIndex_5(t *testing.T) {
	const tmp = "column:name;index: idx_myname2;index: idx_myname;"

	sdx, edx := ExtractFieldEqualsValueIndex(tmp, "index", "idx_myname")
	t.Log(sdx, edx)
	require.Equal(t, " idx_myname", tmp[sdx:edx])
}

func TestExtractFieldEqualsValueIndexV2(t *testing.T) {
	const tmp = "column:name;index: idx_myname;index: idx_myname2,priority:2;"
	//不仅以分号或结尾为分割，还以逗号为分隔符
	sdx, edx := ExtractFieldEqualsValueIndexV2(tmp, "index", "idx_myname2", []string{","})
	t.Log(sdx, edx)
	require.Equal(t, " idx_myname2", tmp[sdx:edx])
}
