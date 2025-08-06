package attributes

import (
	"becmi/localization"
	"fmt"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

var AttributeIndices = [6]string{"STR", "INT", "WIS", "DEX", "CON", "CHA"}

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
	var msg *i18n.Message
	switch {
	case a.Value >= 2 && a.Value <= 3:
		msg = &i18n.Message{
			ID:          "cannot_speak",
			Description: "Cannot speak",
			Other:       "Has trouble speaking, cannot read or write",
		}
	case a.Value >= 4 && a.Value <= 5:
		msg = &i18n.Message{
			ID:          "cannot_read_write",
			Description: "Cannot read or write",
			Other:       "Cannot read or write Common",
		}
	case a.Value >= 6 && a.Value <= 8:
		msg = &i18n.Message{
			ID:          "can_write_simple",
			Description: "Can write simple",
			Other:       "Can write simple common words",
		}
	case a.Value >= 9:
		msg = &i18n.Message{
			ID:          "can_write",
			Description: "Can read and write",
			Other:       "Reads and writes (usually two) native languages",
		}
	default:
		return "Error: Value out of bounds!"
	}

	output := localization.Locale[localization.LanguageSetting].MustLocalize(&i18n.LocalizeConfig{DefaultMessage: msg})
	return output
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
	outputMessage := &i18n.Message{
		ID:          "Attributes",
		Description: "Character Attributes",
		Other: "" +
			"STR Strength      %2d \t (%s Attack Roll [Melee, Unarmed], Damage Roll [Melee, Thrown], Open Doors, optional: Save: Paralysis, Turn to Stone)\n" +
			"INT Intelligence  %2d \t (%s General Skills, %s, optional: Save vs. Mind Attacks)\n" +
			"WIS Wisdom        %2d \t (%s Saving Throw vs. Spells)\n" +
			"DEX Dexterity     %2d \t (%s Attack Roll [Thrown, Missiles], Armor Class, optional: Save: Wands, Dragon Breath)\n" +
			"CON Constitution  %2d \t (%s Hit Points per XP Level, optional: Save: Poison)\n" +
			"CHA Charisma      %2d \t (%s Reaction Adjustment from NPCs, %d max. Retainers, %d Retainer Morale)\n",
	}

	translation := localization.Locale[localization.LanguageSetting].MustLocalize(&i18n.LocalizeConfig{DefaultMessage: outputMessage})

	return fmt.Sprintf(translation,
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
