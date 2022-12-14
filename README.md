# PDGEN - Data Dictionary Generator

[![Build](https://github.com/gsdenys/pdgen/actions/workflows/main.yml/badge.svg)](https://github.com/gsdenys/pdgen/actions/workflows/main.yml) [![Coverage](https://sonarcloud.io/api/project_badges/measure?project=gsdenys_pdgen&metric=coverage)](https://sonarcloud.io/summary/new_code?id=gsdenys_pdgen) [![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=gsdenys_pdgen&metric=vulnerabilities)](https://sonarcloud.io/summary/new_code?id=gsdenys_pdgen) [![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=gsdenys_pdgen&metric=code_smells)](https://sonarcloud.io/summary/new_code?id=gsdenys_pdgen) [![GitHub release](https://img.shields.io/github/release/gsdenys/pdgen.svg)](https://GitHub.com/gsdenys/pdgen/releases/) [![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)


This software has as objective to resolve the problem of maintain the data dictionary actualized when work in agile project that use the migrations to create, update or delete tables, views and schemas.

> :pencil2: &nbsp; _Before start use this software is strongly recommended to have a well documented database, that are comments for all elements like columns, tables, views, schemas and database._


Basically this is an automatic data dictionary generator that use the informations present at the database to export well organized documents. Actually the output formats are:

* __Console:__ This is de default one that outputs directly over the operation system stdout;
* __TXT:__ Generate a TXT file containing the organized information;
* __HTML:__ Generate an HTML file containing the organized information;
* __Markdown (MD):__ Generate a MD file containing the organized information;
* **_DOCX:_** :construction_worker: _Previewed at version 2.0.0_ :construction:

The output can be done by now in the follow languages:

* __en:__ _English (default alias for en-US);_
* __en-US:__ _American english;_
* __fr:__ _French (default alias for fr-CA);_
* __fr-CA:__ _Canadian French;_
* __pt:__ _Portuguese (default alias for pt-BR);_
* __pt-BR:__ _Brazilian portuguese._


The Install and Usage documentation are disponible at: [PDGEN Home Page](https://gsdenys.github.io/pdgen)

