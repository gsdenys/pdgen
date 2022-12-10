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
