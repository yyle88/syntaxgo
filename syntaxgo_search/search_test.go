package syntaxgo_search

import (
	"fmt"
	"go/token"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/done"
	"github.com/yyle88/rese"
	"github.com/yyle88/runpath"
	"github.com/yyle88/runpath/runtestpath"
	"github.com/yyle88/syntaxgo/syntaxgo_ast"
	"github.com/yyle88/syntaxgo/syntaxgo_astnode"
)

func TestAstBundle_Print_Search(t *testing.T) {
	path := runtestpath.SrcPath(t)
	astBundle, err := syntaxgo_ast.NewAstBundleV3(token.NewFileSet(), path)
	require.NoError(t, err)
	done.Done(astBundle.Print())
}

func TestFindFunctionByName(t *testing.T) {
	path := runtestpath.SrcPath(t)
	astBundle := rese.P1(syntaxgo_ast.NewAstBundleV3(token.NewFileSet(), path))

	astFile, _ := astBundle.GetBundle()

	astFunc := FindFunctionByName(astFile, "FindFunctionByName")
	if astFunc == nil {
		return
	}
	for k, s := range astFunc.Doc.List {
		fmt.Println("-----", k, "-----")
		fmt.Println(s.Text)
		fmt.Println("-----", "-", "-----")
	}
}

type Example struct {
	Name string
}

type Examples []*Example

func TestFindArrayTypeByName(t *testing.T) {
	examples := make(Examples, 0, 3)
	examples = append(examples, &Example{Name: "a"})
	examples = append(examples, &Example{Name: "b"})
	examples = append(examples, &Example{Name: "c"})
	t.Log(examples)

	path := runpath.CurrentPath()
	srcData := done.VAE(os.ReadFile(path)).Nice()
	astBundle := rese.P1(syntaxgo_ast.NewAstBundleV2(token.NewFileSet(), srcData))

	astFile, _ := astBundle.GetBundle()

	astFunc := FindArrayTypeByName(astFile, "Examples")
	require.NotNil(t, astFunc)
	t.Log(string(syntaxgo_astnode.GetCode(srcData, astFunc)))
}

func TestFindStructTypeByName(t *testing.T) {
	path := runpath.CurrentPath()
	srcData := done.VAE(os.ReadFile(path)).Nice()
	astBundle := rese.P1(syntaxgo_ast.NewAstBundleV2(token.NewFileSet(), srcData))

	astFile, _ := astBundle.GetBundle()

	astStruct, ok := FindStructTypeByName(astFile, "Example")
	require.True(t, ok)
	require.NotNil(t, astStruct)
	t.Log(string(syntaxgo_astnode.GetCode(srcData, astStruct)))
}

func TestFindStructDeclarationByName(t *testing.T) {
	path := runpath.CurrentPath()
	srcData := done.VAE(os.ReadFile(path)).Nice()
	astBundle := rese.P1(syntaxgo_ast.NewAstBundleV2(token.NewFileSet(), srcData))

	astFile, _ := astBundle.GetBundle()

	astStruct, ok := FindStructDeclarationByName(astFile, "Example")
	require.True(t, ok)
	require.NotNil(t, astStruct)
	t.Log(string(syntaxgo_astnode.GetCode(srcData, astStruct)))
}
