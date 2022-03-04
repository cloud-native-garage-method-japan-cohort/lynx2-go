# FROM golang:1.17.8-bullseye
# WORKDIR /app
# COPY . .
# CMD export GOCACHE=/app;go run main.go
# EXPOSE 8080

# FROM golang:latest

# USER 1001

# RUN mkdir test
# COPY . test/

# EXPOSE 8080

# #CMD ["sh", "-c", "GOCACHE=off", "go", "run", "/go/test/main.go"]
# CMD export GOCACHE=off; cd /go/test/; go run main.go


FROM golang:latest

RUN mkdir /test
COPY . /test

RUN chgrp -R 0 /test && \
    chmod -R g=u /test

USER 1001
EXPOSE 8080

CMD export GOCACHE=/test/; cd /test; go run main.go


# FROM golang:1.17.7-alpine3.15
# WORKDIR /app
# COPY . .
# RUN go build main.go

# FROM alpine:3.15
# COPY --from=build /app/main /
# CMD ["/main"]
