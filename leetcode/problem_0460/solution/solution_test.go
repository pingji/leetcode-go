package solution

import (
	"testing"
)

func TestProblem(t *testing.T) {
	obj := Constructor(2)
	p := &obj
	p.Put(1, 1)   // cache=[1,_], cnt(1)=1
	p.Put(2, 2)   // cache=[2,1], cnt(2)=1, cnt(1)=1
	v := p.Get(1) // 返回 1
	//cache=[1,2], cnt(2)=1, cnt(1)=2
	if v != 1 {
		t.Error("get 1 error")
	}
	p.Put(3, 3) // 去除键 2 ，因为 cnt(2)=1 ，使用计数最小
	// cache=[3,1], cnt(3)=1, cnt(1)=2
	v = p.Get(2) // 返回 -1 (未找到)
	if v != -1 {
		t.Error("get 2 error")
	}
	v = p.Get(3) // 返回 3
	// cache=[3,1], cnt(3)=2, cnt(1)=2
	if v != 3 {
		t.Error("get 3 error")
	}
	p.Put(4, 4) // 去除键 1 ，1 和 3 的 cnt 相同，但 1 最久未使用
	// cache=[4,3], cnt(4)=1, cnt(3)=2
	v = p.Get(1) // 返回 -1 (未找到)
	if v != -1 {
		t.Error("get 1 error")
	}
	v = p.Get(3) // 返回 3
	// cache=[3,4], cnt(4)=1, cnt(3)=3
	if v != 3 {
		t.Error("get 3 error")
	}
	v = p.Get(4) // 返回 4
	// cache=[3,4], cnt(4)=2, cnt(3)=3
	if v != 4 {
		t.Error("get 4 error")
	}
}

func TestProblem2(t *testing.T) {
	obj := Constructor(0)
	p := &obj
	p.Put(0, 0)
	v := p.Get(0) // 返回 -1
	if v != -1 {
		t.Error("get 1 error")
	}
}
