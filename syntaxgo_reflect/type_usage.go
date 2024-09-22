package syntaxgo_reflect

import (
	"path/filepath"
	"reflect"
)

// GetUsageCode 获取在其它包调用某包类型的代码，比如包名是 abc 而类型名是 Demo 则在其它包调用时就是 abc.Demo 这样的，因此这个操作也是非常重要的
func GetUsageCode(a reflect.Type) string {
	name := a.Name()
	if pkgPath := a.PkgPath(); pkgPath != "" {
		pkgName := filepath.Base(pkgPath)
		return pkgName + "." + name
	} else {
		return name
	}
}
