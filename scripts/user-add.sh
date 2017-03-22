echo Creating new user: $3
curl -i -XPOST $5/query -u $1:$2 --data-urlencode "q=CREATE USER $3 WITH PASSWORD '$4'"
