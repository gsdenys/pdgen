package translate

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var T *message.Printer

const (
	defaultLang string = "en"
	defaultLoc  string = "US"
)

type LangRegister interface {
	Register()
}

func getLocaleWindows() (string, string) {
	cmd := exec.Command("powershell", "Get-Culture | select -exp Name")
	output, err := cmd.Output()
	if err == nil {
		langLocRaw := strings.TrimSpace(string(output))
		langLoc := strings.Split(langLocRaw, "-")
		lang := langLoc[0]
		loc := langLoc[1]

		return lang, loc
	}

	return defaultLang, defaultLang
}

func getLocaleDarwin() (string, string) {
	cmd := exec.Command("sh", "osascript -e 'user locale of (get system info)'")
	output, err := cmd.Output()
	if err == nil {
		langLocRaw := strings.TrimSpace(string(output))
		langLoc := strings.Split(langLocRaw, "_")
		lang := langLoc[0]
		loc := langLoc[1]

		return lang, loc
	}

	return defaultLang, defaultLoc
}

func getLocaleLinux() (string, string) {
	envlang, ok := os.LookupEnv("LANG")
	if ok {
		langLocRaw := strings.Split(envlang, ".")[0]
		langLoc := strings.Split(langLocRaw, "_")
		lang := langLoc[0]
		loc := langLoc[1]

		return lang, loc
	}

	return defaultLang, defaultLoc
}

func getLocale() (string, string) {
	osHost := runtime.GOOS
	switch osHost {
	case "windows":
		return getLocaleWindows()
	case "darwin":
		return getLocaleDarwin()
	case "linux":
		return getLocaleLinux()
	}

	return defaultLang, defaultLoc
}

func InitLanguage() {
	lang, loc := getLocale()

	SetTranslation(fmt.Sprintf("%s_%s", lang, loc))
}

func SetTranslation(l string) *message.Printer {
	switch l {
	case "pt", "pt_BR":
		RegisterPT()
		T = message.NewPrinter(language.BrazilianPortuguese)
	case "fr", "fr_CA":
		RegisterFR()
		T = message.NewPrinter(language.CanadianFrench)
	default:
		RegisterEN()
		T = message.NewPrinter(language.AmericanEnglish)
	}

	return T
}
