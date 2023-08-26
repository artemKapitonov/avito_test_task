FROM golang:latest

LABEL version='1.0'

COPY ./ ./

RUN go version

ENV GOPATH=/

# build go app
RUN go mod download

RUN go build -o avito-test-task ./cmd/avito-test-task

CMD [ "./avito-test-task" ]