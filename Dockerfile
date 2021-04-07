FROM mcr.microsoft.com/dotnet/core/aspnet:3.1-buster-slim
RUN mkdir /app
RUN mkdir /data
RUN rm -f /etc/localtime \
&& ln -sv /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
&& echo "Asia/Shanghai" > /etc/timezone
COPY ./main /app
WORKDIR /app
RUN go build -o main . 
CMD ["/app/main"]
