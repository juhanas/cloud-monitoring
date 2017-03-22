echo Dropping user $3
curl -i -XPOST $4/query -u $1:$2 --data-urlencode "q=DROP USER $3"
