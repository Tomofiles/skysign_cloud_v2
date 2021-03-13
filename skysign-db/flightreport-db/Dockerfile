FROM postgres:latest

RUN localedef -i ja_JP -c -f UTF-8 -A /usr/share/locale/locale.alias ja_JP.UTF-8

ENV LANG ja_JP.UTF-8

ENV TZ Asia/Tokyo

ENV POSTGRES_USER tomofiles
ENV POSTGRES_PASSWORD pc+tomofiles
ENV POSTGRES_DB flightreport

COPY *.sql /docker-entrypoint-initdb.d/