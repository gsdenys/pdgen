import os
from sqlalchemy import create_engine as ce
import pandas as pd


__sql_column_dict = '''
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


def create_engine(url: str):
    """Function simple that just call de create engine with postgresql driver with the
    objetive of encapsulate the database iteration in this package in order to avoid
    direct iteration by others.

    Args:
        url (str): the connection url 

    Returns:
        engine : the generated engine
    """
    return ce(url)


def check_connection(url: str) -> bool:
    """Check if the url can generate a funcional connection

    Args:
        url (str): the connection URL

    Returns:
        bool: True if the connection can be generated
    """
    try:
        engine = create_engine(url)
        engine.connect().close()

        return True
    except:
        return False


def get_columns_dict() -> pd.DataFrame:
    conn = create_engine(os.environ['DATABASE_URL'])
    df = pd.read_sql(__sql_column_dict, conn)

    return df
