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
