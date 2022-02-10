FROM golang:1.17 AS builder

WORKDIR /go/src/

COPY . .

RUN make build

FROM scratch

COPY --from=builder /go/bin/transactions /go/bin/transactions

ENTRYPOINT [ "/go/bin/transactions" ]