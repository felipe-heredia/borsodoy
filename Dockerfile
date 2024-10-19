FROM golang:1.23-alpine

WORKDIR /app

RUN apk add --no-cache gcc musl-dev

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./

RUN go mod download

RUN go env -w CGO_ENABLED=1

COPY . /app

CMD ["air", "-c", ".air.toml"]
