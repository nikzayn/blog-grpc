FROM golang:1.19

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /blog-grpc/bin/blog

EXPOSE 8080

CMD [ "/blog-grpc/bin/blog" ]
