FROM golang:alpine as builder
WORKDIR $GOPATH/src/github.com/demget/casbin-example

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /go/bin/example

FROM alpine
WORKDIR /app
COPY model.conf .
COPY --from=builder /go/bin/example .
ENTRYPOINT ["/app/example"]
