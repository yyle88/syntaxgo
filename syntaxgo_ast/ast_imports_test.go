package syntaxgo_ast

import (
	"testing"
)

func TestAddImportsOfPackages(t *testing.T) {
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

	var newSrc = AddImportsOfPackages([]byte(code), []string{
		"fmt",
		"strconv",
	})
	t.Log(string(newSrc)) //虽然结果的代码看着可能有点别扭，但是没关系的，只要再配合golang代码的format就会完美的
}
