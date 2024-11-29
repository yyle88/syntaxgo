package syntaxgo

import (
	"github.com/yyle88/runpath"
	"github.com/yyle88/syntaxgo/syntaxgo_ast"
)

// CurrentPackageName 这里的当前不是当前的意思，而是调用者调用时，就能得到调用位置的包名
func CurrentPackageName() string {
	return syntaxgo_ast.GetPackageNameFromPath(runpath.Skip(1))
}

// GetPkgName 获得某个go文件的包名
func GetPkgName(path string) string {
	return syntaxgo_ast.GetPackageNameFromPath(path)
}
