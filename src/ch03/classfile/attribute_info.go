package classfile

/**
attribute_info {
	u2 attribute_name_index;
	u4 attribute_length;
	u1 info[attribute_length];
}
*/

type AttributeInfo interface {
	readInfo(reader *ClassReader)
}

func readAttributes(reader *ClassReader, cp ConstantPool) []AttributeInfo {
	attrbutesCount := reader.readUint16()
	attributes := make([]AttributeInfo, attrbutesCount)
	for i := range attributes {
		attributes[i] = readAttribute(reader, cp)
	}

	return attributes
}

/**
 * 创建具体的属性实例
 */
func readAttribute(reader *ClassReader, cp ConstantPool) AttributeInfo {
	// 先读取属性名索引
	attrNameIndex := reader.readUint16()
	// 再从常量池中获取属性名
	attrName := cp.getUtf8(attrNameIndex)
	// 获取属性长度
	attrLen := reader.readUint32()
	attrInfo := newAttributeInfo(attrName, attrLen, cp)
	attrInfo.readInfo(reader)
	return attrInfo
}

func newAttributeInfo(attrName string, attrLen uint32, cp ConstantPool) AttributeInfo {
	switch attrName {
	case "Code":
		return &CodeAttribute{cp: cp}
	case "ConstantValue":
		return &ConstantValueAttribute{}
	case "Deprecated":
		return &DeprecatedAttribute{}
	case "Exceptions":
		return &ExceptionsAttribute{}
	case "LineNumberTable":
		return &LineNumberTableAttribute{}
	case "LocalVariableTable":
		return &LocalVariableTableAttribute{}
	case "SourceFile":
		return &SourceFileAttribute{cp: cp}
	case "Synthetic":
		return &SyntheticAttribute{}
	default:
		return &UnparsedAttribute{attrName, attrLen, nil}
	}
}
