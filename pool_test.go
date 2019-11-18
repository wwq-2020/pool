package pool

import (
	"fmt"
	"testing"
	"time"
)

func testTaskHandler(task interface{}) {
	fmt.Println(task.(string))
	time.Sleep(time.Minute)
}

func TestPool(t *testing.T) {
	p := New(10, testTaskHandler)
	p.Submit("helloword")
	p.Submit("helloword")
	p.Submit("helloword")
	p.Submit("helloword")
	time.Sleep(time.Second * 10)
	p.Close()
}
