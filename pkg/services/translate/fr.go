package translate

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func RegisterFR() {
	message.SetString(language.CanadianFrench, "title", "Dictionnaire de données")

	message.SetString(language.CanadianFrench, "title-db", "Base de données %s")
	message.SetString(language.CanadianFrench, "title-schema", "Schéma %s")
	message.SetString(language.CanadianFrench, "title-tables", "Description des tableaux")

	message.SetString(language.CanadianFrench, "desc-tables", "La base de données %s, au schéma %s, contient %d tables décrites ci-dessous. Pour chaque tableau est présenté leur nom, leur description, et un tableau contenant la description de chaque colonne contenant leur nom, leur type et leur description. Dans les cas où le type de données est un type de données personnalisé, les options sont imprimées sous forme d'énumération autorisée")

	message.SetString(language.CanadianFrench, "table-title-name", "Nom")
	message.SetString(language.CanadianFrench, "table-title-type", "Type")
	message.SetString(language.CanadianFrench, "table-title-allow", "Accepte")
	message.SetString(language.CanadianFrench, "table-title-comment", "Commentaire")
}
