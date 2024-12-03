FROM golang:1.23-alpine

WORKDIR /app

COPY . .

ENV CGO_ENABLED=0
RUN go build -o /bin/hook360 ./cmd

FROM scratch

COPY --from=0 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=0 /bin/hook360 /bin/hook360

EXPOSE 8080

ENTRYPOINT ["hook360"]
CMD ["serve"]
