echo Starting to install InfluxDB
kubectl --kubeconfig=kubeconfig run --image=influxdb influxdb --port=8086 --env="INFLUXDB_HTTP_AUTH_ENABLED=true"
kubectl --kubeconfig=kubeconfig expose deployment influxdb --type=LoadBalancer --port=8086
echo Getting private IP address for the container
ipDatabase
while [ $? -ne 0 ]; do
  ipDatabase=$(kubectl --kubeconfig=kubeconfig describe pod influxdb | grep IP | sed -E 's/IP:[[:space:]]+//')
done
echo Initializing database
echo Creating admin user: admin
curl -s -S -i -XPOST $ipDatabase/query --data-urlencode "q=CREATE USER admin WITH PASSWORD 'password' WITH ALL PRIVILEGES" >/dev/null
echo Creating database data
curl -s -S -i -XPOST $ipDatabase/query -u admin:password --data-urlencode "q=CREATE DATABASE data" >/dev/null
echo Getting public IP address for the container
kubectl --kubeconfig=kubeconfig describe service influxdb | grep .elb.amazonaws.com
echo Influxdb installed
