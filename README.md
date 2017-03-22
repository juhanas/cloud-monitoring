cloud-monitoring
================

Cloud-based monitoring for business metrics using [Kubernetes](http://kubernetes.io/), [Minikube](https://github.com/kubernetes/minikube), [Influxdb](https://docs.influxdata.com/influxdb/v1.2/) and [Grafana](http://grafana.org/).

[Go](https://golang.org/)-script is provided for testing the system.

Requirements
-----
[Minikube](https://github.com/kubernetes/minikube) must be installed in order to run the monitoring.
[GoLang](https://golang.org/) is needed to run the provided test script.

Installation
-----
Start minikube with `minikube start` then run the install.sh script in the scripts-folder. This will install Indluxdb and Grafana on the VM, create admin user and database for Influxdb as well as display the url for the database and Grafana.

Next, open Grafana on your browser (user admin, password as specified in scripts) and add Influxdb as the data source.

Usage
-----
Data can be added to Influxdb via the http API. To test the connection, a go-script is provided. Build and run it as usual, giving the required parameters:
```go build && ./cloud-monitoring [influxdb IP address] [database name] [username] [password]
```
The script will run until stopped (Ctrl+C), adding values to the database every second.

To see the values posted by the go-script, add a graph to a dashboard in Grafana, select the correct data source and use query:
```SELECT "value" FROM "measurement.count";
```

Scripts
-----
### database-add-user
Gives the user access to the database
```./database-add-user.sh [adminName] [adminPassword] [username] [databaseName] [privileges: read|write|all] [databaseIP]
```

### database-create
Creates a new database
```./database-create.sh [adminName] [adminPassword] [databaseName] [databaseIP]
```

### database-revoke-user
Removes privileges from the user
```./database-revoke-user.sh [adminName] [adminPassword] [username] [databaseName] [privilegeToRevoke: read|write|all] [databaseIP]
```

### delete-all
Deletes all Influxdb and Grafana instances from Minikube
```./delete-all.sh
```

### install.sh
Installs the Influxdb and Grafana instances on Minikube
```./install [adminName] [adminPassword] [defaultDatabaseName] [grafanaAdminPassword]
```

### user-add
Creates a new normal user with the given password
```./user-create.sh [adminName] [adminPassword] [username] [password] [databaseIP]
```

### user-drop
Removes the given user
```./user-drop.sh [adminName] [adminPassword] [username] [databaseIP]
```

License
-----
MIT, see the License-file.
