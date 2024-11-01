# installing official golang image from dockerhub 
FROM golang:1.23.2

# setting working directory in docker container
WORKDIR /app

# copying the source code to working directory
COPY . .

# setting app port
EXPOSE ${APP_PORT}

# building the application using entry point for make a executable main
RUN go build -o main cmd/main.go

# running the application using the built executable 
CMD ["./main"]