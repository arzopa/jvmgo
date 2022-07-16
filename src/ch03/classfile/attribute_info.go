package classfile

/**
attribute_info {
	u2 attribute_name_index;
	u4 attribute_length;
	u1 info[attribute_length];
}
*/

type AttributeInfo interface {
	readinfo(reader *ClassReader)
}

func readAttributes(reader *ClassReader, cp ConstantPool) []AttributeInfo {
	attrbutesCount := reader.readUint16()
	attributes := make([]AttributeInfo, attrbutesCount)
	for i := range attributes {
		attributes[i] = readAttribute(reader, cp)
	}

	return attributes
}

func readAttribute(reader *ClassReader, cp ConstantPool) AttributeInfo {
	attrnameIndex := reader.readUint16()
	attrName := cp.getUtf8(attrnameIndex)
	attrLen := reader.readUint32()
	attrInfo := newAttributeInfo(attrName, attrLen, cp)
	attrInfo.readInfo(reader)
}

func newAttributeInfo(attrName string, attrLen uint32, cp ConstantPool) AttributeInfo {

}
