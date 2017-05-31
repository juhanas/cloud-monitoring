cloud-monitoring
================

Cloud-based monitoring for business metrics using [Kubernetes](http://kubernetes.io/), [Minikube](https://github.com/kubernetes/minikube), [Influxdb](https://docs.influxdata.com/influxdb/v1.2/) and [Grafana](http://grafana.org/).

[Go](https://golang.org/)- and [Node.js](https://nodejs.org)-scripts are provided for testing the system. Dashboards can be generated from a template with another go-script.

Requirements
-----
[Minikube](https://github.com/kubernetes/minikube) must be installed in order to run the monitoring locally. The tool can also be used in a cloud environment. Instructions are provided for AWS.
[GoLang](https://golang.org/) or [Node.js](https://nodejs.org) is needed to run the provided test script.

Installation
-----

### Local
Start minikube with `minikube start` then run the install.sh- script in the scripts-folder. This will install Indluxdb and Grafana on the VM, create admin user and database for Influxdb as well as display the url for the database and Grafana.

Next, open Grafana on your browser (user and password as specified in the scripts) and add Influxdb as the data source. Finally, create a new Dashboard or upload the generated file (see below).

### Cloud (AWS)

Install a server with Kubernetes. See for example Heptio's [Quick start quide](https://aws.amazon.com/quickstart/architecture/heptio-kubernetes/).
SSH to the server and run the install.sh- script in the scripts/aws-folder. The script will output the public ip-addresses for the Influxdb and Grafana instances.

Next, open Grafana on your browser (user and password as specified in the scripts) and add Influxdb as the data source. Finally, create a new Dashboard or upload the generated file (see below).

Usage
-----
Data can be added to Influxdb via the http API. To test the connection, go- and node.js- scripts are provided. Build and run one of them as usual, giving the required parameters:
```
go build && ./go-script [influxdb IP address] [database name] [username] [password]
```
```
npm install && node index.js [influxdb IP address] [influxdb port] [database name] [username] [password]
```
The script will run until stopped (Ctrl+C), adding values to the database every 1-5 seconds.

To see the values posted by the script, an example Grafana dashboard is provided for each script. Import Dashboard[Go/Node].json to Grafana and select the correct data source.

Dashboard Generation
-----
Custom Dashboards can be generated using the provided script. Start by filling out the definitions.json- file. Next, run the script
```
go build && ./generateDashboard
```
This will generate a dashboardName.json file, which can be imported to Grafana.

### Troubleshooting
If Grafana is not accepting the generated file, the definitions-file most likely has errors.
Check at least the following:
- All the required fields have been entered correctly
- Row and Column numbers are correct
  * They should start from 1, and not skip numbers
- The widths are correct
  * Each element can have a width of 1-12
  * The total width for all elements in a row can not exceed 12

Scripts
-----
### database-add-user
Gives the user access to the database
```
./database-add-user.sh [adminName] [adminPassword] [username] [databaseName] [privileges: read|write|all] [databaseIP]
```

### database-create
Creates a new database
```
./database-create.sh [adminName] [adminPassword] [databaseName] [databaseIP]
```

### database-revoke-user
Removes privileges from the user
```
./database-revoke-user.sh [adminName] [adminPassword] [username] [databaseName] [privilegeToRevoke: read|write|all] [databaseIP]
```

### delete-all
Deletes all Influxdb and Grafana instances from Minikube
```
./delete-all.sh
```

### install.sh
Installs the Influxdb and Grafana instances on Minikube
```
./install [adminName] [adminPassword] [defaultDatabaseName] [grafanaAdminPassword]
```

### aws/install.sh
Installs the Influxdb and Grafana instances on AWS
```
./install [adminName] [adminPassword] [defaultDatabaseName] [grafanaAdminPassword]
```

### user-add
Creates a new normal user with the given password
```
./user-create.sh [adminName] [adminPassword] [username] [password] [databaseIP]
```

### user-drop
Removes the given user
```
./user-drop.sh [adminName] [adminPassword] [username] [databaseIP]
```

License
-----
MIT, see the License-file.
