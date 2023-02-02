FROM golang:1.19.5-alpine3.17
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o /main
EXPOSE 8081
CMD ["/main"]