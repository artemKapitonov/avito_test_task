FROM golang:alpine AS builder

WORKDIR /usr/local/src

COPY ["go.mod", "go.sum", "./"]

RUN go mod download

COPY  ./ ./

RUN go build -o ./bin/avito-test-task ./cmd/avito-test-task

FROM alpine AS runner

COPY --from=builder /usr/local/src/bin/avito-test-task /

COPY configs/config.yml configs/config.yml

COPY migrations/schema migrations/schema

COPY .env .env

CMD [ "./avito-test-task" ]