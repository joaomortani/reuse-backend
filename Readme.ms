# Go Backend Boilerplate

Este é um boilerplate em Go pensado para aplicações backend modernas, com estrutura baseada em boas práticas e pronto para ser usado com frontends como React Native.

## ✨ Tecnologias
- Go
- Gin
- GORM + PostgreSQL
- JWT
- godotenv

## 📦 Estrutura
```
├── cmd/server         # Ponto de entrada do app
├── config             # Carrega variáveis de ambiente
├── database           # Conexão e migrations
├── models             # Structs GORM
├── dto                # Entrada e saída da API
├── repositories       # Acesso ao banco
├── services           # Regras de negócio
├── handlers           # Controllers HTTP
├── routes/v1          # Rotas organizadas por versão
├── .env               # Variáveis de ambiente
├── go.mod / go.sum
```

## 🚀 Como rodar
```bash
cp .env.example .env
# edite seu .env com dados reais

go mod tidy
go run ./cmd/server
```

## 📮 Endpoints
- `POST /api/v1/auth/register`
- `POST /api/v1/auth/login`

## 📌 To-do futuro
- Autenticação com middleware
- Hash de senha
- Refresh Token
- Testes automatizados
- Docker e CI/CD

---

## 🧑‍💻 Por João Mortani