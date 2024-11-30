package syntaxgo_aktnorm

import (
	"go/ast"
	"strconv"

	"github.com/yyle88/tern"
)

// NewPrefixedNameTypeElements creates a NameTypeElements instance using a naming function that adds a prefix and an index.
// NewPrefixedNameTypeElements 使用添加前缀和序号的命名函数创建一个 NameTypeElements 实例。
func NewPrefixedNameTypeElements(
	fieldList *ast.FieldList,
	prefix string,
	sourceCode []byte,
	pkgName string,
	genericTypeParams map[string]ast.Expr,
) NameTypeElements {
	// Creates a new NameTypeElements instance with the specified field list and naming function.
	// 使用指定的字段列表和命名函数创建新的 NameTypeElements 实例。
	return NewNameTypeElements(
		fieldList,
		MakePrefixedNameFunction(prefix),
		sourceCode,
		pkgName,
		genericTypeParams,
	)
}

// MakePrefixedNameFunction returns a function that generates unique names with a prefix and index.
// MakePrefixedNameFunction 返回一个通过前缀和序号生成唯一名称的函数。
func MakePrefixedNameFunction(prefix string) MakeNameFunction {
	return func(ident *ast.Ident, kind string, nameIndex int, anonymousIndex int) string {
		// If the identifier has a name, combine it with the prefix and index.
		// 如果标识符有名称，则将其与前缀和序号组合。
		return tern.BFF(ident != nil && ident.Name != "",
			// When the identifier has a name, the function generates a name by concatenating:
			// 1. The specified prefix.
			// 2. The numerical index (nameIndex).
			// 3. The original name of the identifier (ident.Name).
			//
			// 当标识符有名称时，生成的名称由以下部分组成：
			// 1. 指定的前缀。
			// 2. 数字序号 (nameIndex)。
			// 3. 标识符的原始名称 (ident.Name)。
			func() string { return prefix + strconv.Itoa(nameIndex) + ident.Name },
			// If the identifier is anonymous (no name), generate a name using:
			// 1. The specified prefix.
			// 2. The numerical index (nameIndex).
			//
			// 当标识符是匿名的（没有名称）时，生成的名称由以下部分组成：
			// 1. 指定的前缀。
			// 2. 数字序号 (nameIndex)。
			func() string { return prefix + strconv.Itoa(nameIndex) },
		)
	}
}
