echo Removing privilege $5 of user $3 from database $4
curl -i -XPOST $6/query -u $1:$2 --data-urlencode "q=REVOKE $5 ON $4 FROM $3"
