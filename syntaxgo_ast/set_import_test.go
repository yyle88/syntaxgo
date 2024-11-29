package syntaxgo_ast

import (
	"go/format"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInjectImports(t *testing.T) {
	const code = `package main

	import "time"

	//这是main函数的注释
	func main() {
		fmt.Println("abc") //随便打印个字符串
		fmt.Println(time.Now()) //随便打印当前时间
		fmt.Println(strconv.Itoa(1))
	}
`
	t.Log(code)

	var newSrc = InjectImports([]byte(code), []string{
		"fmt",
		"strconv",
	})
	t.Log(string(newSrc)) //待格式化的数据

	resSrc, err := format.Source(newSrc)
	require.NoError(t, err)
	t.Log(string(resSrc)) //需要微调引用包
}
