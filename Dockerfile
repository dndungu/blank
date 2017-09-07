FROM golang:1.9 as builder

WORKDIR /app

COPY main.go .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM scratch

WORKDIR /app

COPY --from=builder /app/app .

COPY page.tpl .

ENV BLANKPAGE_PORT 80

EXPOSE ${BLANKPAGE_PORT}

CMD ["./app"]
