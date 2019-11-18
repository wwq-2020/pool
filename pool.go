package pool

import (
	"github.com/wwq1988/datastructure/lockfree/queue"
)

// Pool goroutine池
type Pool struct {
	queue *queue.Queue
}

// New 初始化Pool
func New(size uint, taskHandler func(interface{})) *Pool {
	queue := queue.New()
	return NewWithQueue(size, queue, taskHandler)
}

// NewWithQueue 初始化pool,带queue
func NewWithQueue(size uint, queue *queue.Queue, taskHandler func(interface{})) *Pool {
	if size == 0 {
		panic("zero pool size")
	}
	for i := size; i > 0; i-- {
		go queue.Iter(taskHandler)
	}
	pool := &Pool{queue: queue}
	return pool
}

// Submit 提交任务
func (p *Pool) Submit(task interface{}) {
	p.queue.Push(task)
}

// Close 关闭Pool
func (p *Pool) Close() {
	p.queue.Close()
}
