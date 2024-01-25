package util

import (
	"encoding/base64"
	"ldacs_sim_sgw/internal/global"
)

func Base64Encode(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}

func Base64Decode(src string) []byte {
	dec, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		global.LOGGER.Error("base64 decode error")
		return nil
	}
	return dec
}
