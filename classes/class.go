package classes

import (
	"becmi/attributes"
	"becmi/magic"
	"becmi/savingthrows"
	"strings"
)

var ClassIndices = []string{}
var Classes map[string]Class

type XPLevel [36]int

type Class interface {
	Name() string
	Requirement(attr attributes.Attributes) bool
	Level(xp int) int
	NextLevelAt(currentLevel int) int
	CheckXPModifier(a attributes.Attributes) int
	HitDice() (dice, point int)
	BaseMovement() int
	MaxLevel() int
	ArmorProficiency() string
	WeaponProficiency() string
	SavingThrows(currentLevel int) savingthrows.SavingThrows
	SpecialAbilities(currentLevel int) ClassAbilities
	ThAC0(currentLevel int) int
	ThAC0Table(currentLevel int) string
	Race() string
	Magic() string
	Grimoire(currentLevel int) *magic.Spellbook
	SpellList(currentLevel int, spellbook *magic.Spellbook) string
	SpellDescriptions(currentLevel int, spellbook *magic.Spellbook) string
}

type ClassAbility struct {
	Name        string
	MinLevel    int
	Table       string
	Description string
}

type ClassAbilities []ClassAbility

func (ca *ClassAbilities) Add(name string, minLevel int, table, description string) {
	var a ClassAbility = ClassAbility{name, minLevel, table, description}
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
	out = c.Name
	if c.Table != "" { // Ability does contain a table
		out += "\n"
		out += strings.Repeat("-", len(c.Name)) + "\n"
		out += c.Table + "\n"
	}
	return out
}

func (c ClassAbility) DescriptionString() string {
	var out string
	if c.Description != "" {
		out = c.Name + "\n"
		out += strings.Repeat("-", len(c.Name)) + "\n"
		out += c.Description + "\n"
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
