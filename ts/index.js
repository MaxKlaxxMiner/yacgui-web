import Chessground from "./chessground/index.js";
var wg = {};
window.wg = wg;
function getDestsFromFEN(fen) {
    var moves = wg.getMoveMapFromFEN(fen);
    wg.loglog("moves: " + moves);
    var sp = moves.split(",");
    var dests = new Map();
    for (var i = 0; i < sp.length; i += 2) {
        dests.set(sp[i], sp[i + 1].split("|"));
    }
    return dests;
}
wg.ready = function () {
    console.log("ts: ready()");
    var startFEN = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1";
    var config = {
        fen: startFEN,
        movable: {
            free: false,
            dests: getDestsFromFEN(startFEN),
            events: {}
        }
    };
    config.movable.events.after = function (orig, dest, metadata) {
        wg.loglog(orig + "-" + dest);
        var newFen = wg.doMove(config.fen, orig, dest);
        wg.loglog("fen: " + newFen);
        config.fen = newFen;
        config.movable.dests = getDestsFromFEN(config.fen);
        ground.set(config);
    };
    var ground = Chessground(document.getElementById("board"), config);
    window.ground = ground;
    // let lastFen = ground.getFen();
    // setInterval(() => {
    //     const fen = ground.getFen();
    //     if (fen !== lastFen) {
    //         lastFen = fen;
    //         wg.loglog(fen);
    //     }
    // }, 10);
    window.addEventListener("keydown", function (e) {
        if (wg.keyDown(e.code, e.defaultPrevented)) {
            e.preventDefault();
        }
    });
};
wg.loglog = function (line) {
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
};
// @ts-ignore
var go = new Go();
WebAssembly.instantiateStreaming(fetch("wasm/main.wasm"), go.importObject).then(function (result) {
    go.run(result.instance);
});
//# sourceMappingURL=index.js.map