package localization

import (
	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var LanguageBundle *i18n.Bundle
var Locale map[string]*i18n.Localizer

var LanguageSetting string

func init() {
	LanguageBundle = i18n.NewBundle(language.English)

	LanguageBundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	LanguageBundle.MustLoadMessageFile("active.de.toml")

	// Step 2: Create localizer for that bundle using one or more language tags
	Locale = make(map[string]*i18n.Localizer)
	Locale["de"] = i18n.NewLocalizer(LanguageBundle, language.German.String())
	Locale["en"] = i18n.NewLocalizer(LanguageBundle, language.English.String())
}

func GenerateMessage(id, desc, msg string) *i18n.Message {
	outputMessage := &i18n.Message{
		ID:          id,
		Description: desc,
		Other:       msg,
	}

	return outputMessage
}
