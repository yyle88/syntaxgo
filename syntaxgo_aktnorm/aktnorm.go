package syntaxgo_aktnorm

import (
	"go/ast"
	"strings"

	"github.com/yyle88/syntaxgo/internal/utils"
	"github.com/yyle88/syntaxgo/syntaxgo_astnode"
)

type NameTypeElement struct {
	Name       string   //字段名称
	Kind       string   //类型的字符串描述比如 int string A utils.A *B *pkg.B 这样的类型
	Type       ast.Expr //这是go内部的类型描述
	IsEllipsis bool     //是否为 ...int 或者 ...string这样的变长参数类型的
}

// NewNameTypeElement 通过解析源代码的参数/返回列表，获得某个参数/返回位置的具体信息(名称和类型)，把数据规范化，以便于在做其它代码生成器 gen 的时候使用
func NewNameTypeElement(astField *ast.Field, name string, kind string, isEllipsis bool, packageName string, genericsMap map[string]int) *NameTypeElement {
	param := &NameTypeElement{
		Name:       name, //因为单个 *ast.Field 里面的name是个数组，这里更进一步就是把数组的每个元素都展开来，比如 a,b int 就是 a int 和 b int 两个元素
		Kind:       kind,
		Type:       astField.Type,
		IsEllipsis: isEllipsis,
	}
	if packageName != "" {
		param.SetPkgUsage(packageName, genericsMap)
	}
	return param
}

// SetPkgUsage 通过在包内定义的函数的类型，以及包名，就能得到外部调用的参数类型 比如内部是 *A 则在外部则需要传 *pkg.A 的对象
func (T *NameTypeElement) SetPkgUsage(packageName string, genericsMap map[string]int) {
	if strings.Contains(T.Kind, ".") {
		return //当包含“点”的时候表示是调用的其它包的内容
	}
	if s := T.Kind[0]; string(s) == "*" { //it's a pointer type *A *B
		if vcc := T.Kind[1]; vcc >= 'A' && vcc <= 'Z' { //the type is beginning with a capital letter A-Z
			className := T.Kind[1:]
			if _, ok := genericsMap[className]; ok {
				return //说明是泛型定义的函数
			}
			T.Kind = "*" + packageName + "." + className
		}
	} else {
		if s >= 'A' && s <= 'Z' { //the type is beginning with a capital letter A-Z
			className := T.Kind
			if _, ok := genericsMap[className]; ok {
				return //说明是泛型定义的函数
			}
			T.Kind = packageName + "." + className
		}
	}
}

type MakeNameFunction func(name *ast.Ident, kind string, idx int, anonymousIdx int) string

type NameTypeElements []*NameTypeElement

func NewNameTypeElements(astFieldList *ast.FieldList, nameFunc MakeNameFunction, source []byte, packageName string, genericsMap map[string]int) NameTypeElements {
	if astFieldList == nil {
		return make(NameTypeElements, 0) //其实返回空指针也是一样的
	}
	return GetNameTypeElements(astFieldList.List, nameFunc, source, packageName, genericsMap)
}

func GetNameTypeElements(astFields []*ast.Field, nameFunc MakeNameFunction, source []byte, packageName string, genericsMap map[string]int) NameTypeElements {
	var elements = make(NameTypeElements, 0)
	var anonymousCount = 0
	for _, one := range astFields {
		var sKind string
		var isEllipsis bool
		if ellipsis, ok := one.Type.(*ast.Ellipsis); ok {
			sKind = string(source[ellipsis.Ellipsis-1 : one.Type.End()-1])
			isEllipsis = true
		} else {
			sKind = strings.TrimSpace(string(syntaxgo_astnode.GetCode(source, one.Type)))
		}
		if len(one.Names) > 0 { //通常参数都是有名字的，但也会存在参数没有名字的情况，而返回值通常都是只有类型
			for _, name := range one.Names {
				count := len(elements)
				sName := nameFunc(name, sKind, count, 0)
				elem := NewNameTypeElement(one, sName, sKind, isEllipsis, packageName, genericsMap)
				elements = append(elements, elem)
			}
		} else {
			count := len(elements)
			sName := nameFunc(nil, sKind, count, anonymousCount)
			elem := NewNameTypeElement(one, sName, sKind, isEllipsis, packageName, genericsMap)
			elements = append(elements, elem)
			anonymousCount++
		}
	}
	return elements
}

func (xs NameTypeElements) Names() StatementParts {
	var names = make([]string, 0, len(xs))
	for _, x := range xs {
		names = append(names, x.Name)
	}
	return names
}

func (xs NameTypeElements) Kinds() []string {
	var kinds = make([]string, 0, len(xs))
	for _, node := range xs {
		kinds = append(kinds, node.Kind)
	}
	return kinds
}

func (xs NameTypeElements) GetNamesAddressesStats() StatementParts {
	return utils.SetPrefix2Strings("&", xs.Names())
}

func (xs NameTypeElements) GetFunctionParamsStats() StatementParts {
	var params = make([]string, 0, len(xs))
	for _, x := range xs {
		if x.IsEllipsis {
			params = append(params, x.Name+"...")
		} else {
			params = append(params, x.Name)
		}
	}
	return params
}

func (xs NameTypeElements) GetNamesKindsStats() StatementParts {
	var results = make([]string, 0, len(xs))
	for _, x := range xs {
		results = append(results, x.Name+" "+x.Kind)
	}
	return results
}

func (xs NameTypeElements) GetVariablesDefineLines() StatementLines {
	return utils.SetPrefix2Strings("var ", xs.GetNamesKindsStats())
}

func (xs NameTypeElements) GetVariablesGroupByKindsDefineLines() StatementLines {
	var mpKindNames = map[string][]string{}
	var uniqueKinds []string
	for _, one := range xs {
		if names, ok := mpKindNames[one.Kind]; ok {
			mpKindNames[one.Kind] = append(names, one.Name)
		} else {
			mpKindNames[one.Kind] = []string{one.Name}
			uniqueKinds = append(uniqueKinds, one.Kind) //记录新类型，按类型最早出现的次序排列，其实这个slice和这个map共同组成了个 LinklistMap 的数据结构
		}
	}
	var defineLines = make([]string, 0, len(uniqueKinds))
	for _, kindType := range uniqueKinds { //这样是为了让返回的顺序和定义的顺序相同，避免纯使用map导致每次生成的结果都和前一次的不同，gen 的代码总是变动
		names := mpKindNames[kindType]
		parts := strings.Join(names, ", ")
		defineLines = append(defineLines, "var "+parts+" "+kindType)
	}
	return defineLines
}
