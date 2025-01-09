package classes

import (
	"becmi/attributes"
	"becmi/localization"
	"becmi/magic"
	"becmi/savingthrows"
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path"
)

type Halfling struct {
	ID                    string         `yaml:"id"`
	ClassName             string         `yaml:"class"`
	ClassRace             string         `yaml:"race"`
	ClassHD               int            `yaml:"hitdice"`
	ClassHP               int            `yaml:"hitpoints"`
	MaxClassLevel         int            `yaml:"maxlevel"`
	MaxInternalClassLevel int            `yaml:"maxinternallevel"`
	ClassArmor            string         `yaml:"armor"`
	ClassWeapons          string         `yaml:"weapons"`
	Abilities             ClassAbilities `yaml:"abilities"`
}

var HalflingXPLevel XPLevel = XPLevel{-1, 2000, 4000, 8000, 16000, 32000, 64000, 120000, 300000, 600000,
	900000, 1200000, 1500000, 1800000, 2100000, 2400000, 2700000, 3000000}

var HalflingSavingThrows savingthrows.ByLevel = savingthrows.ByLevel{
	// Level 1-3
	{8, 9, 10, 13, 12},
	{8, 9, 10, 13, 12},
	{8, 9, 10, 13, 12},
	// Level 4-6
	{5, 6, 7, 9, 8},
	{5, 6, 7, 9, 8},
	{5, 6, 7, 9, 8},
	// Level 7-8
	{2, 3, 4, 5, 4},
	{2, 3, 4, 5, 4},
	{2, 3, 4, 5, 4},
	{2, 3, 4, 5, 4},
	{2, 3, 4, 5, 4},
	{2, 3, 4, 5, 4},
	{2, 3, 4, 5, 4},
	{2, 3, 4, 5, 4},
	{2, 3, 4, 5, 4},
	{2, 3, 4, 5, 4},
	{2, 3, 4, 5, 4},
	{2, 3, 4, 5, 4},
}

func init() {
	ClassIndices = append(ClassIndices, "Halfling")
	if Classes == nil {
		Classes = make(map[string]Class)
	}
	Classes["Halfling"] = Halfling{}
}

func (c Halfling) Load() Class {
	fileContent, err := os.ReadFile(path.Join("data", "classes", localization.LanguageSetting, "halfling.yaml"))
	if err != nil {
		log.Fatal("Error reading file:", err)
	}

	err = yaml.Unmarshal(fileContent, &c)
	if err != nil {
		log.Fatal("Error unmarshalling YAML:", err)
	}

	//fmt.Println(c)
	return c
}

func (c Halfling) Name() string {
	return c.ClassName
}

func (c Halfling) Race() string { return c.ClassRace }

func (c Halfling) Requirement(attr attributes.Attributes) bool {
	if attr["STR"].Value >= 9 && attr["DEX"].Value >= 9 {
		return true
	}
	return false
}

func (c Halfling) Level(xp int) int {
	for idx := range len(HalflingXPLevel) {
		if xp < HalflingXPLevel[idx] {
			return idx
		}
		if idx >= c.MaxClassLevel {
			return c.MaxClassLevel
		}
	}
	return 0
}

func (c Halfling) Rank(xp int) rune {
	level := c.LevelIncludingRank(xp)
	switch level {
	case 8:
		return 'A'
	case 9:
		return 'B'
	case 10:
		return 'C'
	case 11:
		return 'D'
	case 12:
		return 'E'
	case 13:
		return 'F'
	case 14:
		return 'G'
	case 15:
		return 'H'
	case 16:
		return 'I'
	case 17:
		return 'J'
	case 18:
		return 'K'
	default:
		return ' '
	}
}

func (c Halfling) LevelIncludingRank(xp int) int {
	for idx := range len(HalflingXPLevel) {
		if xp < HalflingXPLevel[idx] {
			return idx
		}
	}
	return 0
}

func (c Halfling) NextLevelAt(xp int) int {
	return HalflingXPLevel[c.LevelIncludingRank(xp)]
}

