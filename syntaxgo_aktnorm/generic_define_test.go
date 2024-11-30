package syntaxgo_aktnorm

import (
	"go/token"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/must"
	"github.com/yyle88/rese"
	"github.com/yyle88/runpath"
	"github.com/yyle88/syntaxgo/syntaxgo_ast"
	"github.com/yyle88/syntaxgo/syntaxgo_search"
)

func checkNiceFunction[A comparable, B comparable](a A, b A, x B) {
	must.Nice(a)
	must.Nice(b)
	must.Nice(x)
}

func TestGetFuncGenericTypeParamsMap(t *testing.T) {
	checkNiceFunction("a", "b", 100)

	path := runpath.Path()
	astBundle := rese.P1(syntaxgo_ast.NewAstBundleV3(token.NewFileSet(), path))

	astFile, _ := astBundle.GetBundle()

	resFunc := syntaxgo_search.FindFunctionByName(astFile, "checkNiceFunction")
	require.NotNil(t, resFunc)

	nameMap := GetFuncGenericTypeParamsMap(resFunc)
	t.Log(nameMap) // output: map[A:comparable B:comparable]
}
