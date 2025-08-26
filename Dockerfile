# Usa la imagen oficial de Go para compilar
FROM golang:1.21 AS builder

# Establece el directorio de trabajo
WORKDIR /app

# Copia los archivos go.mod y go.sum y descarga dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copia el resto del código
COPY . .

# Compila la app en modo release
RUN go build -o main .

# Usa una imagen liviana para producción
FROM gcr.io/distroless/base-debian11

# Copia el binario desde el builder
COPY --from=builder /app/main /app/main

# Expone el puerto que tu app usará
EXPOSE 8080

# Comando para ejecutar la app
ENTRYPOINT ["/app/main"]
