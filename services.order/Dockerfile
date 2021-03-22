FROM golang:alpine
WORKDIR /build
COPY services.order services.order
COPY shared shared
WORKDIR /build/services.order
RUN go build -o main ./src/main.go
WORKDIR /dist
RUN cp /build/services.order/main .
RUN cp /build/services.order/src/app.env .
CMD ["/dist/main"]