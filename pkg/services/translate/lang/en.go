package lang

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func AmericanEnglish(lang language.Tag) language.Tag {
	_ = message.SetString(lang, "title", "Data Dictionary")

	_ = message.SetString(lang, "title-db", "Database %s")
	_ = message.SetString(lang, "title-schema", "Schema %s")
	_ = message.SetString(lang, "title-tables", "Tables Descriptions")

	_ = message.SetString(lang, "desc-tables", "The database %s, at the schema %s, contem %d tables that are described bellow. For each table is presented their name, their description, and a table describing each field. In the cases when the field has a custom data type, the possible options will show at the column allow")

	_ = message.SetString(lang, "table-title-name", "Name")
	_ = message.SetString(lang, "table-title-type", "Type")
	_ = message.SetString(lang, "table-title-allow", "Allow")
	_ = message.SetString(lang, "table-title-comment", "Comment")

	return lang
}
