package test

import (
	"fmt"
	"ldacs_sim_sgw/internal/global"
	"ldacs_sim_sgw/internal/util"
	"testing"
)

func TestGetRandom(t *testing.T) {
	randomString := util.GenerateRandomString(28, util.NumCharset)
	fmt.Println(randomString)

	randomInt := util.GenerateRandomInt(global.UA_LEN)
	fmt.Println(randomInt)
}
