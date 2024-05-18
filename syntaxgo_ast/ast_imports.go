package syntaxgo_ast

import (
	"go/token"
	"reflect"
	"slices"

	"github.com/yyle88/done"
	"github.com/yyle88/syntaxgo/internal/utils"
	"github.com/yyle88/syntaxgo/syntaxgo_reflect"
)

// AddImportsOfPackages 把需要引用的包路径增加到代码里
// 这个函数非常重要，因为有时候就是找不到包名，而有时候有重复的包名，比如"errors"和"github.com/pkg/errors"，而即使是有唯一的包，让代码自动去格式化和找就会非常的耗时
// 因此推荐就是在生成代码时同时也把要引用的都添加进来，这样代码格式化就会非常快
// 因此在这个文件里，我定义了不同的设置包名的函数，因为这个确实是非常的重要
func AddImportsOfPackages(source []byte, packages []string) []byte {
	astFile := done.VCE(NewAstFromSource(source)).Nice()
	utils.AssertBooleanOK(astFile.Package.IsValid()) //没有定义包名的不能使用该功能-即不能补充需要的引用
	utils.AssertBooleanOK(astFile.Name != nil)       //没有定义包名的不能使用该功能-即不能补充需要的引用

	// 把要导入的包设置为map
	var missMap = make(map[string]bool)
	for _, pkgPath := range packages {
		missMap[utils.SetDoubleQuotes(pkgPath)] = true
	}

	// 遍历引用的包，删除已经存在包，map里剩下的包才需要导入到代码里
	for _, one := range astFile.Imports {
		delete(missMap, one.Path.Value)
	}

	if len(missMap) > 0 {
		pkg2quotes := utils.GetMapKeys(missMap)
		slices.Sort(pkg2quotes)

		ptx := utils.NewPTX()
		ptx.Println()         //需要换行符
		if len(missMap) < 2 { //很明显，当是1个的时候，只需要补充一行就行
			//虽然没什么必要循环但循环也是可以的，就是取首个元素就行
			for _, pkg2quote := range pkg2quotes {
				ptx.Println("import", pkg2quote)
			}
		} else {
			ptx.Println("import (")
			for _, pkg2quote := range pkg2quotes {
				ptx.Println("    " + pkg2quote) //把带有双引号的直接写到里面
			}
			ptx.Println(")")
		}
		ptx.Println() //需要换行符

		//假如包名是 package example 就找到包名后面的换行符，把新增的 import 内容写到包名后面，前面已经添加过换行符
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

type AddImportsParam struct {
	Packages   []string       //直接设置包路径列表
	UsingTypes []reflect.Type //设置反射类型，通过类型能找到包路径
	Objects    []any          //设置要引用的对象列表(非指针对象)，通过对象也能找到对象的包路径
}

// AddImports 根据要使用的类型，得到要引用的包路径，把要引用的包设置到代码里，返回修改后的代码
func AddImports(source []byte, param *AddImportsParam) []byte {
	packagePaths := utils.SafeMerge(
		param.Packages,
		syntaxgo_reflect.GetPkgPaths(param.UsingTypes),
		syntaxgo_reflect.GetPkgPaths(syntaxgo_reflect.GetObjectsTypes(param.Objects)),
	)

	return AddImportsOfPackages(source, packagePaths)
}
