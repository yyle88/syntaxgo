package syntaxgo_astnode

import (
	"go/ast"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewNode(t *testing.T) {
	var _ ast.Node = NewNode(1, 100)
	var _ ast.Node = NewNodeV1(NewNode(1, 100))
}

func TestNode_GetCode(t *testing.T) {
	node := NewNode(1, 3)
	code := node.GetCode([]byte("abc"))
	t.Log(string(code))
	require.Equal(t, "ab", string(code))
}

func TestNode_GetText(t *testing.T) {
	node := NewNode(1, 3)
	text := node.GetText([]byte("abc"))
	t.Log(text)
	require.Equal(t, "ab", text)
}
