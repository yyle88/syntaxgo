package syntaxgo_ast

import (
	"go/ast"
	"go/parser"
	"reflect"
	"strings"

	"github.com/yyle88/rese"
	"github.com/yyle88/syntaxgo/internal/utils"
	"github.com/yyle88/syntaxgo/syntaxgo_astnode"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

/*
Go语言语法分析方法使用
通过 ast.File 获得各种元素，比如 包名，引用，结构，函数
*/

func GetFuncDefineCode(source []byte, astFunc *ast.FuncDecl) string {
	return string(source[astFunc.Pos()-1 : astFunc.Body.Lbrace-1])
}

func GetPkgNameXPath(path string) string {
	return rese.P1(NewAstBundleV6(path, parser.PackageClauseOnly)).GetPackageName()
}

func GetPkgNameXFile(astFile *ast.File) (packageName string) {
	return astFile.Name.Name
}

func SeekClasses(astFile *ast.File) (astFunctions []*ast.FuncDecl, astTypes []*ast.TypeSpec, astValues []*ast.ValueSpec) {
	for _, dec := range astFile.Decls {
		switch x := dec.(type) {
		case *ast.FuncDecl:
			astFunctions = append(astFunctions, x)
		case *ast.GenDecl:
			for _, s := range x.Specs {
				switch s := s.(type) {
				case *ast.TypeSpec:
					astTypes = append(astTypes, s)
				case *ast.ValueSpec:
					astValues = append(astValues, s)
				}
			}
		}
	}
	return
}

func SeekTypes(astFile *ast.File) (astTypes []*ast.TypeSpec) {
	for _, dec := range astFile.Decls {
		switch x := dec.(type) {
		case *ast.GenDecl:
			for _, s := range x.Specs {
				switch s := s.(type) {
				case *ast.TypeSpec:
					astTypes = append(astTypes, s)
				}
			}
		}
	}
	return astTypes
}

// SeekArrayXName 查找结构体类型-返回的是结构体的{}里的内容(含括号)
func SeekArrayXName(astFile *ast.File, arrayName string) *ast.ArrayType {
	for _, dec := range astFile.Decls {
		switch x := dec.(type) {
		case *ast.GenDecl:
			for _, s := range x.Specs {
				switch s := s.(type) {
				case *ast.TypeSpec:
					//在这里得到名称
					if s.Name.Name == arrayName {
						switch s := s.Type.(type) {
						case *ast.ArrayType:
							return s
						}
					}
				}
			}
		}
	}
	return nil //在这里返回个空表示没有找到
}

// SeekStructXName 查找结构体类型-返回的是结构体的{}里的内容(含括号)
func SeekStructXName(astFile *ast.File, structName string) *ast.StructType {
	for _, dec := range astFile.Decls {
		switch x := dec.(type) {
		case *ast.GenDecl:
			for _, s := range x.Specs {
				switch s := s.(type) {
				case *ast.TypeSpec:
					//在这里得到名称
					if s.Name.Name == structName {
						switch s := s.Type.(type) {
						case *ast.StructType:
							return s
						}
					}
				}
			}
		}
	}
	return nil //在这里返回个空表示没有找到
}

// SeekMapStructNameTypes 查找所有的结构体类型-返回的是结构体名称和结构体的{}里的内容(含括号)
func SeekMapStructNameTypes(astFile *ast.File) map[string]*ast.StructType {
	astTypes := map[string]*ast.StructType{}
	for _, dec := range astFile.Decls {
		switch x := dec.(type) {
		case *ast.GenDecl:
			for _, s := range x.Specs {
				switch s := s.(type) {
				case *ast.TypeSpec:
					name := s.Name.Name
					switch s := s.Type.(type) {
					case *ast.StructType:
						astTypes[name] = s
					}
				}
			}
		}
	}
	return astTypes
}

// SeekMapStructNameDeclsTypes 查找所有的结构体类型-返回的是完整的定义代码，由type起始到闭大括号结束
func SeekMapStructNameDeclsTypes(astFile *ast.File) map[string]*ast.GenDecl {
	astTypes := map[string]*ast.GenDecl{}
	for _, dec := range astFile.Decls {
		switch x := dec.(type) {
		case *ast.GenDecl:
			for _, s := range x.Specs {
				switch s := s.(type) {
				case *ast.TypeSpec:
					name := s.Name.Name
					switch s.Type.(type) {
					case *ast.StructType:
						astTypes[name] = x
					}
				}
			}
		}
	}
	return astTypes
}

func SeekInterfaceTypes(astFile *ast.File) (mapInterfaceTypes map[string]*ast.InterfaceType) {
	mapInterfaceTypes = map[string]*ast.InterfaceType{}
	for _, dec := range astFile.Decls {
		switch x := dec.(type) {
		case *ast.GenDecl:
			for _, s := range x.Specs {
				switch s := s.(type) {
				case *ast.TypeSpec:
					name := s.Name.String()
					switch s := s.Type.(type) {
					case *ast.InterfaceType:
						mapInterfaceTypes[name] = s
					}
				}
			}
		}
	}
	return mapInterfaceTypes
}

func SeekFunctions(astFile *ast.File) (astFunctions []*ast.FuncDecl) {
	for _, declaration := range astFile.Decls {
		switch x := declaration.(type) {
		case *ast.FuncDecl:
			astFunctions = append(astFunctions, x)
		}
	}
	return astFunctions
}

func SeekFuncXName(astFile *ast.File, name string) (astFunc *ast.FuncDecl) {
	for _, declaration := range astFile.Decls {
		switch x := declaration.(type) {
		case *ast.FuncDecl:
			if x.Name.Name == name {
				return x
			}
		}
	}
	return
}

func SeekFuncXMain(astFile *ast.File) (mainFunction *ast.FuncDecl) {
	return SeekFuncXName(astFile, "main")
}

func GetFunctions(astFile *ast.File) (astFunctions []*ast.FuncDecl) {
	//ast.FileExports(astFile)
	for _, declaration := range astFile.Decls {
		switch x := declaration.(type) {
		case *ast.FuncDecl:
			zaplog.LOG.Debug(x.Name.Name)
			astFunctions = append(astFunctions, x)
			continue
		default:
		}
	}
	return
}

func SeekFunctionXName(astFile *ast.File, funcName string) (res *ast.FuncDecl, ok bool) {
	for _, declaration := range astFile.Decls {
		switch x := declaration.(type) {
		case *ast.FuncDecl:
			if x.Recv == nil { //不是方法（成员函数）而是普通函数
				if x.Name.Name == funcName {
					return x, true
				}
				continue
			}
		default:
			continue
		}
	}
	return nil, false
}

func IsFuncXRecvName(astFunc *ast.FuncDecl, recvName string) bool {
	if astFunc.Recv != nil {
		for _, recv := range astFunc.Recv.List {
			if x, ok := recv.Type.(*ast.StarExpr); ok {
				if t, ok := x.X.(*ast.Ident); ok && t.Name == recvName {
					return true
				}
			} else {
				if t, ok := recv.Type.(*ast.Ident); ok && t.Name == recvName {
					return true
				}
			}
		}
	}
	return false
}

func SeekFuncXRecvName(astFile *ast.File, recvName string, onlyExport bool) (resFunctions []*ast.FuncDecl) {
	for _, declaration := range astFile.Decls {
		switch x := declaration.(type) {
		case *ast.FuncDecl:
			if IsFuncXRecvName(x, recvName) {
				if onlyExport {
					if !utils.C0IsUppercase(x.Name.Name) {
						continue
					}
				}
				resFunctions = append(resFunctions, x)
			}
		default:
		}
	}
	return resFunctions
}

func SeekFuncXRecvNameXFuncName(astFile *ast.File, recvName string, funcName string) (res *ast.FuncDecl, ok bool) {
	for _, declaration := range astFile.Decls {
		switch x := declaration.(type) {
		case *ast.FuncDecl:
			if IsFuncXRecvName(x, recvName) {
				if x.Name.Name == funcName {
					return x, true
				}
				continue
			}
		default:
		}
	}
	return nil, false
}

func GetFunctionsXRecvName(astFunctions []*ast.FuncDecl, recvName string, onlyOut bool) (resFunctions []*ast.FuncDecl) {
	for _, x := range astFunctions {
		if IsFuncXRecvName(x, recvName) {
			if onlyOut {
				if !utils.C0IsUppercase(x.Name.Name) {
					continue
				}
			}
			resFunctions = append(resFunctions, x)
			zaplog.LOG.Debug("IS:", zap.String("name", x.Name.Name))
		} else {
			zaplog.LOG.Debug("NO:", zap.String("name", x.Name.Name))
		}
	}
	return resFunctions
}

func GetFuncDocText(astFunc *ast.FuncDecl) string {
	var docs []string
	if astFunc.Doc != nil {
		for _, i := range astFunc.Doc.List {
			docs = append(docs, i.Text)
		}
	}
	return strings.Join(docs, "\n")
}

func GetFuncRecvNameType(astFunc *ast.FuncDecl, source []byte) (recvName string, recvType string) {
	if astFunc.Recv != nil {
		names := astFunc.Recv.List[0].Names
		if len(names) > 0 {
			recvName = names[0].Name
		}
		nodeRecvType := astFunc.Recv.List[0].Type
		switch node := nodeRecvType.(type) {
		case *ast.Ident:
			recvType = string(syntaxgo_astnode.GetCode(source, node))
		case *ast.StarExpr:
			recvType = string(syntaxgo_astnode.GetCode(source, node.X))
		default:
			zaplog.LOG.Panic("unknown", zap.Any("recv_type", reflect.TypeOf(nodeRecvType)), zap.Any("node_type", reflect.TypeOf(node)))
		}
	}
	return recvName, recvType
}
