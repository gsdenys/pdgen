# DICIONÁRIO DE DADOS PARA O BANCO POSTGRES

standard public database

## ESQUEMA PUBLIC
standard public schema

## DESCRIÇÃO DAS TABELAS
O banco de dados postgres, no que tange ao esquema public, contém 8 tabelas que estão detalhadas abaixo. Para cada tabela é apresentado seu nome, sua descrição, e uma tabela que contém a relação de todas as colunas mostrando o name, o tipo, e a descrição. Nos casos em que o tipo por ma enum customizado, é mostrado também quais os possíveis valores a serem usados

### TEST
table for test propose

| Nome | Tipo | Aceita | Comentário |
| :--- | :--- | :----: | :--- |
| id | integer |  | sequencial unique identifier |
| name | text |  | name of test |

### EMPLOYE


| Nome | Tipo | Aceita | Comentário |
| :--- | :--- | :----: | :--- |
| id | integer |  |  |
| code | integer |  |  |
| ext | integer |  |  |
| nom | text |  |  |
| prenom | text |  |  |
| sexe | genre | FEMININ, MASCULIN |  |
| courriel | text |  |  |

### FORMATION
Table fact pour les formations

| Nome | Tipo | Aceita | Comentário |
| :--- | :--- | :----: | :--- |
| id | integer |  | Itentificateur unique pour la table formation |
| id_employe | integer |  | Itentificateur unique pour la table employe |
| id_emploi | integer |  | Itentificateur unique pour la table emploi |
| id_mosa | integer |  | Itentificateur unique pour la table mosa |
| id_periode | integer |  | Itentificateur unique pour la table periode |
| id_lieu | integer |  | Itentificateur unique pour la table lieu |
| id_ontologie | integer |  | Itentificateur unique pour la table ontologie |

### EMPLOI
Dimension pour les données données sur l'emploi

| Nome | Tipo | Aceita | Comentário |
| :--- | :--- | :----: | :--- |
| id | integer |  | Identificateur unique dans le domaine SGGA |
| code | integer |  | Code d'identification généré par SAGIR et utilisé comme identifiant au systèm |
| embauche | date |  | Date d'embauche d'employé dans cette role |
| status | status_emploi | PERMANENT, OCCASIONEL | Status d'emploi qui peut avoir les valeurs 'P' pour permanent et 'O' pour occasionel |
| taux_horaire | real |  | taux horaire à payer à l'employé |

### MOSA
Dimension pour les données sur ministères, organisme et structure administrative

| Nome | Tipo | Aceita | Comentário |
| :--- | :--- | :----: | :--- |
| id | integer |  | Identificateur unique dans le domaine SGGA |
| code | integer |  | Code d'identification généré par SAGIR et utilisé comme identifiant au systèm |
| nom | text |  | Nom d'organime selon SAGIR |

### PERIODE
Dimension pour les données de référence de temps de formation

| Nome | Tipo | Aceita | Comentário |
| :--- | :--- | :----: | :--- |
| id | integer |  | Identificateur unique dans le domaine SGGA |
| debut | date |  | Date de début de la formation |
| fin | date |  | Date du fin de la formation |
| duree_heures | integer |  | Dureé de la formation en heures |
| duree_jours | integer |  | Durée de la formation en jours |
| duree_reelle | integer |  | Durée réelle de la formation |

### LIEU
Dimention pour les données concernant le lieu de formation

| Nome | Tipo | Aceita | Comentário |
| :--- | :--- | :----: | :--- |
| id | integer |  | Identificateur unique dans le domaine SGGA |
| nom | text |  | Nom du lieu de formation |
| adresse | text |  | Adresse du lieu de formation |

### ONTOLOGIE
Dimension pour l'étude de l'être en tant qu'être en formation, qui permet de comprendre ses propriétés générales

| Nome | Tipo | Aceita | Comentário |
| :--- | :--- | :----: | :--- |
| id | integer |  | Identificateur unique dans le domaine SGGA |
| domaine | text |  | Le domaine de formation |
| sujet | text |  | Le sujet de formation |
| theme | text |  | Le theme de formation |
| sous_theme | text |  | Le sous theme de formation |

