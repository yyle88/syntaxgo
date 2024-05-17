package syntaxgo_ast

import (
	"go/token"
	"reflect"

	"github.com/yyle88/done"
	"github.com/yyle88/syntaxgo/internal/utils"
	"github.com/yyle88/syntaxgo/syntaxgo_reflect"
)

// AddImportsOfTypes 根据要使用的类型，得到要引用的包路径，把要引用的包设置到代码里，返回修改后的代码
func AddImportsOfTypes(filename string, source []byte, usingTypes []reflect.Type) []byte {
	return AddImportsOfPackages(filename, source, syntaxgo_reflect.GetPkgPaths(usingTypes))
}

// AddImportsOfObjects 根据要使用的对象，得到要引用的包路径，把要引用的包设置到代码里，返回修改后的代码
func AddImportsOfObjects(filename string, source []byte, objects []any) []byte {
	return AddImportsOfTypes(filename, source, syntaxgo_reflect.GetObjectsTypes(objects))
}

// AddImportsOfPackages 把需要引用的包路径增加到代码里
// 这个函数非常重要，因为有时候就是找不到包名，而有时候有重复的包名，比如"errors"和"github.com/pkg/errors"，而即使是有唯一的包，让代码自动去格式化和找就会非常的耗时
// 因此推荐就是在生成代码时同时也把要引用的都添加进来，这样代码格式化就会非常快
// 因此在这个文件里，我定义了不同的设置包名的函数，因为这个确实是非常的重要
func AddImportsOfPackages(filename string, source []byte, packages []string) []byte {
	astFile := done.VCE(NewAstFromSource(filename, source)).Nice()
	utils.BooleanOK(astFile.Package.IsValid()) //没有定义包名的不能使用该功能-即不能补充需要的引用
	utils.BooleanOK(astFile.Name != nil)       //没有定义包名的不能使用该功能-即不能补充需要的引用

	// 把要导入的包设置为map
	var wantMap = make(map[string]bool)
	for _, pkgPath := range packages {
		wantMap[utils.SetDoubleQuotes(pkgPath)] = true
	}

	// 遍历引用的包，删除已经存在包，map里剩下的包才需要导入到代码里
	for _, imp := range astFile.Imports {
		delete(wantMap, imp.Path.Value)
	}

	if len(wantMap) > 0 {
		ptx := utils.NewPTX()

		ptx.Println("") //需要换行符

		if len(wantMap) < 2 { //很明显，当是1个的时候，只需要补一行就行
			ptx.Println("import", utils.AnyKeyInMap(wantMap))
		} else {
			ptx.Println("import (")
			for pkg2quote := range wantMap {
				ptx.Println(pkg2quote) //把带有双引号的直接写到里面
			}
			ptx.Println(")")
		}

		ptx.Println("") //需要换行符

		//从包名结尾初开始，向后面找包名定义后面的换行符，假如找到文件末尾也退出
		//正常的是 node 都是 [pos-1:end-1] 而且右边也是不包含的，因此我这里定的其实位置就是包名的正后面
		//注意
		//我们允许包名右边有行注释，假如你后面是个块注释比如 package xx /* xxx 换行 xxx */ 就会有个小问题
		//我就会把代码写到注释里啦
		//我不想再写逻辑单独处理这种情况，因为我感觉这种场景很少见，我建议你修改自己的源代码文件，使其符合规范
		posIdx := int(astFile.Name.End() - 1)
		for posIdx < len(source) && source[posIdx] != ('\n') {
			posIdx++
		}
		//这里退出的时候，这坐标要么是在回车符，要么是在文件末尾，即 len 的位置
		//接下来就是插入新代码
		//因此我们的新 node 的 pos 和 end 也要写成 idx+1 的位置，这样 [pos-1:end-1] 就还是在 [idx:idx] 这个区间
		//假设找不到，这时候 idx==len，让 pos 和 end 都等于 len+1，这样 [pos-1:end-1] 就还是在 [len:len] 的区间
		var node = NewNode(token.Pos(posIdx+1), token.Pos(posIdx+1))
		//我在这里写好几个换行符，这会导致源码不美观，但是没关系的，最后统一执行 format 格式化再写源码就行
		//同样的，在已经有 import 块时，我会再另起一个新的 import 的块，这样虽然是不美观，但依然是需要依靠格式化解决
		//目前对于我来说，这个已经是够用的
		//已经能解决因为没有写 import 而导致的 format 时自动补充引用包，但执行特别慢的问题
		source = ChangeNodeBytesXNewLines(source, node, ptx.Bytes(), 2)
	}
	return source
}

// AddImportsOfPackagesAndTypes 有些包里只有函数或者接口，没有类型，这时候就得不到它的包路径，因此还是得用字符串把包名传进来
func AddImportsOfPackagesAndTypes(filename string, source []byte, packages []string, usingTypes []reflect.Type) []byte {
	return AddImportsOfPackages(filename, source, utils.SafeMerge(packages, syntaxgo_reflect.GetPkgPaths(usingTypes)))
}

// AddImportsOfPackagesAndObjects 有些包里只有函数或者接口，没有类型，这时候就得不到它的包路径，因此还是得用字符串把包名传进来
func AddImportsOfPackagesAndObjects(filename string, source []byte, packages []string, objects []any) []byte {
	return AddImportsOfPackagesAndTypes(filename, source, packages, syntaxgo_reflect.GetObjectsTypes(objects))
}
