package becmi

import (
	"becmi/attributes"
	"becmi/background"
	"becmi/classes"
	"becmi/dice"
	"becmi/localization"
	"becmi/magic"
	"becmi/proficiencies"
	"becmi/savingthrows"
	"becmi/weaponmastery"
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"strings"
	"text/template"
	"unicode"
)

type Character struct {
	Name         string
	Player       string
	Class        classes.Class
	Alignment    string
	XP           int
	Attributes   attributes.Attributes
	SavingThrows savingthrows.SavingThrows
	Movement     int
	ArmorClass   int
	HitPoints    int
	ThAC0        int
	Skills       proficiencies.Proficiencies
	Masteries    weaponmastery.WeaponMastery
	Sex          string
	Age          int
	AgeSpan      string
	Height       string
	Weight       int
	Background   background.Background
	Grimoire     *magic.Spellbook
	GP           int
}

type CharacterRecordSheet struct {
	Name              string   `yaml:"Name,omitempty"`
	Player            string   `yaml:"Player,omitempty"`
	Class             string   `yaml:"Class,omitempty"`
	Alignment         string   `yaml:"Alignment,omitempty"`
	Level             string   `yaml:"Level,omitempty"`
	XP                string   `yaml:"XP,omitempty"`
	NextLevel         int      `yaml:"NextLevel,omitempty"`
	STR               int      `yaml:"STR,omitempty"`
	DEX               int      `yaml:"DEX,omitempty"`
	CON               int      `yaml:"CON,omitempty"`
	INT               int      `yaml:"INT,omitempty"`
	WIS               int      `yaml:"WIS,omitempty"`
	CHA               int      `yaml:"CHA,omitempty"`
	STRMod            string   `yaml:"STR_Mod,omitempty"`
	DEXMod            string   `yaml:"DEX_Mod,omitempty"`
	CONMod            string   `yaml:"CON_Mod,omitempty"`
	INTMod            string   `yaml:"INT_Mod,omitempty"`
	INTModText        string   `yaml:"INT_ModText,omitempty"`
	WISMod            string   `yaml:"WIS_Mod,omitempty"`
	CHAMod            string   `yaml:"CHA_Mod,omitempty"`
	CHAModText        string   `yaml:"CHA_ModText,omitempty"`
	ArmorClass        int      `yaml:"AC,omitempty"`
	HitPoints         int      `yaml:"HP"`
	Movement          string   `yaml:"Move,omitempty"`
	ThAC0             int      `yaml:"THAC0,omitempty"`
	STPoison          int      `yaml:"ST_Poison,omitempty"`
	STWands           int      `yaml:"ST_Wands,omitempty"`
	STPetrification   int      `yaml:"ST_Petrification,omitempty"`
	STDragonBreath    int      `yaml:"ST_DragonBreath,omitempty"`
	STSpells          int      `yaml:"ST_Spells,omitempty"`
	AllowedArmor      string   `yaml:"AllowedArmor,omitempty"`
	AllowedWeapons    string   `yaml:"AllowedWeapons,omitempty"`
	Gender            string   `yaml:"Gender,omitempty"`
	Age               int      `yaml:"Age,omitempty"`
	Height            string   `yaml:"Height,omitempty"`
	Weight            int      `yaml:"Weight,omitempty"`
	Ethnicity         string   `yaml:"Ethnicity,omitempty"`
	Status            string   `yaml:"Status,omitempty"`
	Origin            string   `yaml:"Origin,omitempty"`
	Faith             string   `yaml:"Faith,omitempty"`
	WeaponSkills      []string `yaml:"WeaponSkills"`
	Languages         []string `yaml:"Languages"`
	Skills            []string `yaml:"Skills"`
	Spells            string   `yaml:"Spells,omitempty"`
	Descriptions      string   `yaml:"Descriptions,omitempty"`
	GP                int      `yaml:"GP,omitempty"`
	ClassAbilities    string   `yaml:"ClassAbilities,omitempty"`
	ClassDescriptions string   `yaml:"ClassDescriptions,omitempty"`
	SpellList         string   `yaml:"SpellList,omitempty"`
	SpellDescriptions string   `yaml:"SpellDescriptions,omitempty"`
}

func toUpperFirst(s string) string {
	// Convert the string to a rune slice, so that we deal with
	// runes, and not bytes.
	r := []rune(s)

	// Ensure that there is at least one rune in the slice, and
	// capitalize the first one.
	if len(r) > 0 {
		r[0] = unicode.ToUpper(r[0])
	}

	for i, char := range r {
		if char == '-' && i+1 < len(r) {
			r[i+1] = unicode.ToUpper(r[i+1])
		}
	}

	// Convert the rune slice back to a string and return.
	return string(r)
}

