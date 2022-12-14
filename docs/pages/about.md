# About

The [PDGEN](https://gsdenys.github.io/pdgen) software has as objective to resolve the problem of maintain the data dictionary actualized principally when works in agile project that use the migrations to create, update or delete tables, views and schemas.

Together with the [PDGEN](https://gsdenys.github.io/pdgen) development there is a development software proposition that has focus on a high database documentation during the development phase. It means that any changes that should occur in the database must be documented in the database level. 

__e.g.__ you should include another field in the database, so you must to add the follow command at your migrations:

```sql
COMMENT ON COLUMN new_column_name IS 'Some description of the new column';
```

If this behavior is employed at all database levels for example database, schema, tables, columns, views, etc. You can just run the [PDGEN](https://gsdenys.github.io/pdgen) point to a database that contains all mentioned data and a structured documentation will be exported in the chosen format.

The [Example](pages/example.md) page are dedicated to list many example repositories wrote in different programing languages.


## Supported Databases

In this first version of [PDGEN](https://gsdenys.github.io/pdgen) (v1.0.0), the unique database supported is [PostgreSQL](https://www.postgresql.org), but this cited just below are previewed for the version 2.0.0.

* MySQL
* MariaDB
* MSSQL
* Oracle

If you want to propose another database fell free to create a new issue at the github repository page. TKS.

## Output Formats

For a while, the following types are disponibles for documente exportation, it was chosen due their creation simplicity and the high level of utilization.

* __Console:__ This is de default one that outputs directly over the operation system stdout;
* __TXT:__ Generate a TXT file containing the organized information;
* __HTML:__ Generate an HTML file containing the organized information;
* __Markdown (MD):__ Generate a MD file containing the organized information;

There are another 2 types that has high usability, however, their development spend more time then I should dispend in the first version. So these two showed below also are previewed for de 2.0.0 version.

* __DOCX__
* __PDF__
