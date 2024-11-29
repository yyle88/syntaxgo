package syntaxgo_ast

import (
	"go/ast"
	"go/parser"
	"go/token"

	"github.com/yyle88/erero"
	"github.com/yyle88/syntaxgo/internal/utils"
)

// Deprecated: 需要把逻辑改改不要使用不推荐的包，而且目前看来这个函数几乎没有用途
// MergeAstFilesXRootPath 得到目录中唯一包的语法内容
// 这个函数的命名很差
func MergeAstFilesXRootPath(root string) (*AstBundle, error) {
	fileSet := token.NewFileSet()
	packsMap, err := NewAstPackagesXRootPath(fileSet, root)
	if err != nil {
		return nil, erero.Wro(err)
	}
	if len(packsMap) >= 2 {
		return nil, erero.Errorf("more than one package in root path: %s", root)
	}
	for _, pkg := range packsMap {
		//这样处理意味着只会返回 map 中的第一个包，并且只合并该包中的文件。如果目录中存在多个包，它们将不会被处理。
		//这种设计可能基于假设：你所处理的目录下只会包含一个 Go 包。
		//也就是说，开发者假定根目录下的 Go 文件都属于同一个包，或者只关心第一个包的内容。
		return NewAstBundle(
			fileSet,
			ast.MergePackageFiles(pkg, ast.FilterImportDuplicates),
		), nil
	}
	return nil, erero.Errorf("no packages in root path: %s", root)
}

// Deprecated: 这里暂不知道该怎么修改，但是继续增加个这样的标识，就能在lint时不报错，估计是传导出去啦
// ast.Package has been deprecated
// NewAstPackagesXRootPath 得到整个目录下各个包的语法内容
// 这个函数的命名很差
func NewAstPackagesXRootPath(fset *token.FileSet, root string) (map[string]*ast.Package, error) {
	packsMap, err := parser.ParseDir(
		fset,
		root,
		utils.IsGoSourceFile,
		parser.ParseComments,
	)
	if err != nil {
		return nil, erero.Wro(err)
	}
	return packsMap, nil
}
