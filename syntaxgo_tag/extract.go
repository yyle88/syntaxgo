package syntaxgo_tag

import "regexp"

// ExtractTagValue 提取标签中的特定部分，比如 `gorm:"" json:""` 的 gorm 整体 或者 json 整体
func ExtractTagValue(tag, key string) string {
	// 正则表达式查找键值对
	// \b 确保匹配整个单词，避免部分匹配
	regex := regexp.MustCompile(`\b` + key + `\s*:\s*"([^"]*)"`)
	if matches := regex.FindStringSubmatch(tag); len(matches) > 1 {
		return matches[1]
	}
	return ""
}

// ExtractTagField 提取标签中的特定部分，比如 gorm 里面的 column:name 这部分，Fields Tags
func ExtractTagField(part, fieldName string) string {
	// 正则表达式查找键值对
	// \b 确保匹配整个单词，避免部分匹配
	// 没有引号的键值对情况
	regex := regexp.MustCompile(`\b` + fieldName + `\s*:\s*([^;]+)`)
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

func ExtractTagFieldIndex(part, fieldName string) (string, int, int) {
	// 正则表达式查找键值对
	// \b 确保匹配整个单词，避免部分匹配
	// 没有引号的键值对情况
	regex := regexp.MustCompile(`\b` + fieldName + `\s*:\s*([^;]+)`)
	if indexes := regex.FindStringSubmatchIndex(part); len(indexes) > 3 {
		sdx := indexes[2]
		ex := indexes[3]
		sub := part[sdx:ex]
		return sub, sdx, ex
	}
	return "", -1, -1
}
