FROM golang:1.22.1

WORKDIR /app

COPY go.mod go.sum ./

COPY cmd ./
COPY pkg ./

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -o /booklib

EXPOSE 8080

ENV POSTGRES_HOST: "postgres"
ENV POSTGRES_PORT: "5432"
ENV POSTGRES_DATABASE: "library"
ENV POSTGRES_USER: "go_user"
ENV POSTGRES_PASSWORD: "P@ssw0rd"

CMD ["/booklib"]