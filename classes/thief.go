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

type Thief struct {
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

var ThiefXPLevel XPLevel = XPLevel{-1, 1200, 2400, 4800, 9600, 20000, 40000, 80000, 160000, 280000,
	400000, 520000, 640000, 760000, 880000, 1000000, 1120000, 1240000, 1360000, 1480000,
	1600000, 1720000, 1840000, 1960000, 2080000, 2200000, 2320000, 2440000, 2560000, 2680000,
	2800000, 2920000, 3040000, 3160000, 3280000, 3400000}

type ThiefSkills struct {
	OpenLocks     int
	FindTraps     int
	RemoveTraps   int
	ClimbWalls    int
	MoveSilently  int
	HideInShadows int
	PickPockets   int
	HearNoise     int
}

var ThiefSkillsTable [36]ThiefSkills = [36]ThiefSkills{
	// Level 1-3
	{15, 10, 10, 87, 20, 10, 20, 30},
	{20, 15, 15, 88, 25, 15, 25, 35},
	{25, 20, 20, 89, 30, 20, 30, 40},
	// Level 4-6
	{30, 25, 25, 90, 35, 24, 35, 45},
	{35, 30, 30, 91, 40, 28, 40, 50},
	{40, 35, 34, 92, 44, 32, 45, 54},
	// Level 7-9
	{45, 40, 38, 93, 48, 35, 50, 58},
	{50, 45, 42, 94, 52, 38, 55, 62},
	{54, 50, 46, 95, 55, 41, 60, 66},
	// Level 10-12
	{58, 54, 50, 96, 58, 44, 65, 70},
	{62, 58, 54, 97, 61, 47, 70, 74},
	{66, 62, 58, 98, 64, 50, 75, 78},
	// Level 13-15
	{69, 66, 61, 99, 66, 53, 80, 81},
	{72, 70, 64, 100, 68, 56, 85, 84},
	{75, 73, 67, 101, 70, 58, 90, 87},
	// Level 16-18
	{78, 76, 70, 102, 72, 60, 95, 90},
	{81, 80, 73, 103, 74, 62, 100, 92},
	{84, 83, 76, 104, 76, 64, 105, 94},
	// Level 19-21
	{86, 86, 79, 105, 78, 66, 110, 96},
	{88, 89, 82, 106, 80, 68, 115, 98},
	{90, 92, 85, 107, 82, 70, 120, 100},
	// Level 22-24
	{92, 94, 88, 108, 84, 72, 125, 102},
	{94, 96, 91, 109, 86, 74, 130, 104},
	{96, 98, 94, 110, 88, 76, 135, 106},
	// Level 25-27
	{98, 99, 97, 111, 89, 78, 140, 108},
	{100, 100, 100, 112, 90, 80, 145, 110},
	{102, 101, 103, 113, 91, 82, 150, 112},
	// Level 28-30
	{104, 102, 106, 114, 92, 84, 155, 114},
	{106, 103, 109, 115, 93, 86, 160, 116},
	{108, 104, 112, 116, 94, 88, 165, 118},
	// Level 31-33
	{110, 105, 115, 117, 95, 90, 170, 120},
	{112, 106, 118, 118, 96, 92, 175, 122},
	{114, 107, 121, 118, 97, 94, 180, 124},
	// Level 34-36
	{116, 108, 124, 119, 98, 96, 185, 126},
	{118, 109, 127, 119, 99, 98, 190, 128},
	{120, 110, 130, 120, 100, 100, 195, 130},
}

var ThiefSavingThrows savingthrows.ByLevel = savingthrows.ByLevel{
	// Level 1-4
	{13, 14, 13, 16, 15},
	{13, 14, 13, 16, 15},
	{13, 14, 13, 16, 15},
	{13, 14, 13, 16, 15},
	// Level 5-8
	{11, 12, 11, 14, 13},
	{11, 12, 11, 14, 13},
	{11, 12, 11, 14, 13},
	{11, 12, 11, 14, 13},
	// Level 9-12
	{9, 10, 9, 12, 11},
	{9, 10, 9, 12, 11},
	{9, 10, 9, 12, 11},
	{9, 10, 9, 12, 11},
	// Level 13-16
	{7, 8, 7, 10, 9},
	{7, 8, 7, 10, 9},
	{7, 8, 7, 10, 9},
	{7, 8, 7, 10, 9},
	// Level 17-20
	{5, 6, 5, 8, 7},
	{5, 6, 5, 8, 7},
	{5, 6, 5, 8, 7},
	{5, 6, 5, 8, 7},
	// Level 21-24
	{4, 5, 4, 6, 5},
	{4, 5, 4, 6, 5},
	{4, 5, 4, 6, 5},
	{4, 5, 4, 6, 5},
	// Level 25-28
	{3, 4, 3, 4, 4},
	{3, 4, 3, 4, 4},
	{3, 4, 3, 4, 4},
	{3, 4, 3, 4, 4},
	// Level 29-32
	{2, 3, 2, 3, 3},
	{2, 3, 2, 3, 3},
	{2, 3, 2, 3, 3},
	{2, 3, 2, 3, 3},
	// Level 33-36
	{2, 2, 2, 2, 2},
	{2, 2, 2, 2, 2},
	{2, 2, 2, 2, 2},
	{2, 2, 2, 2, 2},
}

func init() {
	ClassIndices = append(ClassIndices, "Thief")
	if Classes == nil {
		Classes = make(map[string]Class)
	}
	Classes["Thief"] = Thief{}
}

func (c Thief) Load() Class {
	fileContent, err := os.ReadFile(path.Join("data", "classes", localization.LanguageSetting, "thief.yaml"))
	if err != nil {
		log.Fatal("Error reading file:", err)
	}

	err = yaml.Unmarshal(fileContent, &c)
	if err != nil {
		log.Fatal("Error unmarshalling YAML:", err)
	}
	return c
}

func (c Thief) Name() string {
	return c.ClassName
}

func (c Thief) Race() string { return c.ClassRace }

func (c Thief) Requirement(attr attributes.Attributes) bool {
	return true
}

func (c Thief) Level(xp int) int {
	for idx := range len(ThiefXPLevel) {
		if xp < ThiefXPLevel[idx] {
			return idx
		}
	}
	return 0
}

func (c Thief) Rank(xp int) rune {
	return ' '
}

func (c Thief) LevelIncludingRank(xp int) int {
	return c.Level(xp)
}

func (c Thief) NextLevelAt(xp int) int {
	currentLevel := c.Level(xp)
	return ThiefXPLevel[currentLevel]
}

func (c Thief) CheckXPModifier(a attributes.Attributes) int {
	switch {
	case a["DEX"].Value < 6:
		return -20
	case a["DEX"].Value < 9:
		return -10
	case a["DEX"].Value < 13:
		return 0
	case a["DEX"].Value < 16:
		return 5
	case a["DEX"].Value > 17:
		return 10
	}
	return 0
}

func (c Thief) HitDice() (dice, point int) {
	return c.ClassHD, c.ClassHP
}

func (c Thief) MaxLevel() int {
	return c.MaxClassLevel
}

func (c Thief) MaxInternalLevel() int {
	return c.MaxInternalClassLevel
}

func (c Thief) ArmorProficiency() string {
	return c.ClassArmor
}

func (c Thief) WeaponProficiency() string {
	return c.ClassWeapons
}

func (c Thief) SavingThrows(xp int) savingthrows.SavingThrows {
	currentLevel := c.LevelIncludingRank(xp)
	return ThiefSavingThrows[currentLevel-1]
}

func (c Thief) BaseMovement() int {
	return 120
}

func (c Thief) ThAC0(xp int) int {
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

func (c Thief) ThAC0Table(xp int) string {
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

func (c Thief) Magic() string { return "Arcane" }

func (c Thief) Grimoire(xp int) *magic.Spellbook { return nil }

func (c Thief) SpellList(xp int, spellbook *magic.Spellbook) string {
	return ""
}

func (c Thief) SpellDescriptions(xp int, spellbook *magic.Spellbook) string {
	return ""
}

func (s ThiefSkills) String() string {
	return fmt.Sprintf(""+
		"Open Locks:      %d%%\n"+
		"Find Traps:      %d%%\n"+
		"Remove Traps:    %d%%\n"+
		"Climb Walls:     %d%%\n"+
		"Move Silently:   %d%%\n"+
		"Hide In Shadows: %d%%\n"+
		"Pick Pockets:    %d%%\n"+
		"Hear Noise:      %d%%\n", s.OpenLocks, s.FindTraps, s.RemoveTraps, s.ClimbWalls, s.MoveSilently, s.HideInShadows, s.PickPockets, s.HearNoise)
}

func (c Thief) SpecialAbilities(xp int) ClassAbilities {
	currentLevel := c.LevelIncludingRank(xp)
	thiefskills := ThiefSkillsTable[currentLevel-1]
	for ability := range c.Abilities {
		if c.Abilities[ability].Table != "" && c.Abilities[ability].ID == "Thief Skills" {
			formatstr := c.Abilities[ability].Table
			c.Abilities[ability].Table = fmt.Sprintf(formatstr, thiefskills.OpenLocks, thiefskills.FindTraps, thiefskills.RemoveTraps,
				thiefskills.ClimbWalls, thiefskills.MoveSilently, thiefskills.HideInShadows, thiefskills.PickPockets, thiefskills.HearNoise)
			break
		}
	}

	return c.Abilities

	//currentLevel := c.Level(xp)
	//thiefskills := ThiefSkillsTable[currentLevel-1]
	////abilities := []string{thiefskills.String()}
	//abilities := make(ClassAbilities, 0)
	//abilities.Add("Thief Skills", 1, thiefskills.String(), ""+
	//	"Open Locks (OL):\nWith successful use of this special ability, and with professional lockpicks (often called "+
	//	"\"thieves' tools\"), the thief may open locks. The character may try to use this skill only once per lock. The "+
	//	"thief may not try again with that particular lock until he gains another LevelIncludingRank of experience. Without lockpicks, "+
	//	"he may not use this ability.\n\n"+
	//	"Find Trap (FT):\nWith successful use of this special ability, the thief may examine a room or an object and "+
	//	"determine whether it is rigged with traps. He may check only once per trap, and failure prevents the character "+
	//	"from finding any trap in or on the object searched. (Since the DM actually does the rolling, the player doesn't "+
	//	"know how many traps he's rolling to find.) If the thief finds a trap, he may use his Remove Traps ability to "+
	//	"remove or deactivate it.\n\n"+
	//	"Remove Traps (RT):\nWith successful use of this special ability, the thief may remove or deactivate a trap. "+
	//	"He may not roll this ability against a trap unless the trap has been found. The thief may try his ability only "+
	//	"once per trap; failure to remove a trap triggers the trap.\n\n"+
	//	"Climb Walls (CW):\nWith successful use of this special ability, the thief can climb steep surfaces, such as sheer "+
	//	"cliffs, walls, and so forth. The chances for success are good, but if failed, the thief slips at the halfway "+
	//	"point and falls. The DM rolls for success once for every 100' climbed.\n"+
	//	"If the roll is a failure, the thief takes 1-6 (1d6) points of damage per 10' fallen. Falling during a 10' climb "+
	//	"will inflict 1 point of damage.\n\n"+
	//	"Move Silently (MS):\nSuccessful use of this special ability allows the thief to move silently. When the thief "+
	//	"tries to use this skill, he always believes he has been successful, but a failed roll means that someone can "+
	//	"hear his passage. The DM, at his discretion, may modify the thiefs roll at any time: When he tries moving "+
	//	"silently across a field of dried leaves, his percentage chance would go down, while if he does so during a loud "+
	//	"tournament, his chance will be greatly enhanced. Note that it doesn't do the thief any good to use this skill "+
	//	"against someone who is already aware of him.\n\n"+
	//	"Hide in Shadows (HS):\nSuccessful use of this special ability means that the thief moves into and remains in "+
	//	"shadows, making him very hard to see. While the thief is in shadows, observers only get a chance to see him if "+
	//	"they look directly at him, at which time he must roll again; success means that he remains unobserved. While in "+
	//	"shadows, the thief may use his Move Silently ability, but attacking someone reveals the thief. If the thief tries "+
	//	"to hide in shadows but fails, he will not know that his position of concealment is a failure until someone sees "+
	//	"him and announces the fact. Note that if the thief is under direct observation, he can't hide in shadows against "+
	//	"the people watching him; they'll be able to follow his progress with no problem.\n\n"+
	//	"Pick Pockets (PP):\nThis special ability allows the character to steal things from another character's person "+
	//	"without him noticing. It's a very risky skill to use. If the attempt succeeds, the thief is able to pick the "+
	//	"other's pockets without anyone noticing. If the roll is a simple failure, the thief fails to get his hands on "+
	//	"what he's seeking. If the roll is greater than twice what the thief needs to succeed or an 00 in any case, the "+
	//	"thief is caught in the act by his intended victim, and possibly others. When using the skill, subtract 5% per "+
	//	"LevelIncludingRank or HD of victim. (Normal men—men and women who have no adventuring ability at all and do not belong to any "+
	//	"adventuring character class—are treated as being 0 LevelIncludingRank.)\n"+
	//	"Example: A 1st LevelIncludingRank thief tries to pick the pocket of a 1st LevelIncludingRank fighter walking along the street. His chance "+
	//	"is 20% (normal) minus 5 (5 times 1), or 15%. The DM rolls the percentile dice and rolls a 41. This is over twice "+
	//	"what he needed to roll, so the thief is caught in the act.\n\n"+
	//	"Hear Noise (HN):\nThis special ability gives the thief the ability to hear faint noises—such as breathing on the "+
	//	"other side of the door, or the clatter of distant footsteps approaching fast. The DM can rule that any loud situation, "+
	//	"such as a battle, prevents the thief from using this skill.\n")
	//abilities.Add("Backstabbing", 1, "",
	//	"If a thief can sneak up on a victim, completely unnoticed, the thief may backstab "+
	//		"— if he is using a one-handed melee weapon, he may strike at particularly vulnerable points of his target's body. "+
	//		"(Though the ability is called \"backstabbing,\" the weapon doesn't have to be a stabbing weapon. "+
	//		"A thief can use this ability with a club, for example.)\n"+
	//		"When backstabbing, the thief gains a bonus of +4 on the attack roll; if the target is hit, the damage done is "+
	//		"twice normal (roll the damage for the weapon, multiply the result by two, and then add any pertinent modifiers).\n"+
	//		"If the intended victim sees, hears, or is warned of the thief's approach, the thief's attack is not a backstab; "+
	//		"it is an ordinary attack, doing the damage appropriate for the weapon used.\n"+
	//		"When no battle is in progress, a backstab attempt may require a Move Silently ability check. "+
	//		"The DM will make all the necessary decisions on that matter.\n")
	//if currentLevel >= 4 {
	//	abilities.Add("Read Languages", 4, "",
	//		"80% chance to read any normal writing or language (including simple codes, dead languages, treasure maps, "+
	//			"and so on, but not magical writings). If he tries but fails to read a piece of writing, he must gain at "+
	//			"least one experience LevelIncludingRank before trying to read it again.\n")
	//}
	//if currentLevel >= 10 {
	//	abilities.Add("Cast Spells From Magic-User Scrolls", 10, "",
	//		"At 10th LevelIncludingRank, a thief gains the ability to cast magic-user spells from spell scrolls. However, there is "+
	//			"always a 10% chance that the spell will backfire, creating an unexpected result, because of the thief's"+
	//			"imperfect understanding of magical writings. This ability only allows thieves to cast spells from existing "+
	//			"magic scrolls, not to write their own.\n")
	//}
	//out, err := yaml.Marshal(abilities)
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "Unable to marshall abilities to yaml: %s\n", err)
	//} else {
	//	os.WriteFile(path.Join("data", "classes", "thief.yaml"), out, 0640)
	//}
	//return abilities
}
