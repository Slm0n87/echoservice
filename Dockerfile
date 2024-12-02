FROM golang:1.23.3 as builder
WORKDIR /app

COPY . .
RUN go build -o bin/echoservice main.go

FROM gcr.io/distroless/base-debian11@sha256:3c8bc83024156c914b0b9455fa51ae6b0ca5318cfad0508d634e9f1844ef9d80

WORKDIR /app
COPY --from=builder /app/bin/echoservice .

EXPOSE 8080
ENTRYPOINT ["/app/echoservice"]
