# -*- coding: utf-8 -*-
from setuptools import setup

packages = \
['pdgen']

package_data = \
{'': ['*']}

install_requires = \
['click>=8.1.3,<9.0.0',
 'pandas>=1.5.1,<2.0.0',
 'postgres>=4.0,<5.0',
 'psycopg2-binary>=2.9.5,<3.0.0',
 'sqlalchemy>=1.4.44,<2.0.0',
 'tabulate>=0.9.0,<0.10.0']

setup_kwargs = {
    'name': 'pdgen',
    'version': '0.1.0',
    'description': 'Data dictionary generator for postgresql',
    'long_description': '# PostgreSQL Data Dictionary Generator\n\n## Command Line\n\n\n## Web API\n\n\n## Web UI\n\n\npytest --cov-report lcov:coverage/lcov.info tests/',
    'author': 'Denys G. Santos',
    'author_email': 'gsdenys@gmail.com',
    'maintainer': 'None',
    'maintainer_email': 'None',
    'url': 'None',
    'packages': packages,
    'package_data': package_data,
    'install_requires': install_requires,
    'python_requires': '>=3.10,<4.0',
}


setup(**setup_kwargs)
