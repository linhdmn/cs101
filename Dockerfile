# syntax=registry-gitlab.zalopay.vn/docker/images/dockerfile:experimental
FROM golang:1.17 as builder
ENV GO111MODULE=on
# Working directory
WORKDIR /app
# Copy files
COPY go.mod .
COPY go.sum .
# Install app dependencies
RUN go mod download
COPY . .
# Build app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./app

# final stage
FROM alpine:zlp-3.12
# Copy binary from builder
COPY --from=builder /app/app /app
# Run server command
ENV TZ Asia/Saigon
ENTRYPOINT ["/app"]
EXPOSE 8080
