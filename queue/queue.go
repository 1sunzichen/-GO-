package queue

type Queue []interface{}

func (q *Queue) Push(v interface{}) {
	// push 放入 后面
	*q = append(*q, v.(int))
}

func (q *Queue) Pop() int {
	//head 弹出的第一个字符
	head := (*q)[0]
	*q = (*q)[1:] //剩下的 q 数组
	return head.(int)
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
