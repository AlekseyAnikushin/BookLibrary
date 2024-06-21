FROM golang:1.22.1

WORKDIR /app

COPY cmd ./
COPY pkg ./

RUN go mod init github.com/AlekseyAnikushin/book_library
RUN go get github.com/lib/pq

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -o /booklib

EXPOSE 8080

ENV POSTGRES_HOST: "postgres"
ENV POSTGRES_PORT: "5432"
ENV POSTGRES_DATABASE: "library"
ENV POSTGRES_USER: "go_user"
ENV POSTGRES_PASSWORD: "P@ssw0rd"

CMD ["/booklib"]