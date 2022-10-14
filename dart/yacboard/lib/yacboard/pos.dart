import 'board_size.dart';

class Pos {
  static int fromXY(int x, int y) {
    if (x < 0 || x >= BoardSize.width || y < 0 || y >= BoardSize.height) return -1;
    return y * BoardSize.width + x;
  }

  static int fromChars(String c) {
    if (c.length < 2) return -1;
    int x = c.codeUnitAt(0) - 97; // 97 = 'a'
    int y = BoardSize.height + 48 - c.codeUnitAt(1); // 48 = '0'
    return fromXY(x, y);
  }

  static String asString(int pos) {
    if (pos < 0 || pos >= BoardSize.fieldCount) return "-";
    int file = (pos % BoardSize.width) + 97; // 97 = 'a'
    int rank = BoardSize.height - pos ~/ BoardSize.height + 48; // 48 = '0'
    return String.fromCharCodes([file, rank]);
  }
}
