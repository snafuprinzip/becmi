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

type Gnome struct {
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

var GnomeXPLevel XPLevel = XPLevel{-1, 2000, 4000, 8000, 16000, 32000, 60000, 120000, 250000, 510000,
	810000, 1110000, 1310000, 1610000, 1910000, 2210000, 2510000, 2810000, 3110000, 3410000, 3710000, 4010000}

var GnomeSavingThrows savingthrows.ByLevel = savingthrows.ByLevel{
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
	ClassIndices = append(ClassIndices, "Gnome")
	if Classes == nil {
		Classes = make(map[string]Class)
	}
	Classes["Gnome"] = Gnome{}
}

func (c Gnome) Load() Class {
	fileContent, err := os.ReadFile(path.Join("data", "classes", localization.LanguageSetting, "gnome.yaml"))
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

func (c Gnome) Name() string {
	return c.ClassName
}

func (c Gnome) Race() string { return c.ClassRace }

func (c Gnome) Requirement(attr attributes.Attributes) bool {
	if attr["DEX"].Value >= 8 && attr["CON"].Value >= 6 {
		return true
	}
	return false
}

func (c Gnome) Level(xp int) int {
	for idx := range len(GnomeXPLevel) {
		if xp < GnomeXPLevel[idx] {
			return idx - 1
		}
		if idx >= c.MaxClassLevel {
			return c.MaxClassLevel
		}
	}
	return 0
}

func (c Gnome) Rank(xp int) rune {
	return ' '
}

func (c Gnome) LevelIncludingRank(xp int) int {
	for idx := range len(GnomeXPLevel) {
		if xp < GnomeXPLevel[idx] {
			return idx
		}
	}
	return 0
}

func (c Gnome) NextLevelAt(xp int) int {
	return GnomeXPLevel[c.LevelIncludingRank(xp)]
}

func (c Gnome) CheckXPModifier(a attributes.Attributes) int {
	switch {
	case a["DEX"].Value < 13:
		return 0
	case a["DEX"].Value < 16:
		return 5
	case a["DEX"].Value > 15:
		return 10
	}
	return 0
}

func (c Gnome) HitDice() (dice, point int) {
	return c.ClassHD, c.ClassHP
}

func (c Gnome) MaxLevel() int {
	return c.MaxClassLevel
}

func (c Gnome) MaxInternalLevel() int {
	return c.MaxInternalClassLevel
}

func (c Gnome) ArmorProficiency() string {
	return c.ClassArmor
}

func (c Gnome) WeaponProficiency() string {
	return c.ClassWeapons
}

func (c Gnome) SavingThrows(xp int) savingthrows.SavingThrows {
	currentLevel := c.LevelIncludingRank(xp)
	return GnomeSavingThrows[currentLevel-1]
}

func (c Gnome) BaseMovement() int {
	return 120
}

func (c Gnome) ThAC0(xp int) int {
	currentLevel := c.Level(xp)

	switch {
	case currentLevel < 9:
		return 19 - currentLevel
	case currentLevel < 25:
		return 10 - (currentLevel/2 - 4)
	case currentLevel < 35:
		return 2
	case currentLevel >= 35:
		return 1
	}
	return 20
}

func (c Gnome) ThAC0Table(currentLevel int) string {
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

func (c Gnome) Magic() string { return "" }

func (c Gnome) Grimoire(xp int) *magic.Spellbook {
	return nil
}

func (c Gnome) SpellList(xp int, spellbook *magic.Spellbook) string {
	return ""
}

func (c Gnome) SpellDescriptions(xp int, spellbook *magic.Spellbook) string {
	return ""
}

func (c Gnome) SpellDescriptionsObsidian(xp int, spellbook *magic.Spellbook) string {
	return ""
}

func (c Gnome) SpecialAbilities(xp int) ClassAbilities {
	return c.Abilities
}
