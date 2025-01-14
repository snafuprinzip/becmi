package localization

import (
	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var LanguageBundle *i18n.Bundle
var Locale map[string]*i18n.Localizer

var LanguageSetting string // global setting for the output language to be used
var OutputFormat int       // global setting for output format

const (
	OutputFormatText = iota
	OutputFormatHTML
	OutputFormatYAML
	OutputFormatJSON
	OutputFormatObsidian
)

var SupportedLanguages = []string{"de", "en"}

var Translations map[string]map[string]string

func init() {
	LanguageBundle = i18n.NewBundle(language.English)

	LanguageBundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	LanguageBundle.MustLoadMessageFile("active.de.toml")

	// Step 2: Create localizer for that bundle using one or more language tags
	Locale = make(map[string]*i18n.Localizer)
	Locale["de"] = i18n.NewLocalizer(LanguageBundle, language.German.String())
	Locale["en"] = i18n.NewLocalizer(LanguageBundle, language.English.String())

	Translations = make(map[string]map[string]string)

	for _, lang := range SupportedLanguages {
		Translations[lang] = make(map[string]string)
	}

	lawful := &i18n.Message{
		ID:          "lawful",
		Description: "Lawful",
		Other:       "Lawful",
	}

	neutral := &i18n.Message{
		ID:          "neutral",
		Description: "Neutral",
		Other:       "Neutral",
	}

	chaotic := &i18n.Message{
		ID:          "chaotic",
		Description: "Chaotic",
		Other:       "Chaotic",
	}

	male := &i18n.Message{
		ID:          "male",
		Description: "Male",
		Other:       "Male",
	}

	female := &i18n.Message{
		ID:          "female",
		Description: "Female",
		Other:       "Female",
	}

	other := &i18n.Message{
		ID:          "other",
		Description: "Other",
		Other:       "Other",
	}

	ac := &i18n.Message{
		ID:          "ac",
		Description: "AC",
		Other:       "AC",
	}

	hp := &i18n.Message{
		ID:          "hp",
		Description: "HP",
		Other:       "HP",
	}

	roll := &i18n.Message{
		ID:          "roll",
		Description: "Roll",
		Other:       "Roll",
	}

	for _, lang := range SupportedLanguages {
		Translations[lang]["lawful"] = Locale[lang].MustLocalize(&i18n.LocalizeConfig{DefaultMessage: lawful})
		Translations[lang]["neutral"] = Locale[lang].MustLocalize(&i18n.LocalizeConfig{DefaultMessage: neutral})
		Translations[lang]["chaotic"] = Locale[lang].MustLocalize(&i18n.LocalizeConfig{DefaultMessage: chaotic})
		Translations[lang]["male"] = Locale[lang].MustLocalize(&i18n.LocalizeConfig{DefaultMessage: male})
		Translations[lang]["female"] = Locale[lang].MustLocalize(&i18n.LocalizeConfig{DefaultMessage: female})
		Translations[lang]["other"] = Locale[lang].MustLocalize(&i18n.LocalizeConfig{DefaultMessage: other})
		Translations[lang]["ac"] = Locale[lang].MustLocalize(&i18n.LocalizeConfig{DefaultMessage: ac})
		Translations[lang]["hp"] = Locale[lang].MustLocalize(&i18n.LocalizeConfig{DefaultMessage: hp})
		Translations[lang]["roll"] = Locale[lang].MustLocalize(&i18n.LocalizeConfig{DefaultMessage: roll})
	}

}

func GenerateMessage(id, desc, msg string) *i18n.Message {
	outputMessage := &i18n.Message{
		ID:          id,
		Description: desc,
		Other:       msg,
	}

	return outputMessage
}
