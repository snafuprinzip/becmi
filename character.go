package becmi

import (
	"becmi/attributes"
	"becmi/background"
	"becmi/classes"
	"becmi/dice"
	"becmi/localization"
	"becmi/magic"
	"becmi/proficiencies"
	"becmi/savingthrows"
	"becmi/weaponmastery"
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"os"
	"strings"
	"unicode"
)

type Character struct {
	Name         string
	Player       string
	Class        classes.Class
	Alignment    string
	XP           int
	Attributes   attributes.Attributes
	SavingThrows savingthrows.SavingThrows
	Movement     int
	ArmorClass   int
	HitPoints    int
	ThAC0        int
	Skills       proficiencies.Proficiencies
	Masteries    weaponmastery.WeaponMastery
	Sex          string
	Age          int
	AgeSpan      string
	Height       string
	Weight       int
	Background   background.Background
	Grimoire     *magic.Spellbook
}

func toUpperFirst(s string) string {
	// Convert the string to a rune slice, so that we deal with
	// runes, and not bytes.
	r := []rune(s)

	// Ensure that there is at least one rune in the slice, and
	// capitalize the first one.
	if len(r) > 0 {
		r[0] = unicode.ToUpper(r[0])
	}

	// Convert the rune slice back to a string and return.
	return string(r)
}

func NewCharacter(name, player, alignment, sex, class, campaign string, xp int) *Character {
	var ch Character
	ch.Name = name
	ch.Player = player
	ch.Alignment = toUpperFirst(alignment)
	ch.XP = xp
	ch.Sex = strings.ToLower(sex)
	ch.ArmorClass = 10
	class = toUpperFirst(class)
	campaign = toUpperFirst(campaign)

	for { // reroll until class requirements are met
		ch.Attributes = dice.RollAttributes()
		if c, ok := classes.Classes[class]; ok {
			if c.Requirement(ch.Attributes) {
				ch.Class = c.Load()
				//fmt.Println(ch.Class)
				break
			}
		}
	}

	ch.Age, ch.AgeSpan = classes.Age(ch.Class.Race())
	ch.Height, ch.Weight = classes.Body(ch.Class.Race(), ch.Sex)
	ch.Movement = ch.Class.BaseMovement()
	level := ch.Class.Level(ch.XP)
	hd, hpinc := ch.Class.HitDice()
	hp := 0
	for idx := range level {
		if idx <= 9 {
			if con, ok := ch.Attributes["CON"]; ok {
				hproll := dice.RollDice(hd) + con.Modifier()
				if hproll <= 0 {
					hproll = 1
				}
				hp += hproll
			} else {
				fmt.Fprintf(os.Stderr, "Error: No CON attribute found\n")
			}
		} else {
			hp += hpinc
		}
	}
	ch.HitPoints = hp
	ch.SavingThrows = ch.Class.SavingThrows(xp)
	ch.ThAC0 = ch.Class.ThAC0(xp)

	switch campaign {
	case "Karameikos":
		ch.Background = background.NewBGKarameikos(ch.Class.Race(), ch.Class.Name())
	}

	ch.Grimoire = ch.Class.Grimoire(ch.XP)

	return &ch
}

func (c Character) String() string {
	abilities := c.Class.SpecialAbilities(c.XP)
	var classabilities, abilitydescriptions string
	for _, ab := range abilities {
		if ab.MinLevel <= c.Class.Level(c.XP) {
			classabilities += ab.ListString() + "\n"
			abilitydescriptions += ab.DescriptionString() + "\n"
		}
	}

	outputMessage := &i18n.Message{
		ID:          "CharacterRecordSheet",
		Description: "Character record sheet",
		Other: "" +
			"Name:  %-32s \t Player:        %s\n" +
			"Class: %-32s \t Alignment:     %s\n" +
			"Level: %2d %c (max. %2d) \t XP: %8d (%s%%) \t Next Level at: %d\n" +
			"\n" +
			"%s\n" +
			"%s\n" +
			"Armor Class: %3d \t Hitpoints: %d\n" +
			"Movement:    %3d \t ThAC0:     %d\n" +
			"\n" +
			"%s\n" +
			"Armor Proficiencies:  %s\n" +
			"Weapon Proficiencies: %s\n" +
			"\n" +
			"Background\n" +
			"==========\n" +
			"Sex:    %-12s \t Age:    %4d (%s)\n" +
			"Height: %s \t\t Weight: %d cn (%d lb)\n" +
			"%s" +
			"\n" +
			"Special Abilities\n" +
			"=================\n" +
			"%s\n" +
			"%s\n" +
			"Descriptions\n" +
			"============\n" +
			"%s\n" +
			"%s\n",
	}

	translation := localization.Locale[localization.LanguageSetting].MustLocalize(&i18n.LocalizeConfig{DefaultMessage: outputMessage})

	output := fmt.Sprintf(translation,
		c.Name,
		c.Player,
		c.Class.Name(), c.Alignment,
		c.Class.Level(c.XP), c.Class.Rank(c.XP), c.Class.MaxLevel(), c.XP, attributes.SignedInt(c.Class.CheckXPModifier(c.Attributes)), c.Class.NextLevelAt(c.XP),
		c.Attributes,
		c.SavingThrows,
		c.ArmorClass, c.HitPoints,
		c.Movement, c.ThAC0,
		c.Class.ThAC0Table(c.XP),
		c.Class.ArmorProficiency(),
		c.Class.WeaponProficiency(),
		c.Sex, c.Age, c.AgeSpan,
		c.Height, c.Weight, c.Weight/10,
		c.Background,
		classabilities,
		c.Class.SpellList(c.XP, c.Grimoire),
		abilitydescriptions,
		c.Class.SpellDescriptions(c.XP, c.Grimoire),
	)
	return output
}
