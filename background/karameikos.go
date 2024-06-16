package background

import (
	"becmi/dice"
	"fmt"
	"strconv"
)

type Karameikos struct {
	Ethnicity    string
	SocialStatus string
	Titled       bool
	Hometown     string
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
	socialRoll := dice.RollDice(100) + 1
	k.SocialStatus, k.Titled = k.SocialStatusTable(race, socialRoll)

	ethnicityRoll := dice.RollDice(100) + 1 + socialRoll/2
	k.Ethnicity = k.EthnicityTable(race, ethnicityRoll)

	homeRoll := dice.RollDice(20) + 1
	k.Hometown = k.HometownTable(race, class, k.SocialStatus, homeRoll)

	return k
}

func (k Karameikos) Nation() string {
	return "Grandduchy of Karameikos"
}

func (k Karameikos) Nationality() string { return "Karameikan" }

//func (k Karameikos) Ethnicity() string { return k.ethnicity }
//
//func (k Karameikos) SocialStatus() (string, bool) { return k.socialStatus, k.titled }
//
//func (k Karameikos) Hometown() string { return k.hometown }

func (k Karameikos) SocialStatusTable(race string, roll int) (status string, titled bool) {
	if race == "Human" {
		switch {
		case roll < 31:
			status = "Penniless"
		case roll < 61:
			status = "Struggling"
		case roll < 76:
			status = "Comfortable"
		case roll < 86:
			status = "Wealthy"
		case roll < 96:
			status = "Wealthy"
			titled = true
		case roll < 98:
			status = "Very Wealthy"
		case roll < 100:
			status = "Very Wealthy"
			titled = true
		case roll == 100:
			status = "Royal Family"
		}
	} else if race == "Darf" || race == "Gnome" {
		switch {
		case roll < 31:
			status = "Struggling"
		case roll < 61:
			status = "Comfortable"
		case roll < 96:
			status = "Wealthy"
		case roll < 98:
			status = "Very Wealthy"
		case roll < 100:
			status = "Ruling"
		}
	} else if race == "Elf" {
		switch {
		case roll < 91:
			status = "Common"
		case roll < 101:
			status = "Lord"
		}
	} else if race == "Halfling" {
		switch {
		case roll < 21:
			status = "Penniless"
		case roll < 51:
			status = "Struggling"
		case roll < 96:
			status = "Comfortable"
		case roll < 100:
			status = "Wealthy"
		}
	}
	return status + " (" + strconv.Itoa(roll) + ")", titled
}

func (k Karameikos) EthnicityTable(race string, roll int) (ethnicity string) {
	if race == "Human" {
		switch {
		case roll < 71:
			ethnicity = "Traladaran"
		case roll < 91:
			ethnicity = "Mixed Traladaran/Thyatian"
		case roll >= 91:
			ethnicity = "Thyatian"
		}
	} else if race == "Dwarf" || race == "Gnome" {
		ethnicity = "Stronghollow Clan"
	} else if race == "Elf" {
		switch {
		case roll < 96:
			ethnicity = "Callarii Clan"
		case roll > 95:
			ethnicity = "Vyalia Clan"
		}
	} else if race == "Halfling" {
		ethnicity = "Hin"
	}
	return ethnicity + " (" + strconv.Itoa(roll) + ")"
}

func (k Karameikos) HometownTable(race, class, status string, roll int) (hometown string) {
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
			hometown = "Black Eagle Barony"
		case roll < 11:
			hometown = "Homestead"
		case roll < 14:
			hometown = "Village or Town (choice)"
		case roll < 17:
			hometown = "Kelvin"
		default:
			hometown = "Specularum"
		}
	} else if race == "Darf" || race == "Gnome" {
		hometown = "Highforge"
	} else if race == "Elf" {
		hometown = "Rifflian or Specularum (Callarii) or Woods (Vyalia)"
	} else if race == "Halfling" {
		hometown = "Any town, village or city"
	}
	return hometown + " (" + strconv.Itoa(roll) + ")"
}

func (k Karameikos) String() string {
	var titled string
	if k.Titled {
		titled = "(Titled)"
	} else {
		titled = ""
	}
	return fmt.Sprintf("Ethnicity: %s\nSocial Status: %s %s\nHome: %s\n",
		k.Ethnicity, k.SocialStatus, titled, k.Hometown)
}
