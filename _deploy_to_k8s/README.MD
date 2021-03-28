## Using Mininkube on Windows

### Service deploy using kubectl 

0- Start bash

1- eval $(minikube docker-env) // using minikube registery instead of docker desktop. 

2- docker build -t go_identity_minikube -f services.identity/Dockerfile .

3- kubectl run go-identity --image=go_identity_minikube:latest --port=8081 --image-pull-policy='Never'

4- kubectl expose deployment go-identity --type=NodePort

### KrakenD deploy using kubectl 

0- cd _deploy_to_k8s

1- docker build -t k8s-krakend -f Dockerfile .

2- Create deployment and service ymls.

3- kubectl create -f deployment.yaml

4- kubectl create -f service.yaml