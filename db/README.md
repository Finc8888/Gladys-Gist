# Run mysqldb into container

## Migrations
```bash
./db_run.sh migrations/gistbox_create_1.sql
```

## Seeds
```bash
./db_run.sh seeds/gists_insert_1.sql
```
## Check Users
```sql
SELECT user, host from mysql.user
```
## Show columns
```sql
DESC gists;
```

```sql
SHOW CREATE TABLE gists;
```

```sql
SHOW COLUMNS FROM gists;
```

```sql
EXPLAIN gists;
```

## Check grants
```sql
SHOW GRANTS FOR 'web'@'%';
```
## Run db in terminal
```bash
docker exec -it ill-mysql8 mysql -u <user> --password=<pass>
```
