package classfile

/**
Code_attribute {
	u2 attribute_name_index;
	u4 attribute_length;
    // 操作数栈的最大深度
	u2 max_stack;
    // 局部变量表的大小
	u2 max_locals;
	u4 code_length;
	u1 code[code_length];
	u2 exception_table_length;

	{   u2 start_pc;
		u2 end_pc;
		u2 handler_pc;
		u2 catch_type;
	} exception_table[exception_table_length];

	u2 attributes_count;
	attribute_info attributes[attributes_count];
}

*/
type CodeAttribute struct {
	cp             ConstantPool
	maxStack       uint16
	maxLocals      uint16
	code           []byte
	exceptionTable []*ExceptionTableEntry
	attributes     []AttributeInfo
}

type ExceptionTableEntry struct {
	startPc   uint16
	endPc     uint16
	handlerPc uint16
	catchPc   uint16
}

func (self *CodeAttribute) readInfo(reader *ClassReader) {
	self.maxStack = reader.readUint16()
	self.maxLocals = reader.readUint16()
	codeLength := reader.readUint32()
	self.code = reader.readBytes(codeLength)
	self.exceptionTable = readExceptionTable(reader)
	self.attributes = readAttributes(reader, self.cp)
}

func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry {
	exceptionTableLength := reader.readUint16()
	// 创建一个exceptionTableLength大小的数组
	exceptionTable := make([]*ExceptionTableEntry, exceptionTableLength)

	// 初始化数组
	for i := range exceptionTable {
		exceptionTable[i] = &ExceptionTableEntry{
			startPc:   reader.readUint16(),
			endPc:     reader.readUint16(),
			handlerPc: reader.readUint16(),
			catchPc:   reader.readUint16(),
		}
	}

	return exceptionTable
}
