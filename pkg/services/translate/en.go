package translate

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func RegisterEN() {
	message.SetString(language.AmericanEnglish, "title-db", "Data Dictionary for database %s")
	message.SetString(language.AmericanEnglish, "title-schema", "Schema %s")
	message.SetString(language.AmericanEnglish, "title-tables", "Tables Descriptions")

	message.SetString(language.AmericanEnglish, "desc-tables", "The database %s, at the schema %s, contem %d tables that are described bellow. For each table is presented their name, their description, and a table containing description of each column containing their name, type, and description. In the cases of the data type is a custom dada type, the options is printed as an allow  enum")

	message.SetString(language.AmericanEnglish, "table-title-name", "Name")
	message.SetString(language.AmericanEnglish, "table-title-type", "Type")
	message.SetString(language.AmericanEnglish, "table-title-allow", "Allow")
	message.SetString(language.AmericanEnglish, "table-title-comment", "Comment")
}
