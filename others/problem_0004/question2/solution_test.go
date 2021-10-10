package question2

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"
)

type Workload struct {
	id int32
}

func NewWorkload(id int32) *Workload {
	// fmt.Printf("new a workload, id=%d\n", id)
	return &Workload{id: id}
}

func (w *Workload) Work() {
	time.Sleep(time.Second)
	fmt.Printf("%s, execut Work, id=%d\n", time.Now().Format("2006-01-02_15:04:05"), w.id)
}

type Producer struct {
	capacity int32
	count    int32
}

func NewProducer(capacity int32) *Producer {
	return &Producer{
		capacity: capacity,
		count:    0,
	}
}

func (p *Producer) Produce() IWorkload {
	if p.count < p.capacity {
		atomic.AddInt32(&p.count, 1)
		return NewWorkload(p.count)
	} else {
		return nil
	}
}

func TestQuestion2(t *testing.T) {
	producer := NewProducer(100)
	Question2(producer)
}

// func TestQuestion2Unlimited(t *testing.T) {
// 	producer := NewProducer(100)
// 	Question2Unlimited(producer)
// }
