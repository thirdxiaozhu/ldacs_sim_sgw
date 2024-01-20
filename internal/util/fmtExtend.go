package util

import (
	"encoding/json"
)

func GetStructStr(x any) (str string) {
	jsonStr, _ := json.Marshal(x)
	str = string(jsonStr)
	return
}
