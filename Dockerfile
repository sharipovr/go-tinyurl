# ---------- build ----------
FROM golang:1.22 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /go-tinyurl ./cmd/api

# ---------- runtime ----------
FROM gcr.io/distroless/static:nonroot
ENV PORT=8080
COPY --from=builder /go-tinyurl /go-tinyurl
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/go-tinyurl"]
