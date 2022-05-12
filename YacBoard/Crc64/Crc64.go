package Crc64

import "github.com/MaxKlaxxMiner/yacgui-web/YacBoard"

const (
	CrcStart Value = 0xcbf29ce484222325
	CrcMul   Value = 0x100000001b3
)

type Value uint64

func (crc64 Value) UpdateInt(value int) Value {
	return (crc64 ^ Value(uint(value))) * CrcMul
}

func (crc64 Value) UpdateBool(value bool) Value {
	if value {
		return (crc64 ^ 1) * CrcMul
	}
	return crc64 * CrcMul
}

func (crc64 Value) UpdateString(value string) Value {
	result := crc64
	for i := 0; i < len(value); i++ {
		result = (result ^ Value(value[i])) * CrcMul
	}
	return result
}

func (crc64 Value) UpdatePieces(value []YacBoard.Piece) Value {
	result := crc64
	for i := 0; i < len(value); i++ {
		result = (result ^ Value(value[i])) * CrcMul
	}
	return result
}
