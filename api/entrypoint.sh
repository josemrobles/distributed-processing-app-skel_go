#!/bin/bash

echo $1 $2

go get
go build -o /usr/local/bin/api
supervisord -c /etc/supervisor/supervisord.conf
