package gg

// 断言对象是否为 true
func Assert_isTure(bl bool) {
	Assert_isTure2(bl, "必须为 true")
}

// 断言对象是否为 true
func Assert_isTure2(bl bool, msg string) {
	if !bl {
		panic(msg)
	}
}

// 断言对象是否为 false
func Assert_isFalse(bl bool) {
	Assert_isFalse2(bl, "必须为 false")
}

// 断言对象是否为 false
func Assert_isFalse2(bl bool, msg string) {
	if bl {
		panic(msg)
	}
}

func Assert_isNil(obj interface{}) {
	Assert_isNil2(obj, "必须为 nil")
}

func Assert_isNil2(obj interface{}, msg string) {
	if obj == nil {
		panic(msg)
	}
}

// 断言对象是否为null
func Assert_notNil(obj interface{}) {
	Assert_notNil2(obj, "不能为 nil")
}

// 断言对象是否为null
func Assert_notNil2(obj interface{}, msg string) {
	if obj != nil {
		panic(msg)
	}
}

func Assert_notEmpty(str string) {
	Assert_notEmpty2(str, "不能为空字符串")
}

func Assert_notEmpty2(str string, msg string) {
	if len(str) == 0 {
		panic(msg)
	}
}
