package syntaxgo_astvtnorm

import "go/ast"

func CountGenericsMap(genericsTypeParams *ast.FieldList) map[string]int {
	var genericsMap = map[string]int{}

	if genericsTypeParams != nil {
		for _, x := range genericsTypeParams.List {
			for _, name := range x.Names {
				genericsMap[name.Name]++
			}
		}
	}

	return genericsMap
}
