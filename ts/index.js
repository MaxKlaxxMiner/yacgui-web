import Chessground from "./chessground/index.js";
function loglog(line) {
    var o = document.getElementById("output");
    var html = o.innerHTML;
    html += line + "<br>";
    if (document.getElementById("board").clientHeight - 20 < o.clientHeight) {
        var i = html.indexOf("<br>");
        if (i >= 0) {
            html = html.substring(i + 4);
        }
    }
    if (document.getElementById("board").clientHeight < o.clientHeight) {
        var i = html.indexOf("<br>");
        if (i >= 0) {
            html = html.substring(i + 4);
        }
    }
    o.innerHTML = html;
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
window.ground = ground;
var lastFen = ground.getFen();
setInterval(function () {
    var fen = ground.getFen();
    if (fen !== lastFen) {
        lastFen = fen;
        loglog(fen);
    }
}, 10);
window.addEventListener("keydown", function (e) {
    if (GoKeyDown(e.code, e.defaultPrevented)) {
        e.preventDefault();
    }
});
window.loglog = loglog;
// @ts-ignore
var go = new Go();
WebAssembly.instantiateStreaming(fetch("wasm/main.wasm"), go.importObject).then(function (result) {
    go.run(result.instance);
});
//# sourceMappingURL=index.js.map