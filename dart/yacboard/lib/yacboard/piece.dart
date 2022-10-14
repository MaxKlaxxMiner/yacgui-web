class Piece {
  static const none = 0x00;
  static const blocked = white | black;

  static const white = 0x40;
  static const black = 0x80;
  static const colors = white | black;

  static const king = 0x01;
  static const queen = 0x02;
  static const rook = 0x04;
  static const bishop = 0x08;
  static const knight = 0x10;
  static const pawn = 0x20;
  static const basicMask = king | queen | rook | bishop | knight | pawn;

  static const whiteKing = white | king;
  static const whiteQueen = white | queen;
  static const whiteRook = white | rook;
  static const whiteBishop = white | bishop;
  static const whiteKnight = white | knight;
  static const whitePawn = white | pawn;
  static const blackKing = black | king;
  static const blackQueen = black | queen;
  static const blackRook = black | rook;
  static const blackBishop = black | bishop;
  static const blackKnight = black | knight;
  static const blackPawn = black | pawn;

  static int fromChar(String c) {
    switch (c) {
      case "1":
      case "2":
      case "3":
      case "4":
      case "5":
      case "6":
      case "7":
      case "8":
      case "9":
        return none;
      case 'K':
        return whiteKing;
      case 'k':
        return blackKing;
      case 'Q':
        return whiteQueen;
      case 'q':
        return blackQueen;
      case 'R':
        return whiteRook;
      case 'r':
        return blackRook;
      case 'B':
        return whiteBishop;
      case 'b':
        return blackBishop;
      case 'N':
        return whiteKnight;
      case 'n':
        return blackKnight;
      case 'P':
        return whitePawn;
      case 'p':
        return blackPawn;
      default:
        return blocked;
    }
  }

  static String toChar(int piece) {
    switch (piece) {
      case whiteKing:
        return "K";
      case blackKing:
        return "k";
      case whiteQueen:
        return "Q";
      case blackQueen:
        return "q";
      case whiteRook:
        return "R";
      case blackRook:
        return "r";
      case whiteBishop:
        return "B";
      case blackBishop:
        return "b";
      case whiteKnight:
        return "N";
      case blackKnight:
        return "n";
      case whitePawn:
        return "P";
      case blackPawn:
        return "p";
      default:
        return ".";
    }
  }
}
