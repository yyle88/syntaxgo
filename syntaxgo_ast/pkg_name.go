package syntaxgo_ast

import (
	"go/ast"
	"go/parser"

	"github.com/yyle88/rese"
)

func GetPackageNameFromPath(path string) string {
	return rese.P1(NewAstBundleV6(path, parser.PackageClauseOnly)).GetPackageName()
}

func GetPackageNameFromFile(astFile *ast.File) (packageName string) {
	return astFile.Name.Name
}
