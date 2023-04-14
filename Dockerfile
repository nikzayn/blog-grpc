FROM golang:1.19

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . /app/

RUN go build -o /blog-grpc

EXPOSE 8080

CMD [ "/blog-grpc" ]
