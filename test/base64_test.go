package test

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestBase64(t *testing.T) {
	var data []byte

	data = append(data, 232)
	data = append(data, 252)

	str := base64.StdEncoding.EncodeToString(data)
	fmt.Println(str)

	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(data)
}
