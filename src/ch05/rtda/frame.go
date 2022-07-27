package rtda

type Frame struct {
	// 用来实现链表数据结构
	lower *Frame
	// 局部变量表指针
	localVars LocalVars
	// 操作数栈指针
	operandStack *OperandStack
}

/**
 * 构造函数
 * 执行方法所需的局部变量表大小和操作数栈的深度是编译器提前计算好的，
 * 存储在class文件method_info结构的Code属性中
 */
func newFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}
