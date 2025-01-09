package background

import (
	"becmi/dice"
	"becmi/localization"
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"strconv"
)

type Karameikos struct {
	ethnicity    string
	socialStatus string
	titled       bool
	hometown     string
	faith        string
}

func init() {
	BackgroundIndices = append(BackgroundIndices, "Karameikos")
	if Backgrounds == nil {
		Backgrounds = make(map[string]Background)
	}
	Backgrounds["Karameikos"] = Karameikos{}
}

func NewBGKarameikos(race, class string) Karameikos {
	var k Karameikos
	var ethnicityRoll int

	socialRoll := dice.RollDice(100) + 1
	k.socialStatus, k.titled = k.SocialStatusTable(race, socialRoll)

	if race == "Human" {
		ethnicityRoll = dice.RollDice(100) + 1 + socialRoll/2
	} else {
		ethnicityRoll = dice.RollDice(100) + 1
	}
	k.ethnicity = k.EthnicityTable(race, ethnicityRoll)

	homeRoll := dice.RollDice(20) + 1
	k.hometown = k.HometownTable(race, class, k.socialStatus, homeRoll)

	faithRoll := dice.RollDice(100) + 1
	k.faith = k.FaithTable(race, class, faithRoll)

	return k
}

func (k Karameikos) Nation() string {
	nation := &i18n.Message{
		ID:          "Karameikos",
		Description: "Nation: Karameikos",
		Other:       "Grandduchy of Karameikos",
	}
	nationMsg := localization.Locale[localization.LanguageSetting].MustLocalize(&i18n.LocalizeConfig{DefaultMessage: nation})

	return nationMsg
}

func (k Karameikos) Nationality() string { return "Karameikan" }

func (k Karameikos) Ethnicity() string { return k.ethnicity }

func (k Karameikos) SocialStatus() string {
	output := k.socialStatus
	titledString := &i18n.Message{
		ID:          "titled",
		Description: "Social Status: Titled",
		Other:       "Titled",
	}

	titledMsg := localization.Locale[localization.LanguageSetting].MustLocalize(&i18n.LocalizeConfig{DefaultMessage: titledString})

	if k.titled {
		output += "(" + titledMsg + ")"
	}
	return output
}

func (k Karameikos) Hometown() string { return k.hometown }

func (k Karameikos) SocialStatusTable(race string, roll int) (status string, titled bool) {

	penniless := &i18n.Message{
		ID:          "Penniless",
		Description: "Social Status: Penniless",
		Other:       "Penniless",
	}
	pennilessMsg := localization.Locale[localization.LanguageSetting].MustLocalize(&i18n.LocalizeConfig{DefaultMessage: penniless})

	struggling := &i18n.Message{
		ID:          "Struggling",
		Description: "Social Status: Struggling",
		Other:       "Struggling",
	}
	strugglingMsg := localization.Locale[localization.LanguageSetting].MustLocalize(&i18n.LocalizeConfig{DefaultMessage: struggling})

	comfortable := &i18n.Message{
		ID:          "Comfortable",
		Description: "Social Status: Comfortable",
		Other:       "Comfortable",
	}
	comfortableMsg := localization.Locale[localization.LanguageSetting].MustLocalize(&i18n.LocalizeConfig{DefaultMessage: comfortable})

	wealthy := &i18n.Message{
		ID:          "Wealthy",
		Description: "Social Status: Wealthy",
		Other:       "Wealthy",
	}
	wealthyMsg := localization.Locale[localization.LanguageSetting].MustLocalize(&i18n.LocalizeConfig{DefaultMessage: wealthy})

	veryWealthy := &i18n.Message{
		ID:          "Very Wealthy",
		Description: "Social Status: Very Wealthy",
		Other:       "Very Wealthy",
	}
	veryWealthyMsg := localization.Locale[localization.LanguageSetting].MustLocalize(&i18n.LocalizeConfig{DefaultMessage: veryWealthy})

	royalFamily := &i18n.Message{
		ID:          "Royal Family",
		Description: "Social Status: Royal Family",
		Other:       "Royal Family",
	}
	royalFamilyMsg := localization.Locale[localization.LanguageSetting].MustLocalize(&i18n.LocalizeConfig{DefaultMessage: royalFamily})

	ruling := &i18n.Message{
		ID:          "Ruling",
		Description: "Social Status: Ruling",
		Other:       "Ruling",
	}
	rulingMsg := localization.Locale[localization.LanguageSetting].MustLocalize(&i18n.LocalizeConfig{DefaultMessage: ruling})

	lord := &i18n.Message{
		ID:          "Lord",
		Description: "Social Status: Lord",
		Other:       "Lord",
	}
	lordMsg := localization.Locale[localization.LanguageSetting].MustLocalize(&i18n.LocalizeConfig{DefaultMessage: lord})

	common := &i18n.Message{
		ID:          "Common",
		Description: "Social Status: Common",
		Other:       "Common",
	}
	commonMsg := localization.Locale[localization.LanguageSetting].MustLocalize(&i18n.LocalizeConfig{DefaultMessage: common})

	if race == "Human" {
		switch {
		case roll < 31:
			status = pennilessMsg
		case roll < 61:
			status = strugglingMsg
		case roll < 76:
			status = comfortableMsg
		case roll < 86:
			status = wealthyMsg
		case roll < 96:
			status = wealthyMsg
			titled = true
		case roll < 98:
			status = veryWealthyMsg
		case roll < 100:
			status = veryWealthyMsg
			titled = true
		case roll == 100:
			status = royalFamilyMsg
		}
	} else if race == "Dwarf" || race == "Gnome" {
		switch {
		case roll < 31:
			status = strugglingMsg
		case roll < 61:
			status = comfortableMsg
		case roll < 96:
			status = wealthyMsg
		case roll < 98:
			status = veryWealthyMsg
		case roll < 100:
			status = rulingMsg
		}
	} else if race == "Elf" {
		switch {
		case roll < 91:
			status = commonMsg
		case roll < 101:
			status = lordMsg
		}
	} else if race == "Halfling" {
		switch {
		case roll < 21:
			status = pennilessMsg
		case roll < 51:
			status = strugglingMsg
		case roll < 96:
			status = comfortableMsg
		case roll < 100:
			status = wealthyMsg
		}
	}
	return status + " (" + strconv.Itoa(roll) + ")", titled
}

