import Chessground from "./chessground/index.js";
import {Config} from "./chessground/config";
import {Key} from "./chessground/types";

function loglog(line: string) {
    document.getElementById("output").innerHTML += line + "<br>";
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
