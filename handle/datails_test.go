package handle

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("write setup code here...") // 测试之前的做一些设置
	// 如果 TestMain 使用了 flags，这里应该加上flag.Parse()
	retCode := m.Run()                         // 执行测试
	fmt.Println("write teardown code here...") // 测试之后做一些拆卸工作
	os.Exit(retCode)
}

func TestGetDetails(t *testing.T) {
	var gids = []string{
		"43fnj457p3p9is90h7ilt7ii5re",
	}

	got := GetDetails(gids)
	want := got

	if !reflect.DeepEqual(got, want) {
		t.Error("测试不通过")
	}
}
