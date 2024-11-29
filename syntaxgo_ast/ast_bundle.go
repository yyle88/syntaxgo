package syntaxgo_ast

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/parser"
	"go/printer"
	"go/token"

	"github.com/yyle88/erero"
	"golang.org/x/tools/go/ast/astutil"
)

/*
Go语言语法分析方法使用
通过 ast.File 获得各种元素，比如 包名，引用，结构，函数
*/

type AstBundle struct {
	fset *token.FileSet
	file *ast.File
}

func NewAstBundle(fileSet *token.FileSet, astFile *ast.File) *AstBundle {
	return &AstBundle{
		fset: fileSet,
		file: astFile,
	}
}

func NewAstBundleV1(data []byte) (*AstBundle, error) {
	return NewAstBundleV2(token.NewFileSet(), data)
}

func NewAstBundleV2(fset *token.FileSet, data []byte) (*AstBundle, error) {
	astFile, err := parser.ParseFile(fset, "", data, parser.ParseComments)
	if err != nil {
		return nil, erero.Wro(err)
	}
	return NewAstBundle(fset, astFile), nil
}

func NewAstBundleV3(fset *token.FileSet, path string) (*AstBundle, error) {
	astFile, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		return nil, erero.Wro(err)
	}
	return NewAstBundle(fset, astFile), nil
}

func NewAstBundleV4(path string) (*AstBundle, error) {
	return NewAstBundleV3(token.NewFileSet(), path)
}

func NewAstBundleV5(fset *token.FileSet, path string, mode parser.Mode) (*AstBundle, error) {
	astFile, err := parser.ParseFile(fset, path, nil, mode)
	if err != nil {
		return nil, erero.Wro(err)
	}
	return NewAstBundle(fset, astFile), nil
}

func NewAstBundleV6(path string, mode parser.Mode) (*AstBundle, error) {
	return NewAstBundleV5(token.NewFileSet(), path, mode)
}

func (ab *AstBundle) GetBundle() (*ast.File, *token.FileSet) {
	fileSet := ab.fset
	astFile := ab.file

	return astFile, fileSet
}

func (ab *AstBundle) FormatSource() ([]byte, error) {
	var buf bytes.Buffer
	if err := format.Node(&buf, ab.fset, ab.file); err != nil {
		return nil, erero.Wro(err)
	}
	return buf.Bytes(), nil
}

func (ab *AstBundle) SerializeAst() ([]byte, error) {
	var buf bytes.Buffer
	if err := printer.Fprint(&buf, ab.fset, ab.file); err != nil {
		return nil, erero.Wro(err)
	}
	return buf.Bytes(), nil
}

func (ab *AstBundle) AddImport(path string) bool {
	return astutil.AddImport(ab.fset, ab.file, path)
}

func (ab *AstBundle) AddNamedImport(name string, path string) bool {
	return astutil.AddNamedImport(ab.fset, ab.file, name, path)
}

func (ab *AstBundle) DeleteImport(path string) bool {
	return astutil.DeleteImport(ab.fset, ab.file, path)
}

func (ab *AstBundle) DeleteNamedImport(name string, path string) bool {
	return astutil.DeleteNamedImport(ab.fset, ab.file, name, path)
}

func (ab *AstBundle) Print() error {
	return ast.Print(ab.fset, ab.file)
}

func (ab *AstBundle) GetPackageName() string {
	return ab.file.Name.Name
}
