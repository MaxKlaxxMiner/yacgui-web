package WasmGlobal

import (
	"fmt"
	"syscall/js"
)

// WgTest wg: test()
func WgTest(_ js.Value, _ []js.Value) any {
	fmt.Println("wasm: test()")
	return nil
}
