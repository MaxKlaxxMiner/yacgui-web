module wasm

go 1.18

require (
	github.com/MaxKlaxxMiner/yacgui-web v0.0.0-00010101000000-000000000000
	nhooyr.io/websocket v1.8.7
)

require github.com/klauspost/compress v1.10.3 // indirect

replace github.com/MaxKlaxxMiner/yacgui-web => ../
