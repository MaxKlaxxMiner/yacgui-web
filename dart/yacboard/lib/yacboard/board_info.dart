// ignore_for_file: constant_identifier_names

import 'yac_board.dart';

extension BoardInfo on YacBoard {
  static const BoardInfoNone = 0;
  static const EnPassantNone = 0xff;
  static const EnPassantMask = 0xff;
  static const EnPassantBlackA6 = 16;
  static const EnPassantBlackB6 = 17;
  static const EnPassantBlackC6 = 18;
  static const EnPassantBlackD6 = 19;
  static const EnPassantBlackE6 = 20;
  static const EnPassantBlackF6 = 21;
  static const EnPassantBlackG6 = 22;
  static const EnPassantBlackH6 = 23;
  static const EnPassantWhiteA3 = 40;
  static const EnPassantWhiteB3 = 41;
  static const EnPassantWhiteC3 = 42;
  static const EnPassantWhiteD3 = 43;
  static const EnPassantWhiteE3 = 44;
  static const EnPassantWhiteF3 = 45;
  static const EnPassantWhiteG3 = 46;
  static const EnPassantWhiteH3 = 47;
  static const WhiteCanCastleKingside = 0x0100;
  static const WhiteCanCastleQueenside = 0x0200;
  static const BlackCanCastleKingside = 0x0400;
  static const BlackCanCastleQueenside = 0x0800;
  static const HalfmoveClockMask = 0x7fff0000;

  void setBoardInfo(int boardInfo) {
    enPassantPos = boardInfo & EnPassantMask;
    whiteCanCastleKingside = (boardInfo & WhiteCanCastleKingside) != BoardInfoNone;
    whiteCanCastleQueenside = (boardInfo & WhiteCanCastleQueenside) != BoardInfoNone;
    blackCanCastleKingside = (boardInfo & BlackCanCastleKingside) != BoardInfoNone;
    blackCanCastleQueenside = (boardInfo & BlackCanCastleQueenside) != BoardInfoNone;
    halfmoveClock = (boardInfo & HalfmoveClockMask) >> 16;
  }
}
