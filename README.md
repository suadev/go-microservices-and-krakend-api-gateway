## Go Restful Microservices and KrakenD API Gateway Workshop

This is a shopping basket workshop that shows how to use KrakenD API Gateway.

Consist of 5 microservice and API Gateway.

Token based Authentication.

Raw access token is created by the Identity Service. (json model)

This raw access token is signed and verified by KrakendD API Gateway. (bearer token)

Services communicate with each other asynchronously through Kafka.

Order Service makes an HTTP Request to Customer Service to get basket items.

Customer Service keeps a local copy of a subset of user and product data.

## Run in Release Mode & Test

* Run **'docker-compose up'** in the root directory and wait for all containers to get ready.

* 13 containers will be created. Service containers (customer_service, identity_service, etc.) need to be started manually due to the need to wait for the Postgres container to get ready to accept DB connections. (wait-for-it.sh is used for the only Kafka for waiting zookeeper to get ready.)

* Execute **db_all.sql** to create databases of microservices. ( tables will be created by auto-migration )

* Use the Postman collection to play with KrakenD. ( _postman_collection/Go_Microservices_KrakenD.postman_collection.json )

## Containers

<img src="https://github.com/suadev/go-microservices-private/blob/main/_img/containers.JPG" width="270px" height="300px"></img>

## Postman Collection

Use **KrakenD-*** requests to play with KrakenD API Gateway. If you want to debug an individual service, run the service in debug mode and use the requests under the Direct Access folder as shown below.

<img src="https://github.com/suadev/go-microservices-private/blob/main/_img/postman_collection.JPG" width="300px" height="280"></img>

## Run in Debug Mode

In debug mode, you skip KrakenD, so KrakenD is not going to sign the bearer token and not going to inject in the request header as the **'User_Id'** key. You need to add the **'User_Id'** key as a request header for 'Add to Basket' and 'Create Order' requests. (Check the postman collection requests which are under the 'Direct Access' folder. )

For VsCode Users: Select **'All'** configuration on Debug panel and start debugging. All services will be up and running in debug mode. 

## Tool Set

* Go
* <a href="https://github.com/go-gorm/gorm">GORM</a>
* <a href="https://github.com/gin-gonic/gin">Gin</a>
* <a href="https://github.com/Shopify/sarama">Sarama</a>
* <a href="https://github.com/spf13/viper">Viper</a>
* <a href="https://github.com/dgrijalva/jwt-go">jwt-go</a>
* <a href="https://github.com/pact-foundation/pact-go">pact-go<a/>
* <a href="https://github.com/devopsfaith/krakend">KrakenD API Gateway</a>
* <a href="https://github.com/lpereira/lwan">Lwan HTTP/File Server</a>
* PostgreSQL
* Kafka - Zookeeper
* Docker - Docker Compose