package main

import (
	"becmi/dice"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Errorf("Error: No dice string given as argument.")
		os.Exit(1)
	} else {
		for _, dicestr := range os.Args[1:] {
			dice.Roll(dicestr)
		}
	}
}
