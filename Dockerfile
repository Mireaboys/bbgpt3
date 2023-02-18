FROM golang:1.19-alpine3.16
WORKDIR /opt
COPY . .
RUN go mod tidy
RUN go build -o main main.go
CMD './main'