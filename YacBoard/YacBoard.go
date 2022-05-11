package YacBoard

import (
	"strconv"
	"unsafe"
)

const (
	Width  = 8
	Height = 8
)

type YacBoard struct {
	Fields [Width * Height]Piece

	HalfmoveClock int
	MoveNumber    int
	WhiteKingPos  int
	BlackKingPos  int
	EnPassantPos  int

	WhiteMove               bool
	WhiteCanCastleKingside  bool
	WhiteCanCastleQueenside bool
	BlackCanCastleKingside  bool
	BlackCanCastleQueenside bool
}

type byte64 struct {
	u0 uint64
	u1 uint64
	u2 uint64
	u3 uint64
	u4 uint64
	u5 uint64
	u6 uint64
	u7 uint64
}

func (board *YacBoard) Clear() {
	if strconv.IntSize == 64 && len(board.Fields) == 64 {
		ptr := (*byte64)(unsafe.Pointer(&board.Fields))
		ptr.u0 = 0
		ptr.u1 = 0
		ptr.u2 = 0
		ptr.u3 = 0
		ptr.u4 = 0
		ptr.u5 = 0
		ptr.u6 = 0
		ptr.u7 = 0
	} else {
		for i := 0; i < len(board.Fields); i++ {
			board.Fields[i] = 0
		}
	}

	board.HalfmoveClock = 0
	board.MoveNumber = 1
	board.WhiteKingPos = -1
	board.BlackKingPos = -1
	board.EnPassantPos = -1

	board.WhiteMove = true
	board.WhiteCanCastleKingside = false
	board.WhiteCanCastleQueenside = false
	board.BlackCanCastleKingside = false
	board.BlackCanCastleQueenside = false
}

func (board *YacBoard) SetField(pos int, piece Piece) {
	if uint(pos) >= Width*Height {
		panic("argument out of range")
		return
	}
	board.Fields[pos] = piece

	if piece&King == King {
		if piece == WhiteKing {
			board.WhiteKingPos = pos
		} else {
			board.BlackKingPos = pos
		}
	}
}

func (board *YacBoard) GetField(pos int) Piece {
	if uint(pos) >= Width*Height {
		return Blocked
	}
	return board.Fields[pos]
}

func (board *YacBoard) GetFieldXY(x, y int) Piece {
	if uint(x) >= Width || uint(y) >= Height {
		return Blocked
	}
	return board.GetField(x + y*Width)
}

func (board *YacBoard) GetBoardInfo() BoardInfo {
	result := BoardInfo(uint8(int8(board.EnPassantPos))) | BoardInfo(uint(board.HalfmoveClock)<<16)

	if board.WhiteCanCastleKingside {
		result |= WhiteCanCastleKingside
	}
	if board.WhiteCanCastleQueenside {
		result |= WhiteCanCastleQueenside
	}
	if board.BlackCanCastleKingside {
		result |= BlackCanCastleKingside
	}
	if board.BlackCanCastleQueenside {
		result |= BlackCanCastleQueenside
	}

	return result
}

func (board *YacBoard) SetBoardInfo(boardInfo BoardInfo) {
	board.EnPassantPos = int(int8(uint8(boardInfo & EnPassantMask)))
	board.WhiteCanCastleKingside = (boardInfo & WhiteCanCastleKingside) != BoardInfoNone
	board.WhiteCanCastleQueenside = (boardInfo & WhiteCanCastleQueenside) != BoardInfoNone
	board.BlackCanCastleKingside = (boardInfo & BlackCanCastleKingside) != BoardInfoNone
	board.BlackCanCastleQueenside = (boardInfo & BlackCanCastleQueenside) != BoardInfoNone
	board.HalfmoveClock = (int)(boardInfo&HalfmoveCounterMask) >> 16
}

func (board YacBoard) String() string {
	result := make([]byte, 0, Width*Height+(4+1)*Height+6+FenMaxLength)
	for y := 0; y < Height; y++ {
		result = append(result, ' ', ' ', ' ', ' ')
		for x := 0; x < Width; x++ {
			result = append(result, PieceToChar(board.GetFieldXY(x, y)))
		}
		result = append(result, '\n')
	}
	result = append(result, "\nFEN: "...)
	result = append(result, board.GetFEN()...)
	return string(result)
}

