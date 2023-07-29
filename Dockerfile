FROM golang:1.20.6-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o fast-survey ./cmd/api

RUN chmod +x /app/fast-survey

FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/fast-survey /app

CMD [ "/app/fast-survey" ]