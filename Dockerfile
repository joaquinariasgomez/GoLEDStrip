FROM golang:1.20.3

WORKDIR /build

COPY go.mod .
COPY go.sum .
COPY . .

RUN go build -C src/ -o /app

EXPOSE 8888

CMD ["/app"]
#ENTRYPOINT ["tail", "-f", "/dev/null"]
