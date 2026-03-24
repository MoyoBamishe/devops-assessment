FROM golang:1.23-alpine AS builder

RUN apk add --no-cache gcc musl-dev

WORKDIR /app

#  Copy and download the libraries
COPY go.mod go.sum* ./
RUN go mod download

#  Copy the code
COPY . .

#  Build the app
RUN CGO_ENABLED=1 GOOS=linux go build -o server ./hello.go

FROM alpine:latest
RUN apk add --no-cache ca-certificates libc6-compat

WORKDIR /app

# Pick only the finished "server" file
COPY --from=builder /app/server .
RUN chmod +x ./server

EXPOSE 8080
CMD ["./server"]
