# sample-go-kubernetes
Criando aplicação exemplo para trabalhar com kubernetes

## How to run the server
```
docker build -t caioq/sample-go-kubernetes .
docker push caioq/sample-go-kubernetes
docker run --rm -p 8001:8001 caioq/sample-go-kubernetes
```

## How to create clusters with kind
```
kind create cluster --config=k8s/kind.yaml --name=sample-project-cluster
```
 
 ## How to create/run pods and load balancer service
 ```
 kubectl apply -f k8s/configmap-family.yaml
 kubectl apply -f k8s/deployment.yaml
 kubectl apply -f k8s/service.yaml
 kubectl apply -f k8s/hpa.yaml
 ```
 
 ## Using metric server
 ```
 kubectl apply -f k8s/metrics-server.yaml
 ```

 ## How to access service
 ```
 kubectl port-forward svc/goserver-service 8001:8001
 ```

## Stress Test with fortio
```
kubectl run -it --generator=run-pod/v1 fortio --rm --image=fortio/fortio -- load -qps 800 -t 120s -c 70 "http://goserver-service/healthz"
```