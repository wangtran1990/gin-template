RUN

normal mode
  go run main.go

production mode
  GIN_MODE=release go run main.go

live reloading mode using nodemon
  nodemon --exec go run main.go --signal SIGTERM
(install nodemon first via npm: npm install -g nodemon)

DOCKER

Build image
  docker build -t gingonic/template:0.0.1 -t gingonic/template:latest .
  docker build -t gingonic/template:latest .

Run image
  docker run --name running-gingonic -p 127.0.0.1:2000:2000 gingonic/template:latest

Run docker-compose
  docker-compose up