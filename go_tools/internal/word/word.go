package word

import (
	"strings"
	"unicode"
)

func ToUpper(str string) string {
	return strings.ToUpper(str)
}

func ToLower(str string) string {
	return strings.ToLower(str)
}

func UnderscoreToUpperCamelCase(str string) string {
	str = strings.Replace(str, "_", " ", -1)
	// 所有单词首字符转换为大写
	str = strings.Title(str)
	return strings.Replace(str, " ", "", -1)
}

func UnderscoreToLowerCamelCase(str string) string {
	str = UnderscoreToUpperCamelCase(str)
	return string(unicode.ToLower(rune(str[0]))) + str[1:]
}

func CamelCaseToUnderscore(str string) string {
	var output []rune
	for i, r := range str {
		if i == 0 {
			output = append(output, unicode.ToLower(r))
			continue
		}
		if unicode.IsUpper(r) {
			output = append(output, '_')
		}
		output = append(output, unicode.ToLower(r))
	}
	return string(output)
}
