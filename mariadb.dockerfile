FROM mariadb:latest
COPY ./mariadb/my.cnf /etc/mysql/my.cnf
COPY ./mariadb/init-db/01_create_tables.sql /docker-entrypoint-initdb.d/01_create_tables.sql