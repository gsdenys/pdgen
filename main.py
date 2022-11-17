from pdgen import config
from pdgen import connection as conn
import click


def __print_msg(msg: str):
    click.echo('\n' + msg + '\n')


@click.group()
def cli():
    pass


@cli.group()
def connection():
    pass


@connection.command()
@click.option('-u', '--url', 'url')
@click.option('-n', '--name', 'name')
def add(url, name):
    __print_msg(conn.add(url, name))


@connection.command()
@click.argument('name')
def rm(name):
    __print_msg(conn.rm(name))


@connection.command()
def list():
    __print_msg(conn.lst())


@connection.command()
@click.argument('name')
def use(name):
    __print_msg(conn.use(name=name))


# @cli.command()
# @click.option('-f', '--format', 'format')
# def describe(format: str = 'default'):
#     pass


if __name__ == '__main__':
    cli()
