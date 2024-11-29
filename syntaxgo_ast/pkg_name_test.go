package syntaxgo_ast

import (
	"go/token"
	"testing"

	"github.com/yyle88/done"
	"github.com/yyle88/runpath"
	"github.com/yyle88/runpath/runtestpath"
)

func TestGetPackageNameFromPath(t *testing.T) {
	t.Log(GetPackageNameFromPath(runpath.Path()))
}

func TestGetPackageNameFromFile(t *testing.T) {
	path := runtestpath.SrcPath(t)
	t.Log(path)
	astBundle := done.P1(NewAstBundleV3(token.NewFileSet(), path))
	astFile := astBundle.file
	t.Log(GetPackageNameFromFile(astFile))
}
