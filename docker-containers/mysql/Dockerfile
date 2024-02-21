FROM mysql:5.7


ENV MYSQL_ROOT_PASSWORD=ppassword
ENV MYSQL_DATABASE=dev

COPY ./my.cnf /etc/my.cnf

CMD ["mysqld", "--character-set-server=utf8mb4", "--collation-server=utf8mb4_unicode_ci"]