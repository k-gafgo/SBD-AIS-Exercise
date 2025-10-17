  # todo
 # 1st stage: builder
FROM golang:1.25 AS build
WORKDIR /app
COPY . .

RUN scripts/build-application.sh

FROM alpine
COPY --from=build /ordersystem/script/build-application.sh /ordersystem/script/
RUN go build main.go

