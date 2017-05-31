echo Starting to install Grafana
kubectl --kubeconfig=kubeconfig run --image=grafana/grafana grafana --env="GF_SECURITY_ADMIN_PASSWORD=test"
kubectl --kubeconfig=kubeconfig expose deployment grafana --type=LoadBalancer --port=3000
echo Getting public IP address for the container
kubectl --kubeconfig=kubeconfig describe service grafana | grep .elb.amazonaws.com
echo Grafana installed
