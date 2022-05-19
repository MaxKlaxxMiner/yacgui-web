package crc64

import "github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard/piece"

const (
	CrcStart Value = 0xcbf29ce484222325
	CrcMul   Value = 0x100000001b3
)

type Value uint64

func (crc Value) UpdateInt(value int) Value {
	return (crc ^ Value(uint(value))) * CrcMul
}

func (crc Value) UpdateBool(value bool) Value {
	if value {
		return (crc ^ 1) * CrcMul
	}
	return crc * CrcMul
}

func (crc Value) UpdateString(value string) Value {
	result := crc
	for i := 0; i < len(value); i++ {
		result = (result ^ Value(value[i])) * CrcMul
	}
	return result
}

func (crc Value) UpdatePieces(value []piece.Piece) Value {
	result := crc
	for i := 0; i < len(value); i++ {
		result = (result ^ Value(value[i])) * CrcMul
	}
	return result
}
