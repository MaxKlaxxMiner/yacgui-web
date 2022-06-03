package pos

type Pos int
type PosF byte

func FromXY(x, y int) Pos {
	if uint(x) >= Width || uint(y) >= Height {
		return -1
	}
	return Pos(x + y*Width)
}

func FromChars(chars string) Pos {
	if len(chars) < 2 {
		return -1
	}
	x := int(chars[0]) - 'a'
	y := Height + '0' - int(chars[1])
	return FromXY(x, y)
}

func (pos Pos) String() string {
	if uint(pos) >= FieldCount {
		return "-"
	}
	file := pos%Width + 'a'
	rank := Height - pos/Height + '0'
	return string([]byte{byte(file), byte(rank)})
}

func PToF(pos int) PosF {
	if pos < 0 {
		return 0
	}
	return PosF(pos%Width + 1 + (pos/Width+1)*WidthF)
}

func PToFb(pos byte) PosF {
	return PToF(int(pos))
}

func PToFp(pos Pos) PosF {
	return PToF(int(pos))
}

func FToP(pos PosF) int {
	if pos == 0 {
		return -1
	}
	return int(pos)%WidthF - 1 + (int(pos)/WidthF-1)*Width
}
