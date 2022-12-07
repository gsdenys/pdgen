package lang

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func CanadianFrench(lang language.Tag) language.Tag {
	_ = message.SetString(lang, "title", "Dictionnaire de données")

	_ = message.SetString(lang, "title-db", "Base de données %s")
	_ = message.SetString(lang, "title-schema", "Schéma %s")
	_ = message.SetString(lang, "title-tables", "Description des tableaux")

	_ = message.SetString(lang, "desc-tables", "La base de données %s, au schéma %s, contient %d tables décrites ci-dessous. Pour chaque tableau est présenté leur nom, leur description, et un tableau décrivant chaque champ. Dans les cas où le champ a un type de données personnalisé, les options possibles s'afficheront dans la colonne Accepte")
	_ = message.SetString(lang, "table-title-name", "Nom")
	_ = message.SetString(lang, "table-title-type", "Type")
	_ = message.SetString(lang, "table-title-allow", "Accepte")
	_ = message.SetString(lang, "table-title-comment", "Commentaire")

	return lang
}
