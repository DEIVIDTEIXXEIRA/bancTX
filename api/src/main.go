package main

import (
	"fmt"

	"github.com/bankTX/api/src/router"
)

func main() {
	r := router.Generate()
	fmt.Printf("r: %v\n", r)
}
