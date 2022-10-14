import 'dart:io';
import 'dart:typed_data';

import 'move.dart';
import 'move_generator.dart';
import 'yac_board.dart';

class PerftTest {
  static void testAuto(int trim) {
    // source: https://www.chessprogramming.org/Perft_Results
    testFen("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1", [3195901860, 119060324, 4865609, 197281, 8902, 400, 20].skip(trim).toList());
    testFen("r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1", [8031647685, 193690690, 4085603, 97862, 2039, 48].skip(trim).toList());
    testFen("8/2p5/3p4/KP5r/1R3p1k/8/4P1P1/8 w - - 0 1", [3009794393, 178633661, 11030083, 674624, 43238, 2812, 191, 14].skip(trim).toList());
    testFen("r3k2r/Pppp1ppp/1b3nbN/nP6/BBP1P3/q4N2/Pp1P2PP/R2Q1RK1 w kq - 0 1", [706045033, 15833292, 422333, 9467, 264, 6].skip(trim).toList());
    testFen("r2q1rk1/pP1p2pp/Q4n2/bbp1p3/Np6/1B3NBn/pPPP1PPP/R3K2R b KQ - 0 1", [706045033, 15833292, 422333, 9467, 264, 6].skip(trim).toList());
    testFen("rnbq1k1r/pp1Pbppp/2p5/8/2B5/8/PPP1NnPP/RNBQK2R w KQ - 1 8", [3048196529, 89941194, 2103487, 62379, 1486, 44].skip(trim).toList());
    testFen("r4rk1/1pp1qppp/p1np1n2/2b1p1B1/2B1P1b1/P1NP1N2/1PP1QPPP/R4RK1 w - - 0 10", [6923051137, 164075551, 3894594, 89890, 2079, 46].skip(trim).toList());
  }

  static void testFen(String fen, List<int> nodeCounter) {
    var board = YacBoard(startFen: fen);
    print("");
    print(board);

    Int32List moves = board.getMoves();
    stdout.write("\nMoves [${moves.length}]: ");
    for (int i = 0; i < moves.length; i++) {
      if (i > 0) stdout.write(", ");
      stdout.write(Move.toSan(moves[i], board));
    }
    print("");

    for (int level = 1; level <= nodeCounter.length; level++) {
      stdout.write("Level: $level / ${nodeCounter.length} Nodes: ");
      Stopwatch tim = Stopwatch()..start();
      int count = _moveCounter(board, level);
      stdout.write("$count (${tim.elapsed.inMilliseconds} ms)");
      if (count == nodeCounter[nodeCounter.length - level]) {
        print(" [ok]");
      } else {
        print(" [FAIL] $count != ${nodeCounter[nodeCounter.length - level]}");
        throw Exception("perft fail");
      }
    }
  }

  static final _tmps = [Int32List(256), Int32List(256), Int32List(256), Int32List(256), Int32List(256), Int32List(256), Int32List(256), Int32List(256), Int32List(256), Int32List(256)];

  static int _moveCounter(YacBoard board, int level) {
    var tmp = _tmps[level];
    var moveCount = board.getMovesFast(tmp);
    if (level <= 1) {
      return moveCount;
    }
    level--;
    int totalCount = 0;
    var bi = board.getBoardInfo();
    for (int m = 0; m < moveCount; m++) {
      board.doMove(tmp[m]);
      totalCount += _moveCounter(board, level);
      board.doMoveBackward(tmp[m], bi);
    }
    return totalCount;
  }
}
