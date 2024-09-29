FROM golang:alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o /server

EXPOSE 8000

CMD [ "/server" ]