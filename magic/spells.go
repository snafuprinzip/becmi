package magic

import (
	"becmi/localization"
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"gopkg.in/yaml.v3"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type Spell struct {
	ID          string `yaml:"id"`
	Name        string `yaml:"name"`
	Level       int    `yaml:"level"`
	Reversible  bool   `yaml:"reversible"`
	Range       string `yaml:"range"`
	Duration    string `yaml:"duration"`
	Effect      string `yaml:"effect"`
	Description string `yaml:"description"`
}

type SpellLevelList []Spell

type DivineSpells [7]SpellLevelList
type ArcaneSpells [9]SpellLevelList
type PrimalSpells DivineSpells

type Spellbook [9][]string

var AllDivineSpells DivineSpells
var AllArcaneSpells ArcaneSpells
var AllPrimalSpells PrimalSpells

func loadArcaneSpells() {
	dirName := path.Join("data", "spells", "arcane", localization.LanguageSetting)

	files, err := os.ReadDir(dirName)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, file := range files {
		if strings.ToLower(filepath.Ext(file.Name())) != ".yaml" {
			continue
		}

		data, err := os.ReadFile(path.Join(dirName, file.Name()))
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}

		var spell Spell
		err = yaml.Unmarshal(data, &spell)
		if err != nil {
			fmt.Println("Error unmarshaling YAML:", err)
			return
		}

		AllArcaneSpells[spell.Level-1] = append(AllArcaneSpells[spell.Level-1], spell)
	}
	//for _, level := range AllArcaneSpells {
	//	for idx, spell := range level {
	//		fmt.Printf("%2d: %s (%s)\n", idx+1, spell.ID, spell.Name)
	//	}
	//}
}

func loadDivineSpells() {
	dirName := path.Join("data", "spells", "divine", localization.LanguageSetting)

	files, err := os.ReadDir(dirName)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, file := range files {
		if strings.ToLower(filepath.Ext(file.Name())) != ".yaml" {
			continue
		}

		data, err := os.ReadFile(path.Join(dirName, file.Name()))
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}

		var spell Spell
		err = yaml.Unmarshal(data, &spell)
		if err != nil {
			fmt.Println("Error unmarshaling YAML:", err)
			return
		}

		AllDivineSpells[spell.Level-1] = append(AllDivineSpells[spell.Level-1], spell)
	}
}

func loadPrimalSpells() {
	dirName := path.Join("data", "spells", "primal", localization.LanguageSetting)

	files, err := os.ReadDir(dirName)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, file := range files {
		if strings.ToLower(filepath.Ext(file.Name())) != ".yaml" {
			continue
		}

		data, err := os.ReadFile(path.Join(dirName, file.Name()))
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}

		var spell Spell
		err = yaml.Unmarshal(data, &spell)
		if err != nil {
			fmt.Println("Error unmarshaling YAML:", err)
			return
		}

		AllPrimalSpells[spell.Level-1] = append(AllPrimalSpells[spell.Level-1], spell)
	}
}

func LoadSpells() {
	loadArcaneSpells()
	loadDivineSpells()
	loadPrimalSpells()
}

//func SaveSpell() {
//	spell := Spell{Name: "Cure Light Wounds", Level: 1, Reversible: true, Range: "Touch", Duration: "Permanent",
//		Effect: "Any one living Creature", Description: "" +
//			"This spell either heals damage or removes paralysis. If used to heal, it can cure 2-7 (1d6 + 1) points of damage. " +
//			"It cannot heal damage if used to cure paralysis. The cleric may cast it on himself if desired.\n" +
//			"This spell cannot increase a creature's total hit points above the original amount.\n" +
//			"When reversed, this spell, cause light wounds, causes 1d6+1 (2-7) points of damage to any creature or character " +
//			"touched (no saving throw is allowed). The cleric must make a normal attack roll to inflict this damage."}
//	out, err := yaml.Marshal(spell)
//	if err != nil {
//		fmt.Fprintln(os.Stderr, err)
//		return
//	}
//	err = os.WriteFile("data/spells/cure_light_wounds.yaml", out, 0644)
//	if err != nil {
//		fmt.Fprintln(os.Stderr, err)
//		return
//	}
//	fmt.Println(string(out))
//}

func (s Spell) String() string {
	outputMessage := &i18n.Message{
		ID:          "Spell",
		Description: "Spell Description",
		Other: "" +
			"%s\n" +
			"Range: %s\n" +
			"Duration: %s\n" +
			"Effect: %s\n" +
			"%s\n",
	}

	translation := localization.Locale[localization.LanguageSetting].MustLocalize(&i18n.LocalizeConfig{DefaultMessage: outputMessage})

	return fmt.Sprintf(translation, s.Name, s.Range, s.Duration, s.Effect, s.Description)
}

func (s Spell) ObsidianString() string {
	outputMessage := &i18n.Message{
		ID:          "Spell Obsidian",
		Description: "Spell Description for Obsidian",
		Other: "" +
			"##### %s\n" +
			"|              |      |\n" +
			"| :----------- | ---: |\n" +
			"| **Range**    | %s |\n" +
			"| **Duration** | %s |\n" +
			"| ** Effect**  | %s |\n\n" +
			"%s\n",
	}

	translation := localization.Locale[localization.LanguageSetting].MustLocalize(&i18n.LocalizeConfig{DefaultMessage: outputMessage})

	name := s.Name
	if s.Reversible {
		name += " \\*"
	}

	return fmt.Sprintf(translation, name, s.Range, s.Duration, s.Effect, s.Description)
}
