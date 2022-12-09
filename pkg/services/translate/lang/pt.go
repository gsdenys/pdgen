package lang

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

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
