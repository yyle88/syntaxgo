package syntaxgo_reflect

import (
	"path/filepath"
	"reflect"

	"github.com/yyle88/tern"
)

// GenerateTypeUsageCode generates the code for using a type from another package.
// It constructs the code representation for the type as it would be used in another package,
// including the package name and type name. If the type is from the same package, it returns
// just the type name.
//
// For example, if the type is "Demo" from package "abc", this function will return "abc.Demo".
// If the package path is empty, it simply returns the type name.
//
// GenerateTypeUsageCode 用于生成从其他包调用某个包类型的代码。
// 它构造了类型在其他包中的使用代码，包括包名和类型名。
// 如果类型来自相同的包，则只返回类型名。
//
// 举个例子，如果类型是来自包 "abc" 的 "Demo"，这个函数将返回 "abc.Demo"。
// 如果包路径为空，则只返回类型名。
func GenerateTypeUsageCode(a reflect.Type) string {
	// Get the package path of the type.
	// 获取类型的包路径
	pkgPath := a.PkgPath()

	// Use tern.BFF to conditionally return the fully qualified type name or just the type name
	// based on whether the package path is available.
	// 使用 tern.BFF 条件性地返回完整的类型名或仅返回类型名，取决于是否有包路径。
	return tern.BFF(pkgPath != "", func() string {
		// If package path is available, return "packageName.TypeName".
		// 如果包路径可用，返回 "包名.类型名"。
		return filepath.Base(pkgPath) + "." + a.Name()
	}, func() string {
		// If package path is empty, just return the type name.
		// 如果包路径为空，返回类型名。
		return a.Name()
	})
}
