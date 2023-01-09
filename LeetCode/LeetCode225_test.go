package LeetCode

/**
序号：225
标题：用队列实现栈
日期：2022/06/30
类型：
*/

/**
Homebrew story
*/

type MyStack struct {
	//queue1 是主队列，queue2 是辅助队列
	queue1, queue2 []int
}

func Constructor3() (s MyStack) {
	return
}

//一个队列的时候 这么操作
//func (s *MyStack) Push(x int) {
//	n := len(s.queue)
//	s.queue = append(s.queue, x)
//	for ; n > 0; n-- {
//		s.queue = append(s.queue, s.queue[0])
//		s.queue = s.queue[1:]
//	}
//}

func (s *MyStack) Push(x int) {
	s.queue2 = append(s.queue2, x)
	for len(s.queue1) > 0 {
		//每次都是拿最左边的，一个一个拼接过去
		s.queue2 = append(s.queue2, s.queue1[0])
		s.queue1 = s.queue1[1:]
	}
	//两个队列进行交换
	s.queue1, s.queue2 = s.queue2, s.queue1
}

func (s *MyStack) Pop() int {
	v := s.queue1[0]
	s.queue1 = s.queue1[1:]
	return v
}

func (s *MyStack) Top() int {
	return s.queue1[0]
}

func (s *MyStack) Empty() bool {
	return len(s.queue1) == 0
}
