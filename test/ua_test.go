package test

import (
	"fmt"
	"ldacs_sim_sgw/internal/util"
	"testing"
)

func TestUA(t *testing.T) {
	fmt.Println(util.UAformat(10086))
	fmt.Println(util.UAformat("10086"))
}
