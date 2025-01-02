package gg

import (
	"fmt"
	"regexp"
)

const (
	MOBILE      = "(?:0|86|\\+86)?1[3-9]\\d{9}"
	MOBILE_HK   = "(?:0|852|\\+852)?\\d{8}"
	MOBILE_TW   = "(?:0|886|\\+886)?(?:|-)09\\d{8}"
	MOBILE_MO   = "(?:0|853|\\+853)?(?:|-)6\\d{7}"
	TEL         = "(010|02\\d|0[3-9]\\d{2})-?(\\d{6,8})"
	TEL_400_800 = "0\\d{2,3}[\\- ]?[1-9]\\d{6,7}|[48]00[\\- ]?[1-9]\\d{2}[\\- ]?\\d{4}"
)

func isMatch(str, reg string) bool {
	// 编译正则表达式
	re, err := regexp.Compile(reg)
	if err != nil {
		panic(fmt.Sprintf("Failed to compile regex: %s", err))
	}
	// 执行匹配
	return re.MatchString(str)
}

// 手机号
// 移动电话 eg: 中国大陆： +86 180 4953 1399，2位区域码标示+11位数字 中国大陆 +86 Mainland China
func Phone_isMobile(str string) bool {
	return isMatch(str, MOBILE)
}

// 香港手机号 eg: 香港： +852 9876 5432，3位区域码标示+10位数字 香港 Hong Kong
// 中国香港移动电话 eg: 中国香港： +852 5100 4810， 三位区域码+10位数字, 中国香港手机号码8位数
func Phone_isMobileHk(str string) bool {
	return isMatch(str, MOBILE_HK)
}

// 中国台湾移动电话 eg: 中国台湾： +886 09 60 000000， 三位区域码+号码以数字09开头 + 8位数字, 中国台湾手机号码10位数 中国台湾 +886 Taiwan 国际域名缩写：TW
func Phone_isMobileTw(str string) bool {
	return isMatch(str, MOBILE_TW)
}

// 中国澳门移动电话 eg: 中国澳门： +853 68 00000， 三位区域码 +号码以数字6开头 + 7位数字, 中国澳门手机号码8位数 中国澳门 +853 Macao 国际域名缩写：MO
func Phone_isMobileMo(str string) bool {
	return isMatch(str, MOBILE_MO)
}

// 验证是否为座机号码（中国大陆）
func Phone_isTel(str string) bool {
	return isMatch(str, TEL)
}

// 验证是否为400/800电话 @see <a href="https://baike.baidu.com/item/800">800</a>
func Phone_isTel400800(str string) bool {
	return isMatch(str, TEL_400_800)
}

// 验证是否为电话号码
// 是否为座机号码+手机号码（中国大陆）+手机号码（中国香港）+手机号码（中国台湾）+手机号码（中国澳门）
func Phone_isPhone(value string) bool {
	return Phone_isMobile(value) || Phone_isTel400800(value) || Phone_isMobileHk(value) || Phone_isMobileTw(value) || Phone_isMobileMo(value)
}