//public sealed unsafe class Board : IBoard
//{
//#region # // --- Move ---
///// <summary>
///// prüft die theoretischen Bewegungsmöglichkeiten einer Spielfigur auf einem bestimmten Feld
///// </summary>
///// <param name="pos">Position auf dem Spielfeld mit der zu testenden Figur</param>
///// <returns>Aufzählung der theretisch begehbaren Felder</returns>
//IEnumerable<int> ScanMove(int pos)
//{
//var piece = fields[pos];
//if (piece == Piece.None) yield break; // keine Figur auf dem Spielfeld?
//var color = piece & Piece.Colors;
//Debug.Assert(color == (WhiteMove ? Piece.White : Piece.Black)); // passt die Figur-Farbe zum Zug?
//
//int posX = pos % Width;
//int posY = pos / Width;
//switch (piece & Piece.BasicPieces)
//{
//#region # case Piece.King: // König
//case Piece.King:
//{
//if (posX > 0) // nach links
//{
//if (posY > 0 && (fields[pos - (Width + 1)] & color) == Piece.None) yield return pos - (Width + 1); // links-oben
//if ((fields[pos - 1] & color) == Piece.None) yield return pos - 1; // links
//if (posY < Height - 1 && (fields[pos + (Width - 1)] & color) == Piece.None) yield return pos + (Width - 1); // links-unten
//}
//if (posX < Width - 1) // nach rechts
//{
//if (posY > 0 && (fields[pos - (Width - 1)] & color) == Piece.None) yield return pos - (Width - 1); // rechts-oben
//if ((fields[pos + 1] & color) == Piece.None) yield return pos + 1; // rechts
//if (posY < Height - 1 && (fields[pos + (Width + 1)] & color) == Piece.None) yield return pos + (Width + 1); // rechts-unten
//}
//if (posY > 0 && (fields[pos - Width] & color) == Piece.None) yield return pos - Width; // oben
//if (posY < Height - 1 && (fields[pos + Width] & color) == Piece.None) yield return pos + Width; // unten
//} break;
//#endregion
//
//#region # case Piece.Queen: // Dame
//case Piece.Queen:
//{
//// links
//for (int i = 1; i < Width; i++)
//{
//if (posX - i < 0) break;
//int p = pos - i;
//var f = fields[p];
//if ((f & color) != Piece.None) break;
//yield return p;
//if (f != Piece.None) break;
//}
//// rechts
//for (int i = 1; i < Width; i++)
//{
//if (posX + i >= Width) break;
//int p = pos + i;
//var f = fields[p];
//if ((f & color) != Piece.None) break;
//yield return p;
//if (f != Piece.None) break;
//}
//// oben
//for (int i = 1; i < Height; i++)
//{
//if (posY - i < 0) break;
//int p = pos - Width * i;
//var f = fields[p];
//if ((f & color) != Piece.None) break;
//yield return p;
//if (f != Piece.None) break;
//}
//// unten
//for (int i = 1; i < Height; i++)
//{
//if (posY + i >= Height) break;
//int p = pos + Width * i;
//var f = fields[p];
//if ((f & color) != Piece.None) break;
//yield return p;
//if (f != Piece.None) break;
//}
//// links-oben
//for (int i = 1; i < Math.Max(Width, Height); i++)
//{
//if (posX - i < 0 || posY - i < 0) break;
//int p = pos - (Width * i + i);
//var f = fields[p];
//if ((f & color) != Piece.None) break;
//yield return p;
//if (f != Piece.None) break;
//}
//// links-unten
//for (int i = 1; i < Math.Max(Width, Height); i++)
//{
//if (posX - i < 0 || posY + i >= Height) break;
//int p = pos + (Width * i - i);
//var f = fields[p];
//if ((f & color) != Piece.None) break;
//yield return p;
//if (f != Piece.None) break;
//}
//// rechts-oben
//for (int i = 1; i < Math.Max(Width, Height); i++)
//{
//if (posX + i >= Width || posY - i < 0) break;
//int p = pos - (Width * i - i);
//var f = fields[p];
//if ((f & color) != Piece.None) break;
//yield return p;
//if (f != Piece.None) break;
//}
//// rechts-unten
//for (int i = 1; i < Math.Max(Width, Height); i++)
//{
//if (posX + i >= Width || posY + i >= Height) break;
//int p = pos + (Width * i + i);
//var f = fields[p];
//if ((f & color) != Piece.None) break;
//yield return p;
//if (f != Piece.None) break;
//}
//} break;
//#endregion
//
//#region # case Piece.Rook: // Turm
//case Piece.Rook:
//{
//// links
//for (int i = 1; i < Width; i++)
//{
//if (posX - i < 0) break;
//int p = pos - i;
//var f = fields[p];
//if ((f & color) != Piece.None) break;
//yield return p;
//if (f != Piece.None) break;
//}
//// rechts
//for (int i = 1; i < Width; i++)
//{
//if (posX + i >= Width) break;
//int p = pos + i;
//var f = fields[p];
//if ((f & color) != Piece.None) break;
//yield return p;
//if (f != Piece.None) break;
//}
//// oben
//for (int i = 1; i < Height; i++)
//{
//if (posY - i < 0) break;
//int p = pos - Width * i;
//var f = fields[p];
//if ((f & color) != Piece.None) break;
//yield return p;
//if (f != Piece.None) break;
//}
//// unten
//for (int i = 1; i < Height; i++)
//{
//if (posY + i >= Height) break;
//int p = pos + Width * i;
//var f = fields[p];
//if ((f & color) != Piece.None) break;
//yield return p;
//if (f != Piece.None) break;
//}
//} break;
//#endregion
//
//#region # case Piece.Bishop: // Läufer
//case Piece.Bishop:
//{
//// links-oben
//for (int i = 1; i < Math.Max(Width, Height); i++)
//{
//if (posX - i < 0 || posY - i < 0) break;
//int p = pos - (Width * i + i);
//var f = fields[p];
//if ((f & color) != Piece.None) break;
//yield return p;
//if (f != Piece.None) break;
//}
//// links-unten
//for (int i = 1; i < Math.Max(Width, Height); i++)
//{
//if (posX - i < 0 || posY + i >= Height) break;
//int p = pos + (Width * i - i);
//var f = fields[p];
//if ((f & color) != Piece.None) break;
//yield return p;
//if (f != Piece.None) break;
//}
//// rechts-oben
//for (int i = 1; i < Math.Max(Width, Height); i++)
//{
//if (posX + i >= Width || posY - i < 0) break;
//int p = pos - (Width * i - i);
//var f = fields[p];
//if ((f & color) != Piece.None) break;
//yield return p;
//if (f != Piece.None) break;
//}
//// rechts-unten
//for (int i = 1; i < Math.Max(Width, Height); i++)
//{
//if (posX + i >= Width || posY + i >= Height) break;
//int p = pos + (Width * i + i);
//var f = fields[p];
//if ((f & color) != Piece.None) break;
//yield return p;
//if (f != Piece.None) break;
//}
//} break;
//#endregion
//
//#region # case Piece.Knight: // Springer
//case Piece.Knight:
//{
//if (posX > 0) // 1 nach links
//{
//if (posY > 1 && (fields[pos - (Width * 2 + 1)] & color) == Piece.None) yield return pos - (Width * 2 + 1); // -1, -2
//if (posY < Height - 2 && (fields[pos + (Width * 2 - 1)] & color) == Piece.None) yield return pos + (Width * 2 - 1); // -1, +2
//if (posX > 1) // 2 nach links
//{
//if (posY > 0 && (fields[pos - (Width + 2)] & color) == Piece.None) yield return pos - (Width + 2); // -2, -1
//if (posY < Height - 1 && (fields[pos + (Width - 2)] & color) == Piece.None) yield return pos + (Width - 2); // -2, +1
//}
//}
//if (posX < Width - 1) // 1 nach rechts
//{
//if (posY > 1 && (fields[pos - (Width * 2 - 1)] & color) == Piece.None) yield return pos - (Width * 2 - 1); // +1, -2
//if (posY < Height - 2 && (fields[pos + (Width * 2 + 1)] & color) == Piece.None) yield return pos + (Width * 2 + 1); // +1, +2
//if (posX < Width - 2) // 2 nach rechts
//{
//if (posY > 0 && (fields[pos - (Width - 2)] & color) == Piece.None) yield return pos - (Width - 2); // +2, +1
//if (posY < Height - 1 && (fields[pos + (Width + 2)] & color) == Piece.None) yield return pos + (Width + 2); // +2, -1
//}
//}
//} break;
//#endregion
//
//#region # case Piece.Pawn: // Bauer
//case Piece.Pawn:
//{
//if (posY < 1 || posY >= Height - 1) break; // ungültige Position
//
//if (color == Piece.White) // weißer Bauer = nach oben laufen
//{
//if (fields[pos - Width] == Piece.None) // Laufweg frei?
//{
//yield return pos - Width;
//if (posY == 6 && fields[pos - Width * 2] == Piece.None) yield return pos - Width * 2; // Doppelzug
//}
//if (posX > 0 && (EnPassantPos == pos - (Width + 1) || (fields[pos - (Width + 1)] & Piece.Colors) == Piece.Black)) yield return pos - (Width + 1); // nach links-oben schlagen
//if (posX < Width - 1 && (EnPassantPos == pos - (Width - 1) || (fields[pos - (Width - 1)] & Piece.Colors) == Piece.Black)) yield return pos - (Width - 1); // nach rechts-oben schlagen
//}
//else // schwarzer Bauer = nach unten laufen
//{
//if (fields[pos + Width] == Piece.None) // Laufweg frei?
//{
//yield return pos + Width;
//if (posY == 1 && fields[pos + Width * 2] == Piece.None) yield return pos + Width * 2; // Doppelzug
//}
//if (posX > 0 && (EnPassantPos == pos + (Width - 1) || (fields[pos + (Width - 1)] & Piece.Colors) == Piece.White)) yield return pos + (Width - 1); // nach links-unten schlagen
//if (posX < Width - 1 && (EnPassantPos == pos + (Width + 1) || (fields[pos + (Width + 1)] & Piece.Colors) == Piece.White)) yield return pos + (Width + 1); // nach rechts-unten schlagen
//}
//} break;
//#endregion
//}
//}
//
///// <summary>
///// fragt die aktuelle Position des Königs ab
///// </summary>
///// <param name="kingColor">Farbe des Königs, welche abgefragt werden soll</param>
///// <returns>Position auf dem Spielbrett</returns>
//public override int GetKingPos(Piece kingColor)
//{
//return (kingColor & Piece.White) != Piece.None ? whiteKingPos : blackKingPos;
//}
//
///// <summary>
///// prüft, ob ein bestimmtes Spielfeld unter Schach steht
///// </summary>
///// <param name="pos">Position, welche geprüft werden soll</param>
///// <param name="checkerColor">zu prüfende Spielerfarbe, welche das Schach geben könnte (nur <see cref="Piece.White"/> oder <see cref="Piece.Black"/> erlaubt)</param>
///// <returns>true, wenn das Feld angegriffen wird und unter Schach steht</returns>
//public override bool IsChecked(int pos, Piece checkerColor)
//{
//int posX = pos % Width;
//int posY = pos / Width;
//
//// --- Bauern und König prüfen ---
//if (checkerColor == Piece.White)
//{
//if (posX > 0) // nach links
//{
//if (posY > 0 && pos - (Width + 1) == whiteKingPos) return true; // links-oben
//if (pos - 1 == whiteKingPos) return true; // links
//if (posY < Height - 1 && (pos + (Width - 1) == whiteKingPos || fields[pos + (Width - 1)] == Piece.WhitePawn)) return true; // links-unten
//}
//if (posX < Width - 1) // nach rechts
//{
//if (posY > 0 && pos - (Width - 1) == whiteKingPos) return true; // rechts-oben
//if (pos + 1 == whiteKingPos) return true; // rechts
//if (posY < Height - 1 && (pos + (Width + 1) == whiteKingPos || fields[pos + (Width + 1)] == Piece.WhitePawn)) return true; // rechts-unten
//}
//if (posY > 0 && pos - Width == whiteKingPos) return true; // oben
//if (posY < Height - 1 && pos + Width == whiteKingPos) return true; // unten
//}
//else
//{
//if (posX > 0) // nach links
//{
//if (posY > 0 && (pos - (Width + 1) == blackKingPos || fields[pos - (Width + 1)] == Piece.BlackPawn)) return true; // links-oben
//if (pos - 1 == blackKingPos) return true; // links
//if (posY < Height - 1 && pos + (Width - 1) == blackKingPos) return true; // links-unten
//}
//if (posX < Width - 1) // nach rechts
//{
//if (posY > 0 && (pos - (Width - 1) == blackKingPos || fields[pos - (Width - 1)] == Piece.BlackPawn)) return true; // rechts-oben
//if (pos + 1 == blackKingPos) return true; // rechts
//if (posY < Height - 1 && pos + (Width + 1) == blackKingPos) return true; // rechts-unten
//}
//if (posY > 0 && pos - Width == blackKingPos) return true; // oben
//if (posY < Height - 1 && pos + Width == blackKingPos) return true; // unten
//}
//
//// --- Springer prüfen ---
//{
//var knight = checkerColor | Piece.Knight;
//if (posX > 0) // 1 nach links
//{
//if (posY > 1 && fields[pos - (Width * 2 + 1)] == knight) return true; // -1, -2
//if (posY < Height - 2 && fields[pos + (Width * 2 - 1)] == knight) return true; // -1, +2
//if (posX > 1) // 2 nach links
//{
//if (posY > 0 && fields[pos - (Width + 2)] == knight) return true; // -2, -1
//if (posY < Height - 1 && fields[pos + (Width - 2)] == knight) return true; // -2, +1
//}
//}
//if (posX < Width - 1) // 1 nach rechts
//{
//if (posY > 1 && fields[pos - (Width * 2 - 1)] == knight) return true; // +1, -2
//if (posY < Height - 2 && fields[pos + (Width * 2 + 1)] == knight) return true; // +1, +2
//if (posX < Width - 2) // 2 nach rechts
//{
//if (posY > 0 && fields[pos - (Width - 2)] == knight) return true; // +2, +1
//if (posY < Height - 1 && fields[pos + (Width + 2)] == knight) return true; // +2, -1
//}
//}
//}
//
//// --- horizontale und vertikale Wege prüfen ---
//{
//for (int i = 1; i < Width; i++) // links
//{
//if (posX - i < 0) break;
//var f = fields[pos - i];
//if (f == Piece.None) continue;
//if ((f & (Piece.Rook | Piece.Queen)) != Piece.None && (f & checkerColor) != Piece.None) return true;
//break;
//}
//for (int i = 1; i < Width; i++) // rechts
//{
//if (posX + i >= Width) break;
//var f = fields[pos + i];
//if (f == Piece.None) continue;
//if ((f & (Piece.Rook | Piece.Queen)) != Piece.None && (f & checkerColor) != Piece.None) return true;
//break;
//}
//for (int i = 1; i < Height; i++) // oben
//{
//if (posY - i < 0) break;
//var f = fields[pos - Width * i];
//if (f == Piece.None) continue;
//if ((f & (Piece.Rook | Piece.Queen)) != Piece.None && (f & checkerColor) != Piece.None) return true;
//break;
//}
//for (int i = 1; i < Height; i++) // unten
//{
//if (posY + i >= Height) break;
//var f = fields[pos + Width * i];
//if (f == Piece.None) continue;
//if ((f & (Piece.Rook | Piece.Queen)) != Piece.None && (f & checkerColor) != Piece.None) return true;
//break;
//}
//}
//
//// --- diagonale Wege prüfen ---
//{
//for (int i = 1; i < Math.Max(Width, Height); i++) // links-oben
//{
//if (posX - i < 0 || posY - i < 0) break;
//var f = fields[pos - (Width * i + i)];
//if (f == Piece.None) continue;
//if ((f & (Piece.Bishop | Piece.Queen)) != Piece.None && (f & checkerColor) != Piece.None) return true;
//break;
//}
//for (int i = 1; i < Math.Max(Width, Height); i++) // links-unten
//{
//if (posX - i < 0 || posY + i >= Height) break;
//var f = fields[pos + (Width * i - i)];
//if (f == Piece.None) continue;
//if ((f & (Piece.Bishop | Piece.Queen)) != Piece.None && (f & checkerColor) != Piece.None) return true;
//break;
//}
//for (int i = 1; i < Math.Max(Width, Height); i++) // rechts-oben
//{
//if (posX + i >= Width || posY - i < 0) break;
//var f = fields[pos - (Width * i - i)];
//if (f == Piece.None) continue;
//if ((f & (Piece.Bishop | Piece.Queen)) != Piece.None && (f & checkerColor) != Piece.None) return true;
//break;
//}
//for (int i = 1; i < Math.Max(Width, Height); i++) // rechts-unten
//{
//if (posX + i >= Width || posY + i >= Height) break;
//var f = fields[pos + (Width * i + i)];
//if (f == Piece.None) continue;
//if ((f & (Piece.Bishop | Piece.Queen)) != Piece.None && (f & checkerColor) != Piece.None) return true;
//break;
//}
//}
//
//return false;
//}
//
///// <summary>
///// fragt das gesamte Spielbrett ab
///// </summary>
///// <param name="array">Array, wohin die Daten des Spielbrettes gespeichert werden sollen</param>
///// <param name="ofs">Startposition im Array</param>
///// <returns>Anzahl der geschriebenen Bytes</returns>
//public override int GetFastFen(byte[] array, int ofs)
//{
//int p = 0;
//int gap = 0;
//foreach (var field in fields)
//{
//if (field == Piece.None)
//{
//gap++;
//continue;
//}
//if (gap > 0)
//{
//array[ofs + p++] = (byte)(uint)gap;
//gap = 0;
//}
//array[ofs + p++] = (byte)field;
//}
//if (gap > 0) array[ofs + p++] = (byte)(uint)gap;
//array[ofs + p++] = (byte)((WhiteMove ? 1u : 0) | (WhiteCanCastleKingside ? 2u : 0) | (WhiteCanCastleQueenside ? 4u : 0) | (BlackCanCastleKingside ? 8u : 0) | (BlackCanCastleQueenside ? 16u : 0));
//array[ofs + p++] = (byte)(sbyte)EnPassantPos;
//array[ofs + p++] = (byte)(uint)HalfmoveClock;
//array[ofs + p++] = (byte)(uint)(HalfmoveClock >> 8);
//array[ofs + p++] = (byte)(uint)MoveNumber;
//array[ofs + p++] = (byte)(uint)(MoveNumber >> 8);
//
//return p;
//}
//
///// <summary>
///// setzt das gesamte Spielbrett
///// </summary>
///// <param name="array">Array, worraus die Daten des Spielbrettes gelesen werden sollen</param>
///// <param name="ofs">Startposition im Array</param>
///// <returns>Anzahl der gelesenen Bytes</returns>
//public override int SetFastFen(byte[] array, int ofs)
//{
//int p = 0;
//byte b;
//for (int i = 0; i < fields.Length; i++)
//{
//b = array[ofs + p++];
//if (b < 64) // gap found?
//{
//fields[i] = Piece.None;
//while (--b != 0) fields[++i] = Piece.None;
//continue;
//}
//fields[i] = (Piece)b;
//if (((Piece)b & Piece.King) != Piece.None)
//{
//if ((Piece)b == Piece.WhiteKing) whiteKingPos = i; else blackKingPos = i;
//}
//}
//b = array[ofs + p++];
//WhiteMove = (b & 1) != 0;
//WhiteCanCastleKingside = (b & 2) != 0;
//WhiteCanCastleQueenside = (b & 4) != 0;
//BlackCanCastleKingside = (b & 8) != 0;
//BlackCanCastleQueenside = (b & 16) != 0;
//EnPassantPos = (sbyte)array[ofs + p++];
//HalfmoveClock = array[ofs + p] | array[ofs + p + 1] << 8; p += sizeof(short);
//MoveNumber = array[ofs + p] | array[ofs + p + 1] << 8; p += sizeof(short);
//
//return p;
//}
//
///// <summary>
///// führt einen Zug durch und gibt true zurück, wenn dieser erfolgreich war
///// </summary>
///// <param name="move">Zug, welcher ausgeführt werden soll</param>
///// <param name="onlyCheck">optional: gibt an, dass der Zug nur geprüft aber nicht durchgeführt werden soll (default: false)</param>
///// <returns>true, wenn erfolgreich, sonst false</returns>
//public override bool DoMove(Move move, bool onlyCheck = false)
//{
//var piece = fields[move.fromPos];
//
//Debug.Assert((piece & Piece.BasicPieces) != Piece.None); // ist eine Figur auf dem Feld vorhanden?
//Debug.Assert(fields[move.toPos] == move.capturePiece); // stimmt die zu schlagende Figur mit dem Spielfeld überein?
//Debug.Assert((move.capturePiece & Piece.Colors) != (piece & Piece.Colors)); // wird keine eigene Figur gleicher Farbe geschlagen?
//Debug.Assert((piece & Piece.Colors) == (WhiteMove ? Piece.White : Piece.Black)); // passt die Figur-Farbe zum Zug?
//
//// --- Zug durchführen ---
//fields[move.toPos] = piece;
//fields[move.fromPos] = Piece.None;
//
//if (move.toPos == EnPassantPos && (piece & Piece.Pawn) != Piece.None) // ein Bauer schlägt "en passant"?
//{
//Debug.Assert(move.toPos % Width != move.fromPos % Width); // Spalte muss sich ändern
//Debug.Assert(move.capturePiece == Piece.None); // das Zielfeld enhält keine Figur (der geschlagene Bauer ist drüber oder drunter)
//int removePawnPos = WhiteMove ? move.toPos + Width : move.toPos - Width; // Position des zu schlagenden Bauern berechnen
//Debug.Assert(fields[removePawnPos] == (WhiteMove ? Piece.BlackPawn : Piece.WhitePawn)); // es wird ein Bauer erwartet, welcher geschlagen wird
//fields[removePawnPos] = Piece.None; // Bauer entfernen
//}
//
//if (move.promoPiece != Piece.None) fields[move.toPos] = move.promoPiece;
//
//if ((piece & Piece.King) != Piece.None) // wurde König gezogen?
//{
//if (piece == Piece.WhiteKing) whiteKingPos = move.toPos; else blackKingPos = move.toPos;
//}
//
//// --- prüfen, ob der König nach dem Zug im Schach steht ---
//{
//int kingPos = WhiteMove ? whiteKingPos : blackKingPos;
//if (!onlyCheck && kingPos == move.toPos && Math.Abs(move.toPos - move.fromPos) == 2) // wurde der König mit einer Rochade bewegt (zwei Felder seitlich)?
//{
//switch (kingPos)
//{
//case 2: // lange Rochade mit dem schwarzen König
//{
//Debug.Assert(BlackCanCastleQueenside); // lange Rochade sollte noch erlaubt sein
//Debug.Assert(fields[0] == Piece.BlackRook && fields[1] == Piece.None && fields[2] == Piece.BlackKing && fields[3] == Piece.None && fields[4] == Piece.None); // Felder prüfen
//fields[0] = Piece.None; fields[3] = Piece.BlackRook; // Turm bewegen
//} break;
//case 6: // kurze Rochade mit dem schwarzen König
//{
//Debug.Assert(BlackCanCastleKingside); // kurze Rochade sollte noch erlaubt sein
//Debug.Assert(fields[4] == Piece.None && fields[5] == Piece.None && fields[6] == Piece.BlackKing && fields[7] == Piece.BlackRook); // Felder prüfen
//fields[7] = Piece.None; fields[5] = Piece.BlackRook; // Turm bewegen
//} break;
//case 58: // lange Rochade mit dem weißen König
//{
//Debug.Assert(WhiteCanCastleQueenside); // lange Rochade sollte noch erlaubt sein
//Debug.Assert(fields[56] == Piece.WhiteRook && fields[57] == Piece.None && fields[58] == Piece.WhiteKing && fields[59] == Piece.None && fields[60] == Piece.None); // Felder prüfen
//fields[56] = Piece.None; fields[59] = Piece.WhiteRook; // Turm bewegen
//} break;
//case 62: // kurze Rochade mit dem weißen König
//{
//Debug.Assert(WhiteCanCastleKingside); // kurze Rochade sollte noch erlaubt sein
//Debug.Assert(fields[60] == Piece.None && fields[61] == Piece.None && fields[62] == Piece.WhiteKing && fields[63] == Piece.WhiteRook); // Felder prüfen
//fields[63] = Piece.None; fields[61] = Piece.WhiteRook; // Turm bewegen
//} break;
//default: throw new Exception(); // Rochade war unmöglich
//}
//}
//else if (IsChecked(kingPos, WhiteMove ? Piece.Black : Piece.White)) // prüfen, ob der eigene König vom Gegner angegriffen wird und noch im Schach steht
//{
//// --- Zug rückgängig machen ---
//fields[move.toPos] = move.capturePiece;
//fields[move.fromPos] = piece;
//if (move.toPos == EnPassantPos && (piece & Piece.Pawn) != Piece.None) // ein Bauer hat "en passant" geschlagen?
//{
//if (WhiteMove)
//{
//fields[move.toPos + Width] = Piece.BlackPawn; // schwarzen Bauer wieder zurück setzen
//}
//else
//{
//fields[move.toPos - Width] = Piece.WhitePawn; // weißen Bauer wieder zurück setzen
//}
//}
//if ((piece & Piece.King) != Piece.None) // wurde ein König gezogen?
//{
//if (piece == Piece.WhiteKing) whiteKingPos = move.fromPos; else blackKingPos = move.fromPos;
//}
//return false; // Zug war nicht erlaubt, da der König sonst im Schach stehen würde
//}
//}
//
//if (onlyCheck) // Zug sollte nur geprüft werden?
//{
//// --- Zug rückgängig machen ---
//fields[move.toPos] = move.capturePiece;
//fields[move.fromPos] = piece;
//if (move.toPos == EnPassantPos && (piece & Piece.Pawn) != Piece.None) // ein Bauer hat "en passant" geschlagen?
//{
//if (WhiteMove)
//{
//fields[move.toPos + Width] = Piece.BlackPawn; // schwarzen Bauer wieder zurück setzen
//}
//else
//{
//fields[move.toPos - Width] = Piece.WhitePawn; // weißen Bauer wieder zurück setzen
//}
//}
//if ((piece & Piece.King) != Piece.None) // wurde ein König gezogen?
//{
//if (piece == Piece.WhiteKing) whiteKingPos = move.fromPos; else blackKingPos = move.fromPos;
//}
//return true;
//}
//
//EnPassantPos = -1;
//if ((piece & Piece.Pawn) != Piece.None && Math.Abs(move.toPos - move.fromPos) == Width * 2) // wurde ein Bauer zwei Felder weit gezogen -> "en passant" vormerken
//{
//EnPassantPos = (move.fromPos + move.toPos) / 2;
//int posX = EnPassantPos % Width;
//bool opPawn = false;
//if (WhiteMove)
//{
//if (posX > 0 && fields[EnPassantPos - Width - 1] == Piece.BlackPawn) opPawn = true;
//if (posX < Width - 1 && fields[EnPassantPos - Width + 1] == Piece.BlackPawn) opPawn = true;
//}
//else
//{
//if (posX > 0 && fields[EnPassantPos + Width - 1] == Piece.WhitePawn) opPawn = true;
//if (posX < Width - 1 && fields[EnPassantPos + Width + 1] == Piece.WhitePawn) opPawn = true;
//}
//if (!opPawn) EnPassantPos = -1; // kein "en passant" möglich, da kein gegenerischer Bauer in der Nähe ist
//}
//
//// prüfen, ob durch den Zug Rochaden ungültig werden
//switch (move.fromPos)
//{
//case 0: BlackCanCastleQueenside = false; break; // linker schwarzer Turm wurde mindestens das erste Mal bewegt
//case 4: BlackCanCastleQueenside = false; BlackCanCastleKingside = false; break; // schwarzer König wurde mindestens das erste Mal bewegt
//case 7: BlackCanCastleKingside = false; break; // rechter schwarzer Turm wurde mindestens das erste Mal bewegt
//case 56: WhiteCanCastleQueenside = false; break; // linker weißer Turm wurde mindestens das erste Mal bewegt
//case 60: WhiteCanCastleQueenside = false; WhiteCanCastleKingside = false; break; // weißer König wurde mindestens das erste Mal bewegt
//case 63: WhiteCanCastleKingside = false; break; // rechter weißer Turm wurde mindestens das erste Mal bewegt
//}
//switch (move.toPos)
//{
//case 0: BlackCanCastleQueenside = false; break; // linker schwarzer Turm wurde geschlagen
//case 7: BlackCanCastleKingside = false; break; // rechter schwarzer Turm wurde geschlagen
//case 56: WhiteCanCastleQueenside = false; break; // linker weißer Turm wurde geschlagen
//case 63: WhiteCanCastleKingside = false; break; // rechter weißer Turm wurde geschlagen
//}
//
//WhiteMove = !WhiteMove; // Farbe welchseln, damit der andere Spieler am Zug ist
//HalfmoveClock++;
//if (piece == Piece.Pawn || move.capturePiece != Piece.None) HalfmoveClock = 0; // beim Bauernzug oder Schlagen einer Figur: 50-Züge Regel zurücksetzen
//if (WhiteMove) MoveNumber++; // Züge weiter hochzählen
//
//return true;
//}
//
///// <summary>
///// führt einen Zug direkt durch (ohne auf Gültigkeit zu prüfen)
///// </summary>
///// <param name="move">Zug, welcher ausgeführt werden soll</param>
//public override void DoMoveFast(Move move)
//{
//var piece = fields[move.fromPos];
//
//Debug.Assert((piece & Piece.BasicPieces) != Piece.None); // ist eine Figur auf dem Feld vorhanden?
//Debug.Assert(fields[move.toPos] == move.capturePiece); // stimmt die zu schlagende Figur mit dem Spielfeld überein?
//Debug.Assert((move.capturePiece & Piece.Colors) != (piece & Piece.Colors)); // wird keine eigene Figur gleicher Farbe geschlagen?
//Debug.Assert((piece & Piece.Colors) == (WhiteMove ? Piece.White : Piece.Black)); // passt die Figur-Farbe zum Zug?
//
//// --- Zug durchführen ---
//fields[move.toPos] = piece;
//fields[move.fromPos] = Piece.None;
//
//if (move.toPos == EnPassantPos && (piece & Piece.Pawn) != Piece.None) // ein Bauer schlägt "en passant"?
//{
//Debug.Assert(move.toPos % Width != move.fromPos % Width); // Spalte muss sich ändern
//Debug.Assert(move.capturePiece == Piece.None); // das Zielfeld enhält keine Figur (der geschlagene Bauer ist drüber oder drunter)
//int removePawnPos = WhiteMove ? move.toPos + Width : move.toPos - Width; // Position des zu schlagenden Bauern berechnen
//Debug.Assert(fields[removePawnPos] == (WhiteMove ? Piece.BlackPawn : Piece.WhitePawn)); // es wird ein Bauer erwartet, welcher geschlagen wird
//fields[removePawnPos] = Piece.None; // Bauer entfernen
//}
//
//if (move.promoPiece != Piece.None) fields[move.toPos] = move.promoPiece;
//
//if ((piece & Piece.King) != Piece.None) // wurde König gezogen?
//{
//if (piece == Piece.WhiteKing) whiteKingPos = move.toPos; else blackKingPos = move.toPos;
//}
//
//// --- Rochade ziehen ---
//{
//int kingPos = WhiteMove ? whiteKingPos : blackKingPos;
//if (kingPos == move.toPos && Math.Abs(move.toPos - move.fromPos) == 2) // wurde der König mit einer Rochade bewegt (zwei Felder seitlich)?
//{
//switch (kingPos)
//{
//case 2: // lange Rochade mit dem schwarzen König
//{
//Debug.Assert(BlackCanCastleQueenside); // lange Rochade sollte noch erlaubt sein
//Debug.Assert(fields[0] == Piece.BlackRook && fields[1] == Piece.None && fields[2] == Piece.BlackKing && fields[3] == Piece.None && fields[4] == Piece.None); // Felder prüfen
//fields[0] = Piece.None; fields[3] = Piece.BlackRook; // Turm bewegen
//} break;
//case 6: // kurze Rochade mit dem schwarzen König
//{
//Debug.Assert(BlackCanCastleKingside); // kurze Rochade sollte noch erlaubt sein
//Debug.Assert(fields[4] == Piece.None && fields[5] == Piece.None && fields[6] == Piece.BlackKing && fields[7] == Piece.BlackRook); // Felder prüfen
//fields[7] = Piece.None; fields[5] = Piece.BlackRook; // Turm bewegen
//} break;
//case 58: // lange Rochade mit dem weißen König
//{
//Debug.Assert(WhiteCanCastleQueenside); // lange Rochade sollte noch erlaubt sein
//Debug.Assert(fields[56] == Piece.WhiteRook && fields[57] == Piece.None && fields[58] == Piece.WhiteKing && fields[59] == Piece.None && fields[60] == Piece.None); // Felder prüfen
//fields[56] = Piece.None; fields[59] = Piece.WhiteRook; // Turm bewegen
//} break;
//case 62: // kurze Rochade mit dem weißen König
//{
//Debug.Assert(WhiteCanCastleKingside); // kurze Rochade sollte noch erlaubt sein
//Debug.Assert(fields[60] == Piece.None && fields[61] == Piece.None && fields[62] == Piece.WhiteKing && fields[63] == Piece.WhiteRook); // Felder prüfen
//fields[63] = Piece.None; fields[61] = Piece.WhiteRook; // Turm bewegen
//} break;
//default: throw new Exception(); // Rochade war unmöglich
//}
//}
//}
//
//EnPassantPos = -1;
//if ((piece & Piece.Pawn) != Piece.None && Math.Abs(move.toPos - move.fromPos) == Width * 2) // wurde ein Bauer zwei Felder weit gezogen -> "en passant" vormerken
//{
//EnPassantPos = (move.fromPos + move.toPos) / 2;
//int posX = EnPassantPos % Width;
//bool opPawn = false;
//if (WhiteMove)
//{
//if (posX > 0 && fields[EnPassantPos - Width - 1] == Piece.BlackPawn) opPawn = true;
//if (posX < Width - 1 && fields[EnPassantPos - Width + 1] == Piece.BlackPawn) opPawn = true;
//}
//else
//{
//if (posX > 0 && fields[EnPassantPos + Width - 1] == Piece.WhitePawn) opPawn = true;
//if (posX < Width - 1 && fields[EnPassantPos + Width + 1] == Piece.WhitePawn) opPawn = true;
//}
//if (!opPawn) EnPassantPos = -1; // kein "en passant" möglich, da kein gegenerischer Bauer in der Nähe ist
//}
//
//// prüfen, ob durch den Zug Rochaden ungültig werden
//switch (move.fromPos)
//{
//case 0: BlackCanCastleQueenside = false; break; // linker schwarzer Turm wurde mindestens das erste Mal bewegt
//case 4: BlackCanCastleQueenside = false; BlackCanCastleKingside = false; break; // schwarzer König wurde mindestens das erste Mal bewegt
//case 7: BlackCanCastleKingside = false; break; // rechter schwarzer Turm wurde mindestens das erste Mal bewegt
//case 56: WhiteCanCastleQueenside = false; break; // linker weißer Turm wurde mindestens das erste Mal bewegt
//case 60: WhiteCanCastleQueenside = false; WhiteCanCastleKingside = false; break; // weißer König wurde mindestens das erste Mal bewegt
//case 63: WhiteCanCastleKingside = false; break; // rechter weißer Turm wurde mindestens das erste Mal bewegt
//}
//switch (move.toPos)
//{
//case 0: BlackCanCastleQueenside = false; break; // linker schwarzer Turm wurde geschlagen
//case 7: BlackCanCastleKingside = false; break; // rechter schwarzer Turm wurde geschlagen
//case 56: WhiteCanCastleQueenside = false; break; // linker weißer Turm wurde geschlagen
//case 63: WhiteCanCastleKingside = false; break; // rechter weißer Turm wurde geschlagen
//}
//
//WhiteMove = !WhiteMove; // Farbe welchseln, damit der andere Spieler am Zug ist
//HalfmoveClock++;
//if (piece == Piece.Pawn || move.capturePiece != Piece.None) HalfmoveClock = 0; // beim Bauernzug oder Schlagen einer Figur: 50-Züge Regel zurücksetzen
//if (WhiteMove) MoveNumber++; // Züge weiter hochzählen
//}
//
///// <summary>
///// macht einen bestimmten Zug wieder Rückgängig
///// </summary>
///// <param name="move">Zug, welcher rückgängig gemacht werden soll</param>
///// <param name="lastBoardInfos">Spielbrettinformationen der vorherigen Stellung</param>
//public override void DoMoveBackward(Move move, BoardInfo lastBoardInfos)
//{
//// --- Figur zurückziehen ---
//var piece = fields[move.toPos];
//fields[move.fromPos] = piece; // Figur zurücksetzen
//fields[move.toPos] = move.capturePiece; // eventuell geschlagene Figur wiederherstellen
//
//// --- Bauer Umwandlung: promotion ---
//if (move.promoPiece != Piece.None)
//{
//Debug.Assert(piece == move.promoPiece); // umgewandelte Figur sollte übereinstimmen
//fields[move.fromPos] = (piece & Piece.Colors) | Piece.Pawn; // Figur zum Bauern zurück verwandeln :)
//}
//
//// --- Bauer hat "en passant" geschlagen ---
//if ((piece & Piece.Pawn) != Piece.None
//&& move.fromPos % Width != move.toPos % Width
//&& move.capturePiece == Piece.None) // hat der Bauer seitlich ein "Nichts" geschlagen? -> der gegenerische Bauer wurde dann ein Feld drüber/drunter entfernt
//{
//Debug.Assert((lastBoardInfos & BoardInfo.EnPassantMask) != BoardInfo.EnPassantNone); // war "en passant" vorher erlaubt?
//if (WhiteMove)
//{
//fields[(uint)(lastBoardInfos & BoardInfo.EnPassantMask) - Width] = Piece.WhitePawn;
//}
//else
//{
//fields[(uint)(lastBoardInfos & BoardInfo.EnPassantMask) + Width] = Piece.BlackPawn;
//}
//}
//
//// --- eine Rochade wurde gemacht ---
//if ((piece & Piece.King) != Piece.None) // wurde ein König bewegt?
//{
//if (piece == Piece.WhiteKing) whiteKingPos = move.fromPos; else blackKingPos = move.fromPos;
//
//if (Math.Abs(move.fromPos % Width - move.toPos % Width) > 1) // wurde ein König mehr als 1 Feld seitlich bewegt?
//{
//switch (move.toPos)
//{
//case 2: // schwarze lange Rochade auf der Damen-Seite (O-O-O)
//{
//Debug.Assert(fields[0] == Piece.None && fields[1] == Piece.None && fields[2] == Piece.None && fields[3] == Piece.BlackRook && fields[4] == Piece.BlackKing); // passen die Felder?
//Debug.Assert((lastBoardInfos & BoardInfo.BlackCanCastleQueenside) != BoardInfo.None); // war Rochade vorher erlaubt?
//fields[0] = Piece.BlackRook; fields[3] = Piece.None; // Turm zurück in die Ecke setzen
//} break;
//
//case 6: // schwarze kurze Rochade auf der Königs-Seite (O-O)
//{
//Debug.Assert(fields[4] == Piece.BlackKing && fields[5] == Piece.BlackRook && fields[6] == Piece.None && fields[7] == Piece.None);
//Debug.Assert((lastBoardInfos & BoardInfo.BlackCanCastleKingside) != BoardInfo.None); // war Rochade vorher erlaubt?
//fields[7] = Piece.BlackRook; fields[5] = Piece.None; // Turm zurück in die Ecke setzen
//} break;
//
//case 58: // weiße lange Rochade auf der Damen-Seite (O-O-O)
//{
//Debug.Assert(fields[56] == Piece.None && fields[57] == Piece.None && fields[58] == Piece.None && fields[59] == Piece.WhiteRook && fields[60] == Piece.WhiteKing); // passen die Felder?
//Debug.Assert((lastBoardInfos & BoardInfo.WhiteCanCastleQueenside) != BoardInfo.None); // war Rochade vorher erlaubt?
//fields[56] = Piece.WhiteRook; fields[59] = Piece.None; // Turm zurück in die Ecke setzen
//} break;
//
//case 62: // weiße kurze Rochade auf der Königs-Seite (O-O)
//{
//Debug.Assert(fields[60] == Piece.WhiteKing && fields[61] == Piece.WhiteRook && fields[62] == Piece.None && fields[63] == Piece.None); // passen die Felder?
//Debug.Assert((lastBoardInfos & BoardInfo.WhiteCanCastleKingside) != BoardInfo.None); // war Rochade vorher erlaubt?
//fields[63] = Piece.WhiteRook; fields[61] = Piece.None; // Turm zurück in die Ecke setzen
//} break;
//
//default: throw new Exception("invalid move"); // König hat sich seltsam bewegt
//}
//}
//}
//
//// --- Spielbrett Infos anpassen ---
//if (WhiteMove) MoveNumber--;
//WhiteMove = !WhiteMove;
//BoardInfos = lastBoardInfos;
//}
//
//void GetWhiteMoves(MoveListSingle moveList)
//{
//for (int pos = fields.Length - 1; pos >= 0; pos--)
//{
//var piece = fields[pos];
//if ((piece & Piece.Colors) != Piece.White) continue; // Farbe der Figur passt nicht zum Zug oder das Feld ist leer
//
//if ((piece & Piece.Pawn) != Piece.None && pos < Width * 2)
//{
//// Promotion-Zug gefunden? (ein Bauer hat das Ziel erreicht und wird umgewandelt)
//foreach (int movePos in ScanMove(pos)) // alle theoretischen Bewegungsmöglichkeiten prüfen
//{
//var move = new Move(pos, movePos, fields[movePos], Piece.White | Piece.Queen);
//if (DoMove(move, true)) // gültigen Zug gefunden?
//{
//moveList.Add(move);
//// weitere Wahlmöglichkeiten als gültige Züge zurück geben
//move.promoPiece = Piece.White | Piece.Rook;
//moveList.Add(move);
//move.promoPiece = Piece.White | Piece.Bishop;
//moveList.Add(move);
//move.promoPiece = Piece.White | Piece.Knight;
//moveList.Add(move);
//}
//}
//}
//else
//{
//foreach (int movePos in ScanMove(pos)) // alle theoretischen Bewegungsmöglichkeiten prüfen
//{
//var move = new Move(pos, movePos, fields[movePos], Piece.None);
//if (DoMove(move, true)) // gültigen Zug gefunden?
//{
//moveList.Add(move);
//}
//}
//
//// Rochade-Züge prüfen
//if (pos == 60 && piece == Piece.WhiteKing) // der weiße König steht noch auf der Startposition?
//{
//if (WhiteCanCastleQueenside // lange Rochade O-O-O möglich?
//&& fields[57] == Piece.None && fields[58] == Piece.None && fields[59] == Piece.None // sind die Felder noch frei?
//&& !IsChecked(58, Piece.Black) && !IsChecked(59, Piece.Black) && !IsChecked(60, Piece.Black)) // steht der König und seine Laufwege auch nicht im Schach?
//{
//Debug.Assert(fields[56] == Piece.WhiteRook); // weißer Turm sollte noch in der Ecke stehen
//moveList.Add(new Move(pos, pos - 2, Piece.None, Piece.None)); // König läuft zwei Felder = Rochade
//}
//if (WhiteCanCastleKingside // kurze Rochade O-O möglich?
//&& fields[61] == Piece.None && fields[62] == Piece.None // sind die Felder noch frei?
//&& !IsChecked(60, Piece.Black) && !IsChecked(61, Piece.Black) && !IsChecked(62, Piece.Black)) // steht der König und seine Laufwege auch nicht im Schach?
//{
//Debug.Assert(fields[63] == Piece.WhiteRook); // weißer Turm solle noch in der Ecke stehen
//moveList.Add(new Move(pos, pos + 2, Piece.None, Piece.None)); // König läuft zwei Felder = Rochade
//}
//}
//else if (pos == 4 && piece == Piece.BlackKing) // der schwarze König steht noch auf der Startposition?
//{
//if (BlackCanCastleQueenside // lange Rochade O-O-O möglich?
//&& fields[1] == Piece.None && fields[2] == Piece.None && fields[3] == Piece.None // sind die Felder noch frei?
//&& !IsChecked(2, Piece.White) && !IsChecked(3, Piece.White) && !IsChecked(4, Piece.White)) // steht der König und seine Laufwege auch nicht im Schach?
//{
//Debug.Assert(fields[0] == Piece.BlackRook); // schwarzer Turm sollte noch in der Ecke stehen
//moveList.Add(new Move(pos, pos - 2, Piece.None, Piece.None)); // König läuft zwei Felder = Rochade
//}
//if (BlackCanCastleKingside // kurze Rochade O-O möglich?
//&& fields[5] == Piece.None && fields[6] == Piece.None // sind die Felder noch frei?
//&& !IsChecked(4, Piece.White) && !IsChecked(5, Piece.White) && !IsChecked(6, Piece.White)) // steht der König und seine Laufwege auch nicht im Schach?
//{
//Debug.Assert(fields[7] == Piece.BlackRook); // schwarzer Turm solle noch in der Ecke stehen
//moveList.Add(new Move(pos, pos + 2, Piece.None, Piece.None)); // König läuft zwei Felder = Rochade
//}
//}
//}
//}
//}
//
//void GetBlackMoves(MoveListSingle moveList)
//{
//for (int pos = 0; pos < fields.Length; pos++)
//{
//var piece = fields[pos];
//if ((piece & Piece.Colors) != Piece.Black) continue; // Farbe der Figur passt nicht zum Zug oder das Feld ist leer
//
//if ((piece & Piece.Pawn) != Piece.None && pos >= Height * Width - Width * 2)
//{
//// Promotion-Zug gefunden? (ein Bauer hat das Ziel erreicht und wird umgewandelt)
//foreach (int movePos in ScanMove(pos)) // alle theoretischen Bewegungsmöglichkeiten prüfen
//{
//var move = new Move(pos, movePos, fields[movePos], Piece.Black | Piece.Queen);
//if (DoMove(move, true)) // gültigen Zug gefunden?
//{
//moveList.Add(move);
//// weitere Wahlmöglichkeiten als gültige Züge zurück geben
//move.promoPiece = Piece.Black | Piece.Rook;
//moveList.Add(move);
//move.promoPiece = Piece.Black | Piece.Bishop;
//moveList.Add(move);
//move.promoPiece = Piece.Black | Piece.Knight;
//moveList.Add(move);
//}
//}
//}
//else
//{
//foreach (int movePos in ScanMove(pos)) // alle theoretischen Bewegungsmöglichkeiten prüfen
//{
//var move = new Move(pos, movePos, fields[movePos], Piece.None);
//if (DoMove(move, true)) // gültigen Zug gefunden?
//{
//moveList.Add(move);
//}
//}
//
//// Rochade-Züge prüfen
//if (pos == 60 && piece == Piece.WhiteKing) // der weiße König steht noch auf der Startposition?
//{
//if (WhiteCanCastleQueenside // lange Rochade O-O-O möglich?
//&& fields[57] == Piece.None && fields[58] == Piece.None && fields[59] == Piece.None // sind die Felder noch frei?
//&& !IsChecked(58, Piece.Black) && !IsChecked(59, Piece.Black) && !IsChecked(60, Piece.Black)) // steht der König und seine Laufwege auch nicht im Schach?
//{
//Debug.Assert(fields[56] == Piece.WhiteRook); // weißer Turm sollte noch in der Ecke stehen
//moveList.Add(new Move(pos, pos - 2, Piece.None, Piece.None)); // König läuft zwei Felder = Rochade
//}
//if (WhiteCanCastleKingside // kurze Rochade O-O möglich?
//&& fields[61] == Piece.None && fields[62] == Piece.None // sind die Felder noch frei?
//&& !IsChecked(60, Piece.Black) && !IsChecked(61, Piece.Black) && !IsChecked(62, Piece.Black)) // steht der König und seine Laufwege auch nicht im Schach?
//{
//Debug.Assert(fields[63] == Piece.WhiteRook); // weißer Turm solle noch in der Ecke stehen
//moveList.Add(new Move(pos, pos + 2, Piece.None, Piece.None)); // König läuft zwei Felder = Rochade
//}
//}
//else if (pos == 4 && piece == Piece.BlackKing) // der schwarze König steht noch auf der Startposition?
//{
//if (BlackCanCastleQueenside // lange Rochade O-O-O möglich?
//&& fields[1] == Piece.None && fields[2] == Piece.None && fields[3] == Piece.None // sind die Felder noch frei?
//&& !IsChecked(2, Piece.White) && !IsChecked(3, Piece.White) && !IsChecked(4, Piece.White)) // steht der König und seine Laufwege auch nicht im Schach?
//{
//Debug.Assert(fields[0] == Piece.BlackRook); // schwarzer Turm sollte noch in der Ecke stehen
//moveList.Add(new Move(pos, pos - 2, Piece.None, Piece.None)); // König läuft zwei Felder = Rochade
//}
//if (BlackCanCastleKingside // kurze Rochade O-O möglich?
//&& fields[5] == Piece.None && fields[6] == Piece.None // sind die Felder noch frei?
//&& !IsChecked(4, Piece.White) && !IsChecked(5, Piece.White) && !IsChecked(6, Piece.White)) // steht der König und seine Laufwege auch nicht im Schach?
//{
//Debug.Assert(fields[7] == Piece.BlackRook); // schwarzer Turm solle noch in der Ecke stehen
//moveList.Add(new Move(pos, pos + 2, Piece.None, Piece.None)); // König läuft zwei Felder = Rochade
//}
//}
//}
//}
//}
//
//IEnumerable<Move> GetWhiteMoves()
//{
//for (int pos = fields.Length - 1; pos >= 0; pos--)
//{
//var piece = fields[pos];
//if ((piece & Piece.Colors) != Piece.White) continue; // Farbe der Figur passt nicht zum Zug oder das Feld ist leer
//
//if ((piece & Piece.Pawn) != Piece.None && pos < Width * 2)
//{
//// Promotion-Zug gefunden? (ein Bauer hat das Ziel erreicht und wird umgewandelt)
//foreach (int movePos in ScanMove(pos)) // alle theoretischen Bewegungsmöglichkeiten prüfen
//{
//var move = new Move(pos, movePos, fields[movePos], Piece.White | Piece.Queen);
//if (DoMove(move, true)) // gültigen Zug gefunden?
//{
//yield return move;
//// weitere Wahlmöglichkeiten als gültige Züge zurück geben
//move.promoPiece = Piece.White | Piece.Rook;
//yield return move;
//move.promoPiece = Piece.White | Piece.Bishop;
//yield return move;
//move.promoPiece = Piece.White | Piece.Knight;
//yield return move;
//}
//}
//}
//else
//{
//foreach (int movePos in ScanMove(pos)) // alle theoretischen Bewegungsmöglichkeiten prüfen
//{
//var move = new Move(pos, movePos, fields[movePos], Piece.None);
//if (DoMove(move, true)) // gültigen Zug gefunden?
//{
//yield return move;
//}
//}
//
//// Rochade-Züge prüfen
//if (pos == 60 && piece == Piece.WhiteKing) // der weiße König steht noch auf der Startposition?
//{
//if (WhiteCanCastleQueenside // lange Rochade O-O-O möglich?
//&& fields[57] == Piece.None && fields[58] == Piece.None && fields[59] == Piece.None // sind die Felder noch frei?
//&& !IsChecked(58, Piece.Black) && !IsChecked(59, Piece.Black) && !IsChecked(60, Piece.Black)) // steht der König und seine Laufwege auch nicht im Schach?
//{
//Debug.Assert(fields[56] == Piece.WhiteRook); // weißer Turm sollte noch in der Ecke stehen
//yield return new Move(pos, pos - 2, Piece.None, Piece.None); // König läuft zwei Felder = Rochade
//}
//if (WhiteCanCastleKingside // kurze Rochade O-O möglich?
//&& fields[61] == Piece.None && fields[62] == Piece.None // sind die Felder noch frei?
//&& !IsChecked(60, Piece.Black) && !IsChecked(61, Piece.Black) && !IsChecked(62, Piece.Black)) // steht der König und seine Laufwege auch nicht im Schach?
//{
//Debug.Assert(fields[63] == Piece.WhiteRook); // weißer Turm solle noch in der Ecke stehen
//yield return new Move(pos, pos + 2, Piece.None, Piece.None); // König läuft zwei Felder = Rochade
//}
//}
//else if (pos == 4 && piece == Piece.BlackKing) // der schwarze König steht noch auf der Startposition?
//{
//if (BlackCanCastleQueenside // lange Rochade O-O-O möglich?
//&& fields[1] == Piece.None && fields[2] == Piece.None && fields[3] == Piece.None // sind die Felder noch frei?
//&& !IsChecked(2, Piece.White) && !IsChecked(3, Piece.White) && !IsChecked(4, Piece.White)) // steht der König und seine Laufwege auch nicht im Schach?
//{
//Debug.Assert(fields[0] == Piece.BlackRook); // schwarzer Turm sollte noch in der Ecke stehen
//yield return new Move(pos, pos - 2, Piece.None, Piece.None); // König läuft zwei Felder = Rochade
//}
//if (BlackCanCastleKingside // kurze Rochade O-O möglich?
//&& fields[5] == Piece.None && fields[6] == Piece.None // sind die Felder noch frei?
//&& !IsChecked(4, Piece.White) && !IsChecked(5, Piece.White) && !IsChecked(6, Piece.White)) // steht der König und seine Laufwege auch nicht im Schach?
//{
//Debug.Assert(fields[7] == Piece.BlackRook); // schwarzer Turm solle noch in der Ecke stehen
//yield return new Move(pos, pos + 2, Piece.None, Piece.None); // König läuft zwei Felder = Rochade
//}
//}
//}
//}
//}
//
//IEnumerable<Move> GetBlackMoves()
//{
//for (int pos = 0; pos < fields.Length; pos++)
//{
//var piece = fields[pos];
//if ((piece & Piece.Colors) != Piece.Black) continue; // Farbe der Figur passt nicht zum Zug oder das Feld ist leer
//
//if ((piece & Piece.Pawn) != Piece.None && pos >= Height * Width - Width * 2)
//{
//// Promotion-Zug gefunden? (ein Bauer hat das Ziel erreicht und wird umgewandelt)
//foreach (int movePos in ScanMove(pos)) // alle theoretischen Bewegungsmöglichkeiten prüfen
//{
//var move = new Move(pos, movePos, fields[movePos], Piece.Black | Piece.Queen);
//if (DoMove(move, true)) // gültigen Zug gefunden?
//{
//yield return move;
//// weitere Wahlmöglichkeiten als gültige Züge zurück geben
//move.promoPiece = Piece.Black | Piece.Rook;
//yield return move;
//move.promoPiece = Piece.Black | Piece.Bishop;
//yield return move;
//move.promoPiece = Piece.Black | Piece.Knight;
//yield return move;
//}
//}
//}
//else
//{
//foreach (int movePos in ScanMove(pos)) // alle theoretischen Bewegungsmöglichkeiten prüfen
//{
//var move = new Move(pos, movePos, fields[movePos], Piece.None);
//if (DoMove(move, true)) // gültigen Zug gefunden?
//{
//yield return move;
//}
//}
//
//// Rochade-Züge prüfen
//if (pos == 60 && piece == Piece.WhiteKing) // der weiße König steht noch auf der Startposition?
//{
//if (WhiteCanCastleQueenside // lange Rochade O-O-O möglich?
//&& fields[57] == Piece.None && fields[58] == Piece.None && fields[59] == Piece.None // sind die Felder noch frei?
//&& !IsChecked(58, Piece.Black) && !IsChecked(59, Piece.Black) && !IsChecked(60, Piece.Black)) // steht der König und seine Laufwege auch nicht im Schach?
//{
//Debug.Assert(fields[56] == Piece.WhiteRook); // weißer Turm sollte noch in der Ecke stehen
//yield return new Move(pos, pos - 2, Piece.None, Piece.None); // König läuft zwei Felder = Rochade
//}
//if (WhiteCanCastleKingside // kurze Rochade O-O möglich?
//&& fields[61] == Piece.None && fields[62] == Piece.None // sind die Felder noch frei?
//&& !IsChecked(60, Piece.Black) && !IsChecked(61, Piece.Black) && !IsChecked(62, Piece.Black)) // steht der König und seine Laufwege auch nicht im Schach?
//{
//Debug.Assert(fields[63] == Piece.WhiteRook); // weißer Turm solle noch in der Ecke stehen
//yield return new Move(pos, pos + 2, Piece.None, Piece.None); // König läuft zwei Felder = Rochade
//}
//}
//else if (pos == 4 && piece == Piece.BlackKing) // der schwarze König steht noch auf der Startposition?
//{
//if (BlackCanCastleQueenside // lange Rochade O-O-O möglich?
//&& fields[1] == Piece.None && fields[2] == Piece.None && fields[3] == Piece.None // sind die Felder noch frei?
//&& !IsChecked(2, Piece.White) && !IsChecked(3, Piece.White) && !IsChecked(4, Piece.White)) // steht der König und seine Laufwege auch nicht im Schach?
//{
//Debug.Assert(fields[0] == Piece.BlackRook); // schwarzer Turm sollte noch in der Ecke stehen
//yield return new Move(pos, pos - 2, Piece.None, Piece.None); // König läuft zwei Felder = Rochade
//}
//if (BlackCanCastleKingside // kurze Rochade O-O möglich?
//&& fields[5] == Piece.None && fields[6] == Piece.None // sind die Felder noch frei?
//&& !IsChecked(4, Piece.White) && !IsChecked(5, Piece.White) && !IsChecked(6, Piece.White)) // steht der König und seine Laufwege auch nicht im Schach?
//{
//Debug.Assert(fields[7] == Piece.BlackRook); // schwarzer Turm solle noch in der Ecke stehen
//yield return new Move(pos, pos + 2, Piece.None, Piece.None); // König läuft zwei Felder = Rochade
//}
//}
//}
//}
//}
//
///// <summary>
///// berechnet alle erlaubten Zugmöglichkeiten und gibt diese zurück
///// </summary>
///// <returns>Aufzählung der Zugmöglichkeiten</returns>
//public override IEnumerable<Move> GetMoves()
//{
//return WhiteMove ? GetWhiteMoves() : GetBlackMoves();
//}
//
//readonly MoveListSingle tmpMoveList = new MoveListSingle((byte*)Marshal.AllocHGlobal(sizeof(Move) * 256));
//
///// <summary>
///// gibt alle Züge als Array zurück
///// </summary>
///// <returns>alle gültigen Züge als Array</returns>
//public override Move[] GetMovesArray()
//{
//tmpMoveList.Clear();
//if (WhiteMove)
//{
//GetWhiteMoves(tmpMoveList);
//}
//else
//{
//GetBlackMoves(tmpMoveList);
//}
//return tmpMoveList.ToArray();
//}
//
///// <summary>
///// gibt an, ob irgend ein Zug möglich ist
///// </summary>
//public override bool HasMoves
//{
//get
//{
//var fs = fields;
//if (WhiteMove)
//{
//int posX = whiteKingPos % Width;
//fs[whiteKingPos] = Piece.None;
//
//if (posX > 0 && (fs[whiteKingPos - 1] & Piece.White) == Piece.None && !IsChecked(whiteKingPos - 1, Piece.Black))
//{
//fs[whiteKingPos] = Piece.WhiteKing;
//return true;
//}
//
//if (posX < 7 && (fs[whiteKingPos + 1] & Piece.White) == Piece.None && !IsChecked(whiteKingPos + 1, Piece.Black))
//{
//fs[whiteKingPos] = Piece.WhiteKing;
//return true;
//}
//
//fs[whiteKingPos] = Piece.WhiteKing;
//return GetWhiteMoves().Any();
//}
//else
//{
//int posX = blackKingPos % Width;
//fs[blackKingPos] = Piece.None;
//
//if (posX > 0 && (fs[blackKingPos - 1] & Piece.Black) == Piece.None && !IsChecked(blackKingPos - 1, Piece.White))
//{
//fs[blackKingPos] = Piece.BlackKing;
//return true;
//}
//if (posX < 7 && (fs[blackKingPos + 1] & Piece.Black) == Piece.None && !IsChecked(blackKingPos + 1, Piece.White))
//{
//fs[blackKingPos] = Piece.BlackKing;
//return true;
//}
//
//fs[blackKingPos] = Piece.BlackKing;
//return GetBlackMoves().Any();
//}
//}
//}
//#endregion
//}
