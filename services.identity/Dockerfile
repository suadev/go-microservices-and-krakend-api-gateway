FROM golang:alpine
WORKDIR /build
COPY services.identity services.identity
COPY shared shared
WORKDIR /build/services.identity
RUN go build -o main ./src/main.go
WORKDIR /dist
RUN cp /build/services.identity/main .
RUN cp /build/services.identity/src/app.env .
CMD ["/dist/main"]