package svgpieces

const rookPath = "M9 39" +
	"h27" +
	"v-3" +
	"H9" +
	"v3" +
	"z" +
	"m3-3" +
	"v-4" +
	"h21" +
	"v4" +
	"H12" +
	"z" +
	"m-1-22" +
	"V9" +
	"h4" +
	"v2" +
	"h5" +
	"V9" +
	"h5" +
	"v2" +
	"h5" +
	"V9" +
	"h4" +
	"v5" +
	"M34 14" +
	"l-3 3" +
	"H14" +
	"l-3-3" +
	"M31 17" +
	"v12.5" +
	"H14" +
	"V17" +
	"M31 29.5" +
	"l1.5 2.5" +
	"h-20" +
	"l1.5-2.5" +
	"M11 14" +
	"h23"

const rook2Path = "M12 35.5h21m-20-4h19m-18-2h17m-17-13h17M11 14h23"

var rook = newPath(rookPath)
var rook2 = newPath(rook2Path)
