import os
import psycopg2
from sqlalchemy import create_engine
import pandas as pd

__sql_dict = '''
select 
    c.table_name as "Table",
    c.column_name as "Colonne",
    case 
        WHEN c.data_type = 'USER-DEFINED' THEN c.udt_name
        ELSE c.data_type
    end as "Type",
    case 
        WHEN c.data_type = 'USER-DEFINED' THEN (
            SELECT 
                string_agg(pg_enum.enumlabel::text, ', ')
            FROM pg_type 
                JOIN pg_enum 
                    ON pg_enum.enumtypid = pg_type.oid
            WHERE pg_type.typname = c.udt_name
            Group by pg_type.typname
        )
    end as "permet",
    col_description((c.table_schema||'.'||c.table_name)::regclass::oid, c.ordinal_position) as "Commentaire"
from information_schema.columns as c
WHERE c.table_schema = 'public'
order by "Table"
'''


def data() -> pd.DataFrame:
    conn = create_engine(os.environ['DATABASE_URL'])
    df = pd.read_sql(__sql_dict, conn)

    return df