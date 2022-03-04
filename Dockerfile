FROM golang:1.17.8-bullseye
WORKDIR /app
COPY . .
CMD go run main.go
EXPOSE 8080
