package solution

import (
	"testing"
)

func TestProblem(t *testing.T) {
	obj := Constructor(2)
	p := &obj
	p.Put(1, 1)   // 缓存是 {1=1}
	p.Put(2, 2)   // 缓存是 {1=1, 2=2}
	v := p.Get(1) // 返回 1
	if v != 1 {
		t.Error("get 1 error")
	}
	p.Put(3, 3)  // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
	v = p.Get(2) // 返回 -1 (未找到)
	if v != -1 {
		t.Error("get 1 error")
	}
	p.Put(4, 4)  // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
	v = p.Get(1) // 返回 -1 (未找到)
	if v != -1 {
		t.Error("get 1 error")
	}
	v = p.Get(3) // 返回 3
	if v != 3 {
		t.Error("get 3 error")
	}
	v = p.Get(4) // 返回 4
	if v != 4 {
		t.Error("get 4 error")
	}

}
