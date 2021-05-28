package check

import (
	"reflect"
	"regexp"
)

/**
  @author Bill
*/
const ErrUnSportArray = "ErrorUnSportArray"

//手机号码检测
func CheckIsMobile(mobileNum string) bool {
	var regular = "^1[345789]{1}\\d{9}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}

//判断是否是18或15位身份证
func IsIdCard(cardNo string) bool {
	//18位身份证 ^(\d{17})([0-9]|X)$
	if m, _ := regexp.MatchString(`(^\d{15}$)|(^\d{18}$)|(^\d{17}(\d|X|x)$)`, cardNo); !m {
		return false
	}
	return true
}

// 元素中不允许包含slice 与 map
func InArray(array interface{}, data interface{}) bool {
	sVal := reflect.ValueOf(array)
	kind := sVal.Kind()
	if kind == reflect.Slice || kind == reflect.Array {
		for i := 0; i < sVal.Len(); i++ {
			if sVal.Index(i).Interface() == data {
				return true
			}
		}
		return false
	}
	return false
}
