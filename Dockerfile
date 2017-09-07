FROM golang:1.9 as builder

WORKDIR /app

COPY main.go .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM scratch

ENV BLANKPAGE_PORT 80

COPY --from=builder /app/app .

EXPOSE ${BLANKPAGE_PORT}

CMD ["./app"]
