# TLDR

The easy way to use [PDGEN](https://gsdenys.github.io/pdgen) is to execute the command to export the documentation directly to console. Follow the steps below to a fast start.

!> These steps don't install the [PDGEN](https://gsdenys.github.io/pdgen) in your environment, it just disponibilize the executable in a local directory!

1. Download the binary compatible with yourn environment. For this example  __darwin-arm64__ (Mac OS).

```sh
curl -OL https://github.com/gsdenys/pdgen/releases/download/v1.0.0/pdgen-v1.0.0-darwin-arm64.tar.gz
tar -zxvf pdgen-v1.0.0-darwin-arm64.tar.gz
```

2. Execute the command below to show the database documentation in your console.

```sh
./pdgen describe -u "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
```

3. See the output, it should be something like the following block.

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

4. Explore more command options, execute the commands below to see these options.

```sh
./pdgen --help
```

and / or

```sh
./pdgen describe --help
```