package attributes

import "fmt"

var AttributeIndices [6]string = [6]string{"STR", "INT", "WIS", "DEX", "CON", "CHA"}

type Attribute struct {
	Name  string
	Value int
}

type Attributes map[string]Attribute

func (a Attribute) Modifier() int {
	switch {
	case a.Value >= 2 && a.Value <= 3:
		return -3
	case a.Value >= 4 && a.Value <= 5:
		return -2
	case a.Value >= 6 && a.Value <= 8:
		return -1
	case a.Value >= 13 && a.Value <= 15:
		return 1
	case a.Value >= 16 && a.Value <= 17:
		return 2
	case a.Value >= 18 && a.Value <= 19:
		return 3
	case a.Value >= 9 && a.Value <= 12:
		fallthrough
	default:
		return 0
	}
}

func (a Attribute) ModifierString() string {
	return SignedInt(a.Modifier())
}

func (a Attribute) LanguageProficiency() string {
	switch {
	case a.Value >= 2 && a.Value <= 3:
		return "Has trouble speaking, cannot read or write"
	case a.Value >= 4 && a.Value <= 5:
		return "Cannot read or write Common"
	case a.Value >= 6 && a.Value <= 8:
		return "Can write simple common words"
	case a.Value >= 9 && a.Value <= 12:
		return "Reads and writes (usually two) native languages"
	case a.Value >= 13 && a.Value <= 15:
		return "Reads and writes native languages +1 additional"
	case a.Value >= 16 && a.Value <= 17:
		return "Reads and writes native languages +2 additional"
	case a.Value >= 18 && a.Value <= 19:
		return "Reads and writes native languages +3 additional"
	default:
		return "Error: Value out of bounds!"
	}
}

func (a Attribute) MaxRetainers() int {
	switch {
	case a.Value >= 2 && a.Value <= 3:
		return 1
	case a.Value >= 4 && a.Value <= 5:
		return 2
	case a.Value >= 6 && a.Value <= 8:
		return 3
	case a.Value >= 9 && a.Value <= 12:
		return 4
	case a.Value >= 13 && a.Value <= 15:
		return 5
	case a.Value >= 16 && a.Value <= 17:
		return 6
	case a.Value >= 18 && a.Value <= 19:
		return 7
	default:
		return 0
	}
}

func (a Attribute) RetainerMorale() int {
	switch {
	case a.Value >= 2 && a.Value <= 3:
		return 4
	case a.Value >= 4 && a.Value <= 5:
		return 5
	case a.Value >= 6 && a.Value <= 8:
		return 6
	case a.Value >= 9 && a.Value <= 12:
		return 7
	case a.Value >= 13 && a.Value <= 15:
		return 8
	case a.Value >= 16 && a.Value <= 17:
		return 9
	case a.Value >= 18 && a.Value <= 19:
		return 10
	default:
		return 0
	}
}

func (a Attributes) String() string {
	return fmt.Sprintf(""+
		"STR Strength      %2d \t (%s Attack Roll [Melee, Unarmed], Damage Roll [Melee, Thrown], Open Doors, optional: Save: Paralysis, Turn to Stone)\n"+
		"INT Intelligence  %2d \t (%s General Skills, %s, optional: Save vs. Mind Attacks)\n"+
		"WIS Wisdom        %2d \t (%s Saving Throw vs. Spells)\n"+
		"DEX Dexterity     %2d \t (%s Attack Roll [Thrown, Missiles], Armor Class, optional: Save: Wands, Dragon Breath)\n"+
		"CON Constitution  %2d \t (%s Hit Points per XP Level, optional: Save: Poison)\n"+
		"CHA Charisma      %2d \t (%s Reaction Adjustment from NPCs, %d max. Retainers, %d Retainer Morale)\n",
		a["STR"].Value, a["STR"].ModifierString(),
		a["INT"].Value, a["INT"].ModifierString(), a["INT"].LanguageProficiency(),
		a["WIS"].Value, a["WIS"].ModifierString(),
		a["DEX"].Value, a["DEX"].ModifierString(),
		a["CON"].Value, a["CON"].ModifierString(),
		a["CHA"].Value, a["CHA"].ModifierString(), a["CHA"].MaxRetainers(), a["CHA"].RetainerMorale())
}

func SignedInt(i int) string {
	if i > 0 {
		return fmt.Sprintf(" +%d", i)
	} else if i == 0 {
		return fmt.Sprintf("± %d", i)
	} else {
		return fmt.Sprintf(" %d", i)
	}
}
