FROM golang:alpine
WORKDIR /build
COPY services.customer services.customer
COPY shared shared
WORKDIR /build/services.customer
RUN go mod tidy
RUN go build -o main ./src/main.go
WORKDIR /dist
RUN cp /build/services.customer/main .
RUN cp /build/services.customer/src/app.env .
CMD ["/dist/main"]