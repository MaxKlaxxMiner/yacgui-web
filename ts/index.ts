import Chessground from "./chessground/index.js";
import {Config} from "./chessground/config";
import {Key} from "./chessground/types";

function loglog(line: string) {
    const o = document.getElementById("output");
    let html = o.innerHTML;
    html += line + "<br>";
    if (document.getElementById("board").clientHeight - 20 < o.clientHeight) {
        const i = html.indexOf("<br>");
        if (i >= 0) {
            html = html.substring(i + 4);
        }
    }
    if (document.getElementById("board").clientHeight < o.clientHeight) {
        const i = html.indexOf("<br>");
        if (i >= 0) {
            html = html.substring(i + 4);
        }
    }
    o.innerHTML = html;
}

const dests = new Map<Key, Key[]>();
dests.clear();
dests.set("e2", ["e3", "e4"]);

const config: Config = {
    fen: "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
    movable: {
        free: false,
        dests: dests
    }
};
const ground = Chessground(document.getElementById("board"), config);

let lastFen = ground.getFen();
setInterval(() => {
    const fen = ground.getFen();
    if (fen !== lastFen) {
        lastFen = fen;
        loglog(fen);
    }
}, 10);

window.addEventListener("keydown", e => {
    if (e.defaultPrevented) return
    if (e.code == "KeyF") {
        ground.toggleOrientation();
        e.preventDefault()
    }
});

let socket = new WebSocket("ws://" + location.host + "/ws");
loglog("Attempting Connection...");

socket.onopen = () => {
    loglog("Successfully Connected");
    socket.send("Hi From the Client!")
};

socket.onmessage = event => {
    loglog(event.data);
}

socket.onclose = event => {
    console.log("Socket Closed Connection: ", event);
    loglog("Socket Closed Connection");
    socket.send("Client Closed!")
};

socket.onerror = error => {
    console.log("Socket Error: ", error);
    loglog("Socket Error");
};

(<any>window).send = (x: string) => {
    socket.send(x);
};