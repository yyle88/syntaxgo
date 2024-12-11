package syntaxgo_search

import (
	"go/ast"
	"strings"
)

// GetFunctionComment extracts the documentation comment of the specified function.
// GetFunctionComment 提取指定函数的文档注释。
func GetFunctionComment(funcDecl *ast.FuncDecl) string {
	var commentLines []string
	if funcDecl.Doc != nil {
		for _, comment := range funcDecl.Doc.List {
			commentLines = append(commentLines, comment.Text)
		}
	}
	return strings.Join(commentLines, "\n")
}
