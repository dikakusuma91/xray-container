FROM golang:1.23-alpine AS pembangun
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /server main.go

FROM alpine:3.20
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=pembangun /server .
COPY .env.docker .env
COPY repo-xray/ repo-xray/
EXPOSE 8181
CMD ["./server"]
