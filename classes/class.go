package classes

import (
	"becmi/attributes"
	"becmi/localization"
	"becmi/magic"
	"becmi/savingthrows"
	"strings"
)

var ClassIndices []string
var Classes map[string]Class

type XPLevel [36]int

type Class interface {
	Name() string
	Requirement(attr attributes.Attributes) bool
	Level(xp int) int
	LevelIncludingRank(xp int) int
	Rank(xp int) rune
	NextLevelAt(xp int) int
	CheckXPModifier(a attributes.Attributes) int
	HitDice() (dice, point int)
	BaseMovement() int
	MaxLevel() int
	ArmorProficiency() string
	WeaponProficiency() string
	SavingThrows(xp int) savingthrows.SavingThrows
	SpecialAbilities(xp int) ClassAbilities
	ThAC0(xp int) int
	//ThAC0Table(xp int) string
	Race() string
	Magic() string
	Grimoire(xp int) *magic.Spellbook
	SpellList(xp int, spellbook *magic.Spellbook) string
	SpellDescriptions(xp int, spellbook *magic.Spellbook) string
	Load() Class
}

type ClassAbility struct {
	ID          string
	Name        string
	MinLevel    int
	Table       string
	Description string
}

type ClassAbilities []ClassAbility

func (ca *ClassAbilities) Add(name string, minLevel int, table, description string) {
	var a ClassAbility = ClassAbility{name, name, minLevel, table, description}
	*ca = append(*ca, a)
}

func (c ClassAbility) String() string {
	var out string
	out = c.Name + "\n"
	out += strings.Repeat("-", len(c.Name)) + "\n"
	if c.Table != "" { // Ability does contain a table
		out += c.Table + "\n"
		out += strings.Repeat("-", len(c.Name)) + "\n"
	}
	out += c.Description + "\n"
	return out
}

func (c ClassAbility) ListString() string {
	var out string

	switch localization.OutputFormat {
	case localization.OutputFormatText:
		out = "- " + c.Name
		if c.Table != "" { // Ability does contain a table
			out += "\n"
			out += strings.Repeat("-", len(c.Name)) + "\n"
			out += c.Table + "\n"
		}
	case localization.OutputFormatObsidian:
		if c.Table != "" { // Ability does contain a table
			if c.Name == "Spell Slots" {
				out = c.Table + "\n"
			} else {
				out = "#### " + c.Name + "\n" + c.Table + "\n"
			}
		} else {
			out = "- " + c.Name
		}
	}
	return out
}

func (c ClassAbility) DescriptionString() string {
	var out string
	if c.Description != "" {
		switch localization.OutputFormat {
		case localization.OutputFormatText:
			out = c.Name + "\n"
			out += strings.Repeat("-", len(c.Name)) + "\n"
			out += c.Description + "\n"
		case localization.OutputFormatObsidian:
			out = "#### " + c.Name + "\n" + c.Description + "\n"
		}
	}
	return out
}

func AvailableClasses() string {
	classesstr := "["
	for idx, cl := range ClassIndices {
		classesstr += cl
		if idx != len(ClassIndices)-1 {
			classesstr += ", "
		} else {
			classesstr += "]"
		}
	}
	return classesstr
}

func containsString(slice []string, item string) bool {
	for _, a := range slice {
		if a == item {
			return true
		}
	}
	return false
}
