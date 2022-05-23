window.wg = {};

const go = new Go();
WebAssembly.instantiateStreaming(fetch("../wasm/main.wasm"), go.importObject).then((result) => {
    go.run(result.instance);
});
