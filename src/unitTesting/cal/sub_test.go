package cal

import (
	"testing"
)

func TestSub(t *testing.T) {
	res := sub(10, 5)
	if res != 5 {
		t.Fatalf("计算结果错误，预计值5，实际值%d", res)
	}
	t.Log("计算结果正确....")
}
