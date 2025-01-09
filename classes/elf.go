package classes

import (
	"becmi/attributes"
	"becmi/dice"
	"becmi/localization"
	"becmi/magic"
	"becmi/savingthrows"
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path"
)

type Elf struct {
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
type ElfSpellSlots [5]int

var ElfXPLevel XPLevel = XPLevel{-1, 4000, 8000, 16000, 32000, 64000, 120000, 250000, 400000, 600000,
	850000, 1100000, 1350000, 1600000, 1850000, 2100000, 2350000, 2600000, 2850000, 3100000}

var ElfSpellSlotsPerLevel [20]ElfSpellSlots = [20]ElfSpellSlots{
	// Level 1-3
	{1},
	{2},
	{2, 1},
	// Level 4-6
	{2, 2},
	{2, 2, 1},
	{2, 2, 2},
	// Level 7-9
	{3, 2, 2, 1},
	{3, 3, 2, 2},
	{3, 3, 3, 2, 1},
	// Level 10+
	{3, 3, 3, 3, 2},
	{3, 3, 3, 3, 2},
	{3, 3, 3, 3, 2},
	{3, 3, 3, 3, 2},
	{3, 3, 3, 3, 2},
	{3, 3, 3, 3, 2},
	{3, 3, 3, 3, 2},
	{3, 3, 3, 3, 2},
	{3, 3, 3, 3, 2},
	{3, 3, 3, 3, 2},
	{3, 3, 3, 3, 2},
}

var ElfSavingThrows savingthrows.ByLevel = savingthrows.ByLevel{
	// Level 1-3
	{12, 13, 13, 15, 15},
	{12, 13, 13, 15, 15},
	{12, 13, 13, 15, 15},
	// Level 4-6
	{8, 10, 10, 11, 11},
	{8, 10, 10, 11, 11},
	{8, 10, 10, 11, 11},
	// Level 7-9
	{4, 7, 7, 7, 7},
	{4, 7, 7, 7, 7},
	{4, 7, 7, 7, 7},
	// Level 10+
	{2, 4, 4, 3, 3},
	{2, 4, 4, 3, 3},
	{2, 4, 4, 3, 3},
	{2, 4, 4, 3, 3},
	{2, 4, 4, 3, 3},
	{2, 4, 4, 3, 3},
	{2, 4, 4, 3, 3},
	{2, 4, 4, 3, 3},
	{2, 4, 4, 3, 3},
	{2, 4, 4, 3, 3},
	{2, 4, 4, 3, 3},
}

func init() {
	ClassIndices = append(ClassIndices, "Elf")
	if Classes == nil {
		Classes = make(map[string]Class)
	}
	Classes["Elf"] = Elf{}
}

func (c Elf) Load() Class {
	fileContent, err := os.ReadFile(path.Join("data", "classes", localization.LanguageSetting, "elf.yaml"))
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

func (c Elf) Name() string {
	return c.ClassName
}

func (c Elf) Race() string { return c.ClassRace }

func (c Elf) Requirement(attr attributes.Attributes) bool {
	if attr["INT"].Value >= 9 {
		return true
	}
	return false
}

func (c Elf) Level(xp int) int {
	for idx := range len(ElfXPLevel) {
		if xp < ElfXPLevel[idx] {
			return idx
		}
		if idx >= c.MaxClassLevel {
			return c.MaxClassLevel
		}
	}
	return 0
}

func (c Elf) Rank(xp int) rune {
	level := c.LevelIncludingRank(xp)
	switch level {
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
	case 19:
		return 'L'
	case 20:
		return 'M'
	default:
		return ' '
	}
}

func (c Elf) LevelIncludingRank(xp int) int {
	for idx := range len(ElfXPLevel) {
		if xp < ElfXPLevel[idx] {
			return idx
		}
	}
	return 0
}

func (c Elf) NextLevelAt(xp int) int {
	return ElfXPLevel[c.LevelIncludingRank(xp)]
}

func (c Elf) CheckXPModifier(a attributes.Attributes) int {
	switch {
	case a["INT"].Value < 13 || a["STR"].Value < 13:
		return 0
	case a["INT"].Value < 16:
		return 5
	case a["INT"].Value > 15:
		return 10
	}
	return 0
}

func (c Elf) HitDice() (dice, point int) {
	return c.ClassHD, c.ClassHP
}

func (c Elf) MaxLevel() int {
	return c.MaxClassLevel
}

func (c Elf) MaxInternalLevel() int {
	return c.MaxInternalClassLevel
}

func (c Elf) ArmorProficiency() string {
	return c.ClassArmor
}

func (c Elf) WeaponProficiency() string {
	return c.ClassWeapons
}

func (c Elf) SavingThrows(xp int) savingthrows.SavingThrows {
	currentLevel := c.LevelIncludingRank(xp)
	return ElfSavingThrows[currentLevel-1]
}

func (c Elf) BaseMovement() int {
	return 120
}

func (c Elf) ThAC0(xp int) int {
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

func (c Elf) ThAC0Table(currentLevel int) string {
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

func (c Elf) Magic() string { return "Arcane" }

func (s ElfSpellSlots) String() string {
	return fmt.Sprintf("1st: %d,   2nd: %d,   3rd: %d,   4th: %d,   5th: %d\n", s[0], s[1], s[2], s[3], s[4])
}

func (c Elf) Grimoire(xp int) *magic.Spellbook {
	var book magic.Spellbook
	currentLevel := c.LevelIncludingRank(xp)

	book[0] = append(book[0], "Read Magic")
	for {
		roll := dice.RollDice(len(magic.AllArcaneSpells[0]))
		spell := magic.AllArcaneSpells[0][roll-1].ID
		if !containsString(book[0], spell) {
			book[0] = append(book[0], spell)
			break
		}
	}

	if currentLevel >= 2 {
		for {
			roll := dice.RollDice(len(magic.AllArcaneSpells[0]))
			spell := magic.AllArcaneSpells[0][roll-1].ID
			if !containsString(book[0], spell) {
				book[0] = append(book[0], spell)
				break
			}
		}
	}
	if currentLevel >= 3 {
		roll := dice.RollDice(len(magic.AllArcaneSpells[1]))
		book[1] = append(book[1], magic.AllArcaneSpells[1][roll-1].ID)
	}
	if currentLevel >= 4 {
		for {
			roll := dice.RollDice(len(magic.AllArcaneSpells[1]))
			spell := magic.AllArcaneSpells[1][roll-1].ID
			if !containsString(book[1], spell) {
				book[1] = append(book[1], spell)
				break
			}
		}
	}

	return &book
}

func (c Elf) SpellList(xp int, spellbook *magic.Spellbook) string {
	var spellList string

	for idx := 0; idx < 5; idx++ {
		if len(spellbook[idx]) == 0 { // Max Level of Spells in Spellbook
			break
		}
		if idx == 0 { // Print Header
			spellListHeader := &i18n.Message{
				ID:          "Spell List Header",
				Description: "Header for Spell List",
				Other: "" +
					"Spells\n" +
					"======\n\n",
			}
			translation := localization.Locale[localization.LanguageSetting].MustLocalize(&i18n.LocalizeConfig{DefaultMessage: spellListHeader})
			spellList += translation
		}

		spellLevelHeader := &i18n.Message{
			ID:          "Spell Level Header",
			Description: "Header for Spell Level",
			Other: "" +
				"Level %d\n" +
				"--------\n",
		}
		translation := localization.Locale[localization.LanguageSetting].MustLocalize(&i18n.LocalizeConfig{DefaultMessage: spellLevelHeader})
		spellList += fmt.Sprintf(translation, idx+1)

		for _, spell := range spellbook[idx] {
			for _, spelldesc := range magic.AllArcaneSpells[idx] {
				if spelldesc.ID == spell {
					spellList += spelldesc.Name + "\n"
				}
			}
		}

		spellList += "\n"
	}
	return spellList
}

func (c Elf) SpellDescriptions(xp int, spellbook *magic.Spellbook) string {
	var spells string

	for idx := 0; idx < 5; idx++ {
		if len(spellbook[idx]) == 0 { // Max Level of Spells in Spellbook
			break
		}
		if idx == 0 { // Print Header
			spellListHeader := &i18n.Message{
				ID:          "Spell Description Header",
				Description: "Header for Spell Descriptions",
				Other: "" +
					"Spelldescriptions\n" +
					"=================\n\n",
			}
			translation := localization.Locale[localization.LanguageSetting].MustLocalize(&i18n.LocalizeConfig{DefaultMessage: spellListHeader})
			spells += translation
		}
		spellLevelHeader := &i18n.Message{
			ID:          "Spell Level Header",
			Description: "Header for Spell Level",
			Other: "" +
				"Level %d\n" +
				"--------\n",
		}
		translation := localization.Locale[localization.LanguageSetting].MustLocalize(&i18n.LocalizeConfig{DefaultMessage: spellLevelHeader})
		spells += fmt.Sprintf(translation, idx+1)

		for _, spell := range spellbook[idx] {
			for _, spelldesc := range magic.AllArcaneSpells[idx] {
				if spelldesc.ID == spell {
					spells += spelldesc.String() + "\n"
				}
			}
		}
		spells += "\n"
	}
	return spells
}

func (c Elf) SpellDescriptionsObsidian(xp int, spellbook *magic.Spellbook) string {
	var spells string

	for idx := 0; idx < 9; idx++ {
		if len(spellbook[idx]) == 0 { // Max Level of Spells in Spellbook
			break
		}
		if idx == 0 { // Print Header
			spellListHeader := &i18n.Message{
				ID:          "Spell Description Obsidian Header",
				Description: "Header for Spell Descriptions for Obsidian",
				Other:       "### Spelldescriptions\n\n",
			}
			translation := localization.Locale[localization.LanguageSetting].MustLocalize(&i18n.LocalizeConfig{DefaultMessage: spellListHeader})
			spells += translation
		}

		spellLevelHeader := &i18n.Message{
			ID:          "Spell Level Obsidian Header",
			Description: "Header for Spell Level for Obsidian",
			Other:       "#### Level %d\n\n",
		}
		translation := localization.Locale[localization.LanguageSetting].MustLocalize(&i18n.LocalizeConfig{DefaultMessage: spellLevelHeader})
		spells += fmt.Sprintf(translation, idx+1)

		for _, spell := range spellbook[idx] {
			for _, spelldesc := range magic.AllArcaneSpells[idx] {
				if spelldesc.ID == spell {
					spells += spelldesc.ObsidianString() + "\n"
				}
			}
		}
		spells += "\n"
	}
	return spells
}

func (c Elf) SpecialAbilities(xp int) ClassAbilities {
	currentLevel := c.LevelIncludingRank(xp)
	spellslots := ElfSpellSlotsPerLevel[currentLevel-1]
	for ability := range c.Abilities {
		if c.Abilities[ability].Table != "" && c.Abilities[ability].ID == "Spell Slots" {
			formatstr := c.Abilities[ability].Table
			c.Abilities[ability].Table = fmt.Sprintf(formatstr, spellslots[0], spellslots[1], spellslots[2], spellslots[3], spellslots[4])
			break
		}
	}

	return c.Abilities
}
