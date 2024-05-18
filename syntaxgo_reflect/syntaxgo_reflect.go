package syntaxgo_reflect

import (
	"path/filepath"
	"reflect"

	"github.com/yyle88/done"
)

// GetPkgPath 在 Go 语言中，通过反射是无法直接获取函数的包名的。
// reflect.TypeOf(a).PkgPath() 只适用于结构体对象，而不能用于结构体指针，而不适用于指针/函数/方法类型。
// 接口假如存的是非指针对象也只能解出对象的类型而非接口的类型，接口假如存的是对象指针还是不能得到，解析结果是依据存的数据类型，能不能解出也是看是否符合前面的规则
// 函数本身没有一个与之关联的类型，因此无法通过反射获取其包的完整导入路径。
// 如果你需要获取函数所属的包名，最好的方式是直接在代码中引用函数。
func GetPkgPath(a any) string {
	return reflect.TypeOf(a).PkgPath()
}

// GetPkgPathV2 就是 GetPkgPath 的范型版啦，因为最开始是没有范型的，而其后Go支持泛型
// 认为非泛型版也是很有用的（当对象是从函数传过来的，而外部函数非泛型时，而外部函数非常底层没法逐级加泛型时，或者对象过多时，使用非范型版都是有优势的）
// 因此保留非范型版而把范型版命名为V2，毕竟只添加两个字符就能解决问题，还能保证含义不变，V2也还行
func GetPkgPathV2[T any]() string {
	return reflect.TypeOf(GetObject[T]()).PkgPath()
}

func GetPkgName(a any) string {
	var pkgPath = GetPkgPath(a)
	if pkgPath == "" {
		return ""
	}
	return filepath.Base(pkgPath)
}

func GetPkgNameV2[T any]() string {
	return GetPkgName(GetObject[T]())
}

func GetObject[T any]() (a T) {
	return a //TODO 目前暂时不知道如何在编译阶段就能阻止类型传指针，即"[*A]"这样的，因为通过指针对象得不到类型
}

func GetTypeName(a any) string {
	return reflect.TypeOf(a).Name()
}

func GetTypeNameV2[T any]() string {
	return GetTypeName(GetObject[T]())
}

// GetTypeUsageCode 获取在其它包调用某包类型的代码，比如包名是 abc 而类型名是 Demo 则在其它包调用时就是 abc.Demo 这样的，因此这个操作也是非常重要的
func GetTypeUsageCode(a any) string {
	objectType := reflect.TypeOf(a)
	goTypeName := objectType.Name()
	if pkgPath := objectType.PkgPath(); pkgPath != "" {
		pkgName := filepath.Base(pkgPath)
		return pkgName + "." + goTypeName
	} else {
		return goTypeName
	}
}

func GetTypeUsageCodeV2[T any]() string {
	return GetTypeUsageCode(GetObject[T]())
}

func GetObjectsTypes(objects []any) []reflect.Type {
	var objectsTypes = make([]reflect.Type, 0, len(objects))
	for _, a := range objects {
		objectsTypes = append(objectsTypes, done.Nice(reflect.TypeOf(a)))
	}
	return objectsTypes
}

func GetPkgPaths(objectsTypes []reflect.Type) []string {
	var packagePaths = make([]string, 0, len(objectsTypes))
	for _, a := range objectsTypes {
		packagePaths = append(packagePaths, done.Nice(a.PkgPath()))
	}
	return packagePaths
}

// GetPkgPaths4Imports 就是根据类型获取到需要引用(import)的包的路径，因为import需要带双引号，这里就写个逻辑给它加上双引号吧
// 这个功能还挺常用的，特别是在自动生成代码的时候，假如不能补全import的包，在某些场景里执行代码format就会很慢（比如某些IDE或者某些自动格式化的程序的自动格式化就会很慢）
// 很明显这里的 "4" 就是 "For" 的意思，但我认为写个 "4" 会更加洋气些
// 目前知道的获取包名最好的方案就是通过对象，其次才是通过硬编码字符串，因此这个函数十分滴珍贵，需要好好利用起来
func GetPkgPaths4Imports(objectsTypes []reflect.Type) []string {
	var packagePaths = GetPkgPaths(objectsTypes)

	var results = make([]string, 0, len(packagePaths))
	for _, path := range packagePaths {
		results = append(results, `"`+path+`"`)
	}
	return results
}
