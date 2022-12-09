package translate

import (
	"fmt"
	"sort"

	"github.com/Xuanwo/go-locale"
	"github.com/gsdenys/pdgen/pkg/services/translate/lang"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var T *message.Printer

var RegLang map[string]language.Tag = make(map[string]language.Tag)

func GetKeys() []string {
	var ret []string

	for index := range RegLang {
		ret = append(ret, index)
	}

	sort.Strings(ret)

	return ret
}

func getLocale() language.Tag {
	tag, err := locale.Detect()
	if err == nil {
		return tag
	}

	return language.AmericanEnglish
}

func InitLanguage() {
	lang := getLocale()

	if _, ok := RegLang[lang.String()]; ok {
		T = message.NewPrinter(RegLang[lang.String()])
		return
	}

	fmt.Printf("The language %s is not registered. Using %s\n", lang.String(), language.AmericanEnglish.String())
	T = message.NewPrinter(language.AmericanEnglish)
}

func SetLanguage(lang string) bool {
	for index := range RegLang {
		if lang == RegLang[index].String() {
			T = message.NewPrinter(RegLang[index])
			return true
		}
	}

	return false
}

func Register() {
	RegLang[language.English.String()] = lang.AmericanEnglish(language.English)
	RegLang[language.AmericanEnglish.String()] = lang.AmericanEnglish(language.AmericanEnglish)

	RegLang[language.Portuguese.String()] = lang.BrazilianPortuguese(language.Portuguese)
	RegLang[language.BrazilianPortuguese.String()] = lang.BrazilianPortuguese(language.BrazilianPortuguese)

	RegLang[language.French.String()] = lang.CanadianFrench(language.French)
	RegLang[language.CanadianFrench.String()] = lang.CanadianFrench(language.CanadianFrench)
}
