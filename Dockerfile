FROM golang:alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

# RUN go build -o /server
RUN go install github.com/air-verse/air@latest

EXPOSE 8000

# CMD [ "/server" ]
CMD ["air"]