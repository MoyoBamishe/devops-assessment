FROM golang:1.21-alpine AS builder
WORKDIR /app

#  Copy and download the libraries
COPY go.mod ./
RUN go mod download

#  Copy the code
COPY . .

#  Build the app
RUN go build -o server ./hello.go

FROM alpine:latest
WORKDIR /root/

# Pick only the finished "server" file
COPY --from=builder /app/server .

EXPOSE 8080
CMD ["./server"]