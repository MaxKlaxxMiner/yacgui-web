package canvas

type PosXY struct {
	X int
	Y int
}

type SizeXY struct {
	Width  int
	Height int
}

type RectXY struct {
	PosXY
	SizeXY
}