func (k Karameikos) EthnicityTable(race string, roll int) (ethnicity string) {
	traladaran := &i18n.Message{
		ID:          "Traladaran",
		Description: "Ethnicity: Traladaran",
		Other:       "Traladaran",
	}
	traladaranMsg := localization.Locale[localization.LanguageSetting].MustLocalize(&i18n.LocalizeConfig{DefaultMessage: traladaran})

	thyatian := &i18n.Message{
		ID:          "Thyatian",
		Description: "Ethnicity: Thyatian",
		Other:       "Thyatian",
	}
	thyatianMsg := localization.Locale[localization.LanguageSetting].MustLocalize(&i18n.LocalizeConfig{DefaultMessage: thyatian})

	mixedTraladaranThyatian := &i18n.Message{
		ID:          "Mixed Traladaran/Thyatian",
		Description: "Ethnicity: Mixed Traladaran/Thyatian",
		Other:       "Mixed Traladaran/Thyatian",
	}
	mixedMsg := localization.Locale[localization.LanguageSetting].MustLocalize(&i18n.LocalizeConfig{DefaultMessage: mixedTraladaranThyatian})

	stronghollowClan := &i18n.Message{
		ID:          "Stronghollow Clan",
		Description: "Ethnicity: Stronghollow Clan",
		Other:       "Stronghollow Clan",
	}
	stronghollowMsg := localization.Locale[localization.LanguageSetting].MustLocalize(&i18n.LocalizeConfig{DefaultMessage: stronghollowClan})

	callariiClan := &i18n.Message{
		ID:          "Callarii Clan",
		Description: "Ethnicity: Callarii Clan",
		Other:       "Callarii Clan",
	}
	callariiMsg := localization.Locale[localization.LanguageSetting].MustLocalize(&i18n.LocalizeConfig{DefaultMessage: callariiClan})

	vyaliaClan := &i18n.Message{
		ID:          "Vyalia Clan",
		Description: "Ethnicity: Vyalia Clan",
		Other:       "Vyalia Clan",
	}
	vyaliaMsg := localization.Locale[localization.LanguageSetting].MustLocalize(&i18n.LocalizeConfig{DefaultMessage: vyaliaClan})

	if race == "Human" {
		switch {
		case roll < 71:
			ethnicity = traladaranMsg
		case roll < 91:
			ethnicity = mixedMsg
		case roll >= 91:
			ethnicity = thyatianMsg
		}
	} else if race == "Dwarf" || race == "Gnome" {
		ethnicity = stronghollowMsg
	} else if race == "Elf" {
		switch {
		case roll < 96:
			ethnicity = callariiMsg
		case roll > 95:
			ethnicity = vyaliaMsg
		}
	} else if race == "Halfling" {
		ethnicity = "Hin"
	}
	return ethnicity + " (" + strconv.Itoa(roll) + ")"
}

