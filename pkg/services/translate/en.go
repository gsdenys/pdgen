package translate

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type Lang struct {
	Lang language.Tag
}

func RegisterEN() {
	lang := language.AmericanEnglish

	defaultVerboseMessage := "\nYou also can use -v or --verbose to see more details"

	_ = message.SetString(lang, "connect-error", "error when try to connect to %s."+defaultVerboseMessage)
	_ = message.SetString(lang, "ping-error", "connection was created but ping fail, so no content is accessible."+defaultVerboseMessage)
	_ = message.SetString(lang, "db-not-found", "there no database named %s,"+defaultVerboseMessage)
	_ = message.SetString(lang, "schema-not-found", "there no schema named %s."+defaultVerboseMessage)
	_ = message.SetString(lang, "undefined-error", "I don't know whats happens, but the database query returns error."+defaultVerboseMessage)

	_ = message.SetString(language.AmericanEnglish, "title", "Data Dictionary")

	_ = message.SetString(language.AmericanEnglish, "title-db", "Database %s")
	_ = message.SetString(language.AmericanEnglish, "title-schema", "Schema %s")
	_ = message.SetString(language.AmericanEnglish, "title-tables", "Tables Descriptions")

	_ = message.SetString(language.AmericanEnglish, "desc-tables", "The database %s, at the schema %s, contem %d tables that are described bellow. For each table is presented their name, their description, and a table containing description of each column containing their name, type, and description. In the cases of the data type is a custom data type, the options is printed as an allow enum")

	_ = message.SetString(language.AmericanEnglish, "table-title-name", "Name")
	_ = message.SetString(language.AmericanEnglish, "table-title-type", "Type")
	_ = message.SetString(language.AmericanEnglish, "table-title-allow", "Allow")
	_ = message.SetString(language.AmericanEnglish, "table-title-comment", "Comment")
}
