FROM golang:1.22.1

WORKDIR /app

COPY go.mod go.sum ./

COPY cmd ./
COPY pkg ./

COPY pkg/database /usr/local/go/src/book_library/pkg/database
COPY pkg/services /usr/local/go/src/book_library/pkg/services
COPY pkg/handlers /usr/local/go/src/book_library/pkg/handlers
COPY pkg/entities /usr/local/go/src/book_library/pkg/entities
COPY pkg/my_errors /usr/local/go/src/book_library/pkg/my_errors

RUN go install book_library/pkg/database
RUN go install book_library/pkg/services
RUN go install book_library/pkg/handlers
RUN go install book_library/pkg/entities
RUN go install book_library/pkg/my_errors

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -o /booklib

EXPOSE 8080

ENV POSTGRES_HOST: "postgres"
ENV POSTGRES_PORT: "5432"
ENV POSTGRES_DATABASE: "library"
ENV POSTGRES_USER: "go_user"
ENV POSTGRES_PASSWORD: "P@ssw0rd"

CMD ["/booklib"]