FROM golang:1.22.1

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

RUN go get github.com/lib/pq

COPY cmd ./
COPY pkg ./

COPY pkg/entities /usr/local/go/src/entities
COPY pkg/my_errors /usr/local/go/src/my_errors
COPY pkg/storages /usr/local/go/src/storages
COPY pkg/services /usr/local/go/src/services
COPY pkg/handlers /usr/local/go/src/handlers

RUN go install entities
RUN go install my_errors
RUN go install storages
RUN go install services
RUN go install handlers

RUN CGO_ENABLED=0 GOOS=linux go build -o /booklib

EXPOSE 8080

ENV POSTGRES_HOST: "postgres"
ENV POSTGRES_PORT: "5432"
ENV POSTGRES_DATABASE: "library"
ENV POSTGRES_USER: "go_user"
ENV POSTGRES_PASSWORD: "P@ssw0rd"

CMD ["/booklib"]