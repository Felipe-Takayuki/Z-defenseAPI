FROM golang:1.20

WORKDIR /app

COPY . . 

RUN go mod download

RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o app 

EXPOSE 8080

CMD ["./app"]