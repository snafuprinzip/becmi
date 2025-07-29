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
	AllArcaneSpells = ArcaneSpells{}
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
	AllDivineSpells = DivineSpells{}
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
	AllPrimalSpells = PrimalSpells{}
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

func (s Spell) String() string {
	var outputMessage *i18n.Message
	switch localization.OutputFormat {
	case localization.OutputFormatObsidian:
		outputMessage = &i18n.Message{
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
	case localization.OutputFormatText:
		fallthrough
	default:
		outputMessage = &i18n.Message{
			ID:          "Spell",
			Description: "Spell Description",
			Other: "" +
				"%s\n" +
				"Range: %s\n" +
				"Duration: %s\n" +
				"Effect: %s\n" +
				"%s\n",
		}
	}
	translation := localization.Locale[localization.LanguageSetting].MustLocalize(&i18n.LocalizeConfig{DefaultMessage: outputMessage})

	name := s.Name
	if s.Reversible {
		name += " \\*"
	}

	return fmt.Sprintf(translation, name, s.Range, s.Duration, s.Effect, s.Description)
}
