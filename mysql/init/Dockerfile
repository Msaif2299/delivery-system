FROM mysql:latest

COPY ./conf.d/custom.cnf /etc/mysql/conf.d/custom.cnf
RUN chmod 644 /etc/mysql/conf.d/custom.cnf
RUN chown mysql:mysql /etc/mysql/conf.d/custom.cnf

COPY ./drivers.sql /docker-entrypoint-initdb.d/