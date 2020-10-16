package main

import (
	"github.com/iotaledger/wasplib/client"
)

func main() {
}

//export onLoad
func onLoadHelloWorld() {
	exports := client.NewScExports()
	exports.Add("helloWorld")
}

//export helloWorld
func helloWorld() {
	sc := client.NewScContext()
	sc.Log("Hello, world!")
}