func (k Karameikos) HometownTable(race, class, status string, roll int) (hometown string) {
	blackeagle := &i18n.Message{
		ID:          "Black Eagle Barony",
		Description: "Hometown: Black Eagle Barony",
		Other:       "Black Eagle Barony",
	}
	blackeagleMsg := localization.Locale[localization.LanguageSetting].MustLocalize(&i18n.LocalizeConfig{DefaultMessage: blackeagle})

	homestead := &i18n.Message{
		ID:          "Homestead",
		Description: "Hometown: Homestead",
		Other:       "Homestead",
	}
	homesteadMsg := localization.Locale[localization.LanguageSetting].MustLocalize(&i18n.LocalizeConfig{DefaultMessage: homestead})

	village := &i18n.Message{
		ID:          "Village or Town (choice)",
		Description: "Hometown: Village or Town (choice)",
		Other:       "Village or Town (choice)",
	}
	villageMsg := localization.Locale[localization.LanguageSetting].MustLocalize(&i18n.LocalizeConfig{DefaultMessage: village})

	kelvin := &i18n.Message{
		ID:          "Kelvin",
		Description: "Hometown: Kelvin",
		Other:       "Kelvin",
	}
	kelvinMsg := localization.Locale[localization.LanguageSetting].MustLocalize(&i18n.LocalizeConfig{DefaultMessage: kelvin})

	specularum := &i18n.Message{
		ID:          "Specularum",
		Description: "Hometown: Specularum",
		Other:       "Specularum",
	}
	specularumMsg := localization.Locale[localization.LanguageSetting].MustLocalize(&i18n.LocalizeConfig{DefaultMessage: specularum})

	highforge := &i18n.Message{
		ID:          "Highforge",
		Description: "Hometown: Highforge",
		Other:       "Highforge",
	}
	highforgeMsg := localization.Locale[localization.LanguageSetting].MustLocalize(&i18n.LocalizeConfig{DefaultMessage: highforge})

	rifflian := &i18n.Message{
		ID:          "Rifflian or Specularum",
		Description: "Hometown: Elf",
		Other:       "Rifflian or Specularum (Callarii) or Dymrak Woods (Vyalia)",
	}
	rifflianMsg := localization.Locale[localization.LanguageSetting].MustLocalize(&i18n.LocalizeConfig{DefaultMessage: rifflian})

	anyTown := &i18n.Message{
		ID:          "Any town, village or city",
		Description: "Hometown: Any town, village or city",
		Other:       "Any town, village or city",
	}
	anyTownMsg := localization.Locale[localization.LanguageSetting].MustLocalize(&i18n.LocalizeConfig{DefaultMessage: anyTown})

	if race == "Human" {
		if class == "Cleric" {
			roll += 2
		}
		if class == "Magic-User" {
			roll += 4
		}
		if status == "Comfortable" {
			roll += 2
		} else if status == "Wealthy" {
			roll += 4
		} else if status == "Very Wealthy" {
			roll += 6
		}

		switch {
		case roll < 4:
			hometown = blackeagleMsg
		case roll < 11:
			hometown = homesteadMsg
		case roll < 14:
			hometown = villageMsg
		case roll < 17:
			hometown = kelvinMsg
		default:
			hometown = specularumMsg
		}
	} else if race == "Dwarf" || race == "Gnome" {
		hometown = highforgeMsg
	} else if race == "Elf" {
		hometown = rifflianMsg
	} else if race == "Halfling" {
		hometown = anyTownMsg
	}
	return hometown + " (" + strconv.Itoa(roll) + ")"
}

func (k Karameikos) FaithTable(race, class string, roll int) (faith string) {
	churchTraladara := &i18n.Message{
		ID:          "Church of Traladara",
		Description: "Church of Traladara",
		Other:       "Church of Traladara",
	}
	churchTraladaraMsg := localization.Locale[localization.LanguageSetting].MustLocalize(&i18n.LocalizeConfig{DefaultMessage: churchTraladara})

	churchKarameikos := &i18n.Message{
		ID:          "Church of Karameikos",
		Description: "Church of Karameikos",
		Other:       "Church of Karameikos",
	}
	churchKarameikosMsg := localization.Locale[localization.LanguageSetting].MustLocalize(&i18n.LocalizeConfig{DefaultMessage: churchKarameikos})

	cultHalav := &i18n.Message{
		ID:          "Cult of Halav",
		Description: "Cult of Halav",
		Other:       "Cult of Halav",
	}
	cultHalavMsg := localization.Locale[localization.LanguageSetting].MustLocalize(&i18n.LocalizeConfig{DefaultMessage: cultHalav})

	if race == "Human" {
		switch {
		case roll < 71:
			faith = churchTraladaraMsg
		case roll < 96:
			faith = churchKarameikosMsg
		case roll >= 96:
			faith = cultHalavMsg
		}
	} else if race == "Dwarf" {
		faith = "Kagyar"
	} else if race == "Gnome" {
		faith = "Garal Glitterlode"
	} else if race == "Elf" {
		if k.Ethnicity() == "Vyalia Clan" {
			faith = "Ordana"
		} else {
			faith = "Ilsundal"
		}
	} else if race == "Halfling" {
		faith = "Nob Nar"
	}

	return faith
}

func (k Karameikos) Faith() string {
	return k.faith
}

func (k Karameikos) String() string {
	var titled string

	titledString := &i18n.Message{
		ID:          "titled",
		Description: "Social Status: Titled",
		Other:       "Titled",
	}

	titledMsg := localization.Locale[localization.LanguageSetting].MustLocalize(&i18n.LocalizeConfig{DefaultMessage: titledString})

	if k.titled {
		titled = "(" + titledMsg + ")"
	} else {
		titled = ""
	}

	msg := &i18n.Message{
		ID:          "karameikos_background",
		Description: "Karameikos Background",
		Other:       "Ethnicity: %s\nSocial Status: %s %s\nHome: %s\n",
	}

	output := localization.Locale[localization.LanguageSetting].MustLocalize(&i18n.LocalizeConfig{DefaultMessage: msg})

	return fmt.Sprintf(output, k.ethnicity, k.socialStatus, titled, k.hometown)
}
