FROM golang

RUN rm -rf /app

RUN rm -rf .git

RUN rm -rf README.md

COPY . /app

ENV APP_PATH /app

WORKDIR $APP_PATH

RUN chmod a+x /app/*

ENTRYPOINT ["/app/start.sh"]

EXPOSE 2000