package syntaxgo_ast

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/token"

	"github.com/pkg/errors"
	"golang.org/x/tools/go/ast/astutil"
)

// Deprecated: 这个函数能够从ast的对象中把源代码写出来
// 但是，现在这个函数还不是很好用，因为它会把所有的注释都丢弃到源代码的结尾，很显然这不是我想要的，但目前还没找到特别好的解决方案
// TODO 就静候有缘人来啦
func cvtAstFileToBytes(astFile *ast.File) ([]byte, error) {
	var fileSet = token.NewFileSet()
	var buf bytes.Buffer
	if err := format.Node(&buf, fileSet, astFile); err != nil {
		return nil, errors.WithMessage(err, "format node is wrong")
	}
	return buf.Bytes(), nil
}

// Deprecated: 这是添加import的工具
// 这个不好使，因为导出ast对象中的所有源码的函数不好使，在导出时会把注释错位掉，这也就导致了这个函数没法使用
// 我很期待有人能解决这个问题
// 其次是，当源代码中完全没有 import 这字段时，代码将不能主动添加，而是会报异常，这个我觉得很不科学
// 我的意图，我发现，当我自己用字符串拼接写出源代码时，假如我不写引用，在执行 format 格式化的时候就会特别慢，特别慢，很明显这个问题需要解决
// 因此我手动实现了追加 import 的逻辑
// TODO 就静候有缘人来啦
func addImportOfPkgPath(fileSet *token.FileSet, astFile *ast.File, path string) bool {
	return astutil.AddImport(fileSet, astFile, path)
}
