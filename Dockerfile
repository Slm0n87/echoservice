# FROM golang:1.23.3 as builder
# WORKDIR /app

# COPY . .
# RUN go build -o bin/echoservice main.go

# FROM gcr.io/distroless/base-debian11@sha256:2eebbdaceeec718d203fd32370a1c1960b7d9d550d6acb1bca4957ab8560e3fb

# WORKDIR /app
# COPY --from=builder /app/bin/echoservice .

# EXPOSE 8080
# ENTRYPOINT ["/app/echoservice"]


FROM gcr.io/distroless/static-debian11:nonroot

COPY echoservice /echoservice

ENTRYPOINT ["/echoservice"]
