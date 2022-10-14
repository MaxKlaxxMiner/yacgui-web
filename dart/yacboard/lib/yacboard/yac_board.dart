import 'dart:typed_data';
import 'board_info.dart';
import 'fen.dart';
import 'board_size.dart';
import 'piece.dart';
import 'pos.dart';

class YacBoard {
  final Uint8List fields = Uint8List(BoardSize.fieldCount);

  int halfmoveClock = 0;
  int moveNumber = 1;
  int whiteKingPos = -1;
  int blackKingPos = -1;
  int enPassantPos = BoardInfo.EnPassantNone;

  bool whiteMove = true;
  bool whiteCanCastleKingside = false;
  bool whiteCanCastleQueenside = false;
  bool blackCanCastleKingside = false;
  bool blackCanCastleQueenside = false;

  YacBoard({String startFen = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"}) {
    setFen(startFen);
  }

  void clear() {
    for (int i = 0; i < fields.length; i++) {
      fields[i] = 0;
    }
    halfmoveClock = 0;
    moveNumber = 1;
    whiteKingPos = -1;
    blackKingPos = -1;
    enPassantPos = BoardInfo.EnPassantNone;
    whiteMove = true;
    whiteCanCastleKingside = false;
    whiteCanCastleQueenside = false;
    blackCanCastleKingside = false;
    blackCanCastleQueenside = false;
  }

  void setField(int pos, int piece) {
    if (pos < 0 || pos >= BoardSize.fieldCount) return; // out of range
    fields[pos] = piece;
    if (piece & Piece.king != 0) {
      if (piece == Piece.whiteKing) {
        whiteKingPos = pos;
      } else {
        blackKingPos = pos;
      }
    }
  }

  int getField(int pos) {
    if (pos < 0 || pos >= BoardSize.fieldCount) return Piece.blocked;
    return fields[pos];
  }

  int getBoardInfo() {
    int result = enPassantPos | (halfmoveClock << 16);

    if (whiteCanCastleKingside) result |= BoardInfo.WhiteCanCastleKingside;
    if (whiteCanCastleQueenside) result |= BoardInfo.WhiteCanCastleQueenside;
    if (blackCanCastleKingside) result |= BoardInfo.BlackCanCastleKingside;
    if (blackCanCastleQueenside) result |= BoardInfo.BlackCanCastleQueenside;

    return result;
  }

  @override
  String toString() {
    List<int> result = [];
    for (int y = 0; y < BoardSize.height; y++) {
      result.addAll([32, 32, 32, 32]); // "    "
      for (int x = 0; x < BoardSize.width; x++) {
        result.addAll(Piece.toChar(getField(Pos.fromXY(x, y))).codeUnits);
      }
      result.add(10); // '\n'
    }
    result.addAll(("\nFEN: ${getFen()}").codeUnits);
    return String.fromCharCodes(result);
  }
}
