import 'package:yacboard/yacboard/board_info.dart';

import 'board_size.dart';
import 'piece.dart';
import 'pos.dart';
import 'yac_board.dart';

extension Fen on YacBoard {
  bool setFen(String fen) {
    clear();

    if (fen.length > BoardSize.fenMaxBytes) return false; // fen is too long

    var splits = fen.split(" ");
    if (splits.length != 6) return false; // element-count != 6

    var lines = splits[0].split("/");
    if (lines.length != BoardSize.height) return false; // invalid board height

    // --- 1 / 6 - read pieces ---
    for (int y = 0; y < lines.length; y++) {
      var line = lines[y];
      int x = 0;
      for (int i = 0; i < line.length; i++) {
        int p = Piece.fromChar(line[i]);
        if (p == Piece.blocked) return false; // unknown char
        if (p == Piece.none) {
          x += line.codeUnitAt(i) - 48; // 48 = "0"
          continue;
        }
        if (x < BoardSize.width) {
          setField(Pos.fromXY(x, y), p);
        }
        x++;
      }

      if (x != BoardSize.width) return false; // invalid width
    }

    // --- 2 / 6 - side ---
    switch (splits[1]) {
      case "w":
        whiteMove = true;
        break;
      case "b":
        whiteMove = false;
        break;
      default:
        return false; // unknown move color
    }

    // --- 3 / 6 - castling opportunities ---
    switch (splits[2]) {
      case "-":
        break;
      case "q":
        blackCanCastleQueenside = true;
        break;
      case "k":
        blackCanCastleKingside = true;
        break;
      case "kq":
        blackCanCastleKingside = true;
        blackCanCastleQueenside = true;
        break;
      case "Q":
        whiteCanCastleQueenside = true;
        break;
      case "Qq":
        whiteCanCastleQueenside = true;
        blackCanCastleQueenside = true;
        break;
      case "Qk":
        whiteCanCastleQueenside = true;
        blackCanCastleKingside = true;
        break;
      case "Qkq":
        whiteCanCastleQueenside = true;
        blackCanCastleKingside = true;
        blackCanCastleQueenside = true;
        break;
      case "K":
        whiteCanCastleKingside = true;
        break;
      case "Kq":
        whiteCanCastleKingside = true;
        blackCanCastleQueenside = true;
        break;
      case "Kk":
        whiteCanCastleKingside = true;
        blackCanCastleKingside = true;
        break;
      case "Kkq":
        whiteCanCastleKingside = true;
        blackCanCastleKingside = true;
        blackCanCastleQueenside = true;
        break;
      case "KQ":
        whiteCanCastleKingside = true;
        whiteCanCastleQueenside = true;
        break;
      case "KQq":
        whiteCanCastleKingside = true;
        whiteCanCastleQueenside = true;
        blackCanCastleQueenside = true;
        break;
      case "KQk":
        whiteCanCastleKingside = true;
        whiteCanCastleQueenside = true;
        blackCanCastleKingside = true;
        break;
      case "KQkq":
        whiteCanCastleKingside = true;
        whiteCanCastleQueenside = true;
        blackCanCastleKingside = true;
        blackCanCastleQueenside = true;
        break;
      default:
        return false; // unknown castling
    }

    // --- 4 / 6 - "en passant" ---
    enPassantPos = Pos.fromChars(splits[3]);
    if (enPassantPos > 0) {
      if (whiteMove) {
        if (enPassantPos < Pos.fromChars("a6") || enPassantPos > Pos.fromChars("h6")) {
          enPassantPos = BoardInfo.EnPassantNone;
        }
        if (enPassantPos > 0 && fields[enPassantPos + BoardSize.width] != Piece.blackPawn) {
          enPassantPos = BoardInfo.EnPassantNone;
        }
      } else {
        if (enPassantPos < Pos.fromChars("a3") || enPassantPos > Pos.fromChars("h3")) {
          enPassantPos = BoardInfo.EnPassantNone;
        }
        if (enPassantPos > 0 && fields[enPassantPos - BoardSize.width] != Piece.whitePawn) {
          enPassantPos = BoardInfo.EnPassantNone;
        }
      }
    } else {
      enPassantPos = BoardInfo.EnPassantNone;
    }

    if (enPassantPos == BoardInfo.EnPassantNone && splits[3] != "-") return false; // invalid en passant value

    // --- 5 / 6 - read halfmove clock ---
    halfmoveClock = int.tryParse(splits[4]) ?? -1;
    if (halfmoveClock < 0) return false; // invalid halfmove-counter

    // --- 6 / 6 - read move number ---
    moveNumber = int.tryParse(splits[5]) ?? -1;
    if (moveNumber < 1) return false; // invalid movenumber

    return true;
  }

  String getFen() {
    List<int> result = [];

    for (int y = 0; y < BoardSize.height; y++) {
      result.add(47); // '/'
      for (int x = 0; x < BoardSize.width; x++) {
        var c = Piece.toChar(getField(Pos.fromXY(x, y)));
        if (c == ".") {
          // > '0' && < '9'
          if (result[result.length - 1] > 48 && result[result.length - 1] < 57) {
            result[result.length - 1]++;
          } else {
            result.add(49); // '1'
          }
        } else {
          result.add(c.codeUnitAt(0));
        }
      }
    }

    if (whiteMove) {
      result.addAll(" w ".codeUnits);
    } else {
      result.addAll(" b ".codeUnits);
    }

    if (whiteCanCastleKingside) result.add(75); // 'K'
    if (whiteCanCastleQueenside) result.add(81); // 'Q'
    if (blackCanCastleKingside) result.add(107); // 'k'
    if (blackCanCastleQueenside) result.add(113); // 'q'
    if (!whiteCanCastleKingside && !whiteCanCastleQueenside && !blackCanCastleKingside && !blackCanCastleQueenside) result.add(45); // '-'

    result.add(32); // ' '
    result.addAll(Pos.asString(enPassantPos).codeUnits);
    result.add(32); // ' '
    result.addAll(halfmoveClock.toString().codeUnits);
    result.add(32); // ' '
    result.addAll(moveNumber.toString().codeUnits);

    result.removeAt(0);
    return String.fromCharCodes(result);
  }
}
