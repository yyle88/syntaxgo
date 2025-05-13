package syntaxgo_ast

import (
	"reflect"
	"slices"
	"strings"

	"github.com/yyle88/done"
	"github.com/yyle88/must"
	"github.com/yyle88/syntaxgo/internal/utils"
	"github.com/yyle88/syntaxgo/syntaxgo_astnode"
	"github.com/yyle88/syntaxgo/syntaxgo_reflect"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
	"golang.org/x/exp/maps"
)

/*
Package `syntaxgo_ast` provides tools for managing and manipulating Go source code Abstract Syntax Trees (ASTs). It contains functions for managing package import paths, injecting necessary imports into Go source files, and generating import statements.

Key features include:
  - Managing and merging import paths from various sources such as package paths, referenced types, and inferred objects.
  - Injecting necessary imports into Go source code automatically, ensuring the correct packages are imported.
  - Generating formatted import statements for Go source files.

The package provides flexibility for working with Go source code, enabling automated import management and AST manipulation.

Note: This package relies on internal utilities from `syntaxgo_reflect` and `utils`, and it utilizes AST manipulation techniques to interact with Go source code.
*/

/*
Package `syntaxgo_ast` 提供了管理和操作 Go 源代码抽象语法树（AST）的工具。它包含了管理包导入路径、将必要的导入插入 Go 源代码、以及生成导入语句的功能。

主要功能包括：
  - 从多个源头里（如包路径、引用类型和推断对象）管理和合并导入路径。
  - 自动将需要引用的包插入 Go 源代码，确保正确的包被导入代码。
  - 再为 Go 源文件生成格式化的导入语句。

该包为与 Go 源代码的交互提供了灵活性，能够自动化管理导入和进行 AST 操作。

注意：该包依赖于 `syntaxgo_reflect` 和 `utils` 中的内部工具，利用 AST 操作技术高效地与 Go 源代码交互。
*/

// PackageImportOptions is used to manage the import paths for Go packages.
// PackageImportOptions 用于管理 Go 包的导入路径。
type PackageImportOptions struct {
	packages        []string       // List of package paths. // 直接设置包路径列表
	referencedTypes []reflect.Type // List of referenced types to find package paths. // 设置反射类型，通过类型能找到包路径
	inferredObjects []any          // List of inferred objects to find package paths. // 设置要引用的对象列表(非指针对象)，通过对象也能找到对象的包路径
}

// NewPackageImportOptions creates and returns a new PackageImportOptions instance.
// NewPackageImportOptions 创建并返回一个新的 PackageImportOptions 实例。
func NewPackageImportOptions() *PackageImportOptions {
	return &PackageImportOptions{}
}

// SetPkgPath adds a package path to the list of packages to be imported.
// SetPkgPath 将包路径添加到待导入的包列表中。
func (param *PackageImportOptions) SetPkgPath(pkgPath string) *PackageImportOptions {
	param.packages = append(param.packages, pkgPath)
	return param
}

// SetReferencedType adds a referenced type to the list of referenced types.
// SetReferencedType 将一个引用类型添加到引用类型列表中。
func (param *PackageImportOptions) SetReferencedType(reflectType reflect.Type) *PackageImportOptions {
	param.referencedTypes = append(param.referencedTypes, reflectType)
	return param
}

// SetInferredObject adds an inferred object to the list of inferred objects.
// SetInferredObject 将一个推断对象添加到推断对象列表中。
func (param *PackageImportOptions) SetInferredObject(object any) *PackageImportOptions {
	param.inferredObjects = append(param.inferredObjects, object)
	return param
}

// GetPkgPaths returns a merged list of package paths from packages, referenced types, and inferred objects.
// GetPkgPaths 返回从包路径、引用类型和推断对象中合并得到的包路径列表。
func (param *PackageImportOptions) GetPkgPaths() []string {
	return utils.SafeMerge(
		param.packages,
		syntaxgo_reflect.GetPkgPaths(param.referencedTypes),
		syntaxgo_reflect.GetPkgPaths(syntaxgo_reflect.GetTypes(param.inferredObjects)),
	)
}

// InjectImports adds necessary import paths into the provided Go source code.
// InjectImports 将必要的导入路径添加到提供的 Go 源代码中。
func (param *PackageImportOptions) InjectImports(source []byte) []byte {
	return InjectImports(source, param.GetPkgPaths())
}

