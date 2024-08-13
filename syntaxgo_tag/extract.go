package syntaxgo_tag

import (
	"regexp"

	"github.com/pkg/errors"
)

// ExtractTagValue 提取标签中的特定部分，比如 `gorm:"" json:""` 的 gorm 整体 或者 json 整体
func ExtractTagValue(tag, key string) string {
	// 正则表达式查找键值对
	// \b 确保匹配整个单词，避免部分匹配
	// \s 空白
	// 该正则的小括号没有把空白括起来，因此结果不包含空白，即不包含空格部分
	regex := regexp.MustCompile(`\b` + key + `\s*:\s*"([^"]*)"`)
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
		regex = regexp.MustCompile(`\b` + fieldName + `\s*:\s*([^;]+)`)
	} else if action == INCLUDE_WHITESPACE_PREFIX {
		regex = regexp.MustCompile(`\b` + fieldName + `\s*:([^;]+)`)
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
	regex := regexp.MustCompile(`\b` + key + `\s*:\s*"([^"]*)"`)
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
		regex = regexp.MustCompile(`\b` + fieldName + `\s*:\s*([^;]+)`)
	} else if action == INCLUDE_WHITESPACE_PREFIX {
		regex = regexp.MustCompile(`\b` + fieldName + `\s*:([^;]+)`)
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
