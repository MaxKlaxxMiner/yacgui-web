import 'dart:typed_data';
import 'board_info.dart';
import 'is_checked.dart';
import 'board_size.dart';
import 'move.dart';
import 'piece.dart';
import 'yac_board.dart';

extension MoveGenerator on YacBoard {
  Int32List getMoves() {
    var tmp = Int32List(256);
    int count = getMovesFast(tmp);
    return tmp.buffer.asInt32List(0, count);
  }

  doMove(int move) {
    int p = fields[Move.fromPos(move)];

    fields[Move.toPos(move)] = p;
    fields[Move.fromPos(move)] = Piece.none;

    if (Move.toPos(move) == enPassantPos && p & Piece.pawn != Piece.none) {
      // "en passant"?
      if (whiteMove) {
        fields[Move.toPos(move) + BoardSize.width] = Piece.none;
      } else {
        fields[Move.toPos(move) - BoardSize.width] = Piece.none;
      }
    }

    if (Move.promotionPiece(move) != Piece.none) {
      // pawn move with promotion?
      fields[Move.toPos(move)] = Move.promotionPiece(move);
    }

    if (p & Piece.king != Piece.none) {
      // kingmove?
      if (p == Piece.whiteKing) {
        whiteKingPos = Move.toPos(move);
      } else {
        blackKingPos = Move.toPos(move);
      }
    }

    // --- is the king is in check ---
    {
      int kingPos = whiteMove ? whiteKingPos : blackKingPos;
      if (kingPos == Move.toPos(move) && (Move.toPos(move) - Move.fromPos(move) == 2 || Move.toPos(move) - Move.fromPos(move) == -2)) {
        switch (kingPos) {
          case 2:
            fields[0] = Piece.none;
            fields[3] = Piece.blackRook;
            break;
          case 6:
            fields[7] = Piece.none;
            fields[5] = Piece.blackRook;
            break;
          case 58:
            fields[56] = Piece.none;
            fields[59] = Piece.whiteRook;
            break;
          case 62:
            fields[63] = Piece.none;
            fields[61] = Piece.whiteRook;
            break;
        }
      }
    }

    enPassantPos = BoardInfo.EnPassantNone;
    if (p & Piece.pawn != Piece.none && (Move.toPos(move) - Move.fromPos(move) == BoardSize.width * 2 || Move.fromPos(move) - Move.toPos(move) == BoardSize.width * 2)) {
      enPassantPos = (Move.fromPos(move) + Move.toPos(move)) ~/ 2;
      int posX = enPassantPos % BoardSize.width;
      bool opPawn = false;
      if (whiteMove) {
        if (posX > 0 && fields[enPassantPos - BoardSize.width - 1] == Piece.blackPawn) opPawn = true;
        if (posX < BoardSize.width - 1 && fields[enPassantPos - BoardSize.width + 1] == Piece.blackPawn) opPawn = true;
      } else {
        if (posX > 0 && fields[enPassantPos + BoardSize.width - 1] == Piece.whitePawn) opPawn = true;
        if (posX < BoardSize.width - 1 && fields[enPassantPos + BoardSize.width + 1] == Piece.whitePawn) opPawn = true;
      }
      if (!opPawn) enPassantPos = BoardInfo.EnPassantNone;
    }

    switch (Move.fromPos(move)) {
      case 0:
        blackCanCastleQueenside = false;
        break;
      case 4:
        blackCanCastleQueenside = false;
        blackCanCastleKingside = false;
        break;
      case 7:
        blackCanCastleKingside = false;
        break;
      case 56:
        whiteCanCastleQueenside = false;
        break;
      case 60:
        whiteCanCastleQueenside = false;
        whiteCanCastleKingside = false;
        break;
      case 63:
        whiteCanCastleKingside = false;
        break;
    }
    switch (Move.toPos(move)) {
      case 0:
        blackCanCastleQueenside = false;
        break;
      case 7:
        blackCanCastleKingside = false;
        break;
      case 56:
        whiteCanCastleQueenside = false;
        break;
      case 63:
        whiteCanCastleKingside = false;
        break;
    }

    whiteMove = !whiteMove;
    halfmoveClock++;
    if (p & Piece.pawn == Piece.pawn || Move.capturePiece(move) != Piece.none) {
      halfmoveClock = 0;
    }
    if (whiteMove) moveNumber++;
  }

  void doMoveBackward(int move, int lastBoardInfos) {
    int p = fields[Move.toPos(move)];
    fields[Move.fromPos(move)] = p;
    fields[Move.toPos(move)] = Move.capturePiece(move);

    if (Move.promotionPiece(move) != Piece.none) {
      fields[Move.fromPos(move)] = (p & Piece.colors) | Piece.pawn;
    }

    if (p & Piece.pawn != Piece.none && Move.fromPos(move) % BoardSize.width != Move.toPos(move) % BoardSize.width && Move.capturePiece(move) == Piece.none) {
      if (whiteMove) {
        fields[(lastBoardInfos & BoardInfo.EnPassantMask) - BoardSize.width] = Piece.whitePawn;
      } else {
        fields[(lastBoardInfos & BoardInfo.EnPassantMask) + BoardSize.width] = Piece.blackPawn;
      }
    }

    if (p & Piece.king != Piece.none) {
      if (p == Piece.whiteKing) {
        whiteKingPos = Move.fromPos(move);
      } else {
        blackKingPos = Move.fromPos(move);
      }

      int posXdif = Move.fromPos(move) % BoardSize.width - Move.toPos(move) % BoardSize.width;
      if (posXdif > 1 || posXdif < -1) {
        switch (Move.toPos(move)) {
          case 2: // black O-O-O
            fields[0] = Piece.blackRook;
            fields[3] = Piece.none;
            break;
          case 6: // black O-O
            fields[7] = Piece.blackRook;
            fields[5] = Piece.none;
            break;
          case 58: // white O-O-O
            fields[56] = Piece.whiteRook;
            fields[59] = Piece.none;
            break;
          case 62: // white O-O
            fields[63] = Piece.whiteRook;
            fields[61] = Piece.none;
            break;
        }
      }
    }

    if (whiteMove) moveNumber--;
    whiteMove = !whiteMove;
    setBoardInfo(lastBoardInfos);
  }

