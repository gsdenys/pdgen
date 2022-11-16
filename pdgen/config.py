import os
import configparser
from pdgen.pgsql import check_connection

__config_file = os.path.expanduser('~') + "/.pdgen"
__selected = 'SELECTED'
__conn = 'conn'


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
    """Read the .pdgen config file that are stored at the user home directory.

    Returns:
        configparser.ConfigParser: the config property or None
    """
    if not os.path.exists(__config_file):
        return None

    config = configparser.ConfigParser()
    config.read(__config_file)

    return config


def select_connection(name: str = 'DEFAULT') -> bool:
    """Select the a predefined database connection

    Args:
        name (str, optional): the database connection name. Defaults to 'default'.

    Raises:
        Exception: Has no connection with the name passed by parameter
        
    Returns:
        bool: True case operation successfully, other else False
    """
    cfg = read()
    
    if cfg is None: return False
    if name.upper() not in cfg.keys(): return False

    # Case this is the first time that a database is selected 
    if __selected not in cfg.keys():
        cfg.add_section(__selected)

    # Select the database
    cfg[__selected][__conn] = name.upper()

    writer(cfg)
    
    return True


def get_connection() -> str:
    """Get the selected connection name

    Returns:
        str: the selected connection name
    """
    cfg = read()

    if cfg is None: 
        return None
    
    if __selected not in cfg.keys(): 
        return None

    return cfg[__selected][__conn]


def add_connection(url: str, name: str = 'DEFAULT') -> bool:
    """Add a new connection to the connection base

    Args:
        url (str): The connection URL
        name (str, optional): The connection name. Defaults to 'default'.

    Raises:
        Exception: Database connection error
    """
    if not check_connection(url): return False

    cfg = configparser.ConfigParser()

    # Read the config file if it exists
    if os.path.exists(__config_file):
       cfg = read()

    if name not in cfg.keys():
        cfg.add_section(name.upper())

    cfg[name.upper()][__conn] = url

    writer(cfg)


def remove_connection(name: str) -> bool:
    """Remove connection from configuration

    Args:
        name (str): connection name

    Returns:
        bool: True for successful execution, other else False
    """
    cfg = read()
    
    if cfg is None: return False
    
    if cfg.remove_section(name.upper()) is False: return False 

    # Case the removed connection is the selected connection
    if __selected in cfg.keys() and cfg[__selected][__conn] == name.upper():
        cfg.remove_section(__selected)

    writer(cfg)
    
    return True
