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

// BrazilianPortuguese language definition
func BrazilianPortuguese(lang language.Tag) language.Tag {
	_ = message.SetString(lang, "title", "Dicionário de Dados")

	_ = message.SetString(lang, "title-db", "Banco de dados %s")
	_ = message.SetString(lang, "title-schema", "Esquema %s")
	_ = message.SetString(lang, "title-tables", "Descrição das Tabelas")

	_ = message.SetString(lang, "desc-tables", "O banco de dados %s, no esquema %s, contém %d tabelas que estão descritas abaixo. Para cada tabela é apresentado o seu nome, a sua descrição e uma tabela com a descrição de cada campo. Nos casos em que o campo possui um tipo de dados personalizado, as opções possíveis serão exibidas na coluna Aceita.")

	_ = message.SetString(lang, "table-title-name", "Nome")
	_ = message.SetString(lang, "table-title-type", "Tipo")
	_ = message.SetString(lang, "table-title-allow", "Aceita")
	_ = message.SetString(lang, "table-title-comment", "Comentário")

	return lang
}
