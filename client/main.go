package main

import (
	"client/core"
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	core.InitWg()
	core.InitCanvas()

	//core.RunLineDemo()
	//core.RunMouseDemo()

	core.MainAppStart()

	<-make(chan bool, 0)
}
