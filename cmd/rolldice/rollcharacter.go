package main

import (
	"becmi"
	"becmi/background"
	"becmi/classes"
	"becmi/localization"
	"becmi/magic"
	"flag"
	"fmt"
	"os"
)

func main() {
	var name, player, class, alignment, sex, bg string
	var xp int
	var save, obsidian bool

	language := "en"
	if lang, ok := os.LookupEnv("LANG"); ok {
		language = lang[:2]
	}

	flag.StringVar(&name, "n", "Bargle", "character's name")
	flag.StringVar(&player, "p", "NPC", "player's name")
	flag.StringVar(&class, "c", "Cleric", "character class "+classes.AvailableClasses())
	flag.StringVar(&alignment, "a", "Lawful", "alignment [ Lawful, Neutral, Chaotic ]")
	flag.StringVar(&sex, "s", "male", "sex of the character [male, female, other]")
	flag.StringVar(&bg, "b", "Karameikos", "campaign background "+background.AvailableBackgrounds())
	flag.IntVar(&xp, "xp", 0, "experience points")
	flag.StringVar(&language, "lang", language, "application language [en, de]")
	flag.BoolVar(&save, "save", false, "save character to [name].yaml file")
	flag.BoolVar(&obsidian, "obsidian", false, "save character to [name].md file for Obsidian")
	flag.Parse()

	localization.LanguageSetting = language

	//	classes.LoadClasses()
	magic.LoadSpells()

	char := becmi.NewCharacter(name, player, alignment, sex, class, bg, xp)
	fmt.Printf("%s\n", char)

	if save {
		char.Save()
	}

	if obsidian {
		char.SaveObsidian()
	}
}
