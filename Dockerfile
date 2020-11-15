FROM golang:1.15

LABEL from="Bolin Clement" email="clement.bolin@epitech.eu"

WORKDIR /app

COPY . .

RUN go build -o blockChain
RUN ./blockChain

EXPOSE 8080
