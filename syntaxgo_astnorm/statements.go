package syntaxgo_astnorm

import "strings"

// StatementParts represents a list of strings that are usually separated by commas.
// StatementParts 通常用于表示一个以逗号分隔的字符串列表，常见于参数列表或返回值列表。
type StatementParts []string

// MergeParts joins the elements in StatementParts with a comma and a space, and returns the resulting string.
// MergeParts 将 StatementParts 中的元素用逗号和空格连接，并返回结果字符串。
func (stmts StatementParts) MergeParts() string {
	return strings.Join(stmts, ", ")
}

// StatementLines represents a list of strings that are usually separated by newline characters.
// StatementLines 通常用于表示一个以换行符分隔的字符串列表，常见于赋值语句、返回语句或函数调用语句。
type StatementLines []string

// MergeLines joins the elements in StatementLines with a newline character, and returns the resulting string.
// MergeLines 将 StatementLines 中的元素用换行符连接，并返回结果字符串。
func (stmts StatementLines) MergeLines() string {
	return strings.Join(stmts, "\n")
}
