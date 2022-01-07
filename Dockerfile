#STAGE 1
FROM golang:1.17-alpine
RUN apk add build-base
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .
#CMD ["/app/main"]

#STAGE 2
FROM alpine:latest  
RUN apk --no-cache add ca-certificates
RUN mkdir /app-service
WORKDIR /app-service
COPY --from=0 /app ./
CMD ["./app"]
