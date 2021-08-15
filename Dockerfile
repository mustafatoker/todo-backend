FROM golang:1.16-alpine

LABEL maintainer="Mustafa TOKER <iletisim@mustafatoker.com.tr>"

RUN apk update && apk add --no-cache git

RUN go version

WORKDIR /go/src/todo-backend

COPY go.mod go.sum ./

RUN go mod download

ADD ./ ./

RUN go build -o todo-app ./main.go

EXPOSE 8080

CMD ["./todo-app"]