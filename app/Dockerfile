## Multistage build
FROM golang:1.18-alpine as build
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /src
COPY . .
RUN go mod download
RUN go build -o /app

## Distroless deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /
COPY --from=build /app /app

ENTRYPOINT ["/app"]