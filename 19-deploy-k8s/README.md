## Comandos kubectl

kubectl cluster-info --context kind-goexpert

kubectl get nodes

kubectl apply -f k8s/deployment.yaml

kubectl get pods

kubectl delete pod podID

kubectl get svc

kubectl port-forward svc/serversvc 8080:8080 (redirecionamento de portas)