cloud-monitoring
================

Cloud-based monitoring for business metrics using [Kubernetes](http://kubernetes.io/), [Minikube](https://github.com/kubernetes/minikube), [Influxdb](https://docs.influxdata.com/influxdb/v1.2/) and [Grafana](http://grafana.org/).

Furthermore, a go-script is provided for testing the system.

Requirements
-----
[Minikube](https://github.com/kubernetes/minikube) must be installed in order to run the monitoring.
[GoLang](https://golang.org/) is needed to run the provided test script.

Installation
-----
Start minikube, then run the install.sh script in the scripts- folder. This will install Indluxdb and Grafana on the VM, display the urls for the database and grafana, and create the database 'data' on Influxdb.

Next, open Grafana on your browser, sign in (default user and password: admin, admin) and add Influxdb as the data source. Default user and password: root, root. Database used by the go-script: data.

Usage
-----
Data can be added to Influxdb using the HTTP API. To test the connection, a go-script is provided. Add the database url to it, then build and run it as usual: go build && ./cloud-monitoring. The script will run until stopped (Ctrl+C), adding values to the database every second.
To see the values in Grafana, add a graph to a dashboard with query: SELECT "value" FROM "measurement.count";

License
-----
MIT, see the License-file.
