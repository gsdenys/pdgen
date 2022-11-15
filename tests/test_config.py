from pdgen import config
import configparser
import os
import pytest
from os.path import exists

url_success = 'postgresql+psycopg2://postgres:postgres@localhost:5432/postgres'
url_error = 'postgresql+psycopg2://idontknow:idontknow@thiisnoturl:5432/database'


def drop_conf():
    try:
        directory = os.path.expanduser('~') + '/.pdgen'
        os.remove(directory)
    except:
        pass


def test_writer():
    drop_conf()

    cfg = configparser.ConfigParser()

    cfg.add_section('test')
    cfg['test']['name'] = 'test'

    config.writer(cfg)

    file = open(os.path.expanduser('~') + '/.pdgen', mode='r')
    all_of_it = file.read()
    file.close()

    assert all_of_it == "[test]\nname = test\n\n"

    drop_conf()


def test_read():
    # test read file not existent
    drop_conf()
    with pytest.raises(FileNotFoundError):
        config.read()

    # Write a file
    cfg = configparser.ConfigParser()

    cfg.add_section('test')
    cfg['test']['name'] = 'test'

    config.writer(cfg)

    # rest reading
    dt = config.read()
    assert dt['test']['name'] == 'test'

    drop_conf()


def test_select_connection():
    config.add_connection(url_success, name='test')

    # file with no selection
    config.select_connection('test')
    res = config.read()

    assert res['TEST']['conn'] == url_success
    assert res['selected']['conn'] == 'TEST'

    with pytest.raises(Exception):
        config.select_connection('some_test')


def test_get_connection():
    drop_conf()

    # with connection not defined yet
    with pytest.raises(Exception):
        config.get_connection('test')

    # Write a file
    cfg = configparser.ConfigParser()

    cfg.add_section('test')
    cfg['test']['name'] = 'test'

    config.writer(cfg)

    with pytest.raises(Exception):
        config.get_connection('test')

    # file with selection
    cfg.add_section('selected')
    cfg['selected']['conn'] = 'test'
    config.writer(cfg)

    assert config.get_connection() == 'test'


def test_add_connection():
    drop_conf()

    with pytest.raises(Exception):
        config.add_connection(url_error)

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

    # Create the second connection
    config.add_connection(url_success, 'test')

    # Create the third connection
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

    with pytest.raises(Exception):
        cfg['selected']['conn']

    with pytest.raises(Exception):
        config.remove_connection('invalid_data')
