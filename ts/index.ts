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
(<any>window).ground = ground;

let lastFen = ground.getFen();
setInterval(() => {
    const fen = ground.getFen();
    if (fen !== lastFen) {
        lastFen = fen;
        loglog(fen);
    }
}, 10);

declare function GoKeyDown(key: string, defaultPrevented: boolean): boolean;

window.addEventListener("keydown", e => {
    if (GoKeyDown(e.code, e.defaultPrevented)) {
        e.preventDefault()
    }
});

(<any>window).loglog = loglog;

// @ts-ignore
const go = new Go();
WebAssembly.instantiateStreaming(fetch("wasm/main.wasm"), go.importObject).then((result) => {
    go.run(result.instance);
});
