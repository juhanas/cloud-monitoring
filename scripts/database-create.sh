echo Creating database $3
curl -i -XPOST $4/query -u $1:$2 --data-urlencode "q=CREATE DATABASE $3"
