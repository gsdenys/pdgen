from pdgen import config
import configparser
import os
import pytest

url_success = 'postgresql+psycopg2://postgres:postgres@localhost:5432/postgres'
url_error = 'postgresql+psycopg2://idontknow:idontknow@thiisnoturl:5432/database'
__selected = 'SELECTED'
__conn = 'conn'

def drop_conf():
    try:
        directory = os.path.expanduser('~') + '/.pdgen'
        os.remove(directory)
    except:
        pass


def test_writer():
    drop_conf()

    # Create file using writer function
    cfg = configparser.ConfigParser()

    conn_name = 'test'
    cfg.add_section(conn_name.upper())
    cfg[conn_name.upper()][__conn] = url_success

    config.writer(cfg)

    # Read file and asserts that everything is ok
    file = open(os.path.expanduser('~') + '/.pdgen', mode='r')
    all_of_it = file.read()
    file.close()

    assert all_of_it == "[{}]\n{} = {}\n\n".format(
        conn_name.upper(),
        __conn,
        url_success
    )

    drop_conf()


def test_read():
    drop_conf()
    
    assert config.read() is None 

    conn_name = 'test'
    config.add_connection(url_success, conn_name)

    dt = config.read()
    assert dt[conn_name.upper()][__conn] == url_success


def test_select_connection():
    drop_conf()
    
    conn_name = 'test'
    config.add_connection(url_success, conn_name)
    assert config.select_connection(conn_name)
    
    # Check if function save data correctly at the file 
    res = config.read()
    assert res[conn_name.upper()][__conn] == url_success
    assert res[__selected][__conn] == conn_name.upper()

    # assert the fail when try to set a non existente connection
    assert config.select_connection('some_test') is False


def test_get_connection():
    drop_conf()

    # Config file not existent should return None 
    assert config.get_connection() is None

    # An empty config file should return None 
    cfg = configparser.ConfigParser()
    config.writer(cfg)
        
    assert config.get_connection() is None

    # file with successfully connection selection should return an URL
    config.add_connection(url_success, 'test')
    config.select_connection('test')

    assert config.get_connection() == 'TEST'


def test_add_connection():
    drop_conf()

    assert config.add_connection(url_error) is False

    # Add first connection
    config.add_connection(url_success)
    cfg = config.read()

    assert len(cfg.keys()) == 1

    # Overwrite the default
    config.add_connection(url_success)
    cfg = config.read()

    assert len(cfg.keys()) == 1

    # Create the second connection
    config.add_connection(url_success, name='test')
    cfg = config.read()

    assert len(cfg.keys()) == 2


def test_remove_connection():
    drop_conf()

    url_success = 'postgresql+psycopg2://postgres:postgres@localhost:5432/postgres'

    config.add_connection(url_success)
    config.add_connection(url_success, 'test')
    config.add_connection(url_success, name='more_test')
    
    config.remove_connection('test')
    
    cfg = config.read()
    assert len(cfg.keys()) == 2

    # Create test again
    config.add_connection(url_success, name='test')
    config.select_connection('test')
    config.remove_connection('test')

    cfg = config.read()
    assert len(cfg.keys()) == 2

    assert __selected not in cfg.keys() 

    assert config.remove_connection('invalid_data') is False
