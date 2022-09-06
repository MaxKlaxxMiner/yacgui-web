// ignore_for_file: prefer_const_constructors, prefer_const_constructors_in_immutables

import 'dart:math';

import 'package:flutter/material.dart';

import 'board_page.dart';

void main() {
  runApp(const MainApp());
}

class MainApp extends StatelessWidget {
  const MainApp({Key? key}) : super(key: key);

  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Yacgui',
      theme: ThemeData(primarySwatch: Colors.blue),
      debugShowCheckedModeBanner: false,
      home: BoardPage(),
    );
  }
}
