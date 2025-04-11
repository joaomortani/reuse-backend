FROM golang:1.23.2

WORKDIR /app

# Instala o Air no PATH do sistema
RUN go install github.com/air-verse/air@latest && cp /go/bin/air /usr/local/bin/air

# Copia apenas arquivos essenciais primeiro
COPY go.mod go.sum ./
RUN go mod download

# Copia o restante
COPY . .

EXPOSE 8080

ENTRYPOINT ["/app/entrypoint.sh"]