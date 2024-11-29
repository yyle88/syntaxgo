package utils

import (
	"os"
	"path/filepath"
	"unicode"
)

func SetPrefix2Strings(prefix string, a []string) (results []string) {
	results = make([]string, 0, len(a))
	for _, v := range a {
		results = append(results, prefix+v)
	}
	return results
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

func C0IsUppercase(s string) bool {
	runes := []rune(s)
	if len(runes) > 0 {
		return unicode.IsUpper(runes[0])
	}
	return false
}

func SetDoubleQuotes(s string) string {
	return "\"" + s + "\""
}

func IsGoSourceFile(info os.FileInfo) bool {
	return (!info.IsDir()) && (filepath.Ext(info.Name()) == ".go")
}
