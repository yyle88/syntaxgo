package syntaxgo_astnorm

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/rese"
	"github.com/yyle88/runpath/runtestpath"
	"github.com/yyle88/syntaxgo/syntaxgo_ast"
	"github.com/yyle88/syntaxgo/syntaxgo_search"
)

func TestGetSimpleResElements(t *testing.T) {
	path := runtestpath.SrcPath(t)
	t.Log(path)

	source := rese.A1(os.ReadFile(path))
	astBundle := rese.P1(syntaxgo_ast.NewAstBundleV1(source))
	astFile, _ := astBundle.GetBundle()

	resFunc := syntaxgo_search.FindFunctionByName(astFile, "GetSimpleResElements")
	require.NotNil(t, resFunc)

	elements := GetSimpleResElements(resFunc.Type.Results.List, source)
	for _, elem := range elements {
		t.Log(elem.Name, elem.Kind)
	}
}

func TestGetSimpleArgElements(t *testing.T) {
	path := runtestpath.SrcPath(t)
	t.Log(path)

	source := rese.A1(os.ReadFile(path))
	astBundle := rese.P1(syntaxgo_ast.NewAstBundleV1(source))
	astFile, _ := astBundle.GetBundle()

	resFunc := syntaxgo_search.FindFunctionByName(astFile, "GetSimpleArgElements")
	require.NotNil(t, resFunc)

	elements := GetSimpleArgElements(resFunc.Type.Params.List, source)
	for _, elem := range elements {
		t.Log(elem.Name, elem.Kind)
	}
}
