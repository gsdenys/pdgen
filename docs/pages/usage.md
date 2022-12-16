# Usage

The [PDGEN](https://gsdenys.github.io/pdgen), in the root level has 3 different commands (_describe_, _discovery_ and _version_). These commands are described in the next subsections.

## Describe

This command has objective to connect to the database and generate the data dictionary output in the selected format that by default is a _TXT_ expressed at the standard operational system output.

This command has 7 possible flags/configurations that can be user together and combined.


### URL

### Database

### Schema

### Format

### Output

### Language

### Help

This flag, that can be expressed as __-h__ or __--help__ has the priority against others. To show the help menu you can use the command below:

```bash
pdgen describe --help
```

The output will be

```
Connect to the database and generate the data dictionary output in the selected format that by default is a txt expressed at the standard output.

Usage:
  pdgen describe [flags]

Flags:
  -d, --database string   The database to be described (default "postgres")
  -f, --format string     the output types [default html json md txt] (default "DEFAULT")
  -h, --help              help for describe
  -l, --language string   The language selected to the output file
  -o, --output string     The output file path
  -s, --schema string     The schema to be described (default "public")
  -u, --uri string        The database connection uri (default "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
```

## Discovery

This command follow the same principle and rule of the [describe](#Describe), as their parameters as. Their objective is list all databases and perform the [describe](#Describe) for each schema, creating in the end a full and complete database high level documentation.

There are 2 ways to export the documentation using this command: Single and Multiples documents.

In the first one, Just one document is generate, on that is present all databases and schemas visibles for the provided user.

The second one generates many document, each one is composed by one database and one schema. Basically the Multiple documents way perform many [describe](#Describe) command for each database and schema that the provided user can see.

!> Unfortunately this command is __Not Implemented yet__. It is planned for __3.0.0 version__.


## Version

This command shows the software version, to see this execute the command below.

```bash
pdgen version
```

Their output should be something like.

```bash
pdgen v1.0.0
```
