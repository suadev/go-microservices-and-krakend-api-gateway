FROM golang:alpine
WORKDIR /build
COPY services.notification services.notification
COPY shared shared
WORKDIR /build/services.notification
RUN go build -o main ./src/main.go
WORKDIR /dist
RUN cp /build/services.notification/main .
RUN cp /build/services.notification/src/app.env .
CMD ["/dist/main"]