// ignore_for_file: prefer_const_constructors, prefer_const_constructors_in_immutables

import 'dart:math';

import 'package:flutter/material.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({Key? key}) : super(key: key);

  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Yacgui',
      theme: ThemeData(primarySwatch: Colors.blue),
      debugShowCheckedModeBanner: false,
      home: HomePage(),
    );
  }
}

class HomePage extends StatefulWidget {
  HomePage({Key? key}) : super(key: key);

  @override
  State<HomePage> createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> {
  // Color cellWhite = Color(0xfff0d9b5);
  // Color cellBlack = Color(0xffb58863);

  int x = 0;
  int y = 0;

  static List<String> piecePicNames = [
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

  Map<double, List<Image>> pieceImagesCache = {};

  List<Image> getPieceImages(double cellSize) {
    var result = pieceImagesCache[cellSize];
    if (result != null) return result;
    result = List<Image>.generate(piecePicNames.length, (index) => Image.asset("images/pieces/${piecePicNames[index]}.png", width: cellSize, height: cellSize, fit: BoxFit.fill));
    pieceImagesCache[cellSize] = result;
    return result;
  }

  List<Positioned> getPieces(double cellSize, double ofsX, double ofsY) {
    var pieceImages = getPieceImages(cellSize);

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
    double cellSize = dSize.width < dSize.height ? dSize.width / 8 : (dSize.height - dPadding.top - dPadding.bottom) / 8;
    return Scaffold(
      body: SafeArea(
        child: Stack(
          children: [
            Image.asset("images/board.png", width: cellSize * 8, height: cellSize * 8, fit: BoxFit.fill, filterQuality: FilterQuality.none),
            ...getPieces(cellSize, cellSize / 2, cellSize / 2 - dPadding.top),
          ],
        ),
      ),
    );
  }
}
