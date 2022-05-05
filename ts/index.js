import Chessground from "./chessground/index.js";
function loglog(line) {
    document.getElementById("output").innerHTML += line + "<br>";
}
var dests = new Map();
dests.clear();
dests.set("e2", ["e3", "e4"]);
var config = {
    fen: "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
    movable: {
        free: false,
        dests: dests
    }
};
var ground = Chessground(document.getElementById("board"), config);
var lastFen = ground.getFen();
setInterval(function () {
    var fen = ground.getFen();
    if (fen !== lastFen) {
        lastFen = fen;
        loglog(fen);
    }
}, 10);
window.addEventListener("keydown", function (e) {
    if (e.defaultPrevented)
        return;
    if (e.code == "KeyF") {
        ground.toggleOrientation();
        e.preventDefault();
    }
});
//# sourceMappingURL=index.js.map