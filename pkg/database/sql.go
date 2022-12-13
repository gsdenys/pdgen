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
package database

const selectDatabaseComment string = `
SELECT 
	pg_catalog.shobj_description(d.oid, 'pg_database') AS "description"
FROM pg_catalog.pg_database d 
WHERE  datname = $1`

const selectSchemaComment string = `SELECT obj_description($1::regnamespace, 'pg_namespace')`

const selectAllTables string = `
SELECT 
	table_name AS "name",
	CASE 
		WHEN obj_description((c.table_schema||'.'||c.table_name)::regclass, 'pg_class') is NULL THEN ''
		ELSE  obj_description((c.table_schema||'.'||c.table_name)::regclass, 'pg_class')
	END  AS "comment"
FROM information_schema.tables AS c
WHERE c.table_schema = $1`

const selectTable string = `
SELECT 
    c.column_name AS "colonne",
    CASE 
        WHEN c.data_type = 'USER-DEFINED' THEN c.udt_name
        ELSE c.data_type
    END AS "type",
    CASE 
        WHEN c.data_type = 'USER-DEFINED' THEN (
            SELECT 
                string_agg(pg_enum.enumlabel::text, ', ')
            FROM pg_type 
                JOIN pg_enum 
                    ON pg_enum.enumtypid = pg_type.oid
            WHERE pg_type.typname = c.udt_name
            GROUP BY pg_type.typname
        )
		ELSE ''
    END AS "allow",
    CASE
		WHEN col_description((c.table_schema||'.'||c.table_name)::regclass::oid, c.ordinal_position) IS NULL THEN '' 
		ELSE col_description((c.table_schema||'.'||c.table_name)::regclass::oid, c.ordinal_position)
	END AS "comment"
FROM information_schema.columns AS c
WHERE c.table_schema = $1 AND c.table_name = $2`
