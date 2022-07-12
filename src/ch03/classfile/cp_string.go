package classfile

//=============================== string =====================================
/**
Constant_String_info中本身并不存放字符串，值存放了常量池索引
这个索引指向了Constant_Utf8_info常量
CONSTANT_String_info {
	u1 tag;
	u2 string_index;
}
*/
type ConstantStringInfo struct {
	cp          ConstantPool
	stringIndex uint16
}

func (self *ConstantStringInfo) readInfo(reader *ClassReader) {
	self.stringIndex = reader.readUint16()
}

// 按索引从常量池中查找字符串
func (self *ConstantStringInfo) String() string {
	return self.cp.getUtf8(self.stringIndex)
}
