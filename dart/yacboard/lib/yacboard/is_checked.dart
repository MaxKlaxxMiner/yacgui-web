import 'board_size.dart';
import 'piece.dart';
import 'yac_board.dart';

extension MoveGenerator on YacBoard {
  bool isChecked() {
    if (whiteMove) {
      return isCheckedPos(whiteKingPos, Piece.black);
    } else {
      return isCheckedPos(blackKingPos, Piece.whiteKing);
    }
  }

  int invertedMoveColor() {
    return whiteMove ? Piece.black : Piece.white;
  }

  bool isCheckedPos(int pos, int checkerColor) {
    int posX = pos % BoardSize.width;
    int posY = pos ~/ BoardSize.width;

    // --- check pawn and king ---
    if (checkerColor == Piece.white) {
      if (posX > 0) {
        if (posY > 0 && pos - (BoardSize.width + 1) == whiteKingPos) return true;
        if (pos - 1 == whiteKingPos) return true;
        if (posY < BoardSize.height - 1 && (pos + (BoardSize.width - 1) == whiteKingPos || fields[pos + (BoardSize.width - 1)] == Piece.whitePawn)) return true;
      }
      if (posX < BoardSize.width - 1) {
        if (posY > 0 && pos - (BoardSize.width - 1) == whiteKingPos) return true;
        if (pos + 1 == whiteKingPos) return true;
        if (posY < BoardSize.height - 1 && (pos + (BoardSize.width + 1) == whiteKingPos || fields[pos + (BoardSize.width + 1)] == Piece.whitePawn)) return true;
      }
      if (posY > 0 && pos - BoardSize.width == whiteKingPos) return true;
      if (posY < BoardSize.height - 1 && pos + BoardSize.width == whiteKingPos) return true;
    } else {
      if (posX > 0) {
        if (posY > 0 && (pos - (BoardSize.width + 1) == blackKingPos || fields[pos - (BoardSize.width + 1)] == Piece.blackPawn)) return true;
        if (pos - 1 == blackKingPos) return true;
        if (posY < BoardSize.height - 1 && pos + (BoardSize.width - 1) == blackKingPos) return true;
      }
      if (posX < BoardSize.width - 1) {
        if (posY > 0 && (pos - (BoardSize.width - 1) == blackKingPos || fields[pos - (BoardSize.width - 1)] == Piece.blackPawn)) return true;
        if (pos + 1 == blackKingPos) return true;
        if (posY < BoardSize.height - 1 && pos + (BoardSize.width + 1) == blackKingPos) return true;
      }
      if (posY > 0 && pos - BoardSize.width == blackKingPos) return true;
      if (posY < BoardSize.height - 1 && pos + BoardSize.width == blackKingPos) return true;
    }

    // --- check knight ---
    {
      int knight = checkerColor | Piece.knight;
      if (posX > 0) {
        if (posY > 1 && fields[pos - (BoardSize.width * 2 + 1)] == knight) return true;
        if (posY < BoardSize.height - 2 && fields[pos + (BoardSize.width * 2 - 1)] == knight) return true;
        if (posX > 1) {
          if (posY > 0 && fields[pos - (BoardSize.width + 2)] == knight) return true;
          if (posY < BoardSize.height - 1 && fields[pos + (BoardSize.width - 2)] == knight) return true;
        }
      }
      if (posX < BoardSize.width - 1) {
        if (posY > 1 && fields[pos - (BoardSize.width * 2 - 1)] == knight) return true;
        if (posY < BoardSize.height - 2 && fields[pos + (BoardSize.width * 2 + 1)] == knight) return true;
        if (posX < BoardSize.width - 2) {
          if (posY > 0 && fields[pos - (BoardSize.width - 2)] == knight) return true;
          if (posY < BoardSize.height - 1 && fields[pos + (BoardSize.width + 2)] == knight) return true;
        }
      }
    }

    // --- check vertical and horizontal lines ---
    {
      for (int i = 1; i < BoardSize.width; i++) {
        if (posX - i < 0) break;
        int f = fields[pos - i];
        if (f == Piece.none) continue;

        if (f & (Piece.rook | Piece.queen) != Piece.none && f & checkerColor != Piece.none) return true;
        break;
      }
      for (int i = 1; i < BoardSize.width; i++) {
        if (posX + i >= BoardSize.width) break;
        int f = fields[pos + i];
        if (f == Piece.none) continue;
        if (f & (Piece.rook | Piece.queen) != Piece.none && f & checkerColor != Piece.none) return true;
        break;
      }
      for (int i = 1; i < BoardSize.height; i++) {
        if (posY - i < 0) break;
        int f = fields[pos - BoardSize.width * i];
        if (f == Piece.none) continue;
        if (f & (Piece.rook | Piece.queen) != Piece.none && f & checkerColor != Piece.none) return true;
        break;
      }
      for (int i = 1; i < BoardSize.height; i++) {
        if (posY + i >= BoardSize.height) break;
        int f = fields[pos + BoardSize.width * i];
        if (f == Piece.none) continue;
        if (f & (Piece.rook | Piece.queen) != Piece.none && f & checkerColor != Piece.none) return true;
        break;
      }
    }

    // --- check diagonal lines ---
    {
      for (int i = 1; i < BoardSize.width; i++) {
        if (posX - i < 0 || posY - i < 0) break;
        int f = fields[pos - (BoardSize.width * i + i)];
        if (f == Piece.none) continue;
        if (f & (Piece.bishop | Piece.queen) != Piece.none && f & checkerColor != Piece.none) return true;
        break;
      }
      for (int i = 1; i < BoardSize.width; i++) {
        if (posX - i < 0 || posY + i >= BoardSize.height) break;
        int f = fields[pos + (BoardSize.width * i - i)];
        if (f == Piece.none) continue;
        if (f & (Piece.bishop | Piece.queen) != Piece.none && f & checkerColor != Piece.none) return true;
        break;
      }
      for (int i = 1; i < BoardSize.width; i++) {
        if (posX + i >= BoardSize.width || posY - i < 0) break;
        int f = fields[pos - (BoardSize.width * i - i)];
        if (f == Piece.none) continue;
        if (f & (Piece.bishop | Piece.queen) != Piece.none && f & checkerColor != Piece.none) return true;
        break;
      }
      for (int i = 1; i < BoardSize.width; i++) {
        if (posX + i >= BoardSize.width || posY + i >= BoardSize.height) break;
        int f = fields[pos + (BoardSize.width * i + i)];
        if (f == Piece.none) continue;
        if (f & (Piece.bishop | Piece.queen) != Piece.none && f & checkerColor != Piece.none) return true;
        break;
      }
    }

    return false;
  }
}
