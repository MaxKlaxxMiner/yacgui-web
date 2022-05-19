package pos

import (
	"github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard/boardsize"
)

type Pos int

func FromXY(x, y int) Pos {
	if uint(x) >= boardsize.Width || uint(y) >= boardsize.Height {
		return -1
	}
	return Pos(x + y*boardsize.Width)
}

func FromChars(chars string) Pos {
	if len(chars) < 2 {
		return -1
	}
	x := int(chars[0]) - 'a'
	y := boardsize.Height + '0' - int(chars[1])
	return FromXY(x, y)
}

func (pos Pos) String() string {
	if uint(pos) >= boardsize.FieldCount {
		return "-"
	}
	file := pos%boardsize.Width + 'a'
	rank := boardsize.Height - pos/boardsize.Height + '0'
	return string([]byte{byte(file), byte(rank)})
}
