package syntaxgo_ast

import (
	"fmt"
	"go/ast"
	"go/token"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/runpath/runtestpath"
)

func TestNewAstXFilepath(t *testing.T) {
	path := runtestpath.SrcPath(t)
	astFile, e := NewAstXFilepath(path)
	require.NoError(t, e)
	ast.Print(token.NewFileSet(), astFile)
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
