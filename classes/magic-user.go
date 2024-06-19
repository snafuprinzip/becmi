package classes

import (
	"becmi/attributes"
	"becmi/magic"
	"becmi/savingthrows"
	"fmt"
)

type MagicUser struct {
}
type MagicUserSpellSlots [9]int

var MagicUserXPLevel XPLevel = XPLevel{0, 2500, 5000, 10000, 20000, 40000, 80000, 150000, 300000, 450000,
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

func (c MagicUser) Name() string {
	return "Magic-User"
}

func (c MagicUser) Requirement(attr attributes.Attributes) bool {
	return true
}

func (c MagicUser) Level(xp int) int {
	for idx := range len(MagicUserXPLevel) {
		if xp <= MagicUserXPLevel[idx] {
			return idx + 1
		}
	}
	return 0
}

func (c MagicUser) NextLevelAt(currentLevel int) int {
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
	return 4, 1
}

func (c MagicUser) MaxLevel() int {
	return 36
}

func (c MagicUser) ArmorProficiency() string {
	return "None; no shield permitted."
}

func (c MagicUser) WeaponProficiency() string {
	return "Dagger only. Optional (DM's discretion): staff, blowgun, flaming oil, holy water, net, thrown rock, sling, whip."
}

func (c MagicUser) SavingThrows(currentLevel int) savingthrows.SavingThrows {
	return MagicUserSavingThrows[currentLevel-1]
}

func (c MagicUser) SpecialAbilities(currentLevel int) ClassAbilities {
	spellslots := MagicUserSpellSlotsPerLevel[currentLevel-1]
	abilities := make(ClassAbilities, 0)

	if currentLevel >= 2 {
		abilities.Add("Spell Slots", 2, spellslots.String(), "")
	}

	return abilities
}

func (s MagicUserSpellSlots) String() string {
	return fmt.Sprintf("1st: %d, 2nd: %d, 3rd: %d, 4th: %d, 5th: %d, 6th: %d, 7th: %d, 8th: %d, 9th: %d\n", s[0], s[1], s[2], s[3], s[4], s[5], s[6], s[7], s[8])
}

func (c MagicUser) SpellList(currentLevel int, spellbook *magic.Spellbook) string {
	var spellList string

	slots := MagicUserSpellSlotsPerLevel[currentLevel-1]
	for idx := 0; idx < 9; idx++ {
		if slots[idx] == 0 {
			break
		}
		if idx == 0 {
			spellList = "" +
				"Spells\n" +
				"======\n\n"
		}
		spellList += fmt.Sprintf(""+
			"Level %d\n"+
			"--------\n", idx+1)
		for _, spell := range magic.AllArcaneSpells[idx+1] {
			spellList += spell.Name + "\n"
		}
		spellList += "\n"
	}
	return spellList
}

func (c MagicUser) SpellDescriptions(currentLevel int, spellbook *magic.Spellbook) string {
	var spells string

	slots := MagicUserSpellSlotsPerLevel[currentLevel-1]
	for idx := 0; idx < 7; idx++ {
		if slots[idx] == 0 {
			break
		}
		if idx == 0 {
			spells = "" +
				"Spelldescriptions\n" +
				"=================\n\n"
		}
		spells += fmt.Sprintf(""+
			"Level %d\n"+
			"--------\n", idx+1)
		for _, spell := range magic.AllArcaneSpells[idx+1] {
			spells += spell.String() + "\n"
		}
		spells += "\n"
	}
	return spells
}

func (c MagicUser) BaseMovement() int {
	return 120
}

func (c MagicUser) ThAC0(currentLevel int) int {
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

func (c MagicUser) Race() string { return "Human" }

func (c MagicUser) ThAC0Table(currentLevel int) string {
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

func (c MagicUser) Grimoire(currentLevel int) *magic.Spellbook { return nil }
