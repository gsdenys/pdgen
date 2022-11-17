import os
from pdgen import connection as conn
from pdgen import config

url_success = 'postgresql+psycopg2://postgres:postgres@localhost:5432/postgres'
url_error = 'postgresql+psycopg2://idontknow:idontknow@thiisnoturl:5432/database'

def drop_conf():
    try:
        directory = os.path.expanduser('~') + '/.pdgen'
        os.remove(directory)
    except:
        pass
    

def test_add():
    drop_conf()
    
    res = "Unable to connect using the provided URL.\nurl={}".format(url_error)
    add = conn.add(url=url_error, name='default')
    assert res == add
    
    res = "Connection created successfully.\n(name:{}, url:{})".format('default', url_success)
    add = conn.add(url=url_success, name='default')
    assert res == add
    

def test_rm():
    drop_conf()
    
    msg = 'Unable to {} a connection with URL and Name provided. '
    msg += 'It looks like a permission problem at the .pdgen file.\n{}'
    msg = msg.format('remove', config.config_file)
    res = conn.rm('noconn')
    assert res == msg
    
    msg = "There's no connection named '{}'.\n".format('NOCONN')
    msg += "Use 'pdgen connection list' to list all connections."
    conn.add(url_success, 'test')
    res = conn.rm('noconn')
    assert res == msg
    
    
    msg = 'Connection removed successfully.\n'
    msg += "Use 'pdgen connection list' to list all connections."
    res = conn.rm('test')
    assert res == msg