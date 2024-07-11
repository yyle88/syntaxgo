package syntaxgo_tag

import "regexp"

// ExtractTagValue 提取标签中的特定部分，比如 `gorm:"" json:""` 的 gorm 整体 或者 json 整体
func ExtractTagValue(tag, key string) string {
	// 正则表达式查找键值对
	// \b 确保匹配整个单词，避免部分匹配
	regex := regexp.MustCompile(`\b` + key + `\s*:\s*"([^"]*)"`)
	matches := regex.FindStringSubmatch(tag)
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}

// ExtractTagField 提取标签中的特定部分，比如 gorm 里面的 column:name 这部分，Fields Tags
func ExtractTagField(tag, key string) string {
	// 正则表达式查找键值对
	// \b 确保匹配整个单词，避免部分匹配
	// 没有引号的键值对情况
	regex := regexp.MustCompile(`\b` + key + `\s*:\s*([^;]+)`)
	matches := regex.FindStringSubmatch(tag)
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}
