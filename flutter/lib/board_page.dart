import 'package:flutter/material.dart';

class BoardPage extends StatefulWidget {
  const BoardPage({Key? key}) : super(key: key);

  @override
  State<BoardPage> createState() => _BoardPageState();
}

class _BoardPageState extends State<BoardPage> {
  // Color cellWhite = Color(0xfff0d9b5);
  // Color cellBlack = Color(0xffb58863);

  int x = 0;
  int y = 0;

  static final List<String> _piecePicNames = [
    "w-pawn",
    "b-pawn",
    "w-knight",
    "b-knight",
    "w-bishop",
    "b-bishop",
    "w-rook",
    "b-rook",
    "w-queen",
    "b-queen",
    "w-king",
    "b-king",
  ];

  final Map<double, List<Image>> _pieceImagesCache = {};

  List<Image> _getPieceImages(double cellSize) {
    var result = _pieceImagesCache[cellSize];
    if (result != null) return result;
    result = List<Image>.generate(_piecePicNames.length, (index) => Image.asset("images/pieces/${_piecePicNames[index]}.png", width: cellSize, height: cellSize, fit: BoxFit.fill));
    _pieceImagesCache[cellSize] = result;
    return result;
  }

  List<Positioned> _getPieces(double cellSize, double ofsX, double ofsY) {
    var pieceImages = _getPieceImages(cellSize);

    List<Positioned> p = [];

    p.add(
      Positioned(
        left: cellSize * x,
        top: cellSize * y,
        child: Draggable(
          feedback: pieceImages[8],
          childWhenDragging: SizedBox(width: cellSize, height: cellSize),
          child: pieceImages[8],
          onDragEnd: (details) {
            x = (details.offset.dx + ofsX) ~/ cellSize;
            y = (details.offset.dy + ofsY) ~/ cellSize;
            if (x < 0) x = 0;
            if (x > 7) x = 7;
            if (y < 0) y = 0;
            if (y > 7) y = 7;
            setState(() {});
          },
        ),
      ),
    );

    return p;
  }

  @override
  Widget build(BuildContext context) {
    var dSize = MediaQuery.of(context).size;
    var dPadding = MediaQuery.of(context).viewPadding;
    double subTopHeight = dPadding.top + AppBar().preferredSize.height;
    double subHeight = subTopHeight + dPadding.bottom;
    double cellSize = dSize.width < dSize.height ? dSize.width / 8 : (dSize.height - subHeight) / 8;
    return Scaffold(
      appBar: AppBar(title: const Text("single sad queen")),
      body: SafeArea(
        child: Stack(
          children: [
            Image.asset("images/board.png", width: cellSize * 8, height: cellSize * 8, fit: BoxFit.fill, filterQuality: FilterQuality.none),
            ..._getPieces(cellSize, cellSize / 2, cellSize / 2 - subTopHeight),
          ],
        ),
      ),
    );
  }
}
