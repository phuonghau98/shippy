FROM alpine:latest

RUN mkdir /app
WORKDIR /app
ADD consignment_service /app/consignment_service

ENTRYPOINT ["./consignment_service"]