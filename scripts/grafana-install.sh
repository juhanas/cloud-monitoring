echo Starting to install Grafana
kubectl run --image=grafana/grafana grafana
kubectl expose deployment grafana --type=NodePort --port=3000
echo Getting IP address for the container
ip=$(minikube service grafana --url)
echo IP: $ip
