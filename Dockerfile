FROM scratch
ADD ./bin/gp1_goapi /
CMD ["/gp1_goapi"]

#1. CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-w -s" -o ./bin/gp1_goapi main.go

#2. docker build -t gp1-goapi:1.1 .