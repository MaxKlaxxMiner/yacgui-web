import 'dart:io';

import 'package:yacboard/yacboard/is_checked.dart';

import 'move_generator.dart';
import 'piece.dart';
import 'pos.dart';
import 'yac_board.dart';

class Move {
  static int full({required int fromPos, required int toPos, int capturePiece = Piece.none, int promotionPiece = Piece.none}) {
    return (fromPos << 24) | (toPos << 16) | (capturePiece << 8) | promotionPiece;
  }

  static int fromPos(int move) {
    return move >> 24;
  }

  static int toPos(int move) {
    return (move >> 16) & 0xff;
  }

  static int capturePiece(int move) {
    return (move >> 8) & 0xff;
  }

  static int promotionPiece(int move) {
    return move & 0xff;
  }

  static int setPromotionPiece(int move, int promotionPiece) {
    return (move & 0x7fffff00) | promotionPiece;
  }

  static bool isValid(int move, {YacBoard? optionalBoard}) {
    if (fromPos(move) == toPos(move)) return false;
    if (optionalBoard == null) return true;
    return optionalBoard.moveCheck(move);
  }

  static String toSimpleString(int move) {
    if (!isValid(move)) return "-";

    var result = "${Pos.asString(fromPos(move))}-${Pos.asString(toPos(move))}";
    if (capturePiece(move) != Piece.none) {
      result = "${Pos.asString(fromPos(move))}x${Pos.asString(toPos(move))}";
    }

    if (promotionPiece(move) != Piece.none) {
      result += "->${Piece.toChar(promotionPiece(move))}";
    }

    if (capturePiece(move) != Piece.none) {
      result += " (x${Piece.toChar(capturePiece(move))})";
    }

    return result;
  }

  static String toUci(int move) {
    if (!isValid(move)) return "-";

    var result = Pos.asString(fromPos(move)) + Pos.asString(toPos(move));

    if (promotionPiece(move) != Piece.none) {
      result += Piece.toChar(promotionPiece(move)).toLowerCase();
    }

    return result;
  }

  static String toSan(int move, YacBoard board) {
    if (!isValid(move)) return "-";

    var allMoves = board.getMoves();
    var find = false;
    for (int i = 0; i < allMoves.length; i++) {
      if (allMoves[i] == move) {
        find = true;
        break;
      }
    }
    if (!find) return "-";

    var result = Pos.asString(toPos(move));

    if (capturePiece(move) != Piece.none) {
      result = "x$result";
    }

    if (Move.promotionPiece(move) != Piece.none) {
      result += "=${Piece.toChar(promotionPiece(move)).toUpperCase()}";
    }

    var piece = board.fields[fromPos(move)];
    if (piece & Piece.pawn != Piece.pawn) {
      String ext = "";
      for (int i = 0; i < allMoves.length; i++) {
        if (board.fields[fromPos(allMoves[i])] == piece && allMoves[i] != move && toPos(allMoves[i]) == toPos(move)) {
          var srcBad = Pos.asString(fromPos(allMoves[i]));
          var srcGood = Pos.asString(fromPos(move));
          if (ext == "") {
            ext = srcGood[0] != srcBad[0] ? srcGood[0] : srcGood[1];
          } else {
            ext = srcGood;
          }
        }
      }
      result = Piece.toChar(board.fields[fromPos(move)]).toUpperCase() + ext + result;
    }

    if (piece & Piece.king == Piece.king && (fromPos(move) % 8 - toPos(move) % 8).abs() == 2) {
      switch (Pos.asString(toPos(move))) {
        case "g1":
          result = "O-O";
          break;
        case "g8":
          result = "O-O";
          break;
        case "c1":
          result = "O-O-O";
          break;
        case "c8":
          result = "O-O-O";
          break;
        default:
          return "-";
      }
    }

    if (result[0] == "x") {
      result = "${Pos.asString(fromPos(move))[0]}$result";
    }

    var bi = board.getBoardInfo();
    board.doMove(move);
    if (board.isChecked()) result += board.getMoves().isEmpty ? "#" : "+";
    board.doMoveBackward(move, bi);

    return result;
  }
}
