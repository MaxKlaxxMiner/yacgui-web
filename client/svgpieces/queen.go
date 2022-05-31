package svgpieces

const queenPath = "M8 12" +
	"a2 2 0 1 1-4 0 2 2 0 1 1 4 0" +
	"z" +
	"m16.5-4.5" +
	"a2 2 0 1 1-4 0 2 2 0 1 1 4 0" +
	"z" +
	"M41 12" +
	"a2 2 0 1 1-4 0 2 2 0 1 1 4 0" +
	"z" +
	"M16 8.5" +
	"a2 2 0 1 1-4 0 2 2 0 1 1 4 0" +
	"z" +
	"M33 9" +
	"a2 2 0 1 1-4 0 2 2 0 1 1 4 0" +
	"z" +
	"M9 26" +
	"c8.5-1.5 21-1.5 27 0" +
	"l2-12-7 11" +
	"V11" +
	"l-5.5 13.5-3-15-3 15-5.5-14" +
	"V25" +
	"L7 14" +
	"l2 12" +
	"z" +
	"M9 26" +
	"c0 2 1.5 2 2.5 4 1 1.5 1 1 .5 3.5-1.5 1-1.5 2.5-1.5 2.5-1.5 1.5.5 2.5.5 2.5 6.5 1 16.5 1 23 0 0 0 1.5-1 0-2.5 0 0 .5-1.5-1-2.5-.5-2.5-.5-2 .5-3.5 1-2 2.5-2 2.5-4-8.5-1.5-18.5-1.5-27 0" +
	"z" +
	"M11.5 30" +
	"c3.5-1 18.5-1 22 0" +
	"M12 33.5" +
	"c6-1 15-1 21 0"

const queen2Path = "M11 38.5" +
	"a35 35 1 0 0 23 0"

const queen3Path = "M11 29" +
	"a35 35 1 0 1 23 0" +
	"m-21.5 2.5" +
	"h20" +
	"m-21 3" +
	"a35 35 1 0 0 22 0" +
	"m-23 3" +
	"a35 35 1 0 0 24 0" +
	"M9 26"

var queen = newPath(queenPath)
var queen2 = newPath(queen2Path)
var queen3 = newPath(queen3Path)
