FROM golang:alpine
RUN mkdir /app
RUN mkdir /data
COPY ./main /app
WORKDIR /app
RUN go build -o main . 
CMD ["/app/main"]
