# sample-go-kubernetes
Criando aplicação exemplo para trabalhar com kubernetes

## How to run the server
docker build -t caioq/sample-go-kubernetes .
docker run --rm -p 8001:8001 caioq/sample-go-kubernetes
docker push caioq/sample-go-kubernetes

## How to create clusters with kind
 kind create cluster --config=k8s/kind.yaml --name=sample-project-cluster
 
 ### How to create/run pod
 kubectl apply -f k8s/pod.yaml