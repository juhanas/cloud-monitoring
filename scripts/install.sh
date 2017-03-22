if [ "$#" -eq 4 ]
then
    ./influxdb-install.sh $1 $2 $3
    ./grafana-install.sh $4
else
echo "Wrong number of arguments. Required: influxAdminName, influxAdminPassword, databaseName, grafanaAdminPassword"
exit 1
fi
