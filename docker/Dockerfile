FROM golang:1.17 AS builder

WORKDIR /go/src/

COPY . .

RUN go mod download
RUN CGO_ENABLED=0 go build -o /go/bin/transactions main.go


FROM alpine:3.10

COPY --from=builder /go/bin/transactions /go/bin/transactions
COPY --from=builder /go/src/schemas/v1_schema.sql /go/bin/v1_schema.sql

ENTRYPOINT [ "/go/bin/transactions" ]