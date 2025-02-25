package syntaxgo_astnode

import (
	"go/ast"
	"go/token"
)

// Node represents a syntax node with position information.
// Node 表示一个带有位置信息的语法节点。
type Node struct {
	pos token.Pos // position of first character belonging to the node // 节点第一个字符的位置
	end token.Pos // position of first character immediately after the node // 节点结束后第一个字符的位置
}

// NewNode creates a new Node with the specified start and end positions.
// NewNode 创建一个带有指定起始和结束位置的 Node。
func NewNode(pos, end token.Pos) *Node {
	return &Node{pos: pos, end: end}
}

// NewNodeV1 creates a new Node from an existing ast.Node.
// NewNodeV1 从一个现有的 ast.Node 创建一个新的 Node。
func NewNodeV1(node ast.Node) *Node {
	return NewNode(node.Pos(), node.End())
}

// NewNodeV2 creates a new Node with the specified start and end positions as integers.
// NewNodeV2 创建一个带有指定起始和结束位置（整数）的 Node。
func NewNodeV2(pos, end int) *Node {
	return NewNode(token.Pos(pos), token.Pos(end))
}

// Pos returns the start position of the Node.
// Pos 返回 Node 的起始位置。
func (x *Node) Pos() token.Pos {
	return x.pos
}

// End returns the end position of the Node.
// End 返回 Node 的结束位置。
func (x *Node) End() token.Pos {
	return x.end
}

// SetPos sets the start position of the Node.
// SetPos 设置 Node 的起始位置。
func (x *Node) SetPos(pos token.Pos) {
	x.pos = pos
}

// SetEnd sets the end position of the Node.
// SetEnd 设置 Node 的结束位置。
func (x *Node) SetEnd(end token.Pos) {
	x.end = end
}

// GetCode returns the code corresponding to the Node from the source.
// GetCode 从源代码中返回与 Node 对应的代码。
func (x *Node) GetCode(source []byte) []byte {
	return source[x.Pos()-1 : x.End()-1]
}

// GetText returns the text corresponding to the Node from the source.
// GetText 从源代码中返回与 Node 对应的文本。
func (x *Node) GetText(source []byte) string {
	return string(source[x.Pos()-1 : x.End()-1])
}
