package classfile

import "debug/dwarf"

type MemberInfo struct {
	cp              ConstantPool // 常量池指针
	accessFlag      uint16
	nameIndex       uint16
	descriptorIndex uint16
	attribute       []ArributeInfo
}

// 读取字段表或者方法表
func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	memberCount := reader.readUint16()
	members := make([]*MemberInfo, memberCount)
	for i := range members {
		members[i] = readMember(reader, cp)
	}

	return members
}

// 读取字段或者方法数据
func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		cp:              cp,
		accessFlag:      reader.readUint16(),
		nameIndex:       reader.readUint16(),
		descriptorIndex: reader.readUint16(),
		attribute:       readAttributes(reader, cp),
	}
}

func (self *MemberInfo) AccessFlags() uint16 {
}

// 从常量池查找字段或者方法名
func (self *MemberInfo) Name() uint16 {
	return self.cp.getUtf8(self.nameIndex)
}

// 从常量池查找字段或者方法描述符
func (self *MemberInfo) Descriptor() uint16 {
	return self.cp.getUtf8(self.descriptorIndex)
}
