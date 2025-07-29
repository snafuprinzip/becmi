FROM golang:1.24.5-alpine AS builder

WORKDIR /app

# Kopieren der Go-Module und deren Download
COPY go.mod go.sum ./
RUN go mod download

# Kopieren des Quellcodes
COPY . .

# Statisches Kompilieren des Servers
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server ./cmd/server/.

# Final-Stage
FROM scratch

# Kopieren der benötigten Dateien aus dem Builder
COPY --from=builder /app/server /server
COPY --from=builder /app/data /data
COPY --from=builder /app/active.*.toml /

# Port exponieren
EXPOSE 8080

# Server starten
CMD ["/server"]