package reg

import "regexp"

// regexString 正则表达式
// value 匹配的内容
// 返回匹配内容的值
func RegexFunc(regexString string, value string) string {
	var res string
	regex, _ := regexp.Compile(regexString)
	// 在字符串中查找匹配项
	matches := regex.FindAllString(value, -1)
	// 输出匹配项
	for _, res = range matches {
	}
	return res
}
