package translate

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func RegisterPT() {
	message.SetString(language.BrazilianPortuguese, "title", "Dicionário de Dados")

	message.SetString(language.BrazilianPortuguese, "title-db", "Banco de dados %s")
	message.SetString(language.BrazilianPortuguese, "title-schema", "Esquema %s")
	message.SetString(language.BrazilianPortuguese, "title-tables", "Descrição das Tabelas")

	message.SetString(language.BrazilianPortuguese, "desc-tables", "O banco de dados %s, no que tange ao esquema %s, contém %d tabelas que estão detalhadas abaixo. Para cada tabela é apresentado seu nome, sua descrição, e uma tabela que contém a relação de todas as colunas mostrando o name, o tipo, e a descrição. Nos casos em que o tipo for um enum customizado, é mostrado também quais os possíveis valores a serem usados.")

	message.SetString(language.BrazilianPortuguese, "table-title-name", "Nome")
	message.SetString(language.BrazilianPortuguese, "table-title-type", "Tipo")
	message.SetString(language.BrazilianPortuguese, "table-title-allow", "Aceita")
	message.SetString(language.BrazilianPortuguese, "table-title-comment", "Comentário")
}
