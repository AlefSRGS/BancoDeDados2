FROM golang:1.23.2

# Definindo o diretório de trabalho no contêiner
WORKDIR /go/src/app

# Copiando o código fonte para o contêiner
COPY . .

# Instalando o dockerize
RUN apt-get update && apt-get install -y wget && \
    wget https://github.com/jwilder/dockerize/releases/download/v0.6.1/dockerize-linux-amd64-v0.6.1.tar.gz && \
    tar -xvzf dockerize-linux-amd64-v0.6.1.tar.gz && \
    mv dockerize /usr/local/bin/

# Expondo a porta da aplicação
EXPOSE 8080

# Compilando a aplicação Go
RUN go build -o main cmd/main.go

# Rodando a aplicação com o dockerize para aguardar o PostgreSQL
CMD ["dockerize", "-wait", "tcp://postgres_db:5432", "-timeout", "30s", "./main"]
