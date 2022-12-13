/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/
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

// GetKeys returns a string array containing all defined languages
func GetKeys() []string {
	var ret []string

	for index := range RegLang {
		ret = append(ret, index)
	}

	sort.Strings(ret)

	return ret
}

// getLocale returns the system language. Case not detected, it returns AmericanEnglish
func getLocale() language.Tag {
	tag, err := locale.Detect()
	if err == nil {
		return tag
	}

	return language.AmericanEnglish
}

// InitLanguage starts the processo to select and detect the system language
func InitLanguage() {
	lang := getLocale()

	if _, ok := RegLang[lang.String()]; ok {
		T = message.NewPrinter(RegLang[lang.String()])
		return
	}

	fmt.Printf("The language %s is not registered. Using %s\n", lang.String(), language.AmericanEnglish.String())
	T = message.NewPrinter(language.AmericanEnglish)
}

// SetLanguage enable the user select the language manually
func SetLanguage(lang string) bool {
	for index := range RegLang {
		if lang == RegLang[index].String() {
			T = message.NewPrinter(RegLang[index])
			return true
		}
	}

	return false
}

// Register all disponibles languages
func Register() {
	RegLang[language.English.String()] = lang.AmericanEnglish(language.English)
	RegLang[language.AmericanEnglish.String()] = lang.AmericanEnglish(language.AmericanEnglish)

	RegLang[language.Portuguese.String()] = lang.BrazilianPortuguese(language.Portuguese)
	RegLang[language.BrazilianPortuguese.String()] = lang.BrazilianPortuguese(language.BrazilianPortuguese)

	RegLang[language.French.String()] = lang.CanadianFrench(language.French)
	RegLang[language.CanadianFrench.String()] = lang.CanadianFrench(language.CanadianFrench)

	// In case of new language creation register it here
}
