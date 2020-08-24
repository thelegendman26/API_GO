FROM golang:1.12-alpine

WORKDIR /go/src/github.com/API_GO

COPY . .


RUN ["go", "get", "./..."]

#RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]
#ENTRYPOINT CompileDaemon -log-prefix=false -build="go build ./cmd/api/" -command="./api"

RUN go build -o bin ./cmd/bin/
CMD ["./bin"]