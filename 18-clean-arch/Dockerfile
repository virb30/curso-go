FROM golang AS builder

WORKDIR /go/src

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o ./main ./cmd/ordersystem


FROM scratch

COPY --from=builder /go/src/cmd/ordersystem/.env .
COPY --from=builder /go/src/main .

CMD ["./main"]
