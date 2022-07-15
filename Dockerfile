# build stage
FROM golang:alpine AS build-env
RUN apk --no-cache add git gcc bash

# copy the file instead of add we use copy so caching layer could do it better
WORKDIR /src

COPY go.mod /src
COPY go.sum /src
RUN go mod download

COPY . /src
# server
RUN env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ser /src/main.go


# final stage build using pure alpine image
FROM alpine:3.16
LABEL maintainer="prima"
RUN apk --no-cache add git gcc bash

# copy golang binary
WORKDIR /app
COPY --from=build-env /src/ser /app/ser
RUN /app/ser -s false


ENTRYPOINT [ "/app/ser","-s","true" ]