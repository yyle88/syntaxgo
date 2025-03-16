package syntaxgo_aktnorm

import (
	"go/ast"
	"strconv"

	"github.com/yyle88/tern"
)

// GetSimpleResElements extracts NameTypeElements from the given fields with a "res" prefix for names.
// GetSimpleResElements 从给定字段中提取 NameTypeElements，并为名称添加 "res" 前缀。
func GetSimpleResElements(fields []*ast.Field, sourceCode []byte) NameTypeElements {
	return ExtractNameTypeElements(
		fields,
		SimpleMakeNameFunction("res"),
		sourceCode,
		"",
		nil,
	)
}

// GetSimpleArgElements extracts NameTypeElements from the given fields with an "arg" prefix for names.
// GetSimpleArgElements 从给定字段中提取 NameTypeElements，并为名称添加 "arg" 前缀。
func GetSimpleArgElements(fields []*ast.Field, sourceCode []byte) NameTypeElements {
	return ExtractNameTypeElements(
		fields,
		SimpleMakeNameFunction("arg"),
		sourceCode,
		"",
		nil,
	)
}

// SimpleMakeNameFunction returns a function that generates names with a specified prefix, handling both normal and error cases.
// SimpleMakeNameFunction 返回一个函数，该函数生成带有指定前缀的名称，处理正常和错误的情况。
func SimpleMakeNameFunction(prefix string) MakeNameFunction {
	return func(ident *ast.Ident, typeKind string, nameIndex int, anonymousIndex int) string {
		// If the identifier has a name, return the name directly.
		// 如果标识符有名称，则直接返回该名称。
		return tern.BFF(ident != nil && ident.Name != "",
			func() string { return ident.Name },
			// If the identifier does not have a name, generate a name based on the type (error or normal).
			// 如果标识符没有名称，则根据类型（错误或正常）生成名称。
			func() string {
				// If the type is "error", use "err" or "err" followed by the index.
				// 如果类型是 "error"，则使用 "err" 或 "err" 加上索引。
				return tern.BFF(typeKind == "error", func() string {
					return tern.BVV(nameIndex == 0, "err", "err"+strconv.Itoa(nameIndex))
				}, func() string {
					// If it's a normal case, generate a name with the prefix and index.
					// 如果是正常情况，则使用前缀和索引生成名称。
					return tern.BVV(nameIndex == 0, prefix, prefix+strconv.Itoa(nameIndex))
				})
			},
		)
	}
}
