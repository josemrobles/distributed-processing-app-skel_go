
#!/bin/bash -e
cookbook_name=$1
namespace="$2"
if [ "$1" == "" ]
then
	echo "missing cookbook name as first argument"
	exit 1
fi
if [ "$namespace" != "" ]; then
	destination="s3://kt-opsworks/${namespace}/$cookbook_name.tar.gz"
else
	destination="s3://kt-opsworks/$cookbook_name.tar.gz"
fi
echo $destination
(
cd $cookbook_name
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
