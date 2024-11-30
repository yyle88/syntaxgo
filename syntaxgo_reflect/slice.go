package syntaxgo_reflect

import (
	"reflect"

	"github.com/yyle88/done"
)

// GetTypes extracts the reflect.Type for each objects.
// GetTypes 从给定的对象列表中提取每个对象的 reflect.Type。
func GetTypes(objects []any) []reflect.Type {
	var results = make([]reflect.Type, 0, len(objects))
	for _, a := range objects {
		// Append the reflect.Type of each object to the results slice.
		// 将每个对象的 reflect.Type 添加到结果切片中。
		results = append(results, done.Nice(reflect.TypeOf(a)))
	}
	return results
}

// GetPkgPaths returns the package paths of the provided reflect.Type objects.
// GetPkgPaths 返回给定 reflect.Type 对象的包路径。
func GetPkgPaths(objectsTypes []reflect.Type) []string {
	var results = make([]string, 0, len(objectsTypes))
	for _, a := range objectsTypes {
		// Append the package path of each reflect.Type to the results slice.
		// 将每个 reflect.Type 的包路径添加到结果切片中。
		results = append(results, done.Nice(a.PkgPath()))
	}
	return results
}

// GetQuotedPackageImportPaths generates the import paths for the provided types, with the necessary quotes for import statements.
// GetQuotedPackageImportPaths 生成给定类型的导入路径，再为导入语句加上必要的双引号。
func GetQuotedPackageImportPaths(objectsTypes []reflect.Type) []string {
	// Retrieve the package paths for each type.
	// 获取每个类型的包路径。
	var packagePaths = GetPkgPaths(objectsTypes)

	var results = make([]string, 0, len(packagePaths))
	for _, path := range packagePaths {
		// Add double quotes around each package path for use in import statements.
		// 给每个包路径添加双引号，供导入语句使用。
		results = append(results, `"`+path+`"`)
	}
	return results
}
