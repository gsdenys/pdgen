package translate

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func RegisterLanguages() {
	RegisterEN()
	RegisterPT()
	RegisterFR()
}

func GetTranslation(l string) *message.Printer {
	switch l {
	case "pt", "pt_BR":
		return message.NewPrinter(language.BrazilianPortuguese)
	case "fr", "fr_CA":
		return message.NewPrinter(language.CanadianFrench)
	default:
		return message.NewPrinter(language.AmericanEnglish)
	}
}
