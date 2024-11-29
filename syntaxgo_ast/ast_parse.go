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
Package `syntaxgo_ast` provides tools for parsing, analyzing, and manipulating Go source code using the `go/ast` package.

This package simplifies common tasks such as:
  - Creating and managing AST (Abstract Syntax Tree) bundles.
  - Adding or removing import paths programmatically.
  - Formatting AST nodes back into Go source code.
  - Serializing AST structures into textual representations.
  - Accessing metadata like the package name.
  - Printing the AST structure for debugging or analysis.

`syntaxgo_ast` can be used in the following scenarios:
  - Static code analysis.
  - Automated code generation and refactoring.
  - Programmatic manipulation of Go source files.

By abstracting away low-level operations, this package helps developers generate and enhance golang source code.
*/

/*
Package `syntaxgo_ast` 提供了使用 `go/ast` 包解析、分析和操作、修改 Go 源代码的工具。

这个包简化了以下常见任务：
  - 创建和管理 AST（抽象语法树）集合。
  - 以编程方式添加或删除导入路径。
  - 将 AST 节点格式化为 Go 源代码。
  - 将 AST 结构序列化为文本表示。
  - 访问诸如包名之类的元数据。
  - 打印 AST 结构以便调试或分析。

`syntaxgo_ast` 适用于以下场景：
  - 静态代码分析。
  - 自动化代码生成和重构。
  - Go 源代码的编程化操作。

通过抽象底层操作，该包帮助开发者生成和增强 Golang 源代码。
*/

// AstBundle is a wrapper for an AST file and its corresponding FileSet.
// AstBundle 是一个封装 AST 文件及其对应 FileSet 的结构体。
type AstBundle struct {
	// fset represents the file set that contains position information for the AST nodes.
	// fset 表示包含 AST 节点位置信息的文件集。
	fset *token.FileSet

	// file is the parsed AST file representation.
	// file 是已解析的 AST 文件表示。
	file *ast.File
}

// NewAstBundle creates a new AstBundle.
// NewAstBundle 创建一个新的 AstBundle 实例。
func NewAstBundle(fileSet *token.FileSet, astFile *ast.File) *AstBundle {
	return &AstBundle{
		fset: fileSet,
		file: astFile,
	}
}

// NewAstBundleV1 creates an AstBundle from a byte slice.
// NewAstBundleV1 根据字节切片创建 AstBundle 实例。
func NewAstBundleV1(data []byte) (*AstBundle, error) {
	// Create a new FileSet and parse the provided data into an AST.
	// 创建新的 FileSet 并将提供的数据解析为 AST。
	return NewAstBundleV2(token.NewFileSet(), data)
}

// NewAstBundleV2 creates an AstBundle from a byte slice with a provided FileSet.
// NewAstBundleV2 根据给定的 FileSet 和字节切片创建 AstBundle 实例。
func NewAstBundleV2(fset *token.FileSet, data []byte) (*AstBundle, error) {
	// Parse the provided data and attach comments to the AST.
	// 解析提供的数据并将注释附加到 AST。
	astFile, err := parser.ParseFile(fset, "", data, parser.ParseComments)
	if err != nil {
		return nil, erero.Wro(err)
	}
	return NewAstBundle(fset, astFile), nil
}

// NewAstBundleV3 creates an AstBundle by parsing a file from the given path.
// NewAstBundleV3 根据文件路径解析并创建 AstBundle 实例。
func NewAstBundleV3(fset *token.FileSet, path string) (*AstBundle, error) {
	// Parse the Go source file at the specified path and attach comments to the AST.
	// 解析指定路径的 Go 源文件，并将注释附加到 AST。
	astFile, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		return nil, erero.Wro(err)
	}
	return NewAstBundle(fset, astFile), nil
}

// NewAstBundleV4 creates an AstBundle from a file path using a new FileSet.
// NewAstBundleV4 使用新的 FileSet 从文件路径创建 AstBundle 实例。
func NewAstBundleV4(path string) (*AstBundle, error) {
	// Use a new FileSet to parse the file at the provided path.
	// 使用新的 FileSet 解析提供路径的文件。
	return NewAstBundleV3(token.NewFileSet(), path)
}

// NewAstBundleV5 creates an AstBundle by parsing a file with a specific parser mode.
// NewAstBundleV5 根据文件路径和特定的解析模式创建 AstBundle 实例。
func NewAstBundleV5(fset *token.FileSet, path string, mode parser.Mode) (*AstBundle, error) {
	// Parse the file at the given path using the specified parser mode.
	// 使用指定的解析模式解析给定路径的文件。
	astFile, err := parser.ParseFile(fset, path, nil, mode)
	if err != nil {
		return nil, erero.Wro(err)
	}
	return NewAstBundle(fset, astFile), nil
}

