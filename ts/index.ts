import Chessground from "./chessground/index.js";
import {Config} from "./chessground/config";
import * as cg from "./chessground/types";
import {Key} from "./chessground/types";

interface WasmGlobal {
    ready();

    test();

    keyDown(key: string, defaultPrevented: boolean): boolean;

    loglog(ling: string);

    getMoveMapFromFEN(fen: string): string;

    doMove(fen: string, from: Key, to: Key): string;
}

const wg: WasmGlobal = <any>{};
(<any>window).wg = wg;

function getDestsFromFEN(fen: string): cg.Dests {
    const moves = wg.getMoveMapFromFEN(fen)
    wg.loglog("moves: " + moves);
    const sp = moves.split(",");
    const dests = new Map();
    for (let i = 0; i < sp.length; i += 2) {
        dests.set(sp[i], sp[i + 1].split("|"))
    }
    return dests;
}

wg.ready = () => {
    console.log("ts: ready()");

    const startFEN = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1";
    const config: Config = {
        fen: startFEN,
        movable: {
            free: false,
            dests: getDestsFromFEN(startFEN),
            events: {}
        }
    };


    config.movable.events.after = (orig: cg.Key, dest: cg.Key, metadata: cg.MoveMetadata) => {
        wg.loglog(orig + "-" + dest);
        const newFen = wg.doMove(config.fen, orig, dest);
        wg.loglog("fen: " + newFen);
        config.fen = newFen;
        config.movable.dests = getDestsFromFEN(config.fen);
        ground.set(config)
    };

    const ground = Chessground(document.getElementById("board"), config);
    (<any>window).ground = ground;

    // let lastFen = ground.getFen();
    // setInterval(() => {
    //     const fen = ground.getFen();
    //     if (fen !== lastFen) {
    //         lastFen = fen;
    //         wg.loglog(fen);
    //     }
    // }, 10);

    window.addEventListener("keydown", e => {
        if (wg.keyDown(e.code, e.defaultPrevented)) {
            e.preventDefault()
        }
    });
};

wg.loglog = line => {
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
};

// @ts-ignore
const go = new Go();
WebAssembly.instantiateStreaming(fetch("wasm/main.wasm"), go.importObject).then((result) => {
    go.run(result.instance);
});
