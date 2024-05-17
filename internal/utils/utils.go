package utils

import (
	"unicode"

	"github.com/yyle88/zaplog"
)

func SetPrefixToSlice(prefix string, a []string) (ss []string) {
	ss = make([]string, 0, len(a))
	for _, v := range a {
		ss = append(ss, prefix+v)
	}
	return ss
}

func SafeMerge[V any](a ...[]V) (res []V) {
	res = make([]V, 0, SumLength(a...))
	for _, slice := range a {
		res = append(res, slice...)
	}
	return res
}

func SumLength[V any](a ...[]V) (n int) {
	for _, slice := range a {
		n += len(slice)
	}
	return n
}

func C0IsUpperString(s string) bool {
	runes := []rune(s)
	if len(runes) > 0 {
		return unicode.IsUpper(runes[0])
	}
	return false
}

func SetDoubleQuotes(s string) string {
	return "\"" + s + "\""
}

func BooleanOK(v bool) bool {
	if !v {
		zaplog.ZAPS.P1.LOG.Panic("B IS FALSE")
	}
	return v
}

func AnyKeyInMap[K comparable, V any](m map[K]V) (k K) {
	for xk := range m {
		return xk
	}
	return k //返回默认值比如0或者空字符串等
}
