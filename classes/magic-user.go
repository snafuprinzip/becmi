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

type MagicUser struct {
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
type MagicUserSpellSlots [9]int

var MagicUserXPLevel XPLevel = XPLevel{-1, 2500, 5000, 10000, 20000, 40000, 80000, 150000, 300000, 450000,
	600000, 750000, 900000, 1050000, 1200000, 1350000, 1500000, 1650000, 1900000, 1950000,
	2100000, 2250000, 2400000, 2550000, 2700000, 2850000, 3000000, 3150000, 3300000, 3450000,
	3600000, 3750000, 3900000, 4050000, 4200000, 4350000}

var MagicUserSpellSlotsPerLevel [36]MagicUserSpellSlots = [36]MagicUserSpellSlots{
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
	// Level 10-12
	{3, 3, 3, 3, 2},
	{4, 3, 3, 3, 2, 1},
	{4, 4, 4, 3, 2, 1},
	// Level 13-15
	{4, 4, 4, 3, 2, 2},
	{4, 4, 5, 4, 3, 2},
	{5, 4, 6, 4, 3, 2, 1},
	// Level 16-18
	{5, 5, 5, 4, 3, 2, 2},
	{6, 5, 5, 4, 4, 3, 2},
	{6, 5, 5, 4, 4, 3, 2, 1},
	// Level 19-21
	{6, 5, 5, 5, 4, 3, 2, 2},
	{6, 5, 5, 5, 4, 4, 3, 2},
	{6, 5, 5, 5, 4, 4, 3, 2, 1},
	// Level 22-24
	{6, 6, 5, 5, 5, 4, 3, 2, 2},
	{6, 6, 6, 6, 5, 4, 3, 3, 2},
	{7, 7, 6, 6, 5, 5, 4, 3, 2},
	// Level 25-27
	{7, 7, 6, 6, 5, 5, 4, 4, 3},
	{7, 7, 7, 6, 6, 5, 5, 4, 3},
	{7, 8, 7, 6, 6, 5, 5, 5, 4},
	// Level 28-30
	{8, 8, 7, 7, 6, 6, 6, 5, 4},
	{8, 8, 7, 7, 7, 6, 6, 5, 5},
	{8, 8, 8, 7, 7, 7, 6, 6, 5},
	// Level 31-33
	{8, 8, 8, 7, 7, 7, 7, 6, 6},
	{9, 8, 8, 8, 8, 7, 7, 7, 6},
	{9, 9, 9, 8, 8, 8, 7, 7, 7},
	// Level 34-36
	{9, 9, 9, 9, 8, 8, 8, 8, 7},
	{9, 9, 9, 9, 9, 9, 8, 8, 8},
	{9, 9, 9, 9, 9, 9, 9, 9, 9},
}

var MagicUserSavingThrows savingthrows.ByLevel = savingthrows.ByLevel{
	// Level 1-5
	{13, 14, 13, 16, 15},
	{13, 14, 13, 16, 15},
	{13, 14, 13, 16, 15},
	{13, 14, 13, 16, 15},
	{13, 14, 13, 16, 15},
	// Level 6-10
	{11, 12, 11, 14, 12},
	{11, 12, 11, 14, 12},
	{11, 12, 11, 14, 12},
	{11, 12, 11, 14, 12},
	{11, 12, 11, 14, 12},
	// Level 11-15
	{9, 10, 9, 12, 9},
	{9, 10, 9, 12, 9},
	{9, 10, 9, 12, 9},
	{9, 10, 9, 12, 9},
	{9, 10, 9, 12, 9},
	// Level 16-20
	{7, 8, 7, 10, 6},
	{7, 8, 7, 10, 6},
	{7, 8, 7, 10, 6},
	{7, 8, 7, 10, 6},
	{7, 8, 7, 10, 6},
	// Level 21-24
	{5, 6, 5, 8, 4},
	{5, 6, 5, 8, 4},
	{5, 6, 5, 8, 4},
	{5, 6, 5, 8, 4},
	// Level 25-28
	{4, 4, 4, 6, 3},
	{4, 4, 4, 6, 3},
	{4, 4, 4, 6, 3},
	{4, 4, 4, 6, 3},
	// Level 29-32
	{3, 3, 3, 4, 2},
	{3, 3, 3, 4, 2},
	{3, 3, 3, 4, 2},
	{3, 3, 3, 4, 2},
	// Level 33-36
	{2, 2, 2, 2, 2},
	{2, 2, 2, 2, 2},
	{2, 2, 2, 2, 2},
	{2, 2, 2, 2, 2},
}

func init() {
	ClassIndices = append(ClassIndices, "Magic-User")
	if Classes == nil {
		Classes = make(map[string]Class)
	}
	Classes["Magic-User"] = MagicUser{}
}

func (c MagicUser) Load() Class {
	fileContent, err := os.ReadFile(path.Join("data", "classes", localization.LanguageSetting, "magic-user.yaml"))
	if err != nil {
		log.Fatal("Error reading file:", err)
	}

	err = yaml.Unmarshal(fileContent, &c)
	if err != nil {
		log.Fatal("Error unmarshalling YAML:", err)
	}
	return c
}

func (c MagicUser) Name() string {
	return c.ClassName
}

func (c MagicUser) Race() string { return c.ClassRace }

func (c MagicUser) Requirement(attr attributes.Attributes) bool {
	return true
}

func (c MagicUser) Level(xp int) int {
	for idx := range len(MagicUserXPLevel) {
		if xp < MagicUserXPLevel[idx] {
			return idx
		}
	}
	return 0
}

func (c MagicUser) Rank(xp int) rune {
	return ' '
}

func (c MagicUser) LevelIncludingRank(xp int) int {
	return c.Level(xp)
}

func (c MagicUser) NextLevelAt(xp int) int {
	currentLevel := c.Level(xp)
	return MagicUserXPLevel[currentLevel]
}

func (c MagicUser) CheckXPModifier(a attributes.Attributes) int {
	switch {
	case a["INT"].Value < 6:
		return -20
	case a["INT"].Value < 9:
		return -10
	case a["INT"].Value < 13:
		return 0
	case a["INT"].Value < 16:
		return 5
	case a["INT"].Value > 17:
		return 10
	}
	return 0
}

func (c MagicUser) HitDice() (dice, point int) {
	return c.ClassHD, c.ClassHP
}

func (c MagicUser) MaxLevel() int {
	return c.MaxClassLevel
}

func (c MagicUser) MaxInternalLevel() int {
	return c.MaxInternalClassLevel
}

func (c MagicUser) ArmorProficiency() string {
	return c.ClassArmor
}

func (c MagicUser) WeaponProficiency() string {
	return c.ClassWeapons
}

func (c MagicUser) SavingThrows(xp int) savingthrows.SavingThrows {
	currentLevel := c.LevelIncludingRank(xp)
	return MagicUserSavingThrows[currentLevel-1]
}

func (c MagicUser) BaseMovement() int {
	return 120
}

func (c MagicUser) ThAC0(xp int) int {
	currentLevel := c.Level(xp)
	switch {
	case currentLevel < 6:
		return 19
	case currentLevel < 11:
		return 17
	case currentLevel < 16:
		return 15
	case currentLevel < 21:
		return 13
	case currentLevel < 26:
		return 11
	case currentLevel < 31:
		return 9
	case currentLevel < 36:
		return 7
	case currentLevel < 37:
		return 5
	}
	return 20
}

func (c MagicUser) ThAC0Table(xp int) string {
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

	return fmt.Sprintf(""+
		"10   9   8   7   6   5   4   3   2   1     0    -1   -2   -3   -4   -5   -6   -7   -8   -9  -10\n"+
		"%2d  %2d  %2d  %2d  %2d  %2d  %2d  %2d  %2d  %2d    %2d    %2d   %2d   %2d   %2d   %2d   %2d   %2d   %2d   %2d   %2d\n",
		table[30], table[29], table[28], table[27], table[26], table[25], table[24], table[23], table[22], table[21],
		table[20], table[19], table[18], table[17], table[16], table[15], table[14], table[13], table[12], table[11],
		table[10])
}

func (c MagicUser) Magic() string { return "Arcane" }

func (s MagicUserSpellSlots) String() string {
	slots := &i18n.Message{
		ID:          "Spell Slots Magic-User",
		Description: "Spell Slots for Magic-User",
		Other: "" +
			"Spell Slots\n" +
			"-----------\n" +
			"1st: %2d\n2nd: %2d\n3rd: %2d\n4th: %2d\n5th: %2d\n6th: %2d\n7th: %2d\n8th: %2d\n9th: %2d\n",
	}

	translation := localization.Locale[localization.LanguageSetting].MustLocalize(&i18n.LocalizeConfig{DefaultMessage: slots})

	return fmt.Sprintf(translation, s[0], s[1], s[2], s[3], s[4], s[5], s[6], s[7], s[8])
}

func (s MagicUserSpellSlots) ObsidianString() string {
	slots := &i18n.Message{
		ID:          "Spell Slots Magic-User (Obsidian)",
		Description: "Spell Slots for Magic-User for Obsidian",
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
			">| 7th | %2d |\n" +
			">| 8th | %2d |\n" +
			">| 9th | %2d |\n",
	}

	translation := localization.Locale[localization.LanguageSetting].MustLocalize(&i18n.LocalizeConfig{DefaultMessage: slots})

	return fmt.Sprintf(translation, s[0], s[1], s[2], s[3], s[4], s[5], s[6], s[7], s[8])
}

// Grimoire generates a new grimoire for a spell caster character
func (c MagicUser) Grimoire(xp int) *magic.Spellbook {
	currentLevel := c.Level(xp)
	var book magic.Spellbook

	book[0] = append(book[0], "Read Magic")
	roll := dice.RollDice(4)
	switch roll {
	case 1:
		book[0] = append(book[0], "Charm Person")
	case 2:
		book[0] = append(book[0], "Magic Missile")
	case 3:
		book[0] = append(book[0], "Sleep")
	case 4:
		book[0] = append(book[0], "Shield")
	}
	if currentLevel >= 2 {
		roll := dice.RollDice(4)
		switch roll {
		case 1:
			book[0] = append(book[0], "Floating Disc")
		case 2:
			book[0] = append(book[0], "Hold Portal")
		case 3:
			book[0] = append(book[0], "Read Languages")
		case 4:
			book[0] = append(book[0], "Ventriloquism")
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

	//for _, level := range book {
	//	for idx, spell := range level {
	//		fmt.Printf("%2d: %s\n", idx+1, spell)
	//	}
	//}

	return &book
}

// SpellList returns the list of available spells for a character, either from their grimoire or from the whole available spell list
func (c MagicUser) SpellList(xp int, spellbook *magic.Spellbook) string {
	var spellList string

	for idx := 0; idx < 9; idx++ {
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
					spellList += "- " + spelldesc.Name + "\n"
				}
			}
		}

		spellList += "\n"
	}
	return spellList
}

func (c MagicUser) SpellDescriptions(xp int, spellbook *magic.Spellbook) string {
	var spells string

	for idx := 0; idx < 9; idx++ {
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

func (c MagicUser) SpellDescriptionsObsidian(xp int, spellbook *magic.Spellbook) string {
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

func (c MagicUser) SpecialAbilities(xp int) ClassAbilities {
	currentLevel := c.LevelIncludingRank(xp)
	spellslots := MagicUserSpellSlotsPerLevel[currentLevel-1]
	c.Abilities.Add("Spell Slots", 1, spellslots.String(), "")
	//for ability := range c.Abilities {
	//	if c.Abilities[ability].Table != "" && c.Abilities[ability].ID == "Spell Slots" {
	//		c.Abilities[ability].Table = spellslots.String()
	//		break
	//	}
	//}

	return c.Abilities
	//currentLevel := c.Level(xp)
	//spellslots := MagicUserSpellSlotsPerLevel[currentLevel-1]
	//abilities := make(ClassAbilities, 0)
	//
	//if currentLevel >= 2 {
	//	abilities.Add("Spell Slots", 2, spellslots.String(), "")
	//}
	//
	//return abilities
}
