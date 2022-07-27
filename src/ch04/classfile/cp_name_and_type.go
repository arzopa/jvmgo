package classfile

//=============================== nameAndType =====================================
/**
CONSTANT_NameAndType_info {
	u1 tag;
	u2 name_index;
	u2 descriptor_index;
}

1 类型描述符
	1.1 基本类型byte、short、char、int、long、float、double的描述符是单个字母
		分别对应B、S、C、I、J、F、D（long的描述符是J而不是L）
	1.2 引用类型的描述符是L+累的完全限定名+分号
	1.3	数组类型的描述符是[+数组元素类型的描述符

2 字段描述符就是字段类型的描述符
3 方法描述符是（封号风格的参数类型描述符）+ 返回值类型描述符，其中void返回值由单个字母V表示
*/
type ConstantNameAndTypeInfo struct {
	nameIndex       uint16
	descriptorIndex uint16
}

func (self *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
	self.descriptorIndex = reader.readUint16()
}