func (c Halfling) CheckXPModifier(a attributes.Attributes) int {
	switch {
	case a["STR"].Value < 13 && a["DEX"].Value < 13:
		return 0 // both are smaller than 13
	case a["STR"].Value >= 13 && a["DEX"].Value >= 13:
		return 10 // both are higher than or equals 13
	case (a["STR"].Value >= 13) != (a["DEX"].Value >= 13): // only one of the two is 13 or higher
		return 5
	}
	return 0
}

func (c Halfling) HitDice() (dice, point int) {
	return c.ClassHD, c.ClassHP
}

func (c Halfling) MaxLevel() int {
	return c.MaxClassLevel
}

func (c Halfling) MaxInternalLevel() int {
	return c.MaxInternalClassLevel
}

func (c Halfling) ArmorProficiency() string {
	return c.ClassArmor
}

func (c Halfling) WeaponProficiency() string {
	return c.ClassWeapons
}

func (c Halfling) SavingThrows(xp int) savingthrows.SavingThrows {
	currentLevel := c.LevelIncludingRank(xp)
	return HalflingSavingThrows[currentLevel-1]
}

func (c Halfling) BaseMovement() int {
	return 120
}

func (c Halfling) ThAC0(xp int) int {
	currentLevel := c.Level(xp)
	currentRank := c.Rank(xp)

	switch {
	case currentRank == 'A':
		return 15
	case currentRank == 'B':
		return 14
	case currentRank == 'C':
		return 13
	case currentRank == 'D':
		return 12
	case currentRank == 'E':
		return 11
	case currentRank == 'F':
		return 10
	case currentRank == 'G':
		return 9
	case currentRank == 'H':
		return 8
	case currentRank == 'I':
		return 7
	case currentRank == 'J':
		return 6
	case currentRank == 'K':
		return 5
	case currentRank == 'L':
		return 4
	case currentRank == 'M':
		return 3
	case currentLevel < 4:
		return 19
	case currentLevel < 7:
		return 17
	case currentLevel < 10:
		return 15
	}
	return 20
}

func (c Halfling) ThAC0Table(currentLevel int) string {
	thac0 := c.ThAC0(currentLevel)

	var table [40]int // -20 to 19 == offset 20
	table[20] = thac0
	roll := thac0 - 1
	cnt := 0
	for i := 21; i < len(table); i++ {
		table[i] = roll
		if roll != 20 && roll != 30 {
			roll--
		} else if cnt == 4 {
			cnt = 0
			roll--
			continue
		} else {
			cnt++
			continue
		}
	}
	roll = thac0 + 1
	for i := 19; i >= 0; i-- {
		table[i] = roll
		if roll != 20 && roll != 30 {
			roll++
		} else if cnt == 4 {
			cnt = 0
			roll++
			continue
		} else {
			cnt++
			continue
		}
	}

	return fmt.Sprintf(""+
		"10   9   8   7   6   5   4   3   2   1     0    -1   -2   -3   -4   -5   -6   -7   -8   -9  -10\n"+
		"%2d  %2d  %2d  %2d  %2d  %2d  %2d  %2d  %2d  %2d    %2d    %2d   %2d   %2d   %2d   %2d   %2d   %2d   %2d   %2d   %2d\n",
		table[30], table[29], table[28], table[27], table[26], table[25], table[24], table[23], table[22], table[21],
		table[20], table[19], table[18], table[17], table[16], table[15], table[14], table[13], table[12], table[11],
		table[10])
}

func (c Halfling) Magic() string { return "" }

func (c Halfling) Grimoire(xp int) *magic.Spellbook {
	return nil
}

func (c Halfling) SpellList(xp int, spellbook *magic.Spellbook) string {
	return ""
}

func (c Halfling) SpellDescriptions(xp int, spellbook *magic.Spellbook) string {
	return ""
}

func (c Halfling) SpellDescriptionsObsidian(xp int, spellbook *magic.Spellbook) string {
	return ""
}

func (c Halfling) SpecialAbilities(xp int) ClassAbilities {
	return c.Abilities
}
