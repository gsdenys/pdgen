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
package models

// Basic data structure to define the fields that just has name anda description
type Basic struct {
	Name string `json:"name"`
	Desc string `json:"description"`
}

// Columns data structure to define each columns of table
type Columns struct {
	Column  string `json:"column"`
	Type    string `json:"type"`
	Allow   string `json:"allow"`
	Comment string `json:"comment"`
}

// Table data structure that defines a table
type Table struct {
	Name    string    `json:"name"`
	Desc    string    `json:"description"`
	Columns []Columns `json:"columns"`
}

// Describe is the data structure description to the database and schema
type Describe struct {
	Database Basic   `json:"database"`
	Schema   Basic   `json:"schema"`
	Tables   []Table `json:"tables"`
}
