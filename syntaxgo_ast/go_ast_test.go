package syntaxgo_ast

import (
	"fmt"
	"go/ast"
	"go/token"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/done"
	"github.com/yyle88/runpath"
	"github.com/yyle88/runpath/runtestpath"
)

func TestNewAstXFilepath(t *testing.T) {
	path := runtestpath.SrcPath(t)
	astFile, err := NewAstXFilepath(path)
	require.NoError(t, err)
	done.Done(ast.Print(token.NewFileSet(), astFile))
}

func TestSeekFuncXName(t *testing.T) {
	path := runtestpath.SrcPath(t)
	astFile, _ := NewAstXFilepath(path)
	astFunc := SeekFuncXName(astFile, "NewAstXFilepath")
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

func TestSeekArrayXName(t *testing.T) {
	examples := make(Examples, 0, 3)
	examples = append(examples, &Example{Name: "a"})
	examples = append(examples, &Example{Name: "b"})
	examples = append(examples, &Example{Name: "c"})
	t.Log(examples)

	path := runpath.CurrentPath()
	srcData := done.VAE(os.ReadFile(path)).Nice()
	astFile, err := NewAstFromSource(srcData)
	require.NoError(t, err)
	astFunc := SeekArrayXName(astFile, "Examples")
	require.NotNil(t, astFunc)
	t.Log(GetNodeCode(srcData, astFunc))
}

func TestNewAstPackagesXRootPath(t *testing.T) {
	packsMap, err := NewAstPackagesXRootPath(runpath.PARENT.Path())
	require.NoError(t, err)
	t.Log(packsMap)
}

func TestMergeAstFilesXRootPath(t *testing.T) {
	astFile, err := MergeAstFilesXRootPath(runpath.PARENT.Path())
	require.NoError(t, err)
	done.Done(ast.Print(token.NewFileSet(), astFile))
}
