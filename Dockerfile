FROM golang:alpine
RUN mkdir /app
COPY ./main /app
WORKDIR /app
RUN go build -o main . 
CMD ["/app/main"]