// CreateImports generates a string containing import statements for the given package paths.
// CreateImports 根据给定的包路径生成包含导入语句的字符串。
func (param *PackageImportOptions) CreateImports() string {
	return CreateImports(param.GetPkgPaths())
}

// CreateImports generates the import statements as a string from the provided package paths.
// CreateImports 从提供的包路径生成导入语句的字符串。
func CreateImports(imports []string) string {
	if len(imports) == 0 {
		zaplog.LOG.Debug("imports is none") // If no imports, still proceed. Even an empty "import ()" block is valid. // 如果没有导入，依然执行。即使是空的 "import ()" 块也是有效的。
	}
	var pkg2quotes []string
	var mp = map[string]bool{}
	// Iterate over each import path and format them with quotes if necessary.
	// 遍历每个导入路径，并在必要时用双引号格式化它们。
	for _, sub := range imports {
		if sub == "" {
			continue
		}
		if !strings.HasPrefix(sub, `"`) {
			sub = `"` + sub
		}
		if !strings.HasSuffix(sub, `"`) {
			sub = sub + `"`
		}
		if !mp[sub] {
			mp[sub] = true
			pkg2quotes = append(pkg2quotes, sub)
		}
	}
	ptx := utils.NewPTX()
	ptx.Println("import (")
	// Add each package path to the import block.
	// 将每个包路径添加到导入块中。
	for _, s := range pkg2quotes {
		ptx.Println(s)
	}
	ptx.Println(")")
	return ptx.String()
}

// InjectImports inserts the missing import paths into the provided Go source code.
// InjectImports 将缺失的导入路径插入到提供的 Go 源代码中。
func InjectImports(source []byte, packages []string) []byte {
	astBundle := done.VCE(NewAstBundleV1(source)).Nice()
	astFile := astBundle.file
	must.TRUE(astFile.Package.IsValid()) // Ensure the package is valid for importing. // 确保包是有效的才能进行导入。
	must.TRUE(astFile.Name != nil)       // Ensure the file has a valid package name. // 确保文件具有有效的包名。

	// Initialize a map to track the missing package paths.
	// 初始化一个映射来跟踪缺失的包路径。
	var missMap = make(map[string]bool)
	for _, pkgPath := range packages {
		if pkgPath == "" {
			zaplog.LOG.Warn("skip an empty pkg_path")
			continue
		}
		if strings.Contains(pkgPath, `"`) { // Avoid paths with double quotes in them. // 避免路径中包含双引号。
			zaplog.LOG.Warn("skip an wrong pkg_path contains double quotes", zap.String("pkg_path", pkgPath))
			continue
		}
		missMap[utils.SetDoubleQuotes(pkgPath)] = true
	}

	// Remove any packages that are already present in the imports.
	// 删除已经存在于导入中的包。
	for _, one := range astFile.Imports {
		delete(missMap, one.Path.Value)
	}

	if len(missMap) > 0 {
		pkg2quotes := maps.Keys(missMap)
		slices.Sort(pkg2quotes) // Sort the package paths to maintain stability. // 排序包路径以保持稳定性。

		ptx := utils.NewPTX()
		ptx.Println()         // Print a newline for formatting. // 打印换行符以进行格式化。
		if len(missMap) < 2 { // If there is only one missing import, print it directly. // 如果只有一个缺失的导入，直接打印它。
			for _, pkg2quote := range pkg2quotes {
				ptx.Println("import", pkg2quote)
			}
		} else {
			ptx.Println("import (")
			for _, pkg2quote := range pkg2quotes {
				ptx.Println("    " + pkg2quote) // Indent the imports for better readability. // 缩进导入路径以提高可读性。
			}
			ptx.Println(")")
		}
		ptx.Println() // Print a newline at the end. // 在结尾打印换行符。

		// Find the position to insert the new import statements.
		// 找到插入新导入语句的位置。
		posIdx := int(astFile.Name.End() - 1)
		for posIdx < len(source) && source[posIdx] != ('\n') {
			posIdx++
		}

		// Create a new node to represent the inserted import statements.
		// 创建一个新的节点来表示插入的导入语句。
		var node = syntaxgo_astnode.NewNodeV2(posIdx+1, posIdx+1)
		source = syntaxgo_astnode.ChangeNodeCodeSetSomeNewLines(source, node, ptx.Bytes(), 2)
	}
	return source
}
