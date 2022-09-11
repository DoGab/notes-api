# build stage
FROM golang:1.17-alpine3.14 AS builder
WORKDIR /app
COPY . .
RUN go build -o notes-api main.go

# run stage
FROM alpine:3.14
WORKDIR /app
COPY --from=builder /app/notes-api .

CMD [ "/app/notes-api" ]