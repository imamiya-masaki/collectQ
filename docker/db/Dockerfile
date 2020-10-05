FROM postgres:12.1

RUN apt-get update
RUN apt-get -y install locales-all

RUN localedef -i ja_JP -c -f UTF-8 -A /usr/share/locale/locale.alias ja_JP.UTF-8
ENV LANG ja_JP.utf8

# ENV LANG ja_JP.UTF-8
# ENV LANGUAGE ja_JP:ja
# ENV LC_ALL ja_JP.UTF-8

# COPY mysqld_charset.cnf /etc/mysql/conf.d/mysqld_charset.cnf

# EXPOSE 3306