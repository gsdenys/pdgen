# Usage

The [PDGEN](https://gsdenys.github.io/pdgen), in the root level has 3 different commands (_describe_, _discovery_ and _version_). These commands are described in the next subsections.

## Describe

This command has objective to connect to the database and generate the data dictionary output in the selected format that by default is a _TXT_ expressed at the standard operational system output.

This command has 7 possible flags/configurations that can be used together. All mentioned parameters has their default defined, so, if you opt by do not set, the default value will be used to.

### URL

?> The default value is: **_postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"_**

This parameter represents the connection string to the database, it can be applied using tbe following options:

1. Using the name:  __--url__ _[connection string]_

```bash
pdgen describe --url "postgres://theuser:pwd123@some.url/somedb"
```

2. Using the shorthand:  __-u__ _[connection string]_

```bash
pdgen describe -u "postgres://theuser:pwd123@some.url/somedb"
```

### Database

?> The default value is: **_postgres_**

!> In the next version the __default database__ will change, It'll be extracted from the __URL__

This parameters allow user select the database to be described, it can be applied using tbe following options:

1. The name: __--database__

```bash
pdgen describe --database databasename
```

2. The shorthand: __-d__

```bash
pdgen describe -d databasename
```

### Schema

?> The default value is: **_public_**

This parameters allow user select the schema to be described, it can be applied using tbe following options:

1. The name: __--schema__

```bash
pdgen describe --schema schamaname
```

2. The shorthand: __-s__

```bash
pdgen describe -s schamaname
```

### Output

?> By default it can assume the **_output.\[txt | md | html\]_**, depending of the selected format.

This parameters allow user select the output file, it can be applied using tbe following options:

1. The name: __--output__

```bash
pdgen describe --output /path/to/your/file.txt
```

2. The shorthand: __-o__

```bash
pdgen describe -o /path/to/your/file.txt
```

### Format

?> The default output format is the __Console Output (stdout)__.

This parameters allow user select the output format for the documentation. The disponible types are:

* __console:__ This is de default one that outputs directly over the operation system stdout;
* __txt:__ Generate a TXT file containing the organized information;
* __html:__ Generate an HTML file containing the organized information;
* __md:__ Generate a Markdown file containing the organized information;
* **_DOCX:_** :construction_worker: _Previewed at version 2.0.0_ :construction:.

The usage of this parameter cam be applied using the following options:

1. The name: __--format__

```bash
pdgen describe --format txt
```

2. The shorthand: __-f__

```bash
pdgen describe -f txt
```

### Language

?> The default output language is the __system language__. Case this is not enable, the __en-US__ will be selected.

This parameters allow user overwrite the output language for the documentation. The disponible types are:

* __en:__ English (default alias for en-US);
* __en-US:__ American english;
* __fr:__ French (default value for fr-CA)
* __fr-CA:__ Canadian French;
* __pt:__ Portuguese (default alias for pt-BR)
* __pt-BR:__ Brazilian Portuguese;


The usage of this parameter cam be applied using the following options:

1. The name: __--language__

```bash
pdgen describe --language pt-BR
```

2. The shorthand: __-l__

```bash
pdgen describe -l pt-BR
```

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
