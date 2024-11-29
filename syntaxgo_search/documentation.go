package syntaxgo_search

import (
	"go/ast"
	"strings"
)

// GetFunctionDocumentationText extracts the documentation comment of the specified function.
// GetFunctionDocumentationText 提取指定函数的文档注释。
func GetFunctionDocumentationText(funcDecl *ast.FuncDecl) string {
	// Initialize a slice to hold the comment lines.
	// 初始化一个切片，用于存储注释行。
	var commentLines []string

	// If the function has documentation comments, iterate over them.
	// 如果函数有文档注释，遍历这些注释。
	if funcDecl.Doc != nil {
		for _, comment := range funcDecl.Doc.List {
			// Append each comment text to the commentLines slice.
			// 将每一行注释文本添加到commentLines切片中。
			commentLines = append(commentLines, comment.Text)
		}
	}

	// Join the comment lines into a single string and return it.
	// 将注释行连接成一个字符串并返回。
	return strings.Join(commentLines, "\n")
}
