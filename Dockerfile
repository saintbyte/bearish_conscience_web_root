FROM golang:1.16-alpine
ENV HTTP_PORT="1323"
EXPOSE 1323
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY cmd/server.go ./server.go
RUN go build -o server server.go
CMD [ "./server" ]