  bool moveCheck(int move) {
    var moves = getMoves();
    for (int i = 0; i < moves.length; i++) {
      if (moves[i] == move) return true;
    }
    return false;
  }

  int getMovesFast(Int32List moves) {
    if (whiteMove) {
      return _getWhiteMoves(moves);
    } else {
      return _getBlackMoves(moves);
    }
  }

  bool _simpleMoveCheck(int move) {
    var p = fields[Move.fromPos(move)];

    fields[Move.toPos(move)] = p;
    fields[Move.fromPos(move)] = Piece.none;

    if (Move.toPos(move) == enPassantPos && p & Piece.pawn != Piece.none) {
      // "en passant"?
      if (whiteMove) {
        fields[Move.toPos(move) + BoardSize.width] = Piece.none;
      } else {
        fields[Move.toPos(move) - BoardSize.width] = Piece.none;
      }
    }

    if (Move.promotionPiece(move) != Piece.none) {
      fields[Move.toPos(move)] = Move.promotionPiece(move);
    }

    if (p & Piece.king != Piece.none) {
      // kingmove?
      if (p == Piece.whiteKing) {
        whiteKingPos = Move.toPos(move);
      } else {
        blackKingPos = Move.toPos(move);
      }
    }

    // --- is the king is in check ---
    {
      int kingPos;
      if (whiteMove) {
        kingPos = whiteKingPos;
      } else {
        kingPos = blackKingPos;
      }
      if (isCheckedPos(kingPos, invertedMoveColor())) {
        fields[Move.toPos(move)] = Move.capturePiece(move);
        fields[Move.fromPos(move)] = p;
        if (Move.toPos(move) == enPassantPos && p & Piece.pawn != Piece.none) {
          // "en passant" ?
          if (whiteMove) {
            fields[Move.toPos(move) + BoardSize.width] = Piece.blackPawn;
          } else {
            fields[Move.toPos(move) - BoardSize.width] = Piece.whitePawn;
          }
        }
        if (p & Piece.king != Piece.none) {
          if (p == Piece.whiteKing) {
            whiteKingPos = Move.fromPos(move);
          } else {
            blackKingPos = Move.fromPos(move);
          }
        }
        return false;
      }
    }

    fields[Move.toPos(move)] = Move.capturePiece(move);
    fields[Move.fromPos(move)] = p;
    if (Move.toPos(move) == enPassantPos && p & Piece.pawn != Piece.none) {
      if (whiteMove) {
        fields[Move.toPos(move) + BoardSize.width] = Piece.blackPawn;
      } else {
        fields[Move.toPos(move) - BoardSize.width] = Piece.whitePawn;
      }
    }
    if (p & Piece.king != Piece.none) {
      if (p == Piece.whiteKing) {
        whiteKingPos = Move.fromPos(move);
      } else {
        blackKingPos = Move.fromPos(move);
      }
    }
    return true;
  }

