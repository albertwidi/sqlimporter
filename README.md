# SQLImporter

SQLImporter is not a database migration tools. This is tools is for importing and executing sql files into test database.

This works is inspired By Hendra Huang database integration test: https://github.com/Hendra-Huang/databaseintegrationtest.

## What is this for

To test integration with database. When we have a function that interract with database, the schema is imported and can be dropped afterwards.

## Sql file example

```sql
CREATE TABLE IF NOT EXISTS `something1` (
    `something1_id` bigint(20) NOT NULL
)

CREATE TABLE IF NOT EXISTS `something2` (
    `something2_id` bigint(20) NOT NULL,
    `field1` varchar(10) NOT NULL
)
```

## Command Line Tools

Sqlimporter provide CLI interface

This tools might useful if your environment don't have any postgresql/mysql command installed but have your database running in a container.

### Installing

`go get -u github.com/albert-widi/sqlimporter/cmd/sqlimporter`

### Available Commands

```shell
sqlimporter command line tools

Usage:
  sqlimporter [command]

Available Commands:
  help        Help about any command
  import      import postgresql/mysql schema from directory
  test        test command for sqlimporter

Flags:
  -h, --help      help for sqlimporter
  -v, --verbose   sqlimporter verbose output

Use "sqlimporter [command] --help" for more information about a command.
```

To import a schema into a database:

`sqlimporter import postgres --db book --host localhost --port 5432 -u logistic:logistic -f 'files/dbschema/book/'`

## Use It At Your Own Risk

This tools is nowhere near stable and the API might changed dramatically in the future, so use it at your own risk.