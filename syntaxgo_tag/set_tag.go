package syntaxgo_tag

import (
	"fmt"

	"github.com/yyle88/erero"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

type InsertLocation string

//goland:noinspection GoSnakeCaseUsage
const (
	INSERT_LOCATION_TOP InsertLocation = "TOP"
	INSERT_LOCATION_END InsertLocation = "END"
)

// SetTagFieldValue updates the value of an existing field in a struct tag, identified by its key and field name.
// If the field doesn't exist, it inserts the field at the top or end of the tag.
// SetTagFieldValue 根据给定的 key 和 field name 更新结构体标签中指定字段的值。
// 假如字段不存在，它会将新的字段插入到标签的顶部或末尾。
func SetTagFieldValue(tag, key, field, value string, insertLocation InsertLocation) string {
	zaplog.LOG.Debug("modify-tag-field-value", zap.String("tag", tag))

	tagKeyValue, stx, etx := ExtractTagValueIndex(tag, key)
	if stx < 0 || etx < 0 {
		zaplog.LOG.Panic("IMPOSSIBLE") // 能进到这个函数里的都是已经找到标签的
	}
	zaplog.LOG.Debug("modify-tag-field-value", zap.String("tag-key-value", tagKeyValue), zap.Int("stx", stx), zap.Int("etx", etx))

	fieldValue, sfx, efx := ExtractTagFieldIndex(tagKeyValue, field, INCLUDE_WHITESPACE_PREFIX)
	if sfx < 0 || efx < 0 { // 表示没找到 rule 自定义的内容
		newField := fmt.Sprintf("%s:%s;", field, value)

		switch insertLocation {
		case INSERT_LOCATION_TOP:
			pos := stx // 插在 gorm: 的后面
			return tag[:pos] + newField + tag[pos:]
		case INSERT_LOCATION_END:
			pos := etx
			if pos > 0 {
				if c := tag[pos-1]; c == '"' || c == ';' || c == ' ' {
					// 当是第一个或者前一个已经带分号时，就不需要加分号
				} else {
					newField = ";" + newField // 否则就需要在前面添加个分号
				}
			}
			return tag[:pos] + newField + tag[pos:]
		default:
			panic(erero.New("WRONG"))
		}
	}
	zaplog.LOG.Debug("modify-tag-field-value", zap.String("field-value", fieldValue), zap.Int("sfx", sfx), zap.Int("efx", efx))

	spx := stx + sfx // 把起点坐标补上前面的起始坐标
	epx := stx + efx // 把终点坐标补上前面的起始坐标

	zaplog.LOG.Debug("modify-tag-field-value", zap.String("field-value", tag[spx:epx]))

	newValue := value
	if tag[epx] != ';' {
		newValue += ";" // 当没有分号的时候就补个分号，没有也行但是有的话更安全些
	}
	return tag[:spx] + newValue + tag[epx:]
}
