package syntaxgo_astfieldsflat

import "strings"

type StatementParts []string //当得到参数列表/返回列表时，假如列表字段总以逗号分隔的形式出现在代码里，就使用这个类型

func (parts StatementParts) MergeParts() string {
	return strings.Join(parts, ", ")
}

type StatementLines []string //当得到赋值语句/返回语句/函数调用语句时，这些语句间往往是以换行符分隔的，因此放在这里面

func (lines StatementLines) MergeLines() string {
	return strings.Join(lines, "\n")
}
