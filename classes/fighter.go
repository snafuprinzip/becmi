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

type Fighter struct {
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

var FighterXPLevel XPLevel = XPLevel{-1, 2000, 4000, 8000, 16000, 32000, 64000, 120000, 240000, 360000,
	480000, 600000, 720000, 840000, 960000, 1080000, 1200000, 1320000, 1440000, 1560000,
	1680000, 1800000, 1920000, 2040000, 2160000, 2280000, 2400000, 2520000, 2640000, 2760000,
	2880000, 3000000, 3120000, 3240000, 3360000, 3480000}

var FighterSavingThrows savingthrows.ByLevel = savingthrows.ByLevel{
	// Level 1-3
	{12, 13, 14, 15, 16},
	{12, 13, 14, 15, 16},
	{12, 13, 14, 15, 16},
	// Level 4-6
	{10, 11, 12, 13, 14},
	{10, 11, 12, 13, 14},
	{10, 11, 12, 13, 14},
	// Level 7-9
	{8, 9, 10, 11, 12},
	{8, 9, 10, 11, 12},
	{8, 9, 10, 11, 12},
	// Level 10-12
	{6, 7, 8, 9, 10},
	{6, 7, 8, 9, 10},
	{6, 7, 8, 9, 10},
	// Level 13-15
	{6, 6, 7, 8, 9},
	{6, 6, 7, 8, 9},
	{6, 6, 7, 8, 9},
	// Level 16-18
	{5, 6, 6, 7, 8},
	{5, 6, 6, 7, 8},
	{5, 6, 6, 7, 8},
	// Level 19-21
	{5, 5, 6, 6, 7},
	{5, 5, 6, 6, 7},
	{5, 5, 6, 6, 7},
	// Level 22-24
	{4, 5, 5, 5, 6},
	{4, 5, 5, 5, 6},
	{4, 5, 5, 5, 6},
	// Level 25-27
	{4, 4, 5, 4, 5},
	{4, 4, 5, 4, 5},
	{4, 4, 5, 4, 5},
	// Level 28-30
	{3, 4, 4, 3, 4},
	{3, 4, 4, 3, 4},
	{3, 4, 4, 3, 4},
	// Level 31-33
	{3, 3, 3, 2, 3},
	{3, 3, 3, 2, 3},
	{3, 3, 3, 2, 3},
	// Level 34-36
	{2, 2, 2, 2, 2},
	{2, 2, 2, 2, 2},
	{2, 2, 2, 2, 2},
}

func init() {
	ClassIndices = append(ClassIndices, "Fighter")
	if Classes == nil {
		Classes = make(map[string]Class)
	}
	Classes["Fighter"] = Fighter{}
}

func (c Fighter) Load() Class {
	fileContent, err := os.ReadFile(path.Join("data", "classes", localization.LanguageSetting, "fighter.yaml"))
	if err != nil {
		log.Fatal("Error reading file:", err)
	}

	err = yaml.Unmarshal(fileContent, &c)
	if err != nil {
		log.Fatal("Error unmarshalling YAML:", err)
	}
	return c
}

func (c Fighter) Name() string {
	return c.ClassName
}

func (c Fighter) Race() string { return c.ClassRace }

func (c Fighter) Requirement(attr attributes.Attributes) bool {
	return true
}

func (c Fighter) Level(xp int) int {
	for idx := range len(FighterXPLevel) {
		if xp < FighterXPLevel[idx] {
			return idx
		}
	}
	return 0
}

func (c Fighter) Rank(xp int) rune {
	return ' '
}

func (c Fighter) LevelIncludingRank(xp int) int {
	return c.Level(xp)
}

func (c Fighter) NextLevelAt(xp int) int {
	currentLevel := c.Level(xp)
	return FighterXPLevel[currentLevel]
}

func (c Fighter) CheckXPModifier(a attributes.Attributes) int {
	switch {
	case a["STR"].Value < 6:
		return -20
	case a["STR"].Value < 9:
		return -10
	case a["STR"].Value < 13:
		return 0
	case a["STR"].Value < 16:
		return 5
	case a["STR"].Value > 17:
		return 10
	}
	return 0
}

func (c Fighter) HitDice() (dice, point int) {
	return c.ClassHD, c.ClassHP
}

func (c Fighter) MaxLevel() int {
	return c.MaxClassLevel
}

func (c Fighter) MaxInternalLevel() int {
	return c.MaxInternalClassLevel
}

func (c Fighter) ArmorProficiency() string {
	return c.ClassArmor
}

func (c Fighter) WeaponProficiency() string {
	return c.ClassWeapons
}

func (c Fighter) SavingThrows(xp int) savingthrows.SavingThrows {
	currentLevel := c.LevelIncludingRank(xp)
	return FighterSavingThrows[currentLevel-1]
}

func (c Fighter) BaseMovement() int {
	return 120
}

func (c Fighter) ThAC0(xp int) int {
	currentLevel := c.Level(xp)
	switch {
	case currentLevel < 4:
		return 19
	case currentLevel < 7:
		return 17
	case currentLevel < 10:
		return 15
	case currentLevel < 13:
		return 13
	case currentLevel < 16:
		return 11
	case currentLevel < 19:
		return 9
	case currentLevel < 22:
		return 7
	case currentLevel < 25:
		return 5
	case currentLevel < 28:
		return 3
	case currentLevel < 34:
		return 2
	case currentLevel < 37:
		return 1
	}
	return 20
}

func (c Fighter) ThAC0Table(xp int) string {
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

func (c Fighter) Magic() string                    { return "Divine" }
func (c Fighter) Grimoire(xp int) *magic.Spellbook { return nil }
func (c Fighter) SpellList(xp int, spellbook *magic.Spellbook) string {
	return ""
}
func (c Fighter) SpellDescriptions(xp int, spellbook *magic.Spellbook) string {
	return ""
}

func (c Fighter) SpecialAbilities(xp int) ClassAbilities {
	return c.Abilities

	//currentLevel := c.Level(xp)
	//abilities := make(ClassAbilities, 0)
	//abilities.Add("Lance Attack", 1, "",
	//	"f a character is on a riding steed (such as a horse) and is using a lance, he can perform the lance attack if his "+
	//		"mount runs (flies, swims) for 20 yards or more toward the fighter's target.\n"+
	//		"The character gets his Strength and magic adjustments to the attack roll and damage with the lance attack "+
	//		"maneuver. The lance, if it hits, will inflict double damage with a successful hit—roll the damage for the lance, "+
	//		"multiply the result by 2, and then apply all appropriate adjustments. Without enough room to charge—if the mount"+
	//		"moves less than 20 yards or is stationary—the lance does normal damage only.\n"+
	//		"Fighters, dwarves and elves can use a lance attack, but no other character class can do so. If a character has "+
	//		"the multiple attacks maneuver, he may choose the lance attack maneuver for any attack he makes in a round. "+
	//		"However, he cannot hit the same target time after time; he must choose a new target along his mount's line of "+
	//		"movement for each attack, and therefore he must be capable of hitting each target with an attack roll of 2.")
	//abilities.Add("Set Spear vs. Charge", 1, "",
	//	"A character on foot and carrying a spear, pike, sword shield, or lance can set the weapon vs. a charge. "+
	//		"A charge is when a monster charges the character—that is, runs toward him for 20 or more yards before its attack. "+
	//		"A character can also set his spear vs. another character's lance attack against him.\n"+
	//		"When the character \"sets vs. charge,\" he holds the weapon firm, braced against the ground and toward the "+
	//		"onrushing enemy. The character gets his Strength and magic adjustments to his attack and damage rolls. If the "+
	//		"character's attack hits, he inflicts double damage on his target, adding damage adjustments after doubling.\n"+
	//		"The character must declare a set spear vs. charge before he is in hand-to-hand combat with the creature "+
	//		"charging him. For example, if the character's party wins initiative in the round and the character suspects "+
	//		"the monster will charge, he could declare his set spear vs. charge maneuver. Likewise, the characters might see "+
	//		"a group of charging monsters several rounds before they arrive, and set their spears against the charge one or "+
	//		"more rounds ahead of time.\n"+
	//		"Normally, the character makes his attack on the monster's movement phase, when the monster first moves within "+
	//		"range of the weapon. If his attack hits and kills the monster, the monster cannot hurt him in return. If his "+
	//		"attack fails to kill the monster, the monster will be able to attack on its own hand-to-hand combat phase of "+
	//		"the combat sequence.")
	//
	//if currentLevel >= 9 {
	//	abilities.Add("Smash", 9, "", ""+
	//		"This is a Fighter Combat Option maneuver, first available at 9th Level to fighters and mystics, and at other "+
	//		"experience point totals to demihumans (see their experience tables). With this hand-to-hand maneuver, the "+
	//		"character automatically loses initiative and takes a — 5 penalty to the attack roll (he still gets his Strength "+
	//		"and magic adjustments to his attack roll).\n"+
	//		"If attack hits, the character adds his Strength bonus, magic bonuses, and his entire Strength score to his "+
	//		"weapon's normal damage.")
	//	abilities.Add("Parry", 9, "", ""+
	//		"With this maneuver, the fighter does not "+
	//		"make any attack roll. Instead, he blocks incoming attacks for the entire combat round; all enemies attacking "+
	//		"him suffer a —4 penalty to hit him with melee and thrown (but not missile) weapons.")
	//	abilities.Add("Disarm", 9, "", ""+
	//		"This maneuver can only be used when the fighter attacks a weapon-using opponent. The fighter gets his normal "+
	//		"Strength and magic adjustments to his attack roll. If he hits, he inflicts no damage. Instead, the victim "+
	//		"rolls 1d20, minus his Dexterity bonuses, plus his attacker's Dexterity bonuses. If the final roll is greater "+
	//		"than the victim's Dexterity score, the victim drops his weapon.")
	//}
	//if currentLevel >= 12 {
	//	abilities.Add("Multiple Attacks", 12, "", ""+
	//		"In melee combat, if the fighter can hit his opponent with an attack roll of 2 (modified by all bonuses), he "+
	//		"can make two attacks per round against that target (three per round at LevelIncludingRank 24, four per round at LevelIncludingRank 36).\n"+
	//		"Each attack of a multiple attacks maneuver can be a throw, attack, lance attack, or disarm. A character can "+
	//		"mix and match his maneuvers; for instance, a character with three attacks per round could perform an attack, "+
	//		"disarm, attack combination against his foe, or throw three knives instead of one. This maneuver applies to "+
	//		"ideal circumstances, and the character can use movement or some other action instead of another attack.")
	//}
	//
	//out, err := yaml.Marshal(abilities)
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "Unable to marshall abilities to yaml: %s\n", err)
	//} else {
	//	os.WriteFile(path.Join("data", "classes", "fighter.yaml"), out, 0640)
	//}
	//
	//return abilities
}
