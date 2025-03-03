# Usa a imagem oficial do Golang
FROM golang:1.24

# Define o diretório de trabalho dentro do container
WORKDIR /app

# Copia todos os arquivos do projeto para dentro do container
COPY . .

# Baixa as dependências do Go
RUN go mod tidy

# Compila o projeto
RUN go build -o main .

# Define o comando que será executado ao rodar o container
CMD ["/app/main"]
