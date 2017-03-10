echo Starting to install InfluxDB
kubectl run --image=influxdb influxdb --port=8086
kubectl expose deployment influxdb --type=NodePort --port=8086
echo Getting IP address for the container
ip=$(minikube service influxdb --url)
echo IP: $ip
echo Initializing tables
echo Creating table data
curl -i -XPOST $ip/query --data-urlencode "q=CREATE DATABASE data"
