package svgpieces

import (
	"client/jscore/canvas"
	"github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard/piece"
	"syscall/js"
)

func newPath(pathStr string) js.Value {
	return js.Global().Get("Path2D").New(pathStr)
}

func DrawDirect(c *canvas.Context, p piece.Piece) {
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
	case piece.WhiteQueen:
		c.SetMiterLimit(1)
		c.FillPath(queen)
		c.StrokePath(queen)
	case piece.BlackQueen:
		c.FillPath(queen)
		c.StrokePath(queen)
		c.StrokePath(queen2)
		c.SetStrokeStyle("#ddd")
		c.StrokePath(queen3)
	case piece.WhiteKing:
		c.FillPath(king)
		c.StrokePath(king)
		c.FillPath(king2)
		c.StrokePath(king2)
	case piece.BlackKing:
		c.FillPath(king)
		c.StrokePath(king)
		c.FillPath(king2)
		c.StrokePath(king2)
		c.SetStrokeStyle("#ddd")
		c.StrokePath(king3)
	}
}

var piecesCache *canvas.FixImage
var piecesCacheSize int

func Draw(c *canvas.Context, x, y, fieldSize int, p piece.Piece) {
	if piecesCache == nil || fieldSize != piecesCacheSize {
		b := canvas.NewBitmap(180*6, 180*2)
		scale := float64(fieldSize) / 45

		pieces := []piece.Piece{piece.Pawn, piece.Knight, piece.Bishop, piece.Rook, piece.Queen, piece.King}
		for i, pc := range pieces {
			b.ResetTransform()
			b.ScaleF(scale, scale)
			b.Translate(i*45, 0)
			DrawDirect(&b.Context, pc|piece.White)
			b.Translate(0, 45)
			DrawDirect(&b.Context, pc|piece.Black)
		}

		piecesCache = canvas.NewFixImage(b)
		piecesCacheSize = fieldSize
	}

	img := piecesCache.GetImageElement()
	if !img.Get("complete").Bool() || img.Get("naturalHeight").Int() == 0 { // fallback?
		c.Save()
		c.Translate(x, y)
		scale := float64(fieldSize) / 45
		c.ScaleF(scale, scale)
		DrawDirect(c, p)
		c.Restore()
		return
	}

	var cx, cy int
	switch p & piece.BasicMask {
	case piece.Pawn:
	case piece.Knight:
		cx = 1
	case piece.Bishop:
		cx = 2
	case piece.Rook:
		cx = 3
	case piece.Queen:
		cx = 4
	case piece.King:
		cx = 5
	default:
		return
	}
	if p&piece.Colors == piece.Black {
		cy = 1
	}
	c.DrawSprite(piecesCache, cx*fieldSize, cy*fieldSize, fieldSize, fieldSize, x, y)
}
