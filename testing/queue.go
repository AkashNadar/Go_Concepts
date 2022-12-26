package queue

type Queue struct {
	list     []int
	capacity int
}

func CreateQueue(capaVal int) Queue {
	return Queue{
		list:     make([]int, 0, capaVal),
		capacity: capaVal,
	}
}

func (queue *Queue) AddQueue(val int) bool {
	if len(queue.list) < queue.capacity {
		queue.list = append(queue.list, val)
		return true
	} else {
		return false
	}
}

func (queue *Queue) Next() (int, bool) {
	if len(queue.list) > 0 {
		next := queue.list[0]
		queue.list = queue.list[1:]
		return next, true
	}
	return 0, false
}
