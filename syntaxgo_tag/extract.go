package syntaxgo_tag

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/pkg/errors"
)

// ExtractTagValue 提取标签中的特定部分，比如 `gorm:"" json:""` 的 gorm 整体 或者 json 整体
func ExtractTagValue(tag, key string) string {
	// 正则表达式查找键值对
	// \b 确保匹配整个单词，避免部分匹配
	// \s 空白
	// 该正则的小括号没有把空白括起来，因此结果不包含空白，即不包含空格部分
	regex := regexp.MustCompile(`\b` + regexp.QuoteMeta(key) + `\s*:\s*"([^"]*)"`)
	if matches := regex.FindStringSubmatch(tag); len(matches) > 1 {
		return matches[1]
	}
	return ""
}

type ExtractTagFieldAction string

//goland:noinspection GoSnakeCaseUsage
const (
	EXCLUDE_WHITESPACE_PREFIX ExtractTagFieldAction = "EXCLUDE_WHITESPACE_PREFIX" //以前传 TRUE  的地方
	INCLUDE_WHITESPACE_PREFIX ExtractTagFieldAction = "INCLUDE_WHITESPACE_PREFIX" //以前传 FALSE 的地方
)

// ExtractTagField 提取标签中的特定部分，比如 gorm 里面的 column:name 这部分，Fields Tags
func ExtractTagField(part, fieldName string, action ExtractTagFieldAction) string {
	// 正则表达式查找键值对
	// \b 确保匹配整个单词，避免部分匹配
	// 没有引号的键值对情况
	// \s 空白
	// 认为存在需要完整匹配 和 排除前导空格 两种需求，通常在写标签内部的k:v时不应该添加前导空格
	var regex *regexp.Regexp
	if action == EXCLUDE_WHITESPACE_PREFIX {
		regex = regexp.MustCompile(`\b` + regexp.QuoteMeta(fieldName) + `\s*:\s*([^;]+)`)
	} else if action == INCLUDE_WHITESPACE_PREFIX {
		regex = regexp.MustCompile(`\b` + regexp.QuoteMeta(fieldName) + `\s*:([^;]+)`)
	} else {
		panic(errors.New("WRONG"))
	}
	if matches := regex.FindStringSubmatch(part); len(matches) > 1 {
		return matches[1]
	}
	return ""
}

func ExtractTagValueIndex(tag, key string) (string, int, int) {
	// 正则表达式查找键值对
	regex := regexp.MustCompile(`\b` + regexp.QuoteMeta(key) + `\s*:\s*"([^"]*)"`)
	// 查找键值对及其位置
	if indexes := regex.FindStringSubmatchIndex(tag); len(indexes) > 3 {
		sdx := indexes[2]
		edx := indexes[3]
		sub := tag[sdx:edx]
		return sub, sdx, edx
	}
	return "", -1, -1
}

func ExtractTagFieldIndex(part, fieldName string, action ExtractTagFieldAction) (string, int, int) {
	// 正则表达式查找键值对
	var regex *regexp.Regexp
	if action == EXCLUDE_WHITESPACE_PREFIX {
		regex = regexp.MustCompile(`\b` + regexp.QuoteMeta(fieldName) + `\s*:\s*([^;]+)`)
	} else if action == INCLUDE_WHITESPACE_PREFIX {
		regex = regexp.MustCompile(`\b` + regexp.QuoteMeta(fieldName) + `\s*:([^;]+)`)
	} else {
		panic(errors.New("WRONG"))
	}
	if indexes := regex.FindStringSubmatchIndex(part); len(indexes) > 3 {
		sdx := indexes[2]
		ex := indexes[3]
		sub := part[sdx:ex]
		return sub, sdx, ex
	}
	return "", -1, -1
}

// ExtractNoValueFieldNameIndex 匹配单键标签，比如 index 或者 uniqueIndex 这类标签，在简化情况下可以是没有值的
// 返回的是键名的起止坐标，区间包含键名左右的空格部分
// 目前下游使用这个函数的目的是提取到键名的位置，再往里面补充东西
func ExtractNoValueFieldNameIndex(part, fieldName string) (sdx, edx int) {
	// 保留两个 \b，并在 fieldName 和分号/字符串结尾之间允许空格
	re := regexp.MustCompile(fmt.Sprintf(`(\s*\b%s\b\s*)(?:;|$)`, regexp.QuoteMeta(fieldName)))

	// 查找匹配的位置
	matches := re.FindStringSubmatchIndex(part)
	if len(matches) > 3 {
		return matches[2], matches[3]
	}
	return -1, -1
}

// ExtractFieldEqualsValueIndex 返回的是键值对中值的坐标 比如匹配的是 index:idx_abc 或者 uniqueIndex:udx_xyz 这种键值对
// 返回的是值的起止坐标，同时，区间内包含前后的空格
// 目前下游使用这个函数的目的是，获取到值的坐标，替换值的内容
func ExtractFieldEqualsValueIndex(part, fieldName, fieldValue string) (sdx, edx int) {
	return ExtractFieldEqualsValueIndexV2(part, fieldName, fieldValue, []string{})
}

// ExtractFieldEqualsValueIndexV2 和前面的功能相同，但提供自定义分隔符的功能
func ExtractFieldEqualsValueIndexV2(part string, fieldName string, fieldValue string, terminators []string) (int, int) {
	//首先确保有一个分隔符是分号，这样将来和 $ 或的时候就没有语法错误
	var mergeTerminators = []string{";"}
	for _, v := range terminators {
		//接着把传进来的都进行安全处理
		mergeTerminators = append(mergeTerminators, regexp.QuoteMeta(v))
	}
	// 最后把它们以或连接起来，当然其内容有可能就只是";"这个字符串，表示没有额外的分隔符
	ts := strings.Join(mergeTerminators, "|")

	// 构建正则表达式，匹配 fieldName 和 fieldValue，并确保它们两边的空格被忽略
	re := regexp.MustCompile(fmt.Sprintf(`\s*\b%s\b\s*:(\s*%s\s*)(?:`+ts+`|$)`, regexp.QuoteMeta(fieldName), regexp.QuoteMeta(fieldValue)))

	// 查找匹配的位置
	matches := re.FindStringSubmatchIndex(part)
	if len(matches) > 3 {
		return matches[2], matches[3]
	}
	return -1, -1
}
