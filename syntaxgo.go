package syntaxgo

import (
	"github.com/yyle88/runpath"
	"github.com/yyle88/syntaxgo/syntaxgo_ast"
)

// CurrentPackageName returns the package name of the caller's location.
// CurrentPackageName 返回调用者位置的包名。
func CurrentPackageName() string {
	return syntaxgo_ast.GetPackageNameFromPath(runpath.Skip(1))
}

// GetPkgName returns the package name of a given Go file path.
// GetPkgName 返回给定 Go 文件路径的包名。
func GetPkgName(path string) string {
	return syntaxgo_ast.GetPackageNameFromPath(path)
}
