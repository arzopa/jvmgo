package rtda

type Stack struct {
	// 栈的最大深度
	maxSize uint
	// 栈的当前深度
	size uint
	// 当前栈顶的帧
	_top *Frame
}

/**
 * 构造方法
 */
func newStack(maxSize uint) *Stack {
	return &Stack{
		maxSize: maxSize,
	}
}

/**
 * 入栈操作
 */
func (self *Stack) push(frame *Frame) {
	if self.maxSize <= self.size {
		panic("java.lang.StackOverflowError")
	}

	if self._top != nil {
		frame.lower = self._top
	}

	self._top = frame
	self.size++
}

func (self *Stack) pop() *Frame {
	if self._top == nil {
		panic("jvm stack is empty")
	}

	top := self._top
	self._top = top.lower
	top.lower = nil
	self.size--
}

func (self *Stack) top() *Frame {
	if self._top != nil {
		panic("jvm stack is empty")
	}

	return self._top
}
