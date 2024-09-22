package syntaxgo_reflect

import (
	"reflect"

	"github.com/yyle88/done"
)

func GetTypes(objects []any) []reflect.Type {
	var results = make([]reflect.Type, 0, len(objects))
	for _, a := range objects {
		results = append(results, done.Nice(reflect.TypeOf(a)))
	}
	return results
}

func GetPkgPaths(objectsTypes []reflect.Type) []string {
	var results = make([]string, 0, len(objectsTypes))
	for _, a := range objectsTypes {
		results = append(results, done.Nice(a.PkgPath()))
	}
	return results
}

// GetPkgPathsToImportWithQuotes 就是根据类型获取到需要引用(import)的包的路径，因为import需要带双引号，这里就写个逻辑给它加上双引号吧
// 这个功能还挺常用的，特别是在自动生成代码的时候，假如不能补全import的包，在某些场景里执行代码format就会很慢（比如某些IDE或者某些自动格式化的程序的自动格式化就会很慢）
// 很明显这里的 "4" 就是 "For" 的意思，但我认为写个 "4" 会更加洋气些
// 目前知道的获取包名最好的方案就是通过对象，其次才是通过硬编码字符串，因此这个函数十分滴珍贵，需要好好利用起来
func GetPkgPathsToImportWithQuotes(objectsTypes []reflect.Type) []string {
	var packagePaths = GetPkgPaths(objectsTypes)

	var results = make([]string, 0, len(packagePaths))
	for _, path := range packagePaths {
		results = append(results, `"`+path+`"`)
	}
	return results
}
