FROM golang:1.21.4-alpine3.18 AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o server cmd/job-portal-api/main.go

# CMD [ "./server" ]

FROM alpine

WORKDIR /build

COPY --from=builder /app/server .

CMD [ "./server" ]