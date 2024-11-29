package syntaxgo_astnode

import (
	"go/ast"

	"github.com/yyle88/syntaxgo/internal/utils"
)

func SdxEdx(astNode ast.Node) (sdx, edx int) {
	sdx = int(astNode.Pos() - 1)
	edx = int(astNode.End() - 1)
	return sdx, edx
}

func GetCode(source []byte, astNode ast.Node) []byte {
	return source[astNode.Pos()-1 : astNode.End()-1]
}

func GetText(source []byte, astNode ast.Node) string {
	return string(GetCode(source, astNode))
}

func DeleteNodeCode(source []byte, astNode ast.Node) []byte {
	return utils.SafeMerge(
		source[:astNode.Pos()-1],
		source[astNode.End()-1:],
	)
}

func ChangeNodeCode(source []byte, astNode ast.Node, newCode []byte) []byte {
	return utils.SafeMerge(
		source[:astNode.Pos()-1],
		newCode,
		source[astNode.End()-1:],
	)
}

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
