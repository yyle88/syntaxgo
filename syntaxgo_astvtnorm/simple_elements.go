package syntaxgo_astvtnorm

import (
	"go/ast"
	"strconv"

	"github.com/yyle88/tern"
)

func GetSimpleResElements(astFields []*ast.Field, source []byte) NameTypeElements {
	return GetNameTypeElements(
		astFields,
		SimpleMakeNameFunction("res"),
		source,
		"",
		nil,
	)
}

func GetSimpleArgElements(astFields []*ast.Field, source []byte) NameTypeElements {
	return GetNameTypeElements(
		astFields,
		SimpleMakeNameFunction("arg"),
		source,
		"",
		nil,
	)
}

// SimpleMakeNameFunction 这也是一种方案，即区分是正确还是错误，假如是错误时，就以err开始写名称
func SimpleMakeNameFunction(rightPrefix string) MakeNameFunction {
	return func(name *ast.Ident, kind string, idx int, anonymousIdx int) string {
		return tern.BFF(name != nil && name.Name != "",
			func() string { return name.Name },
			func() string {
				return tern.BFF(kind == "error", func() string {
					return tern.BVV(idx == 0, "err", "err"+strconv.Itoa(idx))
				}, func() string {
					return tern.BVV(idx == 0, rightPrefix, rightPrefix+strconv.Itoa(idx)) //就是正确返回值的前缀+编号
				})
			},
		)
	}
}