  int _getWhiteMoves(Int32List mv) {
    int mi = 0;
    for (int pos = fields.length - 1; pos >= 0; pos--) {
      var field = fields[pos];
      if (field & Piece.colors != Piece.white) continue; // wrong color / no p?
      int posX = pos % BoardSize.width;
      int posY = pos ~/ BoardSize.width;

      if (field == Piece.whitePawn && pos < BoardSize.width * 2) {
        if (fields[pos - BoardSize.width] == Piece.none) {
          var move = Move.full(fromPos: pos, toPos: pos - BoardSize.width, capturePiece: fields[pos - BoardSize.width], promotionPiece: Piece.whiteQueen);
          if (_simpleMoveCheck(move)) {
            mv[mi++] = move;
            mv[mi++] = Move.setPromotionPiece(move, Piece.whiteRook);
            mv[mi++] = Move.setPromotionPiece(move, Piece.whiteBishop);
            mv[mi++] = Move.setPromotionPiece(move, Piece.whiteKnight);
          }
        }
        if (posX > 0 && fields[pos - (BoardSize.width + 1)] & Piece.colors == Piece.black) {
          // capture left-top
          var move = Move.full(fromPos: pos, toPos: pos - (BoardSize.width + 1), capturePiece: fields[pos - (BoardSize.width + 1)], promotionPiece: Piece.whiteQueen);
          if (_simpleMoveCheck(move)) {
            mv[mi++] = move;
            mv[mi++] = Move.setPromotionPiece(move, Piece.whiteRook);
            mv[mi++] = Move.setPromotionPiece(move, Piece.whiteBishop);
            mv[mi++] = Move.setPromotionPiece(move, Piece.whiteKnight);
          }
        }
        if (posX < BoardSize.width - 1 && fields[pos - (BoardSize.width - 1)] & Piece.colors == Piece.black) {
          // capture right-top
          var move = Move.full(fromPos: pos, toPos: pos - (BoardSize.width - 1), capturePiece: fields[pos - (BoardSize.width - 1)], promotionPiece: Piece.whiteQueen);
          if (_simpleMoveCheck(move)) {
            mv[mi++] = move;
            mv[mi++] = Move.setPromotionPiece(move, Piece.whiteRook);
            mv[mi++] = Move.setPromotionPiece(move, Piece.whiteBishop);
            mv[mi++] = Move.setPromotionPiece(move, Piece.whiteKnight);
          }
        }
      } else {
        switch (field) {
          case Piece.whiteKing:
            int movePos;
            if (posX > 0) {
              movePos = pos - (BoardSize.width + 1); // left-up
              if (posY > 0 && fields[movePos] & Piece.white == Piece.none) {
                var move = Move.full(fromPos: pos, toPos: movePos, capturePiece: fields[movePos]);
                if (_simpleMoveCheck(move)) mv[mi++] = move;
              }
              movePos = pos - 1; // left
              if (fields[movePos] & Piece.white == Piece.none) {
                var move = Move.full(fromPos: pos, toPos: movePos, capturePiece: fields[movePos]);
                if (_simpleMoveCheck(move)) mv[mi++] = move;
              }
              movePos = pos + (BoardSize.width - 1); // left-down
              if (posY < BoardSize.height - 1 && fields[movePos] & Piece.white == Piece.none) {
                var move = Move.full(fromPos: pos, toPos: movePos, capturePiece: fields[movePos]);
                if (_simpleMoveCheck(move)) mv[mi++] = move;
              }
            }
            if (posX < BoardSize.width - 1) {
              movePos = pos - (BoardSize.width - 1); // right-up
              if (posY > 0 && fields[movePos] & Piece.white == Piece.none) {
                var move = Move.full(fromPos: pos, toPos: movePos, capturePiece: fields[movePos]);
                if (_simpleMoveCheck(move)) mv[mi++] = move;
              }
              movePos = pos + 1; // right
              if (fields[movePos] & Piece.white == Piece.none) {
                var move = Move.full(fromPos: pos, toPos: movePos, capturePiece: fields[movePos]);
                if (_simpleMoveCheck(move)) mv[mi++] = move;
              }
              movePos = pos + (BoardSize.width + 1); // right-down
              if (posY < BoardSize.height - 1 && fields[movePos] & Piece.white == Piece.none) {
                var move = Move.full(fromPos: pos, toPos: movePos, capturePiece: fields[movePos]);
                if (_simpleMoveCheck(move)) mv[mi++] = move;
              }
            }
            movePos = pos - BoardSize.width; // up
            if (posY > 0 && fields[movePos] & Piece.white == Piece.none) {
              var move = Move.full(fromPos: pos, toPos: movePos, capturePiece: fields[movePos]);
              if (_simpleMoveCheck(move)) mv[mi++] = move;
            }
            movePos = pos + BoardSize.width; // down
            if (posY < BoardSize.height - 1 && fields[movePos] & Piece.white == Piece.none) {
              var move = Move.full(fromPos: pos, toPos: movePos, capturePiece: fields[movePos]);
              if (_simpleMoveCheck(move)) mv[mi++] = move;
            }
            if (pos == 60) {
              if (whiteCanCastleQueenside && fields[57] == Piece.none && fields[58] == Piece.none && fields[59] == Piece.none && !isCheckedPos(58, Piece.black) && !isCheckedPos(59, Piece.black) && !isCheckedPos(60, Piece.black)) {
                mv[mi++] = Move.full(fromPos: pos, toPos: pos - 2);
              }
              if (whiteCanCastleKingside && fields[61] == Piece.none && fields[62] == Piece.none && !isCheckedPos(60, Piece.black) && !isCheckedPos(61, Piece.black) && !isCheckedPos(62, Piece.black)) {
                mv[mi++] = Move.full(fromPos: pos, toPos: pos + 2);
              }
            }
            break;

          case Piece.whiteQueen:
            // left
            for (int i = 1; i < BoardSize.width; i++) {
              if (posX - i < 0) break;
              int p = pos - i;
              int f = fields[p];
              if ((f & Piece.white) != Piece.none) break;
              var move = Move.full(fromPos: pos, toPos: p, capturePiece: f);
              if (_simpleMoveCheck(move)) mv[mi++] = move;
              if (f != Piece.none) break;
            }
            // right
            for (int i = 1; i < BoardSize.width; i++) {
              if (posX + i >= BoardSize.width) break;
              int p = pos + i;
              int f = fields[p];
              if (f & Piece.white != Piece.none) break;
              var move = Move.full(fromPos: pos, toPos: p, capturePiece: f);
              if (_simpleMoveCheck(move)) mv[mi++] = move;
              if (f != Piece.none) break;
            }
            // up
            for (int i = 1; i < BoardSize.height; i++) {
              if (posY - i < 0) break;
              int p = pos - BoardSize.width * i;
              int f = fields[p];
              if (f & Piece.white != Piece.none) break;
              var move = Move.full(fromPos: pos, toPos: p, capturePiece: f);
              if (_simpleMoveCheck(move)) mv[mi++] = move;
              if (f != Piece.none) break;
            }
            // down
            for (int i = 1; i < BoardSize.height; i++) {
              if (posY + i >= BoardSize.height) break;
              int p = pos + BoardSize.width * i;
              int f = fields[p];
              if (f & Piece.white != Piece.none) break;
              var move = Move.full(fromPos: pos, toPos: p, capturePiece: f);
              if (_simpleMoveCheck(move)) mv[mi++] = move;
              if (f != Piece.none) break;
            }
            // left-up
            for (int i = 1; i < BoardSize.width; i++) {
              if (posX - i < 0 || posY - i < 0) break;
              int p = pos - (BoardSize.width * i + i);
              int f = fields[p];
              if (f & Piece.white != Piece.none) break;
              var move = Move.full(fromPos: pos, toPos: p, capturePiece: f);
              if (_simpleMoveCheck(move)) mv[mi++] = move;
              if (f != Piece.none) break;
            }
            // left-down
            for (int i = 1; i < BoardSize.width; i++) {
              if (posX - i < 0 || posY + i >= BoardSize.height) break;
              int p = pos + (BoardSize.width * i - i);
              int f = fields[p];
              if (f & Piece.white != Piece.none) break;
              var move = Move.full(fromPos: pos, toPos: p, capturePiece: f);
              if (_simpleMoveCheck(move)) mv[mi++] = move;
              if (f != Piece.none) break;
            }
            // right-up
            for (int i = 1; i < BoardSize.width; i++) {
              if (posX + i >= BoardSize.width || posY - i < 0) break;
              int p = pos - (BoardSize.width * i - i);
              int f = fields[p];
              if (f & Piece.white != Piece.none) break;
              var move = Move.full(fromPos: pos, toPos: p, capturePiece: f);
              if (_simpleMoveCheck(move)) mv[mi++] = move;
              if (f != Piece.none) break;
            }
            // right-down
            for (int i = 1; i < BoardSize.width; i++) {
              if (posX + i >= BoardSize.width || posY + i >= BoardSize.height) break;
              int p = pos + (BoardSize.width * i + i);
              int f = fields[p];
              if (f & Piece.white != Piece.none) break;
              var move = Move.full(fromPos: pos, toPos: p, capturePiece: f);
              if (_simpleMoveCheck(move)) mv[mi++] = move;
              if (f != Piece.none) break;
            }
            break;

          case Piece.whiteRook:
            // left
            for (int i = 1; i < BoardSize.width; i++) {
              if (posX - i < 0) break;
              int p = pos - i;
              int f = fields[p];
              if ((f & Piece.white) != Piece.none) break;
              var move = Move.full(fromPos: pos, toPos: p, capturePiece: f);
              if (_simpleMoveCheck(move)) mv[mi++] = move;
              if (f != Piece.none) break;
            }
            // right
            for (int i = 1; i < BoardSize.width; i++) {
              if (posX + i >= BoardSize.width) break;
              int p = pos + i;
              int f = fields[p];
              if (f & Piece.white != Piece.none) break;
              var move = Move.full(fromPos: pos, toPos: p, capturePiece: f);
              if (_simpleMoveCheck(move)) mv[mi++] = move;
              if (f != Piece.none) break;
            }
            // up
            for (int i = 1; i < BoardSize.height; i++) {
              if (posY - i < 0) break;
              int p = pos - BoardSize.width * i;
              int f = fields[p];
              if (f & Piece.white != Piece.none) break;
              var move = Move.full(fromPos: pos, toPos: p, capturePiece: f);
              if (_simpleMoveCheck(move)) mv[mi++] = move;
              if (f != Piece.none) break;
            }
            // down
            for (int i = 1; i < BoardSize.height; i++) {
              if (posY + i >= BoardSize.height) break;
              int p = pos + BoardSize.width * i;
              int f = fields[p];
              if (f & Piece.white != Piece.none) break;
              var move = Move.full(fromPos: pos, toPos: p, capturePiece: f);
              if (_simpleMoveCheck(move)) mv[mi++] = move;
              if (f != Piece.none) break;
            }
            break;

          case Piece.whiteBishop:
            // left-up
            for (int i = 1; i < BoardSize.width; i++) {
              if (posX - i < 0 || posY - i < 0) break;
              int p = pos - (BoardSize.width * i + i);
              int f = fields[p];
              if (f & Piece.white != Piece.none) break;
              var move = Move.full(fromPos: pos, toPos: p, capturePiece: f);
              if (_simpleMoveCheck(move)) mv[mi++] = move;
              if (f != Piece.none) break;
            }
            // left-down
            for (int i = 1; i < BoardSize.width; i++) {
              if (posX - i < 0 || posY + i >= BoardSize.height) break;
              int p = pos + (BoardSize.width * i - i);
              int f = fields[p];
              if (f & Piece.white != Piece.none) break;
              var move = Move.full(fromPos: pos, toPos: p, capturePiece: f);
              if (_simpleMoveCheck(move)) mv[mi++] = move;
              if (f != Piece.none) break;
            }
            // right-up
            for (int i = 1; i < BoardSize.width; i++) {
              if (posX + i >= BoardSize.width || posY - i < 0) break;
              int p = pos - (BoardSize.width * i - i);
              int f = fields[p];
              if (f & Piece.white != Piece.none) break;
              var move = Move.full(fromPos: pos, toPos: p, capturePiece: f);
              if (_simpleMoveCheck(move)) mv[mi++] = move;
              if (f != Piece.none) break;
            }
            // right-down
            for (int i = 1; i < BoardSize.width; i++) {
              if (posX + i >= BoardSize.width || posY + i >= BoardSize.height) break;
              int p = pos + (BoardSize.width * i + i);
              int f = fields[p];
              if (f & Piece.white != Piece.none) break;
              var move = Move.full(fromPos: pos, toPos: p, capturePiece: f);
              if (_simpleMoveCheck(move)) mv[mi++] = move;
              if (f != Piece.none) break;
            }
            break;

          case Piece.whiteKnight:
            int movePos;
            if (posX > 0) {
              movePos = pos - (BoardSize.width * 2 + 1); // -1, -2
              if (posY > 1 && fields[movePos] & Piece.white == Piece.none) {
                var move = Move.full(fromPos: pos, toPos: movePos, capturePiece: fields[movePos]);
                if (_simpleMoveCheck(move)) mv[mi++] = move;
              }
              movePos = pos + (BoardSize.width * 2 - 1); // -1, +2
              if (posY < BoardSize.height - 2 && fields[movePos] & Piece.white == Piece.none) {
                var move = Move.full(fromPos: pos, toPos: movePos, capturePiece: fields[movePos]);
                if (_simpleMoveCheck(move)) mv[mi++] = move;
              }
              if (posX > 1) {
                movePos = pos - (BoardSize.width + 2); // -2, -1
                if (posY > 0 && fields[movePos] & Piece.white == Piece.none) {
                  var move = Move.full(fromPos: pos, toPos: movePos, capturePiece: fields[movePos]);
                  if (_simpleMoveCheck(move)) mv[mi++] = move;
                }
                movePos = pos + (BoardSize.width - 2); // -2, +1
                if (posY < BoardSize.height - 1 && fields[movePos] & Piece.white == Piece.none) {
                  var move = Move.full(fromPos: pos, toPos: movePos, capturePiece: fields[movePos]);
                  if (_simpleMoveCheck(move)) mv[mi++] = move;
                }
              }
            }
            if (posX < BoardSize.width - 1) {
              movePos = pos - (BoardSize.width * 2 - 1); // +1, -2
              if (posY > 1 && fields[movePos] & Piece.white == Piece.none) {
                var move = Move.full(fromPos: pos, toPos: movePos, capturePiece: fields[movePos]);
                if (_simpleMoveCheck(move)) mv[mi++] = move;
              }
              movePos = pos + (BoardSize.width * 2 + 1); // +1, +2
              if (posY < BoardSize.height - 2 && fields[movePos] & Piece.white == Piece.none) {
                var move = Move.full(fromPos: pos, toPos: movePos, capturePiece: fields[movePos]);
                if (_simpleMoveCheck(move)) mv[mi++] = move;
              }
              if (posX < BoardSize.width - 2) {
                movePos = pos - (BoardSize.width - 2); // +2, +1
                if (posY > 0 && fields[movePos] & Piece.white == Piece.none) {
                  var move = Move.full(fromPos: pos, toPos: movePos, capturePiece: fields[movePos]);
                  if (_simpleMoveCheck(move)) mv[mi++] = move;
                }
                movePos = pos + (BoardSize.width + 2); // +2, -1
                if (posY < BoardSize.height - 1 && fields[movePos] & Piece.white == Piece.none) {
                  var move = Move.full(fromPos: pos, toPos: movePos, capturePiece: fields[movePos]);
                  if (_simpleMoveCheck(move)) mv[mi++] = move;
                }
              }
            }
            break;

          case Piece.whitePawn:
            if (posY < 1 || posY >= BoardSize.height - 1) break; // invalid pos?
            int movePos;
            movePos = pos - BoardSize.width;
            if (fields[movePos] == Piece.none) {
              var move = Move.full(fromPos: pos, toPos: movePos, capturePiece: fields[movePos]);
              if (_simpleMoveCheck(move)) mv[mi++] = move;
              movePos = pos - BoardSize.width * 2;
              if (posY == BoardSize.height - 2 && fields[movePos] == Piece.none) {
                var move = Move.full(fromPos: pos, toPos: movePos, capturePiece: fields[movePos]);
                if (_simpleMoveCheck(move)) mv[mi++] = move;
              }
            }
            movePos = pos - (BoardSize.width + 1);
            if (posX > 0 && (enPassantPos == movePos || fields[movePos] & Piece.colors == Piece.black)) {
              // capture left-top
              var move = Move.full(fromPos: pos, toPos: movePos, capturePiece: fields[movePos]);
              if (_simpleMoveCheck(move)) mv[mi++] = move;
            }
            movePos = pos - (BoardSize.width - 1);
            if (posX < BoardSize.width - 1 && (enPassantPos == movePos || fields[movePos] & Piece.colors == Piece.black)) {
              // capture right-top
              var move = Move.full(fromPos: pos, toPos: movePos, capturePiece: fields[movePos]);
              if (_simpleMoveCheck(move)) mv[mi++] = move;
            }
            break;
        }
      }
    }
    return mi;
  }

