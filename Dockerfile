FROM golang:1.18 as builder

WORKDIR /go/src/app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .
RUN make build

FROM registry.access.redhat.com/ubi8-minimal

WORKDIR /app

COPY --from=builder /go/src/app/sample-service .
CMD ["/app/sample-service"]
