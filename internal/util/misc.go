package util

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

func UAformat(ua any) string {
	switch reflect.TypeOf(ua) {
	case reflect.TypeOf(uint32(0)), reflect.TypeOf(uint(0)), reflect.TypeOf(uint8(0)), reflect.TypeOf(uint16(0)), reflect.TypeOf(uint64(0)),
		reflect.TypeOf(int(0)), reflect.TypeOf(int8(0)), reflect.TypeOf(int16(0)), reflect.TypeOf(int32(0)), reflect.TypeOf(int64(0)):
		return fmt.Sprintf("%09d", ua)
	case reflect.TypeOf(""):

		uaStr := ua.(string)
		isNumeric := regexp.MustCompile(`^\d+$`).MatchString(uaStr)
		if !isNumeric {
			return ""
		}

		// 检查长度是否为 9 字节
		length := len(uaStr)
		if length > 9 {
			return ""
		} else if length < 9 {
			// 补零
			uaStr = strings.Repeat("0", 9-length) + uaStr
		}
		return uaStr
	default:
		return ""

	}
}

//func UAformat(ua string) string {
//	return fmt.Sprintf("%09d", ua)
//}
