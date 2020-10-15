FROM golang

RUN apt-get update

RUN rm -rf /app

COPY . /app

ENV APP_PATH /app

WORKDIR $APP_PATH

RUN chmod a+x /app/*

RUN \
  go get -d -v \
  && go install -v \
  && go build

ENTRYPOINT ["/app/start.sh"]

EXPOSE 2000