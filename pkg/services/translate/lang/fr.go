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
