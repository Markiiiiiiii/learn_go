package splitdemo

import (
	"reflect"
	"testing"
)

// 单元测试
func Test1Split(t *testing.T) {
	ret := Split("afbfc", "f")
	want := []string{"aa", "b", "c"}
	if !reflect.DeepEqual(ret, want) {
		// 测试失败了
		t.Errorf("want:%v\ngot:%v\n", want, ret)
	}
}

func Test2Split(t *testing.T) {
	ret := Split("a:v:x", ":")
	want := []string{"a", "v", "x"}
	if !reflect.DeepEqual(ret, want) {
		// 测试失败了
		t.Errorf("want:%v\ngot:%v\n", want, ret)
	}
}
