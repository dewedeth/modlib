package main

import (
	"fmt"

	"github.com/dewedeth/modlib"
	"github.com/dewedeth/modlib/internal/auth"
	"github.com/dewedeth/modlib/serverlib"
)

func main() {
	fmt.Println("Running server")
	fmt.Println("Config:", modlib.Config())
	fmt.Println("Auth:", auth.GetAuth())
	fmt.Println(serverlib.Hello())
}
