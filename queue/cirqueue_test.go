package queue_test

import (
	"testing"

	"github.com/charlienet/utils/queue"
)

func TestCircleQueue(t *testing.T) {
	q := queue.NewCircleQueue(10)

	t.Log("Size:", q.Size())
	for i := 0; i < 10; i++ {
		t.Log(q.Push(i))
	}

	t.Log("Size:", q.Size())
	t.Log("IsFull:", q.IsFull())
	t.Log("Show:", q.Show())

	q.Pop()
	q.Push(11)
	q.Pop()

	t.Log("Show:", q.Show())
}
