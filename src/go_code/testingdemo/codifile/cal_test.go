package codifile

import (
	"fmt"
	"testing"
)

func TestAddUpper(t *testing.T) {
	res := AddUpper(10)
	if res != 55 {
		t.Fatalf("输出的结果是%v 但是期望的值是 55", res)
	}
	fmt.Println("测试成功")
}
