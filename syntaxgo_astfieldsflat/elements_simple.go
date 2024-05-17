package syntaxgo_astfieldsflat

import (
	"go/ast"
	"strconv"
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
		if name != nil && name.Name != "" {
			return name.Name
		} else {
			if kind == "error" {
				if idx == 0 {
					return "err"
				} else {
					return "err" + strconv.Itoa(idx)
				}
			} else {
				if idx == 0 {
					return rightPrefix
				} else {
					return rightPrefix + strconv.Itoa(idx)
				}
			}
		}
	}
}
