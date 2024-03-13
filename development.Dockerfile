FROM golang:1.22.1

WORKDIR /app

COPY . .

RUN go mod download

CMD ["bash", "-c", "(cd migrations && go run . up) && go run ."]