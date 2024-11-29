package syntaxgo_astnode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSdxEdx(t *testing.T) {
	node := NewNode(1, 3)
	sdx, edx := SdxEdx(node)
	t.Log(sdx, edx)
	require.Equal(t, 0, sdx)
	require.Equal(t, 2, edx)
}

func TestGetCode(t *testing.T) {
	node := NewNode(1, 3)
	code := GetCode([]byte("abc"), node)
	t.Log(string(code))
	require.Equal(t, "ab", string(code))
}

func TestGetText(t *testing.T) {
	node := NewNode(1, 3)
	text := GetText([]byte("abc"), node)
	t.Log(text)
	require.Equal(t, "ab", text)
}

func TestDeleteNodeCode(t *testing.T) {
	require.Equal(t, "a", string(DeleteNodeCode([]byte("abc"), NewNode(2, 4))))
	require.Equal(t, "ac", string(DeleteNodeCode([]byte("abc"), NewNode(2, 3))))
	require.Equal(t, "c", string(DeleteNodeCode([]byte("abc"), NewNode(1, 3))))
}

func TestChangeNodeCode(t *testing.T) {
	require.Equal(t, "a123", string(ChangeNodeCode([]byte("abc"), NewNode(2, 4), []byte("123"))))
	require.Equal(t, "a88c", string(ChangeNodeCode([]byte("abc"), NewNode(2, 3), []byte("88"))))
	require.Equal(t, "666c", string(ChangeNodeCode([]byte("abc"), NewNode(1, 3), []byte("666"))))
}

func TestChangeNodeCodeSetSomeNewLines(t *testing.T) {
	require.Equal(t, "a\n123\n", string(ChangeNodeCodeSetSomeNewLines([]byte("abc"), NewNode(2, 4), []byte("123"), 1)))
	require.Equal(t, "a\n\n88\n\nc", string(ChangeNodeCodeSetSomeNewLines([]byte("abc"), NewNode(2, 3), []byte("88"), 2)))
	require.Equal(t, "\n\n\n666\n\n\nc", string(ChangeNodeCodeSetSomeNewLines([]byte("abc"), NewNode(1, 3), []byte("666"), 3)))
}
