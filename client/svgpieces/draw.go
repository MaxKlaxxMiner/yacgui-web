package svgpieces

import (
	"client/canvas"
	"github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard/piece"
	"syscall/js"
)

func newPath(pathStr string) js.Value {
	return js.Global().Get("Path2D").New(pathStr)
}

func Draw(c *canvas.CanvasContext, p piece.Piece) {
	c.SetLineCap("round")
	c.SetLineWidth(1.5)
	c.SetStrokeStyle("#000")
	if p&piece.Colors == piece.White {
		c.SetFillStyle("#fff")
	} else {
		c.SetFillStyle("#000")
	}
	switch p {
	case piece.WhitePawn:
		c.FillPath(pawn)
		c.StrokePath(pawn)
	case piece.BlackPawn:
		c.FillPath(pawn)
		c.StrokePath(pawn)
	case piece.WhiteBishop:
		c.FillPath(bishop)
		c.StrokePath(bishop)
		c.StrokePath(bishop2)
	case piece.BlackBishop:
		c.FillPath(bishop)
		c.StrokePath(bishop)
		c.SetStrokeStyle("#ddd")
		c.StrokePath(bishop2)
	case piece.WhiteKnight:
		c.FillPath(knight)
		c.StrokePath(knight)
		c.SetFillStyle("#000")
		c.FillPath(knight2)
		c.StrokePath(knight2)
	case piece.BlackKnight:
		c.FillPath(knight)
		c.StrokePath(knight)
		c.SetFillStyle("#ddd")
		c.SetStrokeStyle("#ddd")
		c.FillPath(knight2)
		c.StrokePath(knight2)
		c.FillPath(knight3)
	case piece.WhiteRook:
		c.FillPath(rook)
		c.StrokePath(rook)
	case piece.BlackRook:
		c.FillPath(rook)
		c.StrokePath(rook)
		c.SetLineWidth(1)
		c.SetStrokeStyle("#ddd")
		c.StrokePath(rook2)
	}
}
