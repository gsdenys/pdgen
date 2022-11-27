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
