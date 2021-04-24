FROM golang:1.14
COPY . .
RUN go build -o yotas .
CMD ["./yotas"]
