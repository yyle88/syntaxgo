package syntaxgo_ast

import (
	"go/ast"
	"go/token"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/done"
	"github.com/yyle88/runpath"
)

func TestMergeAstFilesXRootPath(t *testing.T) {
	astBundle, err := MergeAstFilesXRootPath(runpath.PARENT.Path())
	require.NoError(t, err)
	done.Done(ast.Print(astBundle.fset, astBundle.file))
}

func TestNewAstPackagesXRootPath(t *testing.T) {
	fileSet := token.NewFileSet()
	packsMap, err := NewAstPackagesXRootPath(fileSet, runpath.PARENT.Path())
	require.NoError(t, err)
	for packageName, pkg := range packsMap {
		t.Log(packageName, pkg.Name)

		for path, astFile := range pkg.Files {
			t.Log(path)

			done.Done(ast.Print(fileSet, astFile))
		}
	}
}
