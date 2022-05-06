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
var socket = new WebSocket("ws://" + location.host + "/ws");
loglog("Attempting Connection...");
function send(value) {
    socket.send(JSON.stringify(value));
}
var pingTicker;
socket.onopen = function () {
    loglog("Successfully Connected");
    send({ txt: "huhu" });
    pingTicker = setInterval(function () {
        send({ ping: true });
    }, 3000);
};
socket.onmessage = function (event) {
    loglog(event.data);
};
socket.onclose = function (event) {
    console.log("Socket Closed Connection: ", event);
    loglog("Socket Closed Connection");
    send({ close: true });
    clearInterval(pingTicker);
};
socket.onerror = function (error) {
    console.log("Socket Error: ", error);
    loglog("Socket Error");
    clearInterval(pingTicker);
};
window.send = function (x) {
    send({ txt: x });
};
//# sourceMappingURL=index.js.map