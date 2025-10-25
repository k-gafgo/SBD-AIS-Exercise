FROM golang:1.25 AS build
WORKDIR /app
COPY . .
RUN bash scripts/build-application.sh

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/ordersystem /ordersystem
ENTRYPOINT ["/ordersystem"]