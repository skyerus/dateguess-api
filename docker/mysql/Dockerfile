FROM mysql/mysql-server:8.0.26 AS base

COPY sql/init /docker-entrypoint-initdb.d/
COPY sql/data.sql /docker-entrypoint-initdb.d/
COPY sql/conf.d /etc/mysql/conf.d
