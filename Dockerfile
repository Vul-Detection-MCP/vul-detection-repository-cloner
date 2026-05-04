FROM golang:1.26.0 AS build

LABEL author="biggujo"

WORKDIR "/src"

COPY . .

ENV CGO_ENABLED=0

RUN go build .

FROM alpine/git:v2.52.0

WORKDIR /app

COPY --from=build /src/repositorycloner /app/repositorycloner

ENTRYPOINT ["/app/repositorycloner", "/app/list.csv", "/app/git"]
