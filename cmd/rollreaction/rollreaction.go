package main

import (
	"becmi/encounter"
	"becmi/localization"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var modifier int

	flag.StringVar(&localization.LanguageSetting, "lang", "en", "language [en, de]")
	flag.Parse()

	switch len(flag.Args()) {
	case 0:
		modifier = 0
	case 1:
		var err error
		modifier, err = strconv.Atoi(flag.Args()[0])
		if err != nil {
			log.Fatalf("Error: %s\n", err)
		}
	default:
		fmt.Printf("usage: %s [reaction modifier]\n", os.Args[0])
		os.Exit(1)
	}
	fmt.Println(encounter.RollReaction(modifier))
}
