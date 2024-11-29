package syntaxgo_search

import (
	"go/ast"
)

// ExtractFunctionDefinitionCode extracts the code definition of the specified function from the source byte slice.
// ExtractFunctionDefinitionCode 从源字节切片中提取指定函数的代码定义。
func ExtractFunctionDefinitionCode(source []byte, funcDecl *ast.FuncDecl) string {
	// Return the function's code definition as a string.
	// 返回函数的代码定义作为字符串。
	return string(source[funcDecl.Pos()-1 : funcDecl.Body.Lbrace-1])
}

// IsFunctionReceiverName checks if the specified receiver name matches the receiver of the function.
// IsFunctionReceiverName 检查指定的接收者名称是否与函数的接收者匹配。
func IsFunctionReceiverName(funcDecl *ast.FuncDecl, receiverName string) bool {
	// If the function has a receiver, check its name.
	// 如果函数有接收者，检查接收者名称。
	if funcDecl.Recv != nil {
		// Iterate over the list of receiver declarations.
		// 遍历接收者声明列表。
		for _, item := range funcDecl.Recv.List {
			// Check if the receiver is a pointer type.
			// 检查接收者是否为指针类型。
			if starExpr, ok := item.Type.(*ast.StarExpr); ok {
				// If the receiver is a pointer, check if its type matches the receiver name.
				// 如果接收者是指针类型，检查其类型是否匹配接收者名称。
				if ident, ok := starExpr.X.(*ast.Ident); ok && ident.Name == receiverName {
					return true
				}
			} else {
				// If the receiver is not a pointer, check if its type matches the receiver name.
				// 如果接收者不是指针类型，检查其类型是否匹配接收者名称。
				if ident, ok := item.Type.(*ast.Ident); ok && ident.Name == receiverName {
					return true
				}
			}
		}
	}
	// Return false if no match is found.
	// 如果未找到匹配项，返回false。
	return false
}
