package main

import (
	"becmi"
	"becmi/background"
	"becmi/classes"
	"becmi/magic"
	"flag"
	"fmt"
)

func main() {
	var name, player, class, alignment, sex, bg string
	var xp int

	flag.StringVar(&name, "n", "Bargle", "character's name")
	flag.StringVar(&player, "p", "NPC", "player's name")
	flag.StringVar(&class, "c", "Cleric", "character class "+classes.AvailableClasses())
	flag.StringVar(&alignment, "a", "Lawful", "alignment [ Lawful, Neutral, Chaotic ]")
	flag.StringVar(&sex, "s", "male", "sex of the character [male, female, other]")
	flag.StringVar(&bg, "b", "Karameikos", "campaign background "+background.AvailableBackgrounds())
	flag.IntVar(&xp, "xp", 0, "experience points")
	flag.Parse()

	magic.LoadSpells("en")

	char := becmi.NewCharacter(name, player, alignment, sex, class, bg, xp)
	fmt.Printf("%s\n", char)
}
