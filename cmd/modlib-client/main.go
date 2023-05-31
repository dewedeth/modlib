package main

import (
	"fmt"

	"github.com/dewedeth/modlib"
	"github.com/dewedeth/modlib/clientlib"
)

func main() {
	fmt.Println("Running client")
	fmt.Println("Config:", modlib.Config())
	fmt.Println(clientlib.Hello())
}
