import os
import configparser
from .pgsql import check_connection

__config_file = os.path.expanduser('~') + "/.pdgen"


def writer(cfg: configparser.ConfigParser):
    """Write file to the file system home folder. By default the file name is 
    .pdgen, but it can be change passing the parameter to this function

    Args:
        cfg (configparser.ConfigParser): The properties to be persisted
        name (str, optional): the properties file. Defaults to '.pdgen'.
    """
    with open(__config_file, 'w') as file:
        cfg.write(file)


def read() -> configparser.ConfigParser:
    """Read config file. By default the config file is .pdgen, that is stored at
    the home directory. The name and path can be changed as long as the home 
    directory be the base directory 

    Args:
        name (str, optional): the config file name. Defaults to '.pdgen'.

    Returns:
        configparser.ConfigParser: the config property
    """
    if not os.path.exists(__config_file):
        raise FileNotFoundError(
            "No configuration present. Create a connection and try again"
        )

    config = configparser.ConfigParser()
    config.read(__config_file)

    return config


def select_connection(name: str = 'DEFAULT'):
    """Select the a predefined database connection

    Args:
        name (str, optional): the database connection name. Defaults to 'default'.

    Raises:
        Exception: Has no connection with the name passed by parameter
    """
    cfg = read()

    print([x for x in cfg.keys()], name.upper())

    if name.upper() not in cfg.keys():
        raise Exception(
            "Connection named {} is not defined.".format(name.upper())
        )

    # may be it already exist
    try:
        cfg.add_section('selected')
    except configparser.DuplicateSectionError as e:
        pass

    cfg['selected']['conn'] = name.upper()

    writer(cfg)


def get_connection() -> str:
    """Get the selected connection name

    Returns:
        str: the selected connection name
    """
    cfg = read()

    if 'selected' not in cfg.keys():
        raise Exception(
            "None selected connection or connection is not defined yet."
        )

    return cfg['selected']['conn']


def add_connection(url: str, name: str = 'DEFAULT'):
    """Add a new connection to the connection base

    Args:
        url (str): The connection URL
        name (str, optional): The connection name. Defaults to 'default'.

    Raises:
        Exception: Database connection error
    """
    if not check_connection(url):
        raise Exception("Database connection fail. Check the url and try again.")
    
    cfg = None
    
    try:
        cfg = read()
    except:
        cfg = configparser.ConfigParser()
    
    if name not in cfg.keys():
        cfg.add_section(name.upper())
    
    cfg[name.upper()]['conn'] = url
    
    writer(cfg)


def remove_connection(name: str):
    """Remove connection from configuration

    Args:
        name (str): _description_

    Raises:
        Exception: _description_
    """
    cfg = read()
    
    if cfg.remove_section(name.upper()):
        if 'selected' in cfg.keys() and cfg['selected']['conn'] == name.upper():
            cfg.remove_section('selected')
    else:
        raise Exception("Configuration not found")
    
    writer(cfg)
