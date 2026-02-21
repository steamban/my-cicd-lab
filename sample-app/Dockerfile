# First Stage
FROM golang:1.26.0-alpine as builder
WORKDIR /app
COPY main.go .
RUN RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myapp main.go

# -------------------------------------------------------------------------------------

# Second Stage
# Use scratch image since we have "CGO_ENABLED=0"
FROM scratch
WORKDIR /app
COPY --from=builder /app/myapp .

# Run as a non-root user
USER 12000

EXPOSE 8081
CMD ["./myapp"]