// NewAstBundleV6 creates an AstBundle from a file path using a new FileSet and specific parser mode.
// NewAstBundleV6 使用新的 FileSet 和指定的解析模式从文件路径创建 AstBundle 实例。
func NewAstBundleV6(path string, mode parser.Mode) (*AstBundle, error) {
	// Create a new FileSet and parse the file with the provided mode.
	// 创建新的 FileSet，并使用提供的解析模式解析文件。
	return NewAstBundleV5(token.NewFileSet(), path, mode)
}

// GetBundle retrieves the AST file and FileSet from the AstBundle.
// GetBundle 从 AstBundle 中获取 AST 文件和 FileSet。
func (ab *AstBundle) GetBundle() (*ast.File, *token.FileSet) {
	// Return the file and file set stored in the AstBundle.
	// 返回存储在 AstBundle 中的文件和文件集。
	fileSet := ab.fset
	astFile := ab.file
	return astFile, fileSet
}

// FormatSource formats the AST back into Go source code.
// FormatSource 将 AST 格式化为 Go 源代码。
func (ab *AstBundle) FormatSource() ([]byte, error) {
	// Use format.Node to convert the AST into Go source code and return it.
	// 使用 format.Node 将 AST 转换为 Go 源代码并返回。
	var buf bytes.Buffer
	if err := format.Node(&buf, ab.fset, ab.file); err != nil {
		return nil, erero.Wro(err)
	}
	return buf.Bytes(), nil
}

// SerializeAst serializes the AST into a textual representation.
// SerializeAst 将 AST 序列化为文本表示形式。
func (ab *AstBundle) SerializeAst() ([]byte, error) {
	// Use printer.Fprint to serialize the AST structure into text and return it.
	// 使用 printer.Fprint 将 AST 结构序列化为文本并返回。
	var buf bytes.Buffer
	if err := printer.Fprint(&buf, ab.fset, ab.file); err != nil {
		return nil, erero.Wro(err)
	}
	return buf.Bytes(), nil
}

// AddImport adds a new import to the AST.
// AddImport 向 AST 添加新的导入路径。
func (ab *AstBundle) AddImport(path string) bool {
	// Add the specified import path to the AST using astutil.AddImport.
	// 使用 astutil.AddImport 向 AST 添加指定的导入路径。
	return astutil.AddImport(ab.fset, ab.file, path)
}

// AddNamedImport adds a new named import to the AST.
// AddNamedImport 向 AST 添加带名称的导入路径。
func (ab *AstBundle) AddNamedImport(name string, path string) bool {
	// Add the specified named import to the AST using astutil.AddNamedImport.
	// 使用 astutil.AddNamedImport 向 AST 添加指定的带名称导入路径。
	return astutil.AddNamedImport(ab.fset, ab.file, name, path)
}

// DeleteImport removes an import from the AST.
// DeleteImport 从 AST 中删除指定的导入路径。
func (ab *AstBundle) DeleteImport(path string) bool {
	// Remove the specified import path from the AST using astutil.DeleteImport.
	// 使用 astutil.DeleteImport 从 AST 中删除指定的导入路径。
	return astutil.DeleteImport(ab.fset, ab.file, path)
}

// DeleteNamedImport removes a named import from the AST.
// DeleteNamedImport 从 AST 中删除指定名称的导入路径。
func (ab *AstBundle) DeleteNamedImport(name string, path string) bool {
	// Remove the specified named import from the AST using astutil.DeleteNamedImport.
	// 使用 astutil.DeleteNamedImport 从 AST 中删除指定名称的导入路径。
	return astutil.DeleteNamedImport(ab.fset, ab.file, name, path)
}

// Print outputs the AST structure to the console.
// Print 将 AST 结构打印到控制台。
func (ab *AstBundle) Print() error {
	// Use ast.Print to print the entire AST structure to the console for debugging.
	// 使用 ast.Print 将整个 AST 结构打印到控制台以便调试。
	return ast.Print(ab.fset, ab.file)
}

// GetPackageName retrieves the package name from the AST.
// GetPackageName 从 AST 中获取包名。
func (ab *AstBundle) GetPackageName() string {
	// Access the package name from the root node of the AST.
	// 从 AST 的根节点访问包名。
	return ab.file.Name.Name
}
