FROM golang:1.17
COPY . /backend
WORKDIR /backend
RUN go mod download
RUN go build -o main .
EXPOSE 9001
CMD ["./main"]
