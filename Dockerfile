# Etapa de build
FROM golang:1.22-alpine AS builder

# Configura o diretório de trabalho dentro do container
WORKDIR /app

# Copia os arquivos de configuração do Go
COPY go.mod go.sum ./
RUN go mod download

# Copia o código fonte da aplicação para o container
COPY . .

# Compila a aplicação Go
RUN go build -o goexample cmd/server/main.go

# Etapa de execução
FROM alpine:latest

# Cria um diretório para o app
WORKDIR /app

# Copia o binário compilado da etapa anterior
COPY --from=builder /app/goexample .

# Expõe a porta que será usada pela aplicação
EXPOSE 8080

# Comando que será executado ao iniciar o container
CMD ["./goexample"]