package syntaxgo_ast

import (
	"go/parser"
	"go/token"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/done"
	"github.com/yyle88/rese"
	"github.com/yyle88/runpath"
	"github.com/yyle88/runpath/runtestpath"
)

func TestAstBundle_Print(t *testing.T) {
	path := runtestpath.SrcPath(t)
	t.Log(path)
	astBundle := done.P1(NewAstBundleV3(token.NewFileSet(), path))
	done.Done(astBundle.Print())
}

func TestAstBundle_PrintCurrentFile(t *testing.T) {
	path := runpath.CurrentPath()
	t.Log(path)
	astBundle := rese.P1(NewAstBundleV4(path))
	done.Done(astBundle.Print())
}

func TestNewAstBundleV5(t *testing.T) {
	astBundle, err := NewAstBundleV5(token.NewFileSet(), runtestpath.SrcPath(t), parser.PackageClauseOnly)
	require.NoError(t, err)
	done.Done(astBundle.Print())
	t.Log(string(done.VAE(astBundle.FormatSource()).Nice()))
}

func TestNewAstBundleV6(t *testing.T) {
	astBundle, err := NewAstBundleV6(runpath.Path(), parser.ImportsOnly)
	require.NoError(t, err)
	done.Done(astBundle.Print())
	t.Log(string(done.VAE(astBundle.FormatSource()).Nice()))
}

func TestAstBundle_FormatSource(t *testing.T) {
	const code = `package main
	//这是main函数的注释
	func main() {
		fmt.Println("abc") //随便打印个字符串
		fmt.Println(time.Now()) //随便打印当前时间
	}
`
	astBundle, err := NewAstBundleV2(token.NewFileSet(), []byte(code))
	require.NoError(t, err)

	added := astBundle.AddImport("fmt")
	require.True(t, added)

	added = astBundle.AddImport("time")
	require.True(t, added)

	newSrc, err := astBundle.FormatSource()
	require.NoError(t, err)
	t.Log(string(newSrc))
}

func TestAstBundle_SerializeAst(t *testing.T) {
	const code = `package main

	//go:embed hello.txt
	var s string

	//这是main函数的注释
	func main() {
		fmt.Println(s) //打印整个文件内容
	}
`
	astBundle, err := NewAstBundleV2(token.NewFileSet(), []byte(code))
	require.NoError(t, err)

	added := astBundle.AddImport("fmt")
	require.True(t, added)

	added = astBundle.AddNamedImport("_", "embed")
	require.True(t, added)

	newSrc, err := astBundle.SerializeAst()
	require.NoError(t, err)
	t.Log(string(newSrc))
}

func TestAstBundle_GetPackageName(t *testing.T) {
	astBundle := rese.P1(NewAstBundleV6(runpath.Path(), parser.PackageClauseOnly))
	packageName := astBundle.GetPackageName()
	t.Log(packageName)
}