  int _getBlackMoves(Int32List mv) {
    int mi = 0;
    for (int pos = 0; pos < fields.length; pos++) {
      int field = fields[pos];
      if (field & Piece.colors != Piece.black) continue; // wrong color / no p?
      int posX = pos % BoardSize.width;
      int posY = pos ~/ BoardSize.width;

      if (field == Piece.blackPawn && pos >= BoardSize.height * BoardSize.width - BoardSize.width * 2) {
        if (fields[pos + BoardSize.width] == Piece.none) {
          var move = Move.full(fromPos: pos, toPos: pos + BoardSize.width, capturePiece: fields[pos + BoardSize.width], promotionPiece: Piece.blackQueen);
          if (_simpleMoveCheck(move)) {
            mv[mi++] = move;
            mv[mi++] = Move.setPromotionPiece(move, Piece.blackRook);
            mv[mi++] = Move.setPromotionPiece(move, Piece.blackBishop);
            mv[mi++] = Move.setPromotionPiece(move, Piece.blackKnight);
          }
        }
        if (posX > 0 && fields[pos + (BoardSize.width - 1)] & Piece.colors == Piece.white) {
          // capture left-bottom
          var move = Move.full(fromPos: pos, toPos: pos + (BoardSize.width - 1), capturePiece: fields[pos + (BoardSize.width - 1)], promotionPiece: Piece.blackQueen);
          if (_simpleMoveCheck(move)) {
            mv[mi++] = move;
            mv[mi++] = Move.setPromotionPiece(move, Piece.blackRook);
            mv[mi++] = Move.setPromotionPiece(move, Piece.blackBishop);
            mv[mi++] = Move.setPromotionPiece(move, Piece.blackKnight);
          }
        }
        if (posX < BoardSize.width - 1 && fields[pos + (BoardSize.width + 1)] & Piece.colors == Piece.white) {
          // capture right-bottom
          var move = Move.full(fromPos: pos, toPos: pos + (BoardSize.width + 1), capturePiece: fields[pos + (BoardSize.width + 1)], promotionPiece: Piece.blackQueen);
          if (_simpleMoveCheck(move)) {
            mv[mi++] = move;
            mv[mi++] = Move.setPromotionPiece(move, Piece.blackRook);
            mv[mi++] = Move.setPromotionPiece(move, Piece.blackBishop);
            mv[mi++] = Move.setPromotionPiece(move, Piece.blackKnight);
          }
        }
      } else {
        switch (field) {
          case Piece.blackKing:
            int movePos;
            if (posX > 0) {
              movePos = pos - (BoardSize.width + 1); // left-up
              if (posY > 0 && fields[movePos] & Piece.black == Piece.none) {
                var move = Move.full(fromPos: pos, toPos: movePos, capturePiece: fields[movePos]);
                if (_simpleMoveCheck(move)) mv[mi++] = move;
              }
              movePos = pos - 1; // left
              if (fields[movePos] & Piece.black == Piece.none) {
                var move = Move.full(fromPos: pos, toPos: movePos, capturePiece: fields[movePos]);
                if (_simpleMoveCheck(move)) mv[mi++] = move;
              }
              movePos = pos + (BoardSize.width - 1); // left-down
              if (posY < BoardSize.height - 1 && fields[movePos] & Piece.black == Piece.none) {
                var move = Move.full(fromPos: pos, toPos: movePos, capturePiece: fields[movePos]);
                if (_simpleMoveCheck(move)) mv[mi++] = move;
              }
            }
            if (posX < BoardSize.width - 1) {
              movePos = pos - (BoardSize.width - 1); // right-up
              if (posY > 0 && fields[movePos] & Piece.black == Piece.none) {
                var move = Move.full(fromPos: pos, toPos: movePos, capturePiece: fields[movePos]);
                if (_simpleMoveCheck(move)) mv[mi++] = move;
              }
              movePos = pos + 1; // right
              if (fields[movePos] & Piece.black == Piece.none) {
                var move = Move.full(fromPos: pos, toPos: movePos, capturePiece: fields[movePos]);
                if (_simpleMoveCheck(move)) mv[mi++] = move;
              }
              movePos = pos + (BoardSize.width + 1); // right-down
              if (posY < BoardSize.height - 1 && fields[movePos] & Piece.black == Piece.none) {
                var move = Move.full(fromPos: pos, toPos: movePos, capturePiece: fields[movePos]);
                if (_simpleMoveCheck(move)) mv[mi++] = move;
              }
            }
            movePos = pos - BoardSize.width; // up
            if (posY > 0 && fields[movePos] & Piece.black == Piece.none) {
              var move = Move.full(fromPos: pos, toPos: movePos, capturePiece: fields[movePos]);
              if (_simpleMoveCheck(move)) mv[mi++] = move;
            }
            movePos = pos + BoardSize.width; // down
            if (posY < BoardSize.height - 1 && fields[movePos] & Piece.black == Piece.none) {
              var move = Move.full(fromPos: pos, toPos: movePos, capturePiece: fields[movePos]);
              if (_simpleMoveCheck(move)) mv[mi++] = move;
            }
            if (pos == 4) {
              if (blackCanCastleQueenside && fields[1] == Piece.none && fields[2] == Piece.none && fields[3] == Piece.none && !isCheckedPos(2, Piece.white) && !isCheckedPos(3, Piece.white) && !isCheckedPos(4, Piece.white)) {
                mv[mi++] = Move.full(fromPos: pos, toPos: pos - 2);
              }
              if (blackCanCastleKingside && fields[5] == Piece.none && fields[6] == Piece.none && !isCheckedPos(4, Piece.white) && !isCheckedPos(5, Piece.white) && !isCheckedPos(6, Piece.white)) {
                mv[mi++] = Move.full(fromPos: pos, toPos: pos + 2);
              }
            }
            break;

          case Piece.blackQueen:
            // left
            for (int i = 1; i < BoardSize.width; i++) {
              if (posX - i < 0) break;
              int p = pos - i;
              int f = fields[p];
              if ((f & Piece.black) != Piece.none) break;
              var move = Move.full(fromPos: pos, toPos: p, capturePiece: f);
              if (_simpleMoveCheck(move)) mv[mi++] = move;
              if (f != Piece.none) break;
            }
            // right
            for (int i = 1; i < BoardSize.width; i++) {
              if (posX + i >= BoardSize.width) break;
              int p = pos + i;
              int f = fields[p];
              if (f & Piece.black != Piece.none) break;
              var move = Move.full(fromPos: pos, toPos: p, capturePiece: f);
              if (_simpleMoveCheck(move)) mv[mi++] = move;
              if (f != Piece.none) break;
            }
            // up
            for (int i = 1; i < BoardSize.height; i++) {
              if (posY - i < 0) break;
              int p = pos - BoardSize.width * i;
              int f = fields[p];
              if (f & Piece.black != Piece.none) break;
              var move = Move.full(fromPos: pos, toPos: p, capturePiece: f);
              if (_simpleMoveCheck(move)) mv[mi++] = move;
              if (f != Piece.none) break;
            }
            // down
            for (int i = 1; i < BoardSize.height; i++) {
              if (posY + i >= BoardSize.height) break;
              int p = pos + BoardSize.width * i;
              int f = fields[p];
              if (f & Piece.black != Piece.none) break;
              var move = Move.full(fromPos: pos, toPos: p, capturePiece: f);
              if (_simpleMoveCheck(move)) mv[mi++] = move;
              if (f != Piece.none) break;
            }
            // left-up
            for (int i = 1; i < BoardSize.width; i++) {
              if (posX - i < 0 || posY - i < 0) break;
              int p = pos - (BoardSize.width * i + i);
              int f = fields[p];
              if (f & Piece.black != Piece.none) break;
              var move = Move.full(fromPos: pos, toPos: p, capturePiece: f);
              if (_simpleMoveCheck(move)) mv[mi++] = move;
              if (f != Piece.none) break;
            }
            // left-down
            for (int i = 1; i < BoardSize.width; i++) {
              if (posX - i < 0 || posY + i >= BoardSize.height) break;
              int p = pos + (BoardSize.width * i - i);
              int f = fields[p];
              if (f & Piece.black != Piece.none) break;
              var move = Move.full(fromPos: pos, toPos: p, capturePiece: f);
              if (_simpleMoveCheck(move)) mv[mi++] = move;
              if (f != Piece.none) break;
            }
            // right-up
            for (int i = 1; i < BoardSize.width; i++) {
              if (posX + i >= BoardSize.width || posY - i < 0) break;
              int p = pos - (BoardSize.width * i - i);
              int f = fields[p];
              if (f & Piece.black != Piece.none) break;
              var move = Move.full(fromPos: pos, toPos: p, capturePiece: f);
              if (_simpleMoveCheck(move)) mv[mi++] = move;
              if (f != Piece.none) break;
            }
            // right-down
            for (int i = 1; i < BoardSize.width; i++) {
              if (posX + i >= BoardSize.width || posY + i >= BoardSize.height) break;
              int p = pos + (BoardSize.width * i + i);
              int f = fields[p];
              if (f & Piece.black != Piece.none) break;
              var move = Move.full(fromPos: pos, toPos: p, capturePiece: f);
              if (_simpleMoveCheck(move)) mv[mi++] = move;
              if (f != Piece.none) break;
            }
            break;

          case Piece.blackRook:
            // left
            for (int i = 1; i < BoardSize.width; i++) {
              if (posX - i < 0) break;
              int p = pos - i;
              int f = fields[p];
              if ((f & Piece.black) != Piece.none) break;
              var move = Move.full(fromPos: pos, toPos: p, capturePiece: f);
              if (_simpleMoveCheck(move)) mv[mi++] = move;
              if (f != Piece.none) break;
            }
            // right
            for (int i = 1; i < BoardSize.width; i++) {
              if (posX + i >= BoardSize.width) break;
              int p = pos + i;
              int f = fields[p];
              if (f & Piece.black != Piece.none) break;
              var move = Move.full(fromPos: pos, toPos: p, capturePiece: f);
              if (_simpleMoveCheck(move)) mv[mi++] = move;
              if (f != Piece.none) break;
            }
            // up
            for (int i = 1; i < BoardSize.height; i++) {
              if (posY - i < 0) break;
              int p = pos - BoardSize.width * i;
              int f = fields[p];
              if (f & Piece.black != Piece.none) break;
              var move = Move.full(fromPos: pos, toPos: p, capturePiece: f);
              if (_simpleMoveCheck(move)) mv[mi++] = move;
              if (f != Piece.none) break;
            }
            // down
            for (int i = 1; i < BoardSize.height; i++) {
              if (posY + i >= BoardSize.height) break;
              int p = pos + BoardSize.width * i;
              int f = fields[p];
              if (f & Piece.black != Piece.none) break;
              var move = Move.full(fromPos: pos, toPos: p, capturePiece: f);
              if (_simpleMoveCheck(move)) mv[mi++] = move;
              if (f != Piece.none) break;
            }
            break;

          case Piece.blackBishop:
            // left-up
            for (int i = 1; i < BoardSize.width; i++) {
              if (posX - i < 0 || posY - i < 0) break;
              int p = pos - (BoardSize.width * i + i);
              int f = fields[p];
              if (f & Piece.black != Piece.none) break;
              var move = Move.full(fromPos: pos, toPos: p, capturePiece: f);
              if (_simpleMoveCheck(move)) mv[mi++] = move;
              if (f != Piece.none) break;
            }
            // left-down
            for (int i = 1; i < BoardSize.width; i++) {
              if (posX - i < 0 || posY + i >= BoardSize.height) break;
              int p = pos + (BoardSize.width * i - i);
              int f = fields[p];
              if (f & Piece.black != Piece.none) break;
              var move = Move.full(fromPos: pos, toPos: p, capturePiece: f);
              if (_simpleMoveCheck(move)) mv[mi++] = move;
              if (f != Piece.none) break;
            }
            // right-up
            for (int i = 1; i < BoardSize.width; i++) {
              if (posX + i >= BoardSize.width || posY - i < 0) break;
              int p = pos - (BoardSize.width * i - i);
              int f = fields[p];
              if (f & Piece.black != Piece.none) break;
              var move = Move.full(fromPos: pos, toPos: p, capturePiece: f);
              if (_simpleMoveCheck(move)) mv[mi++] = move;
              if (f != Piece.none) break;
            }
            // right-down
            for (int i = 1; i < BoardSize.width; i++) {
              if (posX + i >= BoardSize.width || posY + i >= BoardSize.height) break;
              int p = pos + (BoardSize.width * i + i);
              int f = fields[p];
              if (f & Piece.black != Piece.none) break;
              var move = Move.full(fromPos: pos, toPos: p, capturePiece: f);
              if (_simpleMoveCheck(move)) mv[mi++] = move;
              if (f != Piece.none) break;
            }
            break;

          case Piece.blackKnight:
            int movePos;
            if (posX > 0) {
              movePos = pos - (BoardSize.width * 2 + 1); // -1, -2
              if (posY > 1 && fields[movePos] & Piece.black == Piece.none) {
                var move = Move.full(fromPos: pos, toPos: movePos, capturePiece: fields[movePos]);
                if (_simpleMoveCheck(move)) mv[mi++] = move;
              }
              movePos = pos + (BoardSize.width * 2 - 1); // -1, +2
              if (posY < BoardSize.height - 2 && fields[movePos] & Piece.black == Piece.none) {
                var move = Move.full(fromPos: pos, toPos: movePos, capturePiece: fields[movePos]);
                if (_simpleMoveCheck(move)) mv[mi++] = move;
              }
              if (posX > 1) {
                movePos = pos - (BoardSize.width + 2); // -2, -1
                if (posY > 0 && fields[movePos] & Piece.black == Piece.none) {
                  var move = Move.full(fromPos: pos, toPos: movePos, capturePiece: fields[movePos]);
                  if (_simpleMoveCheck(move)) mv[mi++] = move;
                }
                movePos = pos + (BoardSize.width - 2); // -2, +1
                if (posY < BoardSize.height - 1 && fields[movePos] & Piece.black == Piece.none) {
                  var move = Move.full(fromPos: pos, toPos: movePos, capturePiece: fields[movePos]);
                  if (_simpleMoveCheck(move)) mv[mi++] = move;
                }
              }
            }
            if (posX < BoardSize.width - 1) {
              movePos = pos - (BoardSize.width * 2 - 1); // +1, -2
              if (posY > 1 && fields[movePos] & Piece.black == Piece.none) {
                var move = Move.full(fromPos: pos, toPos: movePos, capturePiece: fields[movePos]);
                if (_simpleMoveCheck(move)) mv[mi++] = move;
              }
              movePos = pos + (BoardSize.width * 2 + 1); // +1, +2
              if (posY < BoardSize.height - 2 && fields[movePos] & Piece.black == Piece.none) {
                var move = Move.full(fromPos: pos, toPos: movePos, capturePiece: fields[movePos]);
                if (_simpleMoveCheck(move)) mv[mi++] = move;
              }
              if (posX < BoardSize.width - 2) {
                movePos = pos - (BoardSize.width - 2); // +2, +1
                if (posY > 0 && fields[movePos] & Piece.black == Piece.none) {
                  var move = Move.full(fromPos: pos, toPos: movePos, capturePiece: fields[movePos]);
                  if (_simpleMoveCheck(move)) mv[mi++] = move;
                }
                movePos = pos + (BoardSize.width + 2); // +2, -1
                if (posY < BoardSize.height - 1 && fields[movePos] & Piece.black == Piece.none) {
                  var move = Move.full(fromPos: pos, toPos: movePos, capturePiece: fields[movePos]);
                  if (_simpleMoveCheck(move)) mv[mi++] = move;
                }
              }
            }
            break;

          case Piece.blackPawn:
            if (posY < 1 || posY >= BoardSize.height - 1) break; // invalid pos?
            int movePos;
            movePos = pos + BoardSize.width;
            if (fields[movePos] == Piece.none) {
              var move = Move.full(fromPos: pos, toPos: movePos, capturePiece: fields[movePos]);
              if (_simpleMoveCheck(move)) mv[mi++] = move;
              movePos = pos + BoardSize.width * 2;
              if (posY == 1 && fields[movePos] == Piece.none) {
                var move = Move.full(fromPos: pos, toPos: movePos, capturePiece: fields[movePos]);
                if (_simpleMoveCheck(move)) mv[mi++] = move;
              }
            }
            movePos = pos + (BoardSize.width - 1);
            if (posX > 0 && (enPassantPos == movePos || fields[movePos] & Piece.colors == Piece.white)) {
              // capture left-top
              var move = Move.full(fromPos: pos, toPos: movePos, capturePiece: fields[movePos]);
              if (_simpleMoveCheck(move)) mv[mi++] = move;
            }
            movePos = pos + (BoardSize.width + 1);
            if (posX < BoardSize.width - 1 && (enPassantPos == movePos || fields[movePos] & Piece.colors == Piece.white)) {
              // capture right-top
              var move = Move.full(fromPos: pos, toPos: movePos, capturePiece: fields[movePos]);
              if (_simpleMoveCheck(move)) mv[mi++] = move;
            }
            break;
        }
      }
    }
    return mi;
  }
}
