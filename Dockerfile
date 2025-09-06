# Etapa de construcción
FROM golang:1.24 AS builder

WORKDIR /app

# Copia los archivos de dependencias y descarga
COPY go.mod go.sum ./
RUN go mod download

# Copia el resto del código
COPY . .

# Compila la aplicación
RUN go build -o main .

# Etapa de ejecución
FROM debian:bookworm-slim

WORKDIR /app
COPY --from=builder /app/main .

# Expone el puerto (usa la variable PORT para servicios como Koyeb/Render/Fly.io)
ENV PORT=8080
EXPOSE 8080

CMD ["./main"]