func NewCharacter(name, player, alignment, sex, class, campaign string, xp int) *Character {
	var ch Character
	ch.Name = name
	ch.Player = player
	ch.Alignment = toUpperFirst(alignment)
	ch.XP = xp
	ch.Sex = strings.ToLower(sex)
	ch.ArmorClass = 10
	class = toUpperFirst(class)
	campaign = toUpperFirst(campaign)

	for { // reroll until class requirements are met
		ch.Attributes = dice.RollAttributes()
		if c, ok := classes.Classes[class]; ok {
			if c.Requirement(ch.Attributes) {
				ch.Class = c.Load()
				//fmt.Println(ch.Class)
				break
			}
		}
	}

	ch.Age, ch.AgeSpan = classes.Age(ch.Class.Race())
	ch.Height, ch.Weight = classes.Body(ch.Class.Race(), ch.Sex)
	ch.Movement = ch.Class.BaseMovement()
	level := ch.Class.LevelIncludingRank(ch.XP)
	hd, hpinc := ch.Class.HitDice()
	hp := 0
	for idx := range level {
		if idx <= 9 {
			if con, ok := ch.Attributes["CON"]; ok {
				hproll := dice.RollDice(hd) + con.Modifier()
				if hproll <= 0 {
					hproll = 1
				}
				hp += hproll
			} else {
				fmt.Fprintf(os.Stderr, "Error: No CON attribute found\n")
			}
		} else {
			hp += hpinc
		}
	}
	ch.HitPoints = hp
	ch.SavingThrows = ch.Class.SavingThrows(xp)
	ch.ThAC0 = ch.Class.ThAC0(xp)

	switch campaign {
	case "Karameikos":
		ch.Background = background.NewBGKarameikos(ch.Class.Race(), ch.Class.Name())
	}

	ch.Grimoire = ch.Class.Grimoire(ch.XP)

	//var testxp, nextxp int
	//for testlevel := range 22 {
	//	nextxp = ch.Class.NextLevelAt(testxp + 1)
	//	fmt.Printf("%2d: %3d\n", testlevel, ch.Class.ThAC0(testxp+1))
	//	testxp = nextxp
	//}

	ch.GP = dice.Roll("3d6") * 10

	return &ch
}

func (c Character) String() string {
	abilities := c.Class.SpecialAbilities(c.XP)
	var classabilities, abilitydescriptions string
	for _, ab := range abilities {
		if ab.MinLevel <= c.Class.LevelIncludingRank(c.XP) {
			classabilities += ab.ListString() + "\n"
			abilitydescriptions += ab.DescriptionString() + "\n"
		}
	}

	fmt.Println("classabilities: ", classabilities)
	fmt.Println("abilitydescriptions: ", abilitydescriptions)

	outputMessage := &i18n.Message{
		ID:          "CharacterRecordSheet",
		Description: "Character record sheet",
		Other: "" +
			"Name:  %-32s \t Player:        %s\n" +
			"Class: %-32s \t Alignment:     %s\n" +
			"Level: %2d %c (max. %2d) \t XP: %8d (%s%%) \t Next Level at: %d\n" +
			"\n" +
			"%s\n" +
			"%s\n" +
			"Armor Class: %3d \t Hitpoints: %d\n" +
			"Movement:    %3d \t ThAC0:     %d\n" +
			"\n" +
			"%s\n" +
			"Armor Proficiencies:  %s\n" +
			"Weapon Proficiencies: %s\n" +
			"\n" +
			"Background\n" +
			"==========\n" +
			"Sex:    %-12s \t Age:    %4d (%s)\n" +
			"Height: %s \t\t Weight: %d cn (%d lb)\n" +
			"%s" +
			"Starting Money: %d GP\n\n" +
			"Special Abilities\n" +
			"=================\n" +
			"%s\n" +
			"%s\n" +
			"Descriptions\n" +
			"============\n" +
			"%s\n" +
			"%s\n",
	}

	translation := localization.Locale[localization.LanguageSetting].MustLocalize(&i18n.LocalizeConfig{DefaultMessage: outputMessage})

	output := fmt.Sprintf(translation,
		c.Name,
		c.Player,
		c.Class.Name(), c.Alignment,
		c.Class.Level(c.XP), c.Class.Rank(c.XP), c.Class.MaxLevel(), c.XP, attributes.SignedInt(c.Class.CheckXPModifier(c.Attributes)), c.Class.NextLevelAt(c.XP),
		c.Attributes,
		c.SavingThrows,
		c.ArmorClass, c.HitPoints,
		c.Movement, c.ThAC0,
		c.Class.ThAC0Table(c.XP),
		c.Class.ArmorProficiency(),
		c.Class.WeaponProficiency(),
		c.Sex, c.Age, c.AgeSpan,
		c.Height, c.Weight, c.Weight/10,
		c.Background,
		c.GP,
		classabilities,
		c.Class.SpellList(c.XP, c.Grimoire),
		abilitydescriptions,
		c.Class.SpellDescriptions(c.XP, c.Grimoire),
	)
	return output
}

func (c Character) Save() {
	file, err := os.Create(c.Name + ".yaml")
	if err != nil {
		log.Printf("Error creating file: %s\n", err)
		return
	}
	defer file.Close()

	encoder := yaml.NewEncoder(file)
	encoder.SetIndent(2)

	if err := encoder.Encode(c); err != nil {
		log.Printf("Error encoding character to YAML: %s\n", err)
		return
	}

	log.Printf("Character saved to %s.yaml\n", c.Name)
}

