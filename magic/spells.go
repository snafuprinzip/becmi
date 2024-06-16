package magic

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type Spell struct {
	Name        string
	Level       int
	Reversible  bool
	Range       string
	Duration    string
	Effect      string
	Description string
}

type SpellLevelList []Spell

type DivineSpells [7]SpellLevelList
type ArcaneSpells [9]SpellLevelList
type PrimalSpells DivineSpells

type Spellbook [9][]string

var AllDivineSpells DivineSpells
var AllArcaneSpells ArcaneSpells
var AllPrimalSpells PrimalSpells

func loadArcaneSpells(languageCode string) {
	var dirName string
	if languageCode == "" || languageCode == "en" {
		dirName = path.Join("data", "spells", "arcane")
	} else {
		dirName = path.Join("data", "spells", "arcane", languageCode)
	}

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

		AllArcaneSpells[spell.Level] = append(AllArcaneSpells[spell.Level], spell)
	}
}

func loadDivineSpells(languageCode string) {
	var dirName string
	if languageCode == "" || languageCode == "en" {
		dirName = path.Join("data", "spells", "divine")
	} else {
		dirName = path.Join("data", "spells", "divine", languageCode)
	}

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

		AllDivineSpells[spell.Level] = append(AllDivineSpells[spell.Level], spell)
	}
}

func loadPrimalSpells(languageCode string) {
	var dirName string
	if languageCode == "" || languageCode == "en" {
		dirName = path.Join("data", "spells", "primal")
	} else {
		dirName = path.Join("data", "spells", "primal", languageCode)
	}

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

		AllPrimalSpells[spell.Level] = append(AllPrimalSpells[spell.Level], spell)
	}
}

func LoadSpells(languageCode string) {
	loadArcaneSpells(languageCode)
	loadDivineSpells(languageCode)
	loadPrimalSpells(languageCode)
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
	return fmt.Sprintf(""+
		"%s\n"+
		"Range: %s\n"+
		"Duration: %s\n"+
		"Effect: %s\n"+
		"%s\n", s.Name, s.Range, s.Duration, s.Effect, s.Description)
}
