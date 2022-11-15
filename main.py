import configparser
import click

from pdgen.config import set_connection, read

@click.group()
@click.option('--debug/--no-debug')
def cli(debug):
    click.echo(f"Debug mode is {'on' if debug else 'off'}")

@cli.command()
@click.option('--username')
def greet(username):
    click.echo(f"Hello {username}!")

@cli.command()
@click.option('-n', '--name', 'name')
@click.option('-u', '--url', 'url')
def connect(name, url):
    set_connection(name, url)
    

@cli.command()
@click.option('-s', '--show') 
def connect():
    print(read())



if __name__ == '__main__':
    cli(auto_envvar_prefix='GREETER')