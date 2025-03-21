FROM golang:1.22-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/bin/filestruc cmd/main/main.go 

FROM debian:bookworm-slim

WORKDIR /app
COPY --from=builder /app/bin/filestruc .  

CMD ["/app/filestruc"] 
