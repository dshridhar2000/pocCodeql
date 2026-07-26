# Stage 1: Build
FROM golang:1.25 AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o server .

# Stage 2: Run
FROM gcr.io/distroless/base-debian12
COPY --from=builder /app/server /server
CMD ["/server"]
