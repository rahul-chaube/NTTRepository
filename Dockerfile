FROM golang:1.22.3 AS baseImage
WORKDIR /app
COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o ./exoplanet

FROM alpine:latest

WORKDIR /app

COPY --from=baseImage /app/exoplanet ./
ENTRYPOINT ["./exoplanet"]