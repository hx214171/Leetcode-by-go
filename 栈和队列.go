// 											简单题
// # 剑指 Offer 09. 用两个栈实现队列
// # 用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，
// # 分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead 操作返回 -1 )
// # 思路：利用两个栈，这里调用container/list的包，list.New()得到一个双向链表。
// 入队，stack1直接push元素。出队则判断情况，stack2为空，将stack1的元素都出栈并加入stack1。
// stack2不为空，则stack2出栈并返回元素值。如果stack1和stack2都为空，则返回-1。
type CQueue struct {
	stack1, stack2 *list.List
}

func Constructor() CQueue {
	return CQueue{list.New(), list.New()}
}

func (this *CQueue) AppendTail(value int) {
	this.stack1.PushBack(value)
}

func (this *CQueue) DeleteHead() int {
	if this.stack2.Len() == 0 {
		for this.stack1.Len() != 0 {
			this.stack2.PushBack(this.stack1.Remove(this.stack1.Back()))
		}
	}
	if this.stack2.Len() != 0 {
		e := this.stack2.Back()
		this.stack2.Remove(e)
		return e.Value.(int)
		//e.Value是 interface{} 接口类型，.(int)将其断言为int
	}
	return -1
}

// # 剑指 Offer 30. 包含min函数的栈
// # 定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的 min 函数在该栈中，
// # 调用 min、push 及 pop 的时间复杂度都是 O(1)。
// 思路:使用两个栈实现，栈A stack用来正常实现push、pop、top方法，栈B minStack用来存栈A的非严格降序元素，
// 栈B的栈顶元素始终对应栈A的最小元素。push方法,栈A直接在末尾添加元素，栈B为空或末尾元素大于要添加的元素时才添加元素。
// pop方法,栈A直接pop()，栈B考虑末尾元素是否与栈A末尾元素相等，若相等栈B pop(),min直接返回栈B末尾元素，
// top返回栈A末尾元素。
type MinStack struct {
	stack, minStack *list.List
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{list.New(), list.New()}
}

func (this *MinStack) Push(x int) {
	this.stack.PushBack(x)
	if this.minStack.Len() == 0 || x <= this.minStack.Back().Value.(int) {
		this.minStack.PushBack(x)
	}
}

func (this *MinStack) Pop() {
	if this.stack.Back().Value.(int) == this.minStack.Back().Value.(int) {
		this.minStack.Remove(this.minStack.Back())
	}
	this.stack.Remove(this.stack.Back())
}

func (this *MinStack) Top() int {
	return this.stack.Back().Value.(int)
}

func (this *MinStack) Min() int {
	return this.minStack.Back().Value.(int)
}

//                                           中等题
// # 剑指 Offer 31. 栈的压入、弹出序列
// # 输入两个整数序列，第一个序列表示栈的压入顺序，请判断第二个序列是否为该栈的弹出顺序。假设压入栈的所有数字均不相等。
// # 例如，序列 {1,2,3,4,5} 是某栈的压栈序列，序列 {4,5,3,2,1} 是该压栈序列对应的一个弹出序列，
// # 但 {4,3,5,1,2} 就不可能是该压栈序列的弹出序列。
// 思路：创建一个辅助栈来模拟入栈 出栈
func validateStackSequences(pushed []int, popped []int) bool {
	stack := []int{}
	index := 0
	for _, num := range pushed {
		stack = append(stack, num)
		for len(stack) != 0 && popped[index] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			index += 1
		}
	}
	return len(stack) == 0
}

// # 剑指 Offer 59 - I. 滑动窗口的最大值
// # 给定一个数组 nums 和滑动窗口的大小 k，请找出所有滑动窗口里的最大值。
思路1：暴力法 时间复杂度O(nk)
func maxSlidingWindow(nums []int, k int) []int {
    if len(nums) == 0 || k == 0 {return nil}
    res := []int{}
    for i := 0; i < len(nums)-k+1; i++ {
        res = append(res, max(nums[i:i+k]))
    }
    return res
}

func max(arr []int) int{
    maxValue := arr[0]
    for i:=1; i < len(arr); i++ {
        if arr[i] > maxValue {
            maxValue = arr[i]
        }
    }
    return maxValue
}

思路2：类似于最小栈，维护一个单调递减的双端队列，双端队列的首元素即为最大值，
遍历数组的同时，去除队列尾部更小的值， 并在尾部添加数组元素，当遍i>=k时，
判断队列首元素能否删除，i>=k-1 滑动窗口达到k个元素时，将队列首元素加入res。
func maxSlidingWindow(nums []int, k int) []int {
    res := []int{}
    deque := []int{}
    for i:= 0; i < len(nums); i++ {
        for len(deque) > 0 && deque[len(deque)-1] < nums[i] {
            deque = deque[:len(deque)-1]
        }
        deque = append(deque, nums[i])
        if i >= k && deque[0] == nums[i-k]{
            deque = deque[1:]
        }
        if i >= k-1 {
            res = append(res, deque[0])
        }
    }
    return res
}

# 剑指 Offer 59 - II. 队列的最大值
# 请定义一个队列并实现函数 max_value 得到队列里的最大值，要求函数max_value、push_back 
# 和 pop_front 的均摊时间复杂度都是O(1)。若队列为空，pop_front 和 max_value 需要返回 -1
思路：跟前一题类似，维护一个单调递减的双端队列。创建队列和一个单调队列。返回最大值：如果单调队列为空，返回-1，
否则返回单调队列首元素。入队：队列为空，返回-1，否则将元素加入队列末尾，同时去除单调队列末尾
小于该元素的元素，然后加入单调队列末尾。出队：如果队列为空，返回-1。如果单调队列首元素等于队列首元素，
则同时删去单调队列和队列首元素，否则只删除队列首元素，最后返回该值。
type MaxQueue struct {
    q []int
    max []int
}

func Constructor() MaxQueue {
    return MaxQueue{[]int{}, []int{}}
}

func (this *MaxQueue) Max_value() int {
    if len(this.max) == 0 {
        return -1
    } else {
        return this.max[0]
    }
}

func (this *MaxQueue) Push_back(value int)  {
    this.q = append(this.q, value)
    for len(this.max) > 0 && value > this.max[len(this.max)-1] {
        this.max = this.max[:len(this.max)-1]
    }
    this.max = append(this.max, value)
}

func (this *MaxQueue) Pop_front() int {
    if len(this.q) == 0 {
        return -1
    }
    pop := this.q[0]
    if pop == this.max[0] {
        this.max = this.max[1:]
    }
    this.q = this.q[1:]
    return pop
}