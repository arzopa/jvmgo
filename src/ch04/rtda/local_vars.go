package rtda

import "math"

type LocalVars []Slot

func newLocalVars(maxLocals uint) LocalVars {
	if maxLocals > 0 {
		return make([]Slot, maxLocals)
	}

	return nil
}

/**
 * 没有对bool、byte、short、char类型定义存取方法
 * 这些类型的值都可以转成int值来处理
 */
func (self LocalVars) SetInt(index uint, val int32) {
	self[index].num = val
}

func (self LocalVars) GetInt(index uint) int32 {
	return self[index].num
}

func (self LocalVars) SetFloat(index uint, val float32) {
	bits := math.Float32bits(val)
	self[index].num = int32(bits)
}

func (self LocalVars) GetFloat(index uint) float32 {
	bits := uint32(self[index].num)
	return math.Float32frombits(bits)
}

func (self LocalVars) SetLong(index uint, val int64) {
	self[index].num = int32(val)
	self[index+1].num = int32(val >> 32)
}

func (self LocalVars) GetLong(index uint) int64 {
	low := uint32(self[index].num)
	hight := uint32(self[index].num)

	return int64(hight)<<32 | int64(low)
}

func (self LocalVars) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	self.SetLong(index, int64(bits))
}

func (self LocalVars) GetDouble(index uint) float64 {
	bits := uint64(self[index].num)
	return math.Float64frombits(bits)
}
