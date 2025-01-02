package gg

import "unicode"

// 判断字符是否为 ASCII 字符，即 0 到 127 之间的字符
func Char_isAscii(ch rune) bool {
	return ch >= 0 && ch <= 127
}

// 判断字符是否为可见 ASCII 字符，即 32 到 126 之间的字符
func Char_isAsciiPrintable(ch rune) bool {
	return ch >= 32 && ch <= 126
}

// 判断字符是否为字母
func Char_isLetter(ch rune) bool {
	return unicode.IsLetter(ch)
}

// 判断字符是否为大写字母
func Char_isLetterUpper(ch rune) bool {
	return unicode.IsUpper(ch)
}

// 判断字符是否为小写字母
func Char_isLetterLower(ch rune) bool {
	return unicode.IsLower(ch)
}

// 判断字符是否为数字
func Char_isNumber(ch rune) bool {
	return unicode.IsDigit(ch)
}

// 判断字符是否为阿拉伯数字
func Char_isArabicDigit(ch rune) bool {
	return ch >= '0' && ch <= '9'
}

// 判断字符是否为字母或数字
func Char_isLetterOrNumber(ch rune) bool {
	return unicode.IsLetter(ch) || unicode.IsDigit(ch)
}

// IsBlankChar 判断字符是否为空白符
func Char_isBlankChar(ch rune) bool {
	// 检查是否为空格、制表符、全角空格或不间断空格
	return ch == ' ' || ch == '\t' || ch == '　' || ch == '\u00A0'
}

// 判断字符是否为空白符（支持 int 类型）
func Char_isBlankCharInt(ch int) bool {
	return Char_isBlankChar(rune(ch))
}

// 判断字符是否为 Emoji 表情符
func Char_isEmoji(ch rune) bool {
	// Emoji 的 Unicode 范围
	return (ch >= 0x1F600 && ch <= 0x1F64F) || // 表情符号
		(ch >= 0x1F300 && ch <= 0x1F5FF) || // 杂项符号和象形文字
		(ch >= 0x1F680 && ch <= 0x1F6FF) || // 交通和地图符号
		(ch >= 0x2600 && ch <= 0x26FF) || // 杂项符号
		(ch >= 0x2700 && ch <= 0x27BF) || // 装饰符号
		(ch >= 0xFE00 && ch <= 0xFE0F) || // 变体选择器
		(ch >= 0x1F900 && ch <= 0x1F9FF) || // 补充符号和象形文字
		(ch >= 0x1F1E6 && ch <= 0x1F1FF) // 地区指示符号
}

// 判断字符是否为文件分隔符
func Char_isFileSeparator(ch rune) bool {
	// Windows 文件分隔符为 '\'，Linux（Unix）文件分隔符为 '/'
	return ch == '\\' || ch == '/'
}
