package syntaxgo_astfieldsflat

import (
	"go/ast"
	"strconv"
)

func NewUsePrefixGetElements(
	astFieldList *ast.FieldList,
	prefix string,
	source []byte,
	packageName string,
	genericsMap map[string]int,
) NameTypeElements {
	return NewNameTypeElements(
		astFieldList,
		UsePrefixMakeNameFunction(prefix),
		source,
		packageName,
		genericsMap,
	)
}

// UsePrefixMakeNameFunction 这是默认的一种方案，即无论如何我都添加前缀和序号，接着再添加原来的名称，这样能确保不重复
func UsePrefixMakeNameFunction(prefix string) MakeNameFunction {
	return func(name *ast.Ident, kind string, idx int, anonymousIdx int) string {
		if name != nil && name.Name != "" {
			return prefix + strconv.Itoa(idx) + name.Name
		} else {
			return prefix + strconv.Itoa(idx)
		}
	}
}