func (c Character) SaveObsidian() {
	var crs CharacterRecordSheet

	chaModMessage := &i18n.Message{
		ID:          "RetainerStats",
		Description: "Charisma Retainer Statistics",
		Other:       "%d max. Retainers, %d Retainer Morale",
	}
	translation := localization.Locale[localization.LanguageSetting].MustLocalize(&i18n.LocalizeConfig{DefaultMessage: chaModMessage})

	abilities := c.Class.SpecialAbilities(c.XP)
	var classabilities, abilitydescriptions string
	for _, ab := range abilities {
		if ab.MinLevel <= c.Class.LevelIncludingRank(c.XP) {
			classabilities += ab.ListString() + "\n"
			abilitydescriptions += ab.DescriptionString() + "\n"
		}
	}

	crs.Name = c.Name
	crs.Player = c.Player
	crs.Class = c.Class.Name()
	crs.Alignment = c.Alignment
	crs.Level = fmt.Sprintf("%d %c", c.Class.Level(c.XP), c.Class.Rank(c.XP))
	crs.NextLevel = c.Class.NextLevelAt(c.XP)
	crs.XP = fmt.Sprintf("%d (%s%%)", c.XP, attributes.SignedInt(c.Class.CheckXPModifier(c.Attributes)))
	crs.STR = c.Attributes["STR"].Value
	crs.DEX = c.Attributes["DEX"].Value
	crs.CON = c.Attributes["CON"].Value
	crs.INT = c.Attributes["INT"].Value
	crs.WIS = c.Attributes["WIS"].Value
	crs.CHA = c.Attributes["CHA"].Value
	crs.INTModText = c.Attributes["INT"].LanguageProficiency()
	crs.CHAModText = fmt.Sprintf(translation, c.Attributes["CHA"].MaxRetainers(), c.Attributes["CHA"].RetainerMorale())
	crs.STRMod = attributes.SignedInt(c.Attributes["STR"].Modifier())
	crs.DEXMod = attributes.SignedInt(c.Attributes["DEX"].Modifier())
	crs.CONMod = attributes.SignedInt(c.Attributes["CON"].Modifier())
	crs.INTMod = attributes.SignedInt(c.Attributes["INT"].Modifier())
	crs.WISMod = attributes.SignedInt(c.Attributes["WIS"].Modifier())
	crs.CHAMod = attributes.SignedInt(c.Attributes["CHA"].Modifier())
	crs.ArmorClass = c.ArmorClass
	crs.HitPoints = c.HitPoints
	crs.Movement = fmt.Sprintf("%de (%de)", c.Movement, c.Movement/3)
	crs.ThAC0 = c.ThAC0
	crs.STPoison = c.SavingThrows[0]
	crs.STWands = c.SavingThrows[1]
	crs.STPetrification = c.SavingThrows[2]
	crs.STDragonBreath = c.SavingThrows[3]
	crs.STSpells = c.SavingThrows[4]
	crs.AllowedArmor = c.Class.ArmorProficiency()
	crs.AllowedWeapons = c.Class.WeaponProficiency()
	crs.Gender = c.Sex
	crs.Age = c.Age
	crs.Height = c.Height
	crs.Weight = c.Weight
	crs.Ethnicity = c.Background.Ethnicity()
	crs.Status = c.Background.SocialStatus()
	crs.Origin = c.Background.Hometown()
	crs.Faith = c.Background.Faith()
	crs.GP = c.GP
	crs.WeaponSkills = []string{}
	crs.Languages = []string{}
	crs.Skills = []string{}
	for _, sk := range c.Skills.Proficiencies {
		crs.Skills = append(crs.Skills, sk.Name)
	}
	crs.ClassAbilities = classabilities
	crs.ClassDescriptions = abilitydescriptions

	crs.Spells = c.Class.SpellList(c.XP, c.Grimoire)
	crs.Descriptions = c.Class.SpellDescriptionsObsidian(c.XP, c.Grimoire)

	//fmt.Println("classabilities: ", classabilities)
	//fmt.Println("abilitydescriptions: ", abilitydescriptions)
	//fmt.Println("spells: ", c.Class.SpellList(c.XP, c.Grimoire))
	//fmt.Println("spelldescriptions: ", c.Class.SpellDescriptions(c.XP, c.Grimoire))

	file, err := os.Create(c.Name + ".md")
	if err != nil {
		log.Printf("Error creating file: %s\n", err)
	}
	defer file.Close()

	fmt.Fprintln(file, "---\ncssclasses:\n  - pathfinder\n  - table\n  - hide-header-underline")

	encoder := yaml.NewEncoder(file)
	encoder.SetIndent(2)

	if err := encoder.Encode(crs); err != nil {
		log.Printf("Error encoding character to YAML: %s\n", err)
		return
	}

	fmt.Fprintln(file, "---")

	templ := template.Must(template.ParseFiles("data/templates/character.template"))

	err = templ.ExecuteTemplate(file, "de/character", crs)
	if err != nil {
		log.Printf("Error executing template: %s\n", err)
	}

	log.Printf("Character saved to %s.md\n", c.Name)
}
