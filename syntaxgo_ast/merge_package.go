package syntaxgo_ast

import (
	"go/ast"
	"go/parser"
	"go/token"

	"github.com/yyle88/erero"
	"github.com/yyle88/syntaxgo/internal/utils"
)

/*
Package `syntaxgo_ast` provides tools for working with Go source code Abstract Syntax Trees (AST). This file primarily defines functions for parsing Go packages within a directory and merging their files into a single `AstBundle` instance.

Key features include:
  - Parsing all Go packages in a directory and generating a mapping with package information.
  - Merging all files of a specified package into a single AST, supporting operations for a single package or subpackages.
  - Providing convenient interfaces to handle collections of Go files in a directory/package.

It is important to note that some of the logic in this file relies on the deprecated `ast.Package` type and related methods (such as `ast.MergePackageFiles`). This means that future versions of Go may remove these features, and therefore, gradual refactoring is needed to accommodate newer versions.
*/

/*
Package `syntaxgo_ast` 提供了处理 Go 源代码语法树 (AST) 的工具。当前文件主要定义了一些函数，用于解析目录中的 Go 包，将其中的文件合并为一个 `AstBundle` 实例。

主要功能包括：
  - 解析目录下的所有 Go 包，生成包含包信息的映射。
  - 合并指定包的所有文件为一个语法树，既支持单一包也支持选择子包的操作。
  - 提供便捷接口以处理目录下的 Go 文件集合。

需要注意的是，本文件的部分逻辑使用了已被标记为过时的 `ast.Package` 类型以及相关方法（例如 `ast.MergePackageFiles`）。这意味着未来的 Go 版本可能会移除这些功能，因此需要逐步重构以适应新版本。
*/

// MergeOnePackageFiles merges all Go source files of a specific package within a given directory.
// MergeOnePackageFiles 合并指定目录下某个包的所有 Go 源文件，生成一个语法树。
// Deprecated: This function uses the deprecated `ast.Package` and `ast.MergePackageFiles` methods.
// Note: This has limited usefulness and may need refactoring in the future.
func MergeOnePackageFiles(root string, packageName string) (*AstBundle, error) {
	var fileSet = token.NewFileSet()
	// Parse the directory and retrieve a map of package names to package information.
	// 解析目录，获取包名到包信息的映射。
	packagesMap, err := ParseRootGetPackages(fileSet, root)
	if err != nil {
		return nil, erero.Wro(err)
	}

	// Check if the specified package name exists in the parsed data.
	// 检查解析结果中是否存在指定的包名。
	pkg, ok := packagesMap[packageName]
	if !ok {
		return nil, erero.Errorf("no package name = %s in root path: %s", packageName, root)
	}

	// Merge the package files into a single AST bundle.
	// 将该包的所有文件合并为单一语法树。
	res := NewAstBundle(fileSet, ast.MergePackageFiles(pkg, ast.FilterImportDuplicates))
	return res, nil
}

// ParseRootGetPackages parses the entire directory and retrieves a map of package names to package information.
// ParseRootGetPackages 解析指定目录下的所有 Go 包，返回包名到包信息的映射。
// Deprecated: This function uses the deprecated `ast.Package` type and should be updated in the future.
// Note: The function name could be more descriptive of its purpose.
func ParseRootGetPackages(fset *token.FileSet, root string) (map[string]*ast.Package, error) {
	packagesMap, err := parser.ParseDir(
		fset,
		root,
		utils.IsGoSourceFile, // Filters to include only Go source files.
		parser.ParseComments,
	)
	if err != nil {
		return nil, erero.Wro(err)
	}
	return packagesMap, nil
}

// MergeSubPackageFiles merges all Go source files of the only package in the given directory.
// If there is more than one package, it returns an error.
// MergeSubPackageFiles 合并指定目录下唯一包的所有 Go 源文件，若目录中存在多个包则返回错误。
// Deprecated: This function uses the deprecated `ast.Package` type and `ast.MergePackageFiles` method.
func MergeSubPackageFiles(root string) (*AstBundle, error) {
	var fileSet = token.NewFileSet()
	// Parse the directory and retrieve a map of package names to package information.
	// 解析目录，获取包名到包信息的映射。
	packagesMap, err := ParseRootGetPackages(fileSet, root)
	if err != nil {
		return nil, erero.Wro(err)
	}

	// If multiple packages are found, return an error.
	// 如果目录中存在多个包，则返回错误。
	if len(packagesMap) >= 2 {
		return nil, erero.Errorf("more than one package in root path: %s", root)
	}

	// Iterate over the map and merge the first package's files into a single AST bundle.
	// 遍历映射，将第一个包的所有文件合并为单一语法树。
	for _, pkg := range packagesMap {
		res := NewAstBundle(fileSet, ast.MergePackageFiles(pkg, ast.FilterImportDuplicates))
		return res, nil
	}

	// If no packages are found, return an error.
	// 如果没有找到任何包，则返回错误。
	return nil, erero.Errorf("no packages in root path: %s", root)
}
