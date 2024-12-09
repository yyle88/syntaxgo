package syntaxgo_aktnorm

import (
	"go/ast"
	"strings"

	"github.com/yyle88/must"
	"github.com/yyle88/syntaxgo/internal/utils"
	"github.com/yyle88/syntaxgo/syntaxgo_astnode"
)

// NameTypeElement represents a single element with a name, type, and associated information.
// NameTypeElement 代表一个包含名称、类型和相关信息的元素
type NameTypeElement struct {
	Name       string   // Field name / 字段名称
	Kind       string   // Type description (e.g., int, string, A, utils.A, *B, *pkg.B) / 类型描述 (例如: int, string, A, utils.A, *B, *pkg.B)
	Type       ast.Expr // Go's internal type representation / Go 的内部类型表示
	IsEllipsis bool     // Whether the type is a variadic (e.g., ...int or ...string) / 是否是变参类型 (例如: ...int 或 ...string)
}

// NewNameTypeElement creates a new NameTypeElement by parsing the source code for parameter/return information and normalizing the data for further code generation.
// NewNameTypeElement 通过解析源代码中的参数/返回信息并规范化数据来创建一个新的 NameTypeElement，以供进一步的代码生成。
func NewNameTypeElement(
	field *ast.Field, // AST field representing the parameter/return / AST 字段表示参数/返回值
	paramName string, // Name of the parameter / 参数名称
	paramType string, // Type of the parameter / 参数类型
	isEllipsis bool, // Whether the parameter is a variadic / 是否为变参
	packageName string, // The package name for external types / 外部类型的包名
	genericTypeParams map[string]ast.Expr, // Map of generic type parameters / 泛型类型参数的映射
) *NameTypeElement {
	elem := &NameTypeElement{
		Name:       paramName,
		Kind:       paramType,
		Type:       field.Type,
		IsEllipsis: isEllipsis,
	}
	if packageName != "" {
		elem.AdjustTypeWithPackage(packageName, genericTypeParams)
	}
	return elem
}

// AdjustTypeWithPackage adjusts the type to include the package name if needed (for external types).
// AdjustTypeWithPackage 如果需要（对于外部类型），将类型调整为包含包名。
func (element *NameTypeElement) AdjustTypeWithPackage(
	packageName string, // The package name / 包名
	genericTypeParams map[string]ast.Expr, // Map of generic type parameters / 泛型类型参数的映射
) {
	if strings.Contains(element.Kind, ".") {
		return // Already includes package name / 已经包含包名
	}

	shortKind := strings.TrimSpace(element.Kind)

	if element.IsEllipsis {
		must.True(strings.HasPrefix(shortKind, "..."))
		shortKind = strings.TrimPrefix(shortKind, "...")
		shortKind = strings.TrimSpace(shortKind)
	}

	if s := shortKind[0]; string(s) == "*" {
		if vcc := shortKind[1]; vcc >= 'A' && vcc <= 'Z' {
			className := shortKind[1:]
			if _, ok := genericTypeParams[className]; ok {
				return // It's a generic type / 是泛型类型
			}
			shortKind = "*" + packageName + "." + className
		} else {
			return // basic-type(int string float64) || // not-exportable-type
		}
	} else {
		if s := shortKind[0]; s >= 'A' && s <= 'Z' {
			className := shortKind
			if _, ok := genericTypeParams[className]; ok {
				return // It's a generic type / 是泛型类型
			}
			shortKind = packageName + "." + className
		} else {
			return // basic-type(int string float64) || // not-exportable-type
		}
	}

	if element.IsEllipsis {
		shortKind = "..." + shortKind
	}

	element.Kind = shortKind
}

// MakeNameFunction generates a name for a parameter or return value.
// MakeNameFunction 用于为参数或返回值生成名称。
type MakeNameFunction func(name *ast.Ident, kind string, idx int, anonymousIdx int) string

// NameTypeElements is a collection of NameTypeElement.
// NameTypeElements 是 NameTypeElement 的集合。
type NameTypeElements []*NameTypeElement

// NewNameTypeElements creates a list of NameTypeElements based on AST field list.
// NewNameTypeElements 根据 AST 字段列表创建 NameTypeElements 的列表。
func NewNameTypeElements(
	fieldList *ast.FieldList, // AST field list for function parameters / AST 字段列表，表示函数参数
	nameFunc MakeNameFunction, // Function to generate parameter names / 用于生成参数名称的函数
	source []byte, // Source code for extracting information / 用于提取信息的源代码
	packageName string, // Package name for external types / 外部类型的包名
	genericTypeParams map[string]ast.Expr, // Map of generic type parameters / 泛型类型参数的映射
) NameTypeElements {
	if fieldList == nil {
		return make(NameTypeElements, 0) // Return an empty list / 返回一个空列表
	}
	return ExtractNameTypeElements(fieldList.List, nameFunc, source, packageName, genericTypeParams)
}

