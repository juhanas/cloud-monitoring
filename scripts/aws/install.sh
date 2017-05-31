if [ "$#" -eq 4 ]
then
    ./influxdb-install-aws.sh $1 $2 $3
    ./grafana-install-aws.sh $4
else
echo "Wrong number of arguments. Required: influxAdminName, influxAdminPassword, databaseName, grafanaAdminPassword"
exit 1
fi
