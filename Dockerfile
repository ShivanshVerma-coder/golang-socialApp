FROM golang:1.12.0-alpine3.9
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o /cmd/main/main .
CMD ["/app/cmd/main/main"]