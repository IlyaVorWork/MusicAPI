FROM golang:1.23.6-alpine3.21 AS build
WORKDIR /MusicAPI
COPY go.mod go.sum ./
RUN go mod download
COPY . .
WORKDIR /MusicAPI/cmd/app
RUN go build -o build

FROM alpine:3.18 AS prod
COPY --from=build /MusicAPI /.
EXPOSE 8080
ENTRYPOINT ["/cmd/app/build"]