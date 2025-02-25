package syntaxgo_astnode

import (
	"go/ast"

	"github.com/yyle88/syntaxgo/internal/utils"
)

// SdxEdx returns the start and end positions of the given AST node as integers.
// SdxEdx 返回给定 AST 节点的起始和结束位置（整数）。
func SdxEdx(astNode ast.Node) (sdx, edx int) {
	sdx = int(astNode.Pos() - 1)
	edx = int(astNode.End() - 1)
	return sdx, edx
}

// GetCode returns the code corresponding to the given AST node from the source.
// GetCode 从源代码中返回与给定 AST 节点对应的代码。
func GetCode(source []byte, astNode ast.Node) []byte {
	return source[astNode.Pos()-1 : astNode.End()-1]
}

// GetText returns the text corresponding to the given AST node from the source.
// GetText 从源代码中返回与给定 AST 节点对应的文本。
func GetText(source []byte, astNode ast.Node) string {
	return string(GetCode(source, astNode))
}

// DeleteNodeCode removes the code corresponding to the given AST node from the source.
// DeleteNodeCode 从源代码中删除与给定 AST 节点对应的代码。
func DeleteNodeCode(source []byte, astNode ast.Node) []byte {
	return utils.SafeMerge(
		source[:astNode.Pos()-1],
		source[astNode.End()-1:],
	)
}

// ChangeNodeCode replaces the code corresponding to the given AST node with new code in the source.
// ChangeNodeCode 用新代码替换源代码中与给定 AST 节点对应的代码。
func ChangeNodeCode(source []byte, astNode ast.Node, newCode []byte) []byte {
	return utils.SafeMerge(
		source[:astNode.Pos()-1],
		newCode,
		source[astNode.End()-1:],
	)
}

// ChangeNodeCodeSetSomeNewLines replaces the code corresponding to the given AST node with new code and adds new lines in the source.
// ChangeNodeCodeSetSomeNewLines 用新代码替换源代码中与给定 AST 节点对应的代码，并添加新行。
func ChangeNodeCodeSetSomeNewLines(source []byte, astNode ast.Node, newCode []byte, numLine int) []byte {
	var newLines = make([]byte, 0, numLine)
	for idx := 0; idx < numLine; idx++ {
		newLines = append(newLines, '\n')
	}
	return utils.SafeMerge(
		source[:astNode.Pos()-1],
		newLines,
		newCode,
		newLines,
		source[astNode.End()-1:],
	)
}
