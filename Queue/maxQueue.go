type MaxQueue struct {
	Enqueue *Queue
	Dequeue *Queue
}

type Queue struct {
	Set    []int
	Length int
}

func (q *Queue) Push(num int) {
	q.Set = append(q.Set, num)
	q.Length++
}

func (q *Queue) Pop() int {
	if q.Length == 0 {
		return -1
	}
	result := q.Set[0]
	q.Set = q.Set[1:]
	q.Length--
	return result
}

func (q *Queue) Get() int {
	if q.Length == 0 {
		return -1
	}
	return q.Set[0]
}

func Constructor() MaxQueue {
	return MaxQueue{
		Enqueue: &Queue{},
		Dequeue: &Queue{},
	}
}

func (this *MaxQueue) Max_value() int {
	return this.Dequeue.Get()
}

// 通过辅助队列，实现max_value、push_back、pop_front的时间复杂度都是O(1)
// 1、不能通过每次push记录max，因为队列pop最大值后，这个记录就失效了
// 2、通过辅助队列先入先出的特性，保证了，当一个max进入时，前面的最小值就失效了
// 3、例如[5, 4, 2, 3]，5入队，同时5入辅助队列；4、2入队入辅助队；3入队，因为3比2大，所以删除2，填入3
// 4、在5、4出队后，3出队之前，3都是队列的最大值，所以辅助队列无须保存2
func (this *MaxQueue) Push_back(value int) {
	this.Enqueue.Push(value)
	if this.Dequeue.Length == 0 {
		this.Dequeue.Push(value)
		return
	}
	if this.Dequeue.Set[this.Dequeue.Length-1] < value {
		for i := 0; i < this.Dequeue.Length; i++ {
			if this.Dequeue.Set[i] >= value {
				continue
			}
			this.Dequeue.Set[i] = value
			this.Dequeue.Length = i + 1
			this.Dequeue.Set = this.Dequeue.Set[0:this.Dequeue.Length]
			return
		}
	}
	this.Dequeue.Push(value)
}

func (this *MaxQueue) Pop_front() int {
	if this.Enqueue.Length == 0 {
		return -1
	}
	result := this.Enqueue.Pop()
	max := this.Max_value()
	if result == max {
		this.Dequeue.Pop()
	}
	return result
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */