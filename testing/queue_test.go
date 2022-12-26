package queue

import "testing"

func TestAddQueue(t *testing.T) {
	queue := CreateQueue(3)

	for i := 0; i < 3; i++ {
		if len(queue.list) != i {
			t.Errorf("Incorrect queue element count: %v, wanted %v", len(queue.list), i)
		}
		if !queue.AddQueue(i) {
			t.Errorf("Not appending into the queue: %v", i)
		}
	}

	if queue.AddQueue(4) {
		t.Errorf("added more than the capacity! capacity: %v", len(queue.list))
	}
}

func TestNext(t *testing.T) {

	q := CreateQueue(3)
	for i := 0; i < 3; i++ {
		q.AddQueue(i)
	}

	for i := 0; i < 3; i++ {
		val, isPresent := q.Next()
		if !isPresent {
			t.Errorf("Expected a value!!")
		}
		if val != i {
			t.Errorf("value does not match val:%v i:%v", val, i)
		}
	}

	_, isPresent := q.Next()
	if isPresent {
		t.Errorf("queue should be empty")
	}
}
