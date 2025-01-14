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

type Dwarf struct {
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

var DwarfXPLevel XPLevel = XPLevel{-1, 2200, 4400, 8800, 17000, 35000, 70000, 140000, 270000, 400000,
	530000, 660000, 800000, 1000000, 1200000, 1400000, 1600000, 1800000, 2000000, 2200000, 2400000, 2600000}

var DwarfSavingThrows savingthrows.ByLevel = savingthrows.ByLevel{
	// Level 1-3
	{8, 9, 10, 13, 12},
	{8, 9, 10, 13, 12},
	{8, 9, 10, 13, 12},
	// Level 4-6
	{6, 7, 8, 10, 9},
	{6, 7, 8, 10, 9},
	{6, 7, 8, 10, 9},
	// Level 7-9
	{4, 5, 6, 7, 6},
	{4, 5, 6, 7, 6},
	{4, 5, 6, 7, 6},
	// Level 10+
	{2, 3, 4, 4, 3},
	{2, 3, 4, 4, 3},
	{2, 3, 4, 4, 3},
	{2, 3, 4, 4, 3},
	{2, 3, 4, 4, 3},
	{2, 3, 4, 4, 3},
	{2, 3, 4, 4, 3},
	{2, 3, 4, 4, 3},
	{2, 3, 4, 4, 3},
	{2, 3, 4, 4, 3},
	{2, 3, 4, 4, 3},
	{2, 3, 4, 4, 3},
	{2, 3, 4, 4, 3},
}

func init() {
	ClassIndices = append(ClassIndices, "Dwarf")
	if Classes == nil {
		Classes = make(map[string]Class)
	}
	Classes["Dwarf"] = Dwarf{}
}

func (c Dwarf) Load() Class {
	fileContent, err := os.ReadFile(path.Join("data", "classes", localization.LanguageSetting, "dwarf.yaml"))
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

func (c Dwarf) Name() string {
	return c.ClassName
}

func (c Dwarf) Race() string { return c.ClassRace }

func (c Dwarf) Requirement(attr attributes.Attributes) bool {
	if attr["CON"].Value >= 9 {
		return true
	}
	return false
}

func (c Dwarf) Level(xp int) int {
	for idx := range len(DwarfXPLevel) {
		if xp < DwarfXPLevel[idx] {
			return idx
		}
		if idx >= c.MaxClassLevel {
			return c.MaxClassLevel
		}
	}
	return 0
}

func (c Dwarf) Rank(xp int) rune {
	level := c.LevelIncludingRank(xp)
	switch level {
	case 12:
		return 'C'
	case 13:
		return 'D'
	case 14:
		return 'E'
	case 15:
		return 'F'
	case 16:
		return 'G'
	case 17:
		return 'H'
	case 18:
		return 'I'
	case 19:
		return 'J'
	case 20:
		return 'K'
	case 21:
		return 'L'
	case 22:
		return 'M'
	default:
		return ' '
	}
}

func (c Dwarf) LevelIncludingRank(xp int) int {
	for idx := range len(DwarfXPLevel) {
		if xp < DwarfXPLevel[idx] {
			return idx
		}
	}
	return 0
}

func (c Dwarf) NextLevelAt(xp int) int {
	return DwarfXPLevel[c.LevelIncludingRank(xp)]
}

func (c Dwarf) CheckXPModifier(a attributes.Attributes) int {
	switch {
	case a["STR"].Value < 13:
		return 0
	case a["STR"].Value < 16:
		return 5
	case a["STR"].Value > 15:
		return 10
	}
	return 0
}

func (c Dwarf) HitDice() (dice, point int) {
	return c.ClassHD, c.ClassHP
}

func (c Dwarf) MaxLevel() int {
	return c.MaxClassLevel
}

func (c Dwarf) MaxInternalLevel() int {
	return c.MaxInternalClassLevel
}

func (c Dwarf) ArmorProficiency() string {
	return c.ClassArmor
}

func (c Dwarf) WeaponProficiency() string {
	return c.ClassWeapons
}

func (c Dwarf) SavingThrows(xp int) savingthrows.SavingThrows {
	currentLevel := c.LevelIncludingRank(xp)
	return DwarfSavingThrows[currentLevel-1]
}

func (c Dwarf) BaseMovement() int {
	return 120
}

func (c Dwarf) ThAC0(xp int) int {
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
	case currentLevel < 13:
		return 13
	case currentLevel < 21:
		return 11
	}
	return 20
}

func (c Dwarf) ThAC0Table(currentLevel int) string {
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

	var formatString string
	switch localization.OutputFormat {
	case localization.OutputFormatText:
		formatString = "10   9   8   7   6   5   4   3   2   1     0    -1   -2   -3   -4   -5   -6   -7   -8   -9  -10\n" +
			"%2d  %2d  %2d  %2d  %2d  %2d  %2d  %2d  %2d  %2d    %2d    %2d   %2d   %2d   %2d   %2d   %2d   %2d   %2d   %2d   %2d\n"
	case localization.OutputFormatObsidian:
		formatString = "" +
			"| 10  | 9   | 8   | 7   | 6   | 5   |  4  |  3  |  2  |  1  | **0** | -1  | -2  | -3  | -4  | -5  | -6  | -7  | -8  | -9  | -10 |\n" +
			"| :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: |\n" +
			"| %2d | %2d | %2d | %2d | %2d | %2d | %2d | %2d | %2d | %2d |  %2d  | %2d | %2d | %2d | %2d | %2d | %2d | %2d | %2d | %2d | %2d |\n"
	}

	return fmt.Sprintf(formatString,
		table[30], table[29], table[28], table[27], table[26], table[25], table[24], table[23], table[22], table[21],
		table[20], table[19], table[18], table[17], table[16], table[15], table[14], table[13], table[12], table[11],
		table[10])
}

func (c Dwarf) Magic() string { return "" }

func (c Dwarf) Grimoire(xp int) *magic.Spellbook {
	return nil
}

func (c Dwarf) SpellList(xp int, spellbook *magic.Spellbook) string {
	return ""
}

func (c Dwarf) SpellDescriptions(xp int, spellbook *magic.Spellbook) string {
	return ""
}

func (c Dwarf) SpecialAbilities(xp int) ClassAbilities {
	return c.Abilities
}
