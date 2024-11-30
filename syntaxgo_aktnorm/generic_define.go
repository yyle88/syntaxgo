package syntaxgo_aktnorm

import "go/ast"

// GetFuncGenericTypeParamsMap extracts a map of custom generic type parameter names and their associated types from a function declaration's type parameters.
// GetFuncGenericTypeParamsMap 从函数声明 (FuncDecl) 的泛型类型参数列表中提取自定义泛型类型参数名称及其关联的类型，返回一个以名称为键、类型表达式 (ast.Expr) 为值的映射。
func GetFuncGenericTypeParamsMap(funcDecl *ast.FuncDecl) map[string]ast.Expr {
	if funcDecl == nil || funcDecl.Type == nil {
		return map[string]ast.Expr{}
	}
	return GetGenericTypeParamsMap(funcDecl.Type.TypeParams)
}

// GetGenericFuncTypeParamsMap extracts a map of custom generic type parameter names and their associated types from a function type's type parameters.
// GetGenericFuncTypeParamsMap 从函数类型 (FuncType) 的泛型类型参数列表中提取自定义泛型类型参数名称及其关联的类型，返回一个以名称为键、类型表达式 (ast.Expr) 为值的映射。
func GetGenericFuncTypeParamsMap(funcType *ast.FuncType) map[string]ast.Expr {
	if funcType == nil {
		return map[string]ast.Expr{}
	}
	return GetGenericTypeParamsMap(funcType.TypeParams)
}

// GetGenericTypeParamsMap extracts a map of custom generic type parameter names and their associated types from a type parameter field list.
// GetGenericTypeParamsMap 从泛型类型参数字段列表 (FieldList) 中提取自定义泛型类型参数名称及其关联的类型，返回一个以名称为键、类型表达式 (ast.Expr) 为值的映射。
func GetGenericTypeParamsMap(fields *ast.FieldList) map[string]ast.Expr {
	nameMap := make(map[string]ast.Expr)
	if fields != nil {
		for _, field := range fields.List {
			for _, name := range field.Names {
				// The key is the name of the custom generic type parameter, and the value is its type expression.
				// 键是自定义泛型类型参数的名称，值是其类型表达式类型。
				nameMap[name.Name] = field.Type
			}
		}
	}
	return nameMap
}
