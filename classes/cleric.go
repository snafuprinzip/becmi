package classes

import (
	"becmi/attributes"
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

type Cleric struct {
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
type ClericSpellSlots [7]int
type TurnUndeadAbilities [14]string

var ClericXPLevel XPLevel = XPLevel{-1, 1500, 3000, 6000, 12000, 25000, 50000, 100000, 200000, 300000, 400000, 500000, 600000, 700000,
	800000, 900000, 1000000, 1100000, 1200000, 1300000, 1400000, 1500000, 1600000, 1700000, 1800000, 1900000,
	2000000, 2100000, 2200000, 2300000, 2400000, 2500000, 2600000, 2700000, 2800000, 2900000}
var ClericSpellSlotsPerLevel [36]ClericSpellSlots = [36]ClericSpellSlots{
	// Level 1-3
	{0, 0, 0, 0, 0, 0, 0},
	{1, 0, 0, 0, 0, 0, 0},
	{2, 0, 0, 0, 0, 0, 0},
	// Level 4-6
	{2, 1, 0, 0, 0, 0, 0},
	{2, 2, 0, 0, 0, 0, 0},
	{2, 2, 1, 0, 0, 0, 0},
	// Level 7-9
	{3, 2, 2, 0, 0, 0, 0},
	{3, 3, 2, 1, 0, 0, 0},
	{3, 3, 3, 2, 0, 0, 0},
	// Level 10-12
	{4, 4, 3, 2, 1, 0, 0},
	{4, 4, 3, 3, 2, 0, 0},
	{4, 4, 4, 3, 2, 1, 0},
	// Level 13-15
	{5, 5, 4, 3, 2, 2, 0},
	{5, 5, 5, 3, 3, 2, 0},
	{6, 5, 5, 3, 3, 3, 0},
	// Level 16-18
	{6, 5, 5, 4, 4, 3, 0},
	{6, 6, 5, 4, 4, 3, 1},
	{6, 6, 5, 4, 4, 3, 2},
	// Level 19-21
	{7, 6, 5, 4, 4, 4, 2},
	{7, 6, 5, 4, 4, 4, 3},
	{7, 6, 5, 5, 5, 4, 3},
	// Level 22-24
	{7, 6, 5, 5, 5, 4, 4},
	{7, 7, 6, 6, 5, 4, 4},
	{8, 7, 6, 6, 5, 5, 4},
	// Level 25-27
	{8, 7, 6, 6, 5, 5, 5},
	{8, 7, 7, 6, 6, 5, 5},
	{8, 8, 7, 6, 6, 6, 5},
	// Level 28-30
	{8, 8, 7, 7, 7, 6, 5},
	{8, 8, 7, 7, 7, 6, 6},
	{8, 8, 8, 7, 7, 7, 6},
	// Level 31-33
	{8, 8, 8, 8, 8, 7, 6},
	{9, 8, 8, 8, 8, 7, 7},
	{9, 9, 8, 8, 8, 8, 7},
	// Level 34-36
	{9, 9, 9, 8, 8, 8, 8},
	{9, 9, 9, 9, 9, 8, 8},
	{9, 9, 9, 9, 9, 9, 9},
}
var TurnUndeadAbilitiesPerLevel [36]TurnUndeadAbilities = [36]TurnUndeadAbilities{
	// Level 1-8
	{"7", "9", "11"},
	{"T", "7", "9", "11"},
	{"T", "T", "7", "9", "11"},
	{"D", "T", "T", "7", "9", "11"},
	{"D", "D", "T", "T", "7", "9", "11"},
	{"D", "D", "D", "T", "T", "7", "9", "11"},
	{"D", "D", "D", "D", "T", "T", "7", "9", "11"},
	{"D", "D", "D", "D", "D", "T", "T", "7", "9", "11"},
	// Level 9-10
	{"D", "D", "D", "D", "D", "D", "T", "T", "7", "9", "11"},
	{"D", "D", "D", "D", "D", "D", "T", "T", "7", "9", "11"},
	// Level 11-12
	{"D+", "D", "D", "D", "D", "D", "D", "T", "T", "7", "9", "11"},
	{"D+", "D", "D", "D", "D", "D", "D", "T", "T", "7", "9", "11"},
	// Level 13-14
	{"D+", "D+", "D", "D", "D", "D", "D", "D", "T", "T", "7", "9", "11"},
	{"D+", "D+", "D", "D", "D", "D", "D", "D", "T", "T", "7", "9", "11"},
	// Level 15-16
	{"D+", "D+", "D+", "D", "D", "D", "D", "D", "D", "T", "T", "7", "9", "11"},
	{"D+", "D+", "D+", "D", "D", "D", "D", "D", "D", "T", "T", "7", "9", "11"},
	// Level 17-20
	{"D+", "D+", "D+", "D+", "D", "D", "D", "D", "D", "D", "T", "T", "7", "9"},
	{"D+", "D+", "D+", "D+", "D", "D", "D", "D", "D", "D", "T", "T", "7", "9"},
	{"D+", "D+", "D+", "D+", "D", "D", "D", "D", "D", "D", "T", "T", "7", "9"},
	{"D+", "D+", "D+", "D+", "D", "D", "D", "D", "D", "D", "T", "T", "7", "9"},
	// Level 21-24
	{"D+", "D+", "D+", "D+", "D+", "D", "D", "D", "D", "D", "D", "T", "T", "7"},
	{"D+", "D+", "D+", "D+", "D+", "D", "D", "D", "D", "D", "D", "T", "T", "7"},
	{"D+", "D+", "D+", "D+", "D+", "D", "D", "D", "D", "D", "D", "T", "T", "7"},
	{"D+", "D+", "D+", "D+", "D+", "D", "D", "D", "D", "D", "D", "T", "T", "7"},
	// Level 25-28
	{"D#", "D+", "D+", "D+", "D+", "D", "D", "D", "D", "D", "D", "T", "T", "7"},
	{"D#", "D+", "D+", "D+", "D+", "D", "D", "D", "D", "D", "D", "T", "T", "7"},
	{"D#", "D+", "D+", "D+", "D+", "D", "D", "D", "D", "D", "D", "T", "T", "7"},
	// Level 29-32
	{"D#", "D#", "D+", "D+", "D+", "D+", "D+", "D", "D", "D", "D", "D", "T", "T"},
	{"D#", "D#", "D+", "D+", "D+", "D+", "D+", "D", "D", "D", "D", "D", "T", "T"},
	{"D#", "D#", "D+", "D+", "D+", "D+", "D+", "D", "D", "D", "D", "D", "T", "T"},
	// Level 33-36
	{"D#", "D#", "D#", "D+", "D+", "D+", "D+", "D+", "D", "D", "D", "D", "T", "T"},
	{"D#", "D#", "D#", "D+", "D+", "D+", "D+", "D+", "D", "D", "D", "D", "T", "T"},
	{"D#", "D#", "D#", "D+", "D+", "D+", "D+", "D+", "D", "D", "D", "D", "T", "T"},
}
var ClericSavingThrows savingthrows.ByLevel = savingthrows.ByLevel{
	// Level 1-4
	{11, 12, 14, 16, 15},
	{11, 12, 14, 16, 15},
	{11, 12, 14, 16, 15},
	{11, 12, 14, 16, 15},
	// Level 5-8
	{9, 10, 12, 14, 13},
	{9, 10, 12, 14, 13},
	{9, 10, 12, 14, 13},
	{9, 10, 12, 14, 13},
	// Level 9-12
	{7, 8, 10, 12, 11},
	{7, 8, 10, 12, 11},
	{7, 8, 10, 12, 11},
	{7, 8, 10, 12, 11},
	// Level 13-16
	{6, 7, 8, 10, 9},
	{6, 7, 8, 10, 9},
	{6, 7, 8, 10, 9},
	{6, 7, 8, 10, 9},
	// Level 17-20
	{5, 6, 6, 8, 7},
	{5, 6, 6, 8, 7},
	{5, 6, 6, 8, 7},
	{5, 6, 6, 8, 7},
	// Level 21-24
	{4, 5, 5, 6, 5},
	{4, 5, 5, 6, 5},
	{4, 5, 5, 6, 5},
	{4, 5, 5, 6, 5},
	// Level 25-28
	{3, 4, 4, 4, 4},
	{3, 4, 4, 4, 4},
	{3, 4, 4, 4, 4},
	{3, 4, 4, 4, 4},
	// Level 29-32
	{2, 3, 3, 3, 3},
	{2, 3, 3, 3, 3},
	{2, 3, 3, 3, 3},
	{2, 3, 3, 3, 3},
	// Level 33-36
	{2, 2, 2, 2, 2},
	{2, 2, 2, 2, 2},
	{2, 2, 2, 2, 2},
	{2, 2, 2, 2, 2},
}

func init() {
	ClassIndices = append(ClassIndices, "Cleric")
	if Classes == nil {
		Classes = make(map[string]Class)
	}
	Classes["Cleric"] = Cleric{}
}

func (c Cleric) Load() Class {
	fileContent, err := os.ReadFile(path.Join("data", "classes", localization.LanguageSetting, "cleric.yaml"))
	if err != nil {
		log.Fatal("Error reading file:", err)
	}

	err = yaml.Unmarshal(fileContent, &c)
	if err != nil {
		log.Fatal("Error unmarshalling YAML:", err)
	}
	return c
}

func (c Cleric) Name() string {
	return c.ClassName
}

func (c Cleric) Race() string { return c.ClassRace }

func (c Cleric) Requirement(attr attributes.Attributes) bool {
	return true
}

func (c Cleric) Level(xp int) int {
	for idx := range len(ClericXPLevel) {
		if xp < ClericXPLevel[idx] {
			return idx
		}
	}
	return 0
}

func (c Cleric) Rank(xp int) rune {
	return ' '
}

func (c Cleric) LevelIncludingRank(xp int) int {
	return c.Level(xp)
}

func (c Cleric) NextLevelAt(xp int) int {
	currentLevel := c.Level(xp)
	return ClericXPLevel[currentLevel]
}

func (c Cleric) CheckXPModifier(a attributes.Attributes) int {
	switch {
	case a["WIS"].Value < 6:
		return -20
	case a["WIS"].Value < 9:
		return -10
	case a["WIS"].Value < 13:
		return 0
	case a["WIS"].Value < 16:
		return 5
	case a["WIS"].Value > 17:
		return 10
	}
	return 0
}

func (c Cleric) HitDice() (dice, point int) {
	return c.ClassHD, c.ClassHP
}

func (c Cleric) MaxLevel() int {
	return c.MaxClassLevel
}

func (c Cleric) MaxInternalLevel() int {
	return c.MaxInternalClassLevel
}

func (c Cleric) ArmorProficiency() string {
	return c.ClassArmor
}

func (c Cleric) WeaponProficiency() string {
	return c.ClassWeapons
}

func (c Cleric) SavingThrows(xp int) savingthrows.SavingThrows {
	currentLevel := c.LevelIncludingRank(xp)
	return ClericSavingThrows[currentLevel-1]
}

func (c Cleric) BaseMovement() int {
	return 120
}

func (c Cleric) ThAC0(xp int) int {
	currentLevel := c.Level(xp)

	switch {
	case currentLevel < 5:
		return 19
	case currentLevel < 9:
		return 17
	case currentLevel < 13:
		return 15
	case currentLevel < 17:
		return 13
	case currentLevel < 21:
		return 11
	case currentLevel < 25:
		return 9
	case currentLevel < 29:
		return 7
	case currentLevel < 33:
		return 5
	case currentLevel < 36:
		return 3
	case currentLevel < 37:
		return 2
	}
	return 20
}

func (c Cleric) ThAC0Table(xp int) string {
	currentLevel := c.Level(xp)
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

func (c Cleric) Magic() string                    { return "Divine" }
func (c Cleric) Grimoire(xp int) *magic.Spellbook { return nil }

func (s ClericSpellSlots) String() string {
	var slots *i18n.Message

	switch localization.OutputFormat {
	case localization.OutputFormatText:
		slots = &i18n.Message{
			ID:          "Spell Slots Cleric",
			Description: "Spell Slots for Cleric",
			Other: "" +
				"Spell Slots\n" +
				"-----------\n" +
				"1st: %2d\n2nd: %2d\n3rd: %2d\n4th: %2d\n5th: %2d\n6th: %2d\n7th: %2d\n",
		}
	case localization.OutputFormatObsidian:
		slots = &i18n.Message{
			ID:          "Spell Slots Cleric (Obsidian)",
			Description: "Spell Slots for Cleric for Obsidian",
			Other: "" +
				"> [!infobox]\n" +
				"> ### Spell Slots per Level\n" +
				">| Level    | \\#Spells |\n" +
				">| --- | ---: |\n" +
				">| 1st | %2d |\n" +
				">| 2nd | %2d |\n" +
				">| 3rd | %2d |\n" +
				">| 4th | %2d |\n" +
				">| 5th | %2d |\n" +
				">| 6th | %2d |\n" +
				">| 7th | %2d |\n",
		}
	}
	translation := localization.Locale[localization.LanguageSetting].MustLocalize(&i18n.LocalizeConfig{DefaultMessage: slots})

	return fmt.Sprintf("\n"+translation, s[0], s[1], s[2], s[3], s[4], s[5], s[6])
}

func (c Cleric) SpellList(xp int, spellbook *magic.Spellbook) string {
	var spellList string
	currentLevel := c.Level(xp)

	slots := ClericSpellSlotsPerLevel[currentLevel-1]
	for idx := 0; idx < 7; idx++ {
		if slots[idx] == 0 { // Max Level of available Spells
			break
		}
		switch localization.OutputFormat {
		case localization.OutputFormatText:
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
		case localization.OutputFormatObsidian:
			if idx == 0 { // Print Header
				spellListHeader := &i18n.Message{
					ID:          "Spell List Header (Obsidian)",
					Description: "Header for Spell List for Obsidian",
					Other: "" +
						"### Spells\n\n",
				}
				translation := localization.Locale[localization.LanguageSetting].MustLocalize(&i18n.LocalizeConfig{DefaultMessage: spellListHeader})
				spellList += translation
			}

			spellLevelHeader := &i18n.Message{
				ID:          "Spell Level Header (Obsidian)",
				Description: "Header for Spell Level for Obsidian",
				Other: "" +
					"#### Level %d\n",
			}
			translation := localization.Locale[localization.LanguageSetting].MustLocalize(&i18n.LocalizeConfig{DefaultMessage: spellLevelHeader})
			spellList += fmt.Sprintf(translation, idx+1)

		}

		for _, spell := range magic.AllDivineSpells[idx] {
			spellList += "- " + spell.Name + "\n"
		}
		spellList += "\n"
	}
	return spellList
}

func (c Cleric) SpellDescriptions(xp int, spellbook *magic.Spellbook) string {
	var spells string
	currentLevel := c.Level(xp)

	slots := ClericSpellSlotsPerLevel[currentLevel-1]
	for idx := 0; idx < 7; idx++ {
		if slots[idx] == 0 {
			break
		}
		switch localization.OutputFormat {
		case localization.OutputFormatText:
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
		case localization.OutputFormatObsidian:
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
		}

		for _, spell := range magic.AllDivineSpells[idx] {
			spells += spell.String() + "\n"
		}
		spells += "\n"
	}
	return spells
}

func (a TurnUndeadAbilities) String() string {
	var turnUndead, turnUndeadLegend *i18n.Message

	switch localization.OutputFormat {
	case localization.OutputFormatText:
		turnUndead = &i18n.Message{
			ID:          "Turn Undead",
			Description: "Turn Undead Table",
			Other: "" +
				"Skeleton   \t %2s\n" +
				"Zombie     \t %2s\n" +
				"Ghoul      \t %2s\n" +
				"Wight      \t %2s\n" +
				"Wraith     \t %2s\n" +
				"Mummy      \t %2s\n" +
				"Spectre    \t %2s\n" +
				"Vampire    \t %2s\n" +
				"Phantom    \t %2s\n" +
				"Haunt      \t %2s\n" +
				"Spirit     \t %2s\n" +
				"Nightshade \t %2s\n" +
				"Lich       \t %2s\n" +
				"Special    \t %2s\n",
		}

		turnUndeadLegend = &i18n.Message{
			ID:          "Turn Undead",
			Description: "Turn Undead Legend Table",
			Other: "\n" +
				"7, 9, or 11 number needed to turn successfully\n" +
				"T           automatic turn, 2d6 Hit Dice of undead\n" +
				"D           automatic Destroy, 2d6 Hit Dice of undead\n" +
				"D+          automatic Destroy, 3d6 Hit Dice of undead\n" +
				"D#          automatic Destroy, 4d6 Hit Dice of undead\n\n",
		}
	case localization.OutputFormatObsidian:
		turnUndead = &i18n.Message{
			ID:          "Turn Undead (Obsidian)",
			Description: "Turn Undead Table for Obsidian",
			Other: "" +
				"|           |     |\n" +
				"| :--- | :---: |\n" +
				"| Skeleton   | %2s |\n" +
				"| Zombie     | %2s |\n" +
				"| Ghoul      | %2s |\n" +
				"| Wight      | %2s |\n" +
				"| Wraith     | %2s |\n" +
				"| Mummy      | %2s |\n" +
				"| Spectre    | %2s |\n" +
				"| Vampire    | %2s |\n" +
				"| Phantom    | %2s |\n" +
				"| Haunt      | %2s |\n" +
				"| Spirit     | %2s |\n" +
				"| Nightshade | %2s |\n" +
				"| Lich       | %2s |\n" +
				"| Special    | %2s |\n",
		}
		turnUndeadLegend = &i18n.Message{
			ID:          "Turn Undead Legend (Obsidian)",
			Description: "Turn Undead Table Legend for Obsidian",
			Other: "\n" +
				">[!infobox]\n" +
				">| | |\n" +
				">| :---: | :--- |\n" +
				">| 7, 9, or 11 | number needed to turn successfully |\n" +
				">| T           | automatic turn, 2d6 Hit Dice of undead |\n" +
				">| D           | automatic Destroy, 2d6 Hit Dice of undead |\n" +
				">| D+          | automatic Destroy, 3d6 Hit Dice of undead |\n" +
				">| D#          | automatic Destroy, 4d6 Hit Dice of undead |\n\n",
		}
	}
	turnUndeadMsg := localization.Locale[localization.LanguageSetting].MustLocalize(&i18n.LocalizeConfig{DefaultMessage: turnUndead})
	turnUndeadLegendMsg := localization.Locale[localization.LanguageSetting].MustLocalize(&i18n.LocalizeConfig{DefaultMessage: turnUndeadLegend})

	return fmt.Sprintf(turnUndeadLegendMsg+turnUndeadMsg, a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], a[10], a[11], a[12], a[13])
}

func (c Cleric) SpecialAbilities(xp int) ClassAbilities {
	currentLevel := c.LevelIncludingRank(xp)
	spellslots := ClericSpellSlotsPerLevel[currentLevel-1]
	c.Abilities.Add("Spell Slots", 2, spellslots.String(), "")
	for ability := range c.Abilities {
		if c.Abilities[ability].Table != "" && c.Abilities[ability].ID == "Turn Undead" {
			tua := TurnUndeadAbilitiesPerLevel[currentLevel-1]
			c.Abilities[ability].Table = tua.String()
		}

	}

	return c.Abilities
	//currentLevel := c.Level(xp)
	//tua := TurnUndeadAbilitiesPerLevel[currentLevel-1]
	//tuastr := tua.String() + "----------\n" +
	//	"7, 9, or 11 number needed to turn successfully\n" +
	//	"T           automatic turn, 2d6 Hit Dice of undead\n" +
	//	"D           automatic Destroy, 2d6 Hit Dice of undead\n" +
	//	"D+          automatic Destroy, 3d6 Hit Dice of undead\n" +
	//	"D#          automatic Destroy, 4d6 Hit Dice of undead\n"
	//spellslots := ClericSpellSlotsPerLevel[currentLevel-1]
	//abilities := make(ClassAbilities, 0)
	//abilities.Add("Turn Undead", 1, tuastr, ""+
	//	"A cleric has the power to force certain monsters called the \"undead\" (skeletons, zombies, ghouls, wights, "+
	//	"and other types) to run away, or even to perish. This special ability is called \"turning undead.\"\n"+
	//	"When a cleric encounters an undead monster, the cleric may either attack it normally (with a weapon or spell), or "+
	//	"try to turn it. The cleric cannot both attack and turn undead in one round.\n"+
	//	"When you want your cleric to try to turn undead, just tell your Dungeon Master \"I'll try to turn undead this round.\" "+
	//	"The power to turn undead is inherent in the cleric; he does not need the symbol of his faith or any other "+
	//	"device to do it, unless the DM declares otherwise.\n"+
	//	"Undead monsters are not automatically turned by the cleric. When the encounter occurs, the player must refer to "+
	//	"the cleric's Turning Undead table to find the effect the cleric has.\n"+
	//	"When the cleric tries to turn an undead monster, find the cleric's LevelIncludingRank of experience across the top of the "+
	//	"table. Read down the left column until you find the name of the undead monster.\n"+
	//	"If you see a \"—\"in the column, then you can not turn the monster. If you see anything else, you have a chance "+
	//	"to turn the monster, or perhaps several monsters. See immediately below, under \"Explanation of Results,\" to "+
	//	"learn how to find out if you have turned the monster.\n"+
	//	"Apply the results immediately. If the attempt succeeds, one or more of the undead monsters will retreat or be "+
	//	"destroyed. But don't forget, if the monster is turned, it hasn't been destroyed; it may decide to return soon...\n"+
	//	"If you try to turn a specific undead monster (for instance, one specific vampire) and fail, you cannot try again "+
	//	"to turn it in the same fight. At some later encounter, you can try to turn that individual again.\n\n"+
	//	"7, 9, or 11: Whenever a number is listed, the cleric has a chance to turn the undead monsters. The player rolls "+
	//	"2d6 (two six-sided dice). If the total is equal to or greater than the number given, the attempt at turning undead "+
	//	"is successful.\n"+
	//	"When the attempt at turning undead is successful, the Dungeon Master (not the player) will roll 2d6 to determine "+
	//	"the number of Hit Dice of undead monsters that turn away. At least one monster will be turned, regardless of what "+
	//	"the DM rolls on his dice.\n"+
	//	"T: The attempt at turning the undead automatically succeeds; the cleric's player does not need to roll for success. "+
	//	"To determine how many undead will be turned, the DM rolls 2d6 as described above; regardless of his roll, at least "+
	//	"one undead will be turned.\n"+
	//	"D: The attempt at turning the undead automatically succeeds—in fact, it succeeds so well that the affected monsters "+
	//	"are destroyed instead of merely turned. To determine how many Hit Dice of undead will be destroyed, the DM rolls "+
	//	"2d6 as described above; regardless of his roll, at least one undead will be destroyed. (The DM decides what "+
	//	"happens when the monsters are destroyed: They might fade away, burst into flame and crumble away, or disintegrate "+
	//	"like a vampire in sunlight, for instance.)\n"+
	//	"D + : This is the same as the \" D \" result above, except that the DM rolls 3d6 to find out how many Hit Dice of "+
	//	"undead will be destroyed. Regardless of the roll, at least one undead will be destroyed.\n"+
	//	"D#: This is the same as the \"D\" result above, except that the DM rolls 4d6 to find out how many Hit Dice of "+
	//	"undead will be destroyed. Regardless of the roll, at least one undead will be destroyed.\n")
	//
	//if currentLevel >= 2 {
	//	abilities.Add("Spell Slots", 2, spellslots.String(), "")
	//}
	//
	//out, err := yaml.Marshal(abilities)
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "Unable to marshall abilities to yaml: %s\n", err)
	//} else {
	//	os.WriteFile(path.Join("data", "classes", "cleric.yaml"), out, 0640)
	//}
	//
	//return abilities
}
