package syntaxgo_ast

import (
	"go/ast"
	"go/token"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/done"
	"github.com/yyle88/rese"
	"github.com/yyle88/runpath"
)

func TestMergeOnePackageFiles(t *testing.T) {
	packageName := rese.P1(NewAstBundleV4(runpath.Path())).GetPackageName()
	t.Log(packageName)

	astBundle, err := MergeOnePackageFiles(runpath.PARENT.Path(), packageName)
	require.NoError(t, err)
	done.Done(ast.Print(astBundle.fset, astBundle.file))
}

func TestParseRootGetPackages(t *testing.T) {
	fileSet := token.NewFileSet()
	packsMap, err := ParseRootGetPackages(fileSet, runpath.PARENT.Path())
	require.NoError(t, err)
	for packageName, pkg := range packsMap {
		t.Log(packageName, pkg.Name)

		for path, astFile := range pkg.Files {
			t.Log(path)

			done.Done(ast.Print(fileSet, astFile))
		}
	}
}

func TestMergeSubPackageFiles(t *testing.T) {
	astBundle, err := MergeSubPackageFiles(runpath.PARENT.Path())
	require.NoError(t, err)
	done.Done(ast.Print(astBundle.fset, astBundle.file))
}
