FROM golang:alpine
WORKDIR /build
COPY services.product services.product
COPY shared shared
WORKDIR /build/services.product
RUN go build -o main ./src/main.go
WORKDIR /dist
RUN cp /build/services.product/main .
RUN cp /build/services.product/src/app.env .
CMD ["/dist/main"]