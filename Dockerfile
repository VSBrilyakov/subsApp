FROM golang:1.25

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN apt-get update
RUN apt-get -y install postgresql-client

RUN go mod download
RUN go build -o subsApp ./cmd/main.go

CMD ["./subsApp"]