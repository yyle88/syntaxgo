package syntaxgo_reflect

import (
	"path/filepath"
	"reflect"
)

// GetPkgPath 在 Go 语言中，通过反射是无法直接获取函数的包名的。
// reflect.TypeOf(a).PkgPath() 只适用于结构体对象，而不能用于结构体指针，而不适用于指针/函数/方法类型。
// 接口假如存的是非指针对象也只能解出对象的类型而非接口的类型，接口假如存的是对象指针还是不能得到，解析结果是依据存的数据类型，能不能解出也是看是否符合前面的规则
// 函数本身没有一个与之关联的类型，因此无法通过反射获取其包的完整导入路径。
// 如果你需要获取函数所属的包名，最好的方式是直接在代码中引用函数。
func GetPkgPath(a any) string {
	return reflect.TypeOf(a).PkgPath()
}

// GetPkgPathV2 就是 GetPkgPath 的泛型版啦，因为最开始是没有泛型的，而其后Go支持泛型
// 认为非泛型版也是很有用的（当对象是从函数传过来的，而外部函数非泛型时，而外部函数非常底层没法逐级加泛型时，或者对象过多时，使用非泛型版都是有优势的）
// 因此保留非泛型版而把泛型版命名为V2，毕竟只添加两个字符就能解决问题，还能保证含义不变，V2也还行
func GetPkgPathV2[T any]() string {
	return reflect.TypeOf(GetObject[T]()).PkgPath()
}

func GetPkgPathV3(a any) string {
	return GetTypeV3(a).PkgPath()
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

func GetPkgNameV3(a any) string {
	var pkgPath = GetPkgPathV3(a)
	if pkgPath == "" {
		return ""
	}
	return filepath.Base(pkgPath)
}

func GetType(a any) reflect.Type {
	return reflect.TypeOf(a)
}

func GetTypeV2[T any]() reflect.Type {
	return GetType(GetObject[T]())
}

// GetTypeV3 获取类型信息，由于有的时候会传对象而有的时候会传指针，因此这里做个简单的适配
func GetTypeV3(object any) reflect.Type {
	if vtp := reflect.TypeOf(object); vtp.Kind() == reflect.Ptr {
		// Elem() panics if the type's Kind is not Array, Chan, Map, Pointer, or Slice.
		return vtp.Elem()
	} else {
		return vtp
	}
}

func GetTypeName(a any) string {
	return reflect.TypeOf(a).Name()
}

func GetTypeNameV2[T any]() string {
	return GetTypeName(GetObject[T]())
}

// GetTypeNameV3 获取类型名称，由于有的时候会传对象而有的时候会传指针，因此这里做个简单的适配
func GetTypeNameV3(object any) string {
	return GetTypeV3(object).Name()
}
