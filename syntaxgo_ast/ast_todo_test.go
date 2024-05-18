package syntaxgo_ast

import (
	"go/token"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_addImportOfPkgPath(t *testing.T) {
	const code = `package main

	import "time"

	//这是main函数的注释
	func main() {
		fmt.Println("abc") //随便打印个字符串
		fmt.Println(time.Now()) //随便打印当前时间
	}
`
	astFile, err := NewAstFromSource([]byte(code))
	require.NoError(t, err)

	//问题：假如代码中完全没有 import 这个块时，代码会直接报异常
	added := addImportOfPkgPath(token.NewFileSet(), astFile, "fmt")
	require.True(t, added)

	//问题：最后的注释不能回到正确的位置，这个也是挺坑的
	newSrc, err := cvtAstFileToBytes(astFile)
	require.NoError(t, err)

	t.Log(string(newSrc))
}
