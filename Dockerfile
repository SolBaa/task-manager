FROM golang:1.17-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd/server


# Etapa de producción
FROM alpine:latest

# Crear un directorio en el contenedor
WORKDIR /root/

# Copiar el binario de la etapa de construcción
COPY --from=builder /app/main .

# Exponer el puerto en el que corre la aplicación
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./main"]
