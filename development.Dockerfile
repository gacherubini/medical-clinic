FROM golang:1.21.6

WORKDIR /app

COPY . .

RUN go mod download

CMD ["bash", "-c", "(cd migrations && go run . up) && go run ."]