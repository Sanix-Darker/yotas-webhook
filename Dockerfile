FROM golang:1.14

COPY . .

RUN mv .env.example .env

RUN go build -o yotas .

CMD ["./yotas"]
