# Get Started

[PDGEN](https://gsdenys.github.io/pdgen) is a data dictionary exporter that allow generate your database documentation using reverse engineering.

The easy way to use [PDGEN](https://gsdenys.github.io/pdgen) is to execute the command to export the _database_ documentation directly to console. Follow the steps below to a fast start.

!> These steps don't install the [PDGEN](https://gsdenys.github.io/pdgen) in your environment, it just disponibilize the executable in a local directory!

## Download

1. Download the binary compatible with yourn environment. For this example  __darwin-arm64__ (Mac OS).

```bash
curl -OL https://github.com/gsdenys/pdgen/releases/download/v1.0.0/pdgen-v1.0.0-darwin-arm64.tar.gz
tar -zxvf pdgen-v1.0.0-darwin-arm64.tar.gz
```

## Execute

Execute the command below to show the database documentation in your console.

```bash
./pdgen describe -u "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
```

Now your can see the output, it should be something like the following block.

```
DATA DICTIONARY

DATABASE POSTGRES
standard public database

SCHEMA PUBLIC
standard public schema

TABLES DESCRIPTIONS
The database postgres, at the schema public, contem 1 tables that are described bellow. For each table is presented their name, their description, and a table describing each field. In the cases when the field has a custom data type, the possible options will show at the column allow

TEST
table for test propose

Name  Type     Allow  Comment                       
id    integer         sequencial unique identifier  
name  text            name of test                  
```

## Explore

Explore more command options, execute the commands below to see these options.

```bash
./pdgen --help
```

and / or

```bash
./pdgen describe --help
```