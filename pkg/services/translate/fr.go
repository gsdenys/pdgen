package translate

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func RegisterFR() {
	_ = message.SetString(language.CanadianFrench, "title", "Dictionnaire de données")

	_ = message.SetString(language.CanadianFrench, "title-db", "Base de données %s")
	_ = message.SetString(language.CanadianFrench, "title-schema", "Schéma %s")
	_ = message.SetString(language.CanadianFrench, "title-tables", "Description des tableaux")

	_ = message.SetString(language.CanadianFrench, "desc-tables", "La base de données %s, au schéma %s, contient %d tables décrites ci-dessous. Pour chaque tableau est présenté leur nom, leur description, et un tableau contenant la description de chaque colonne contenant leur nom, leur type et leur description. Dans les cas où le type de données est un type de données personnalisé, les options sont imprimées sous forme d'énumération autorisée")

	_ = message.SetString(language.CanadianFrench, "table-title-name", "Nom")
	_ = message.SetString(language.CanadianFrench, "table-title-type", "Type")
	_ = message.SetString(language.CanadianFrench, "table-title-allow", "Accepte")
	_ = message.SetString(language.CanadianFrench, "table-title-comment", "Commentaire")
}
