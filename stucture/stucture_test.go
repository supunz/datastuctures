package stucture

import (
	"testing"
)

//Test - testing
func Test(t *testing.T) {
	queueTest(t)
	linkedlistTest(t)
	arraylistTest(t)
}

func recordTest(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}

func queueTest(t *testing.T) {
	size := 5
	q := NewQueue(size)
	if q == nil {
		t.Errorf("Unable to make queue when size is %d", size)
	}
	d, err := q.Dequeue()
	if d != nil || err == nil {
		t.Errorf("Dequeue from an empty queue is allowed")
	}

	err = q.Enqueue(5)
	if err != nil {
		t.Errorf("Enqueue is not-allowed")
	}

	d, err = q.Dequeue()
	if d == nil || err != nil {
		t.Errorf("Dequeue from an non-empty queue is not-allowed")
	}

	for i := 0; i < (size + 3); i++ {
		err := q.Enqueue(i)
		if (i + 1) > size {
			if !q.IsFull() {
				t.Errorf("queue IsFull() improperly implemented")
			}
			if err == nil {
				t.Errorf("can put elements larger than queue size")
			}
		}
	}

	for i := 0; i < (size + 3); i++ {
		d, err := q.Dequeue()
		if (i + 1) > size {
			if !q.IsEmpty() {
				t.Errorf("queue IsEmpty() improperly implemented")
			}
			if d != nil || err == nil {
				t.Errorf("can remove elements which are not actually in queue")
			}
		}
	}
}

func linkedlistTest(t *testing.T) {
	//todo : implement linkedlist test case
}

func arraylistTest(t *testing.T) {
	size := 5
	arrayList := NewArrayList(size)
	if arrayList == nil {
		t.Error("Unable to create arraylist")
	}

	arrayList.Get(2)
}
