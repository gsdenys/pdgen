from pdgen import config
import click


@click.group()
def cli():
    pass


@cli.command()
@click.option('-u', '--url', 'url')
@click.option('-n', '--name', 'name')
def add(url, name):

    if name is None:
        name = "DEFAULT"

    if not config.check_connection(url):
        click.echo(f'Unable to connect using the provided URL')
        return

    if config.add_connection(url, name) is False:
        click.echo(f'Unable to add a connection with URL and Name provided.')
        return

    click.echo(f'Connection created successfully')


@cli.command()
@click.argument('connection')
def rm(connection):
    if config.remove_connection(connection) is False:
        click.echo(
            'Unable to remove connection. Check the connection name and try again.')
        return

    click.echo(f'Connection removed successfully.')


if __name__ == '__main__':
    cli()
