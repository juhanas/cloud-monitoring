echo Starting to install InfluxDB
kubectl run --image=influxdb influxdb --port=8086 --env="INFLUXDB_HTTP_AUTH_ENABLED=true"
kubectl expose deployment influxdb --type=NodePort --port=8086
echo Getting IP address for the container
ipDatabase=$(minikube service influxdb --url)
echo IP: $ipDatabase
echo Initializing database
echo Creating admin user: $1
curl -s -S -i -XPOST $ipDatabase/query --data-urlencode "q=CREATE USER $1 WITH PASSWORD '$2' WITH ALL PRIVILEGES" >/dev/null
echo Creating database $3
curl -s -S -i -XPOST $ipDatabase/query -u $1:$2 --data-urlencode "q=CREATE DATABASE $3" >/dev/null
echo Influxdb installed
