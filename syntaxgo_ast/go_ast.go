package syntaxgo_ast

import (
	"go/ast"
	"go/parser"
	"go/token"
	"reflect"
	"strings"

	"github.com/yyle88/done"
	"github.com/yyle88/erero"
	"github.com/yyle88/syntaxgo/internal/utils"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

/*
Go语言语法分析方法使用
通过 ast.File 获得各种元素，比如 包名，引用，结构，函数
*/

// NewAstXFilepath
// 这是个实例函数。
// 当然也可以用来做转换。
func NewAstXFilepath(path string) (astFile *ast.File, e error) {
	return parser.ParseFile(token.NewFileSet(), path, nil, parser.ParseComments)
}

// NewAstPackageClauseOnlyXFilepath 有的时候你只想拿到包名，这时候没有必要拿到全部的文件
func NewAstPackageClauseOnlyXFilepath(path string) (astFile *ast.File, e error) {
	return parser.ParseFile(token.NewFileSet(), path, nil, parser.PackageClauseOnly)
}

// NewAstImportsOnlyXFilepath 当然有的时候你不仅要拿包名，还要拿引用包的部分，因此就用这个函数就行
func NewAstImportsOnlyXFilepath(path string) (astFile *ast.File, e error) {
	return parser.ParseFile(token.NewFileSet(), path, nil, parser.ImportsOnly)
}

// NewAstFromSource 这个函数适用于我把文件已经读出来的场景
// 这时候我将会直接使用二进制的 data 内容
// 因此这个时候 filename 只是参考值，也允许传空白字符串，都不影响，看调用者的心情
// 因此我为了简化代码就不设置这个参数啦，只传递文件内容即可
func NewAstFromSource(data []byte) (astFile *ast.File, e error) {
	fileSet := token.NewFileSet()
	astFile, e = parser.ParseFile(fileSet, "", data, parser.ParseComments)
	if e != nil {
		return nil, e
	}
	return astFile, nil
}

func GetNodeIndex(astNode ast.Node) (sdx, edx int) {
	sdx = int(astNode.Pos() - 1)
	edx = int(astNode.End() - 1)
	return sdx, edx
}

func GetNodeCode(source []byte, astNode ast.Node) string {
	return string(source[astNode.Pos()-1 : astNode.End()-1])
}

func GetNodeBytes(source []byte, astNode ast.Node) []byte {
	return source[astNode.Pos()-1 : astNode.End()-1]
}

func DeleteNodeBytes(source []byte, astNode ast.Node) []byte {
	return utils.SafeMerge(
		source[:astNode.Pos()-1],
		source[astNode.End()-1:],
	)
}

func ChangeNodeBytes(source []byte, astNode ast.Node, newCode []byte) []byte {
	return utils.SafeMerge(
		source[:astNode.Pos()-1],
		newCode,
		source[astNode.End()-1:],
	)
}

func ChangeNodeBytesXNewLines(source []byte, astNode ast.Node, newCode []byte, newLinesNum int) []byte {
	var newLines = make([]byte, 0, newLinesNum)
	for idx := 0; idx < newLinesNum; idx++ {
		newLines = append(newLines, '\n')
	}
	return utils.SafeMerge(
		source[:astNode.Pos()-1],
		newLines,
		newCode,
		newLines,
		source[astNode.End()-1:],
	)
}

func GetFuncDefineCode(source []byte, astFunc *ast.FuncDecl) string {
	return string(source[astFunc.Pos()-1 : astFunc.Body.Lbrace-1])
}

func GetPkgNameXPath(path string) string {
	//只需要 package name 就行
	astFile := done.VCE(parser.ParseFile(token.NewFileSet(), path, nil, parser.PackageClauseOnly)).Nice()
	pkgName := astFile.Name.Name
	return pkgName
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

// Deprecated: 这里暂不知道该怎么修改，但是继续增加个这样的标识，就能在lint时不报错，估计是传导出去啦
// ast.Package has been deprecated
// NewAstPackagesXRootPath 得到整个目录下各个包的语法内容
// 这个函数的命名很差
func NewAstPackagesXRootPath(rootPath string) (map[string]*ast.Package, error) {
	packsMap, err := parser.ParseDir(
		token.NewFileSet(),
		rootPath,
		utils.IsFileIsGoFile,
		parser.ParseComments,
	)
	if err != nil {
		return nil, erero.Wro(err)
	}
	return packsMap, nil
}

// Deprecated: 需要把逻辑改改不要使用不推荐的包，而且目前看来这个函数几乎没有用途
// MergeAstFilesXRootPath 得到目录中唯一包的语法内容
// 这个函数的命名很差
func MergeAstFilesXRootPath(rootPath string) (*ast.File, error) {
	packsMap, err := NewAstPackagesXRootPath(rootPath)
	if err != nil {
		return nil, erero.Wro(err)
	}
	if len(packsMap) >= 2 {
		return nil, erero.Errorf("more than one package in root path: %s", rootPath)
	}
	for _, pkg := range packsMap {
		//这样处理意味着只会返回 map 中的第一个包，并且只合并该包中的文件。如果目录中存在多个包，它们将不会被处理。
		//这种设计可能基于假设：你所处理的目录下只会包含一个 Go 包。
		//也就是说，开发者假定根目录下的 Go 文件都属于同一个包，或者只关心第一个包的内容。
		return ast.MergePackageFiles(pkg, ast.FilterImportDuplicates), nil
	}
	return nil, erero.Errorf("no packages in root path: %s", rootPath)
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
					if !utils.C0IsUPPER(x.Name.Name) {
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
				if !utils.C0IsUPPER(x.Name.Name) {
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
			recvType = GetNodeCode(source, node)
		case *ast.StarExpr:
			recvType = GetNodeCode(source, node.X)
		default:
			zaplog.LOG.Panic("unknown", zap.Any("recv_type", reflect.TypeOf(nodeRecvType)), zap.Any("node_type", reflect.TypeOf(node)))
		}
	}
	return recvName, recvType
}
