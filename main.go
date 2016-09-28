package main

import (
	"fmt"

	"github.com/fatih/color"
)

func main() {
	if !color.NoColor {
		fmt.Println("verdor lib test.")
	}
}
