## Build
FROM golang:1.19-buster AS build

WORKDIR /build

COPY ./ ./

RUN go mod tidy

RUN go build -o /app

## Deploy
FROM golang:1.19-buster

WORKDIR /

COPY --from=build /app /app

ENV PORT "8080"

EXPOSE ${PORT}

ENTRYPOINT ["/app"]