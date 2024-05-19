docker image rm secap-gw:latest
docker build -t secap-gw:latest .

minikube image rm secap-gw:latest
minikube image load secap-gw:latest

kubectl apply -f .k8s/deployment.yaml -n secap-compass