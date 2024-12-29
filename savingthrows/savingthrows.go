package savingthrows

import (
	"becmi/localization"
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type SavingThrows [5]int
type ByLevel [36]SavingThrows

func (s SavingThrows) String() string {
	outputMessage := &i18n.Message{
		ID:          "SavingThrows",
		Description: "Character Saving Throws",
		Other: "" +
			"Death Ray, Poison         %2d\n" +
			"Magic Wands               %2d\n" +
			"Paralysis, Turn to Stone  %2d\n" +
			"Dragon Breath             %2d\n" +
			"Rod, Staff, Spells        %2d\n",
	}

	translation := localization.Locale[localization.LanguageSetting].MustLocalize(&i18n.LocalizeConfig{DefaultMessage: outputMessage})

	return fmt.Sprintf(translation, s[0], s[1], s[2], s[3], s[4])
}
