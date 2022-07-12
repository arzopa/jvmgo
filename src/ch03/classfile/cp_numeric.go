package classfile

import "math"

// CONSTANT_Integer_info使用4字节存储整数常量
// CONSTANT_Integer_info和后面将要介绍的其他三种数字常量无论是结构，还是实现，都非常相似，
// 所以把它们定义在同一个文件中
/*
CONSTANT_Integer_info {
	u1 tag;
	u4 bytes;
}
*/
type ConstantIntegerInfo struct {
	val int32
}

// 先读取一个uint32数据，然后把他转型成int32类型
func (self *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	self.val = int32(bytes)
}

//=============================== float =====================================
/**
CONSTANT_Float_info {
	u1 tag;
    u4 bytes;
}
*/
type ConstantFloatInfo struct {
	val float32
}

func (self *ConstantFloatInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	self.val = math.Float32frombits(bytes)
}

//=============================== long =====================================
/**
CONSTANT_Long_info {
	u1 tag;
    u4 high_bytes;
    u4 low_bytes;
}
*/
type ConstantLongInfo struct {
	val int64
}

func (self *ConstantLongInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	self.val = int64(bytes)
}

//=============================== double =====================================
/**
CONSTANT_Double_info {
	u1 tag;
	u4 high_bytes;
	u4 low_bytes;
}
*/
type ConstantDoubleInfo struct {
	val float64
}

func (self *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	self.val = math.Float64frombits(bytes)
}
