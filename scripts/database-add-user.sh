echo Adding user $3 to database $4 with $5 privileges
curl -i -XPOST $6/query -u $1:$2 --data-urlencode "q=GRANT $5 ON $4 TO $3"
