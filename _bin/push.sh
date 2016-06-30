
#!/bin/bash -e
cookbook_name=$1
namespace="$2"
location="$3"
if [ "$1" == "" ]
then
	echo "missing cookbook name as first argument"
	exit 1
fi
if [ "$namespace" != "" ]; then
	destination="s3://${location}/${namespace}/$cookbook_name.tar.gz"
else
	destination="s3://${location}/$cookbook_name.tar.gz"
fi
echo $destination
(
cd ../_cookbooks
if [ -d ../_tmp/ ]; then
	targz="../_tmp/$cookbook_name.tar.gz"
else
	targz="../../_tmp/$cookbook_name.tar.gz"
fi

berks package $targz
s3cmd put $targz "${destination}"
s3cmd get --continue "${destination}" /dev/null
echo "${destination}"
)
