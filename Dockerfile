FROM golang:1.20

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY main.go .
COPY internal ./internal
COPY cmd ./cmd

RUN CGO_ENABLED=0 GOOS=linux go build -o /monitoring-exporter

EXPOSE 8081

ENTRYPOINT ["/monitoring-exporter"]