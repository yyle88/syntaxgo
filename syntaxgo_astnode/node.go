package syntaxgo_astnode

import (
	"go/ast"
	"go/token"
)

type Node struct {
	pos token.Pos // position of first character belonging to the node
	end token.Pos // position of first character immediately after the node
}

func NewNode(pos, end token.Pos) *Node {
	return &Node{pos: pos, end: end}
}

func NewNodeV1(node ast.Node) *Node {
	return NewNode(node.Pos(), node.End())
}
func NewNodeV2(pos, end int) *Node {
	return NewNode(token.Pos(pos), token.Pos(end))
}

func (x *Node) Pos() token.Pos {
	return x.pos
}

func (x *Node) End() token.Pos {
	return x.end
}

func (x *Node) GetCode(source []byte) []byte {
	return source[x.Pos()-1 : x.End()-1]
}

func (x *Node) GetText(source []byte) string {
	return string(source[x.Pos()-1 : x.End()-1])
}
