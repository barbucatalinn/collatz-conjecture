FROM golang:alpine as builder
RUN apk update && apk add --no-cache ca-certificates tzdata git && update-ca-certificates
RUN adduser -D -g '' unprivileged_user
WORKDIR /code
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w' -o "/bin/app" cmd/app/main.go
FROM scratch as production
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /bin/app bin/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
USER unprivileged_user
HEALTHCHECK NONE
ENTRYPOINT ["/bin/app"]