// ExtractNameTypeElements extracts NameTypeElements from the AST fields.
// ExtractNameTypeElements 从 AST 字段中提取 NameTypeElements。
func ExtractNameTypeElements(
	fields []*ast.Field, // List of AST fields / AST 字段列表
	nameFunc MakeNameFunction, // Function to generate names / 用于生成名称的函数
	source []byte, // Source code / 源代码
	packageName string, // Package name / 包名
	genericTypeParams map[string]ast.Expr, // Map of generic type parameters / 泛型类型参数的映射
) NameTypeElements {
	var elements = make(NameTypeElements, 0) // Create an empty elements list / 创建一个空的元素列表
	var anonymousCount = 0                   // Counter for anonymous fields / 匿名字段计数器
	for _, field := range fields {
		var stringType string
		var isVariadic bool
		if ellipsis, ok := field.Type.(*ast.Ellipsis); ok {
			stringType = string(source[ellipsis.Ellipsis-1 : field.Type.End()-1]) // Extract ellipsis type / 提取变参类型
			isVariadic = true
		} else {
			stringType = strings.TrimSpace(string(syntaxgo_astnode.GetCode(source, field.Type))) // Extract regular type / 提取常规类型
		}
		if len(field.Names) > 0 { // Parameters usually have names, but returns often don't / 参数通常有名称，但返回值通常没有
			for _, fieldName := range field.Names {
				count := len(elements)
				paramName := nameFunc(fieldName, stringType, count, 0) // Generate name for the parameter / 为参数生成名称
				elem := NewNameTypeElement(field, paramName, stringType, isVariadic, packageName, genericTypeParams)
				elements = append(elements, elem)
			}
		} else {
			count := len(elements)
			paramName := nameFunc(nil, stringType, count, anonymousCount) // Generate name for anonymous field / 为匿名字段生成名称
			elem := NewNameTypeElement(field, paramName, stringType, isVariadic, packageName, genericTypeParams)
			elements = append(elements, elem)
			anonymousCount++
		}
	}
	return elements
}

// Names returns a list of names of the elements.
// Names 返回元素名称的列表。
func (elements NameTypeElements) Names() StatementParts {
	var names = make([]string, 0, len(elements))
	for _, element := range elements {
		names = append(names, element.Name)
	}
	return names
}

// Kinds returns a list of types of the elements.
// Kinds 返回元素类型的列表。
func (elements NameTypeElements) Kinds() []string {
	var kinds = make([]string, 0, len(elements))
	for _, node := range elements {
		kinds = append(kinds, node.Kind)
	}
	return kinds
}

// FormatAddressableNames returns the names prefixed with "&" (addressable names).
// FormatAddressableNames 返回带有 "&" 前缀的名称（可寻址名称）。
func (elements NameTypeElements) FormatAddressableNames() StatementParts {
	return utils.SetPrefix2Strings("&", elements.Names()) // Add '&' to each name for addressable types / 为每个名称添加 "&" 使其可寻址
}

// GenerateFunctionParams generates the function parameters list.
// GenerateFunctionParams 生成函数参数列表。
func (elements NameTypeElements) GenerateFunctionParams() StatementParts {
	var params = make([]string, 0, len(elements))
	for _, element := range elements {
		if element.IsEllipsis {
			params = append(params, element.Name+"...")
		} else {
			params = append(params, element.Name)
		}
	}
	return params
}

// FormatNamesWithKinds returns the names with their types (e.g., "a int").
// FormatNamesWithKinds 返回名称及其类型（例如："a int"）。
func (elements NameTypeElements) FormatNamesWithKinds() StatementParts {
	var results = make([]string, 0, len(elements)) // Initialize a slice for the results / 初始化一个切片来存放结果
	for _, element := range elements {
		results = append(results, element.Name+" "+element.Kind) // Combine name and type / 合并名称和类型
	}
	return results
}

// GenerateVarDefinitions generates variable definitions (e.g., "var a int").
// GenerateVarDefinitions 生成变量定义（例如："var a int"）。
func (elements NameTypeElements) GenerateVarDefinitions() StatementLines {
	return utils.SetPrefix2Strings("var ", elements.FormatNamesWithKinds()) // Add "var" to each definition / 为每个定义添加 "var"
}

// GroupVarsByKindToLines groups variables by their type and generates the corresponding definition lines.
// GroupVarsByKindToLines 按类型分组变量，并生成相应的定义行。
func (elements NameTypeElements) GroupVarsByKindToLines() StatementLines {
	var typeToNamesMap = map[string][]string{} // Mapping of type to variable names / 类型到名称的映射
	var uniqueTypes []string                   // Store unique types / 存储唯一的类型

	// Iterate over elements and group by type / 遍历所有元素，按类型分组
	for _, element := range elements {
		if names, exists := typeToNamesMap[element.Kind]; exists {
			typeToNamesMap[element.Kind] = append(names, element.Name)
		} else {
			typeToNamesMap[element.Kind] = []string{element.Name}
			uniqueTypes = append(uniqueTypes, element.Kind) // Record type in order of appearance / 按照出现顺序记录类型
		}
	}

	// Generate variable definition lines / 生成变量定义行
	var definitionLines = make([]string, 0, len(uniqueTypes))
	for _, kind := range uniqueTypes {
		names := typeToNamesMap[kind]
		parts := strings.Join(names, ", ") // Join names with commas / 用逗号连接名称
		definitionLines = append(definitionLines, "var "+parts+" "+kind)
	}

	return definitionLines
}
