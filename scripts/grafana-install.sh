echo Starting to install Grafana
kubectl run --image=grafana/grafana grafana --env="GF_SECURITY_ADMIN_PASSWORD=$1"
kubectl expose deployment grafana --type=NodePort --port=3000
echo Getting IP address for the container
ipGrafana=$(minikube service grafana --url)
echo IP: $ipGrafana
echo Grafana installed
