package syntaxgo_ast

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
	"github.com/yyle88/syntaxgo/syntaxgo_astnode"
)

func TestAstBundle_Print_Search(t *testing.T) {
	path := runtestpath.SrcPath(t)
	astBundle, err := NewAstBundleV3(token.NewFileSet(), path)
	require.NoError(t, err)
	done.Done(astBundle.Print())
}

func TestSeekFuncXName(t *testing.T) {
	path := runtestpath.SrcPath(t)
	astBundle := rese.P1(NewAstBundleV3(token.NewFileSet(), path))
	astFunc := SeekFuncXName(astBundle.file, "NewAstXFilepath")
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
	astBundle := rese.P1(NewAstBundleV2(token.NewFileSet(), srcData))
	astFunc := SeekArrayXName(astBundle.file, "Examples")
	require.NotNil(t, astFunc)
	t.Log(string(syntaxgo_astnode.GetCode(srcData, astFunc)))
}
