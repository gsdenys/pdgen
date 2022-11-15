from ..pdgen import pgsql
import pytest


def test_create_engine():
    url_success = 'postgresql+psycopg2://postgres:postgres@localhost:5432/postgres'
    url_error = 'postgresql+psycopg2://idonknow:idontknow@thiisnoturl:5432/database'

    engine = pgsql.create_engine(url_success)
    conn = engine.connect()
    
    assert conn is not None
    conn.close()

    with pytest.raises(Exception):
        engine = pgsql.create_engine(url_error)
        engine.connect()
        

def test_check_connection():
    url_success = 'postgresql+psycopg2://postgres:postgres@localhost:5432/postgres'
    url_error = 'postgresql+psycopg2://idonknow:idontknow@thiisnoturl:5432/database'
    
    assert pgsql.check_connection(url_success) is True
    assert pgsql.check_connection(url_error) is False
