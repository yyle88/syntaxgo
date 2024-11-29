package syntaxgo_search

import (
	"go/ast"
	"reflect"

	"github.com/yyle88/syntaxgo/internal/utils"
	"github.com/yyle88/syntaxgo/syntaxgo_astnode"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

/*
This package provides tools for handling the Go source code Abstract Syntax Tree (AST), helping developers analyze and extract information from Go source code, such as functions, types, variables, etc. With this tool, developers can more easily search and manipulate the source code.

Key features include:
- Searching for functions, types, and variables.
- Finding specific array types and struct types.
- Returning struct and interface types by name.

It simplifies the process of analyzing Go source code, allowing developers to better understand and manipulate the code.
*/

/*
这个包提供了处理 Go 源代码抽象语法树（AST）的工具，帮助开发者分析和提取 Go 源码中的信息，像是函数、类型、变量等。通过这个工具，开发者可以更轻松地查找和操作源码内容。

主要功能包括：
- 查找函数、类型和变量。
- 查找特定的数组类型、结构体类型。
- 按名称返回结构体类型和接口类型。

简化了 Go 源代码分析的操作，使得开发者可以更好地理解和操作代码。
*/

// FindClassesAndFunctions finds all functions, types, and values in the given AST file.
// FindClassesAndFunctions 查找给定AST文件中的所有函数、类型和变量。
func FindClassesAndFunctions(astFile *ast.File) (functions []*ast.FuncDecl, types []*ast.TypeSpec, values []*ast.ValueSpec) {
	for _, decl := range astFile.Decls {
		switch item := decl.(type) {
		case *ast.FuncDecl:
			// Append function declaration to the functions slice.
			// 将函数声明添加到functions切片中。
			functions = append(functions, item)
		case *ast.GenDecl:
			// Iterate through the type and value specifications.
			// 遍历类型和变量声明。
			for _, spec := range item.Specs {
				switch spec := spec.(type) {
				case *ast.TypeSpec:
					// Append type declaration to the types slice.
					// 将类型声明添加到types切片中。
					types = append(types, spec)
				case *ast.ValueSpec:
					// Append value specification to the values slice.
					// 将变量声明添加到values切片中。
					values = append(values, spec)
				}
			}
		}
	}
	// Return functions, types, and values slices.
	// 返回函数、类型和变量切片。
	return
}

// FindTypes finds all type declarations in the given AST file.
// FindTypes 查找给定AST文件中的所有类型声明。
func FindTypes(astFile *ast.File) (types []*ast.TypeSpec) {
	for _, decl := range astFile.Decls {
		switch item := decl.(type) {
		case *ast.GenDecl:
			// Iterate through type specifications.
			// 遍历类型声明。
			for _, spec := range item.Specs {
				switch spec := spec.(type) {
				case *ast.TypeSpec:
					// Append type declaration to the types slice.
					// 将类型声明添加到types切片中。
					types = append(types, spec)
				}
			}
		}
	}
	// Return types slice.
	// 返回类型切片。
	return types
}

// FindArrayTypeByName finds an array type by its name in the given AST file and returns its content (including brackets).
// FindArrayTypeByName 根据名称查找数组类型，返回结构体内数组的内容（含括号）。
func FindArrayTypeByName(astFile *ast.File, arrayName string) *ast.ArrayType {
	for _, decl := range astFile.Decls {
		switch item := decl.(type) {
		case *ast.GenDecl:
			// Iterate through type specifications in the general declaration.
			// 遍历一般声明中的类型规范。
			for _, spec := range item.Specs {
				switch spec := spec.(type) {
				case *ast.TypeSpec:
					// Check if the name matches the given array name.
					// 检查名称是否与给定的数组名称匹配。
					if spec.Name.Name == arrayName {
						switch typeSpec := spec.Type.(type) {
						case *ast.ArrayType:
							// Return the array type.
							// 返回数组类型。
							return typeSpec
						}
					}
				}
			}
		}
	}
	// Return nil if array type is not found.
	// 如果没有找到数组类型，返回nil。
	return nil
}

// FindStructTypeByName finds a struct type by its name in the given AST file and returns its content (including brackets).
// FindStructTypeByName 根据名称查找结构体类型，返回结构体内的内容（含括号）。
func FindStructTypeByName(astFile *ast.File, structName string) *ast.StructType {
	for _, decl := range astFile.Decls {
		switch item := decl.(type) {
		case *ast.GenDecl:
			// Iterate through type specifications in the general declaration.
			// 遍历一般声明中的类型规范。
			for _, spec := range item.Specs {
				switch spec := spec.(type) {
				case *ast.TypeSpec:
					// Check if the name matches the given struct name.
					// 检查名称是否与给定的结构体名称匹配。
					if spec.Name.Name == structName {
						switch typeSpec := spec.Type.(type) {
						case *ast.StructType:
							// Return the struct type.
							// 返回结构体类型。
							return typeSpec
						}
					}
				}
			}
		}
	}
	// Return nil if struct type is not found.
	// 如果没有找到结构体类型，返回nil。
	return nil
}

// MapStructTypesByName finds all struct types in the given AST file and returns a map of struct names to their corresponding struct content (including brackets).
// MapStructTypesByName 查找给定AST文件中的所有结构体类型，并返回一个结构体名称到其对应结构体内容（含括号）的映射。
func MapStructTypesByName(astFile *ast.File) map[string]*ast.StructType {
	structTypes := map[string]*ast.StructType{}
	for _, decl := range astFile.Decls {
		switch item := decl.(type) {
		case *ast.GenDecl:
			// Iterate through type specifications in the general declaration.
			// 遍历一般声明中的类型规范。
			for _, spec := range item.Specs {
				switch spec := spec.(type) {
				case *ast.TypeSpec:
					name := spec.Name.Name
					switch typeSpec := spec.Type.(type) {
					case *ast.StructType:
						// Add struct name and type to the map.
						// 将结构体名称和类型添加到map中。
						structTypes[name] = typeSpec
					}
				}
			}
		}
	}
	// Return the map of struct names and types.
	// 返回结构体名称和类型的映射。
	return structTypes
}

// MapStructTypeDeclarationsByName finds all struct types in the given AST file and returns a map of struct names to their complete declarations (from type to closing bracket).
// MapStructTypeDeclarationsByName 查找给定AST文件中的所有结构体类型，并返回一个结构体名称到其完整声明（从type开始到闭括号结束）的映射。
func MapStructTypeDeclarationsByName(astFile *ast.File) map[string]*ast.GenDecl {
	structDeclarations := map[string]*ast.GenDecl{}
	for _, decl := range astFile.Decls {
		switch item := decl.(type) {
		case *ast.GenDecl:
			// Iterate through type specifications in the general declaration.
			// 遍历一般声明中的类型规范。
			for _, spec := range item.Specs {
				switch spec := spec.(type) {
				case *ast.TypeSpec:
					name := spec.Name.Name
					switch spec.Type.(type) {
					case *ast.StructType:
						// Add struct name and its declaration to the map.
						// 将结构体名称和其声明添加到map中。
						structDeclarations[name] = item
					}
				}
			}
		}
	}
	// Return the map of struct names and their declarations.
	// 返回结构体名称及其声明的映射。
	return structDeclarations
}

// FindInterfaceTypes finds all interface types in the given AST file and returns a map of interface names to their corresponding interface content.
// FindInterfaceTypes 查找给定AST文件中的所有接口类型，并返回一个接口名称到其对应接口内容的映射。
func FindInterfaceTypes(astFile *ast.File) (interfaceTypesMap map[string]*ast.InterfaceType) {
	interfaceTypesMap = map[string]*ast.InterfaceType{}
	for _, decl := range astFile.Decls {
		switch item := decl.(type) {
		case *ast.GenDecl:
			// Iterate through type specifications in the general declaration.
			// 遍历一般声明中的类型规范。
			for _, spec := range item.Specs {
				switch spec := spec.(type) {
				case *ast.TypeSpec:
					name := spec.Name.String()
					switch typeSpec := spec.Type.(type) {
					case *ast.InterfaceType:
						// Add interface name and type to the map.
						// 将接口名称和类型添加到map中。
						interfaceTypesMap[name] = typeSpec
					}
				}
			}
		}
	}
	// Return the map of interface names and types.
	// 返回接口名称和类型的映射。
	return interfaceTypesMap
}

// FindFunctions finds all function declarations in the given AST file.
// FindFunctions 查找给定AST文件中的所有函数声明。
func FindFunctions(astFile *ast.File) (functions []*ast.FuncDecl) {
	for _, decl := range astFile.Decls {
		switch item := decl.(type) {
		case *ast.FuncDecl:
			// Append function declaration to the functions slice.
			// 将函数声明添加到functions切片中。
			functions = append(functions, item)
		}
	}
	// Return functions slice.
	// 返回函数切片。
	return functions
}

// FindFunctionByName finds a function by its name in the given AST file and returns the function declaration.
// FindFunctionByName 根据名称查找函数，返回该函数的声明。
func FindFunctionByName(astFile *ast.File, functionName string) (function *ast.FuncDecl) {
	// Iterate through all declarations in the AST file
	// 遍历AST文件中的所有声明
	for _, declaration := range astFile.Decls {
		switch item := declaration.(type) {
		case *ast.FuncDecl:
			// If the function name matches, return the function declaration
			// 如果函数名称匹配，返回该函数的声明
			if item.Name.Name == functionName {
				function = item
				return function
			}
		}
	}
	// Return nil if the function is not found
	// 如果未找到函数，返回nil
	return nil
}

// FindMainFunction finds the main function in the given AST file.
// FindMainFunction 查找给定AST文件中的main函数
func FindMainFunction(astFile *ast.File) (mainFunction *ast.FuncDecl) {
	// Use FindFunctionByName to find the main function
	// 使用FindFunctionByName查找main函数
	return FindFunctionByName(astFile, "main")
}

// ExtractFunctions extracts all functions from the given AST file.
// ExtractFunctions 从给定的AST文件中提取所有函数
func ExtractFunctions(astFile *ast.File) (functions []*ast.FuncDecl) {
	// Iterate through all declarations in the AST file
	// 遍历AST文件中的所有声明
	for _, declaration := range astFile.Decls {
		switch item := declaration.(type) {
		case *ast.FuncDecl:
			// Log the function name for debugging
			// 记录函数名称以进行调试
			zaplog.LOG.Debug(item.Name.Name)
			// Append the function to the result list
			// 将函数添加到结果列表中
			functions = append(functions, item)
			continue
		default:
		}
	}
	// Return the list of functions
	// 返回函数列表
	return
}

// FindFunctionByNameWithCheck finds a function by name with a check for its existence in the given AST file.
// FindFunctionByNameWithCheck 查找给定AST文件中指定名称的函数，并检查是否存在
func FindFunctionByNameWithCheck(astFile *ast.File, functionName string) (result *ast.FuncDecl, found bool) {
	// Iterate through all declarations in the AST file
	// 遍历AST文件中的所有声明
	for _, declaration := range astFile.Decls {
		switch item := declaration.(type) {
		case *ast.FuncDecl:
			// Check if the function has no receiver (not a method)
			// 检查函数是否没有接收者（不是方法）
			if item.Recv == nil {
				// If the function name matches, return it with 'found' set to true
				// 如果函数名称匹配，则返回该函数并设置'found'为true
				if item.Name.Name == functionName {
					return item, true
				}
				continue
			}
		default:
			continue
		}
	}
	// Return nil and false if the function is not found
	// 如果未找到函数，则返回nil和false
	return nil, false
}

// FindFunctionsByReceiverName finds all functions with the specified receiver name in the given AST file.
// FindFunctionsByReceiverName 查找给定AST文件中具有指定接收者名称的所有函数
func FindFunctionsByReceiverName(astFile *ast.File, receiverName string, onlyExport bool) (matchingFunctions []*ast.FuncDecl) {
	// Iterate through all declarations in the AST file
	// 遍历AST文件中的所有声明
	for _, declaration := range astFile.Decls {
		switch item := declaration.(type) {
		case *ast.FuncDecl:
			// Check if the function has the specified receiver name
			// 检查函数是否具有指定的接收者名称
			if IsFunctionReceiverName(item, receiverName) {
				// If onlyExport is true, ensure the function name is exported
				// 如果onlyExport为true，确保函数名称是已导出的
				if onlyExport {
					if !utils.C0IsUppercase(item.Name.Name) {
						continue
					}
				}
				// Append the function to the result list
				// 将符合条件的函数添加到结果列表中
				matchingFunctions = append(matchingFunctions, item)
			}
		default:
		}
	}
	// Return the list of matching functions
	// 返回符合条件的函数列表
	return matchingFunctions
}

// FindFunctionByReceiverAndName finds a function by both its receiver name and function name in the given AST file.
// FindFunctionByReceiverAndName 查找给定AST文件中具有指定接收者名称和函数名称的函数
func FindFunctionByReceiverAndName(astFile *ast.File, receiverName string, functionName string) (result *ast.FuncDecl, found bool) {
	// Iterate through all declarations in the AST file
	// 遍历AST文件中的所有声明
	for _, declaration := range astFile.Decls {
		switch item := declaration.(type) {
		case *ast.FuncDecl:
			// Check if the function has the specified receiver name
			// 检查函数是否具有指定的接收者名称
			if IsFunctionReceiverName(item, receiverName) {
				// If the function name matches, return it with 'found' set to true
				// 如果函数名称匹配，则返回该函数并设置'found'为true
				if item.Name.Name == functionName {
					return item, true
				}
				continue
			}
		default:
		}
	}
	// Return nil and false if the function is not found
	// 如果未找到函数，则返回nil和false
	return nil, false
}

// ExtractFunctionsByReceiverName extracts functions with the specified receiver name from the provided list of functions.
// ExtractFunctionsByReceiverName 从提供的函数列表中提取具有指定接收者名称的函数
func ExtractFunctionsByReceiverName(functions []*ast.FuncDecl, receiverName string, onlyExport bool) (filteredFunctions []*ast.FuncDecl) {
	// Iterate through the list of functions
	// 遍历函数列表
	for _, function := range functions {
		// Check if the function has the specified receiver name
		// 检查函数是否具有指定的接收者名称
		if IsFunctionReceiverName(function, receiverName) {
			// If onlyExport is true, ensure the function name is exported
			// 如果onlyExport为true，确保函数名称是已导出的
			if onlyExport {
				if !utils.C0IsUppercase(function.Name.Name) {
					continue
				}
			}
			// Append the function to the result list
			// 将符合条件的函数添加到结果列表中
			filteredFunctions = append(filteredFunctions, function)
			zaplog.LOG.Debug("IS:", zap.String("name", function.Name.Name))
		} else {
			// Log the function name if it doesn't match
			// 如果函数名称不匹配，则记录该函数名称
			zaplog.LOG.Debug("NO:", zap.String("name", function.Name.Name))
		}
	}
	// Return the filtered list of functions
	// 返回过滤后的函数列表
	return filteredFunctions
}

// GetFunctionReceiverNameAndType gets the receiver name and type of the given function declaration.
// GetFunctionReceiverNameAndType 获取给定函数声明的接收者名称和类型
func GetFunctionReceiverNameAndType(astFunc *ast.FuncDecl, source []byte) (receiverName string, receiverType string) {
	// Check if the function has a receiver
	// 检查函数是否具有接收者
	if astFunc.Recv != nil {
		names := astFunc.Recv.List[0].Names
		// If the receiver has a name, assign it
		// 如果接收者有名称，则赋值
		if len(names) > 0 {
			receiverName = names[0].Name
		}
		nodeRecvType := astFunc.Recv.List[0].Type
		// Check the type of the receiver
		// 检查接收者的类型
		switch node := nodeRecvType.(type) {
		case *ast.Ident:
			receiverType = string(syntaxgo_astnode.GetCode(source, node))
		case *ast.StarExpr:
			receiverType = string(syntaxgo_astnode.GetCode(source, node.X))
		default:
			// Log and panic if the receiver type is unknown
			// 如果接收者类型未知，则记录并触发panic
			zaplog.LOG.Panic("unknown", zap.Any("recv_type", reflect.TypeOf(nodeRecvType)), zap.String("node", string(syntaxgo_astnode.GetCode(source, nodeRecvType))))
		}
	}
	// Return receiver name and type
	// 返回接收者名称和类型
	return
}
