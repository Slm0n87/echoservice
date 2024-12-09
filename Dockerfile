FROM gcr.io/distroless/static-debian11:nonroot

COPY echoservice /echoservice

EXPOSE 8080
ENTRYPOINT ["/echoservice"